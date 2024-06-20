package pgx

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"main/internal/models"
	"main/internal/utils"
	"slices"
	"strings"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

const sqlUserTable = `users u`
const sqlUserOrder = `order by u.id`

// scan keys (separated by new line by 5 keys)
const sqlUserKeys = `u.id, u.first_name, u.last_name, u.middle_name, u.username, 
	u.password, u.status, u.phone, u.phone_verified_at, u.email, 
	u.email_verified_at, u.last_active, u.created_at, u.updated_at`

func scanUser(rows pgx.Row, model *models.User, addColumns ...interface{}) error {
	scanColumns := append([]interface{}{
		&model.Id, &model.FirstName, &model.LastName, &model.MiddleName, &model.Username,
		&model.Password, &model.Status, &model.Phone, &model.PhoneVerifiedAt, &model.Email,
		&model.EmailVerifiedAt, &model.LastActive, &model.CreatedAt, &model.UpdatedAt}, addColumns...)
	return rows.Scan(scanColumns...)
}

func (a Access) UserFindById(ctx context.Context, id uint) (model *models.User, err error) {
	// logger update ctx
	err = a.runQuery(ctx, func(tx *pgxpool.Conn) (err error) {
		sql := fmt.Sprintf(`select %v from %v where id=$1`, sqlUserKeys, sqlUserTable)
		err = scanUser(tx.QueryRow(ctx, sql, id),
			model)
		return
	})
	if err != nil {
		// describe err and log
	}
	return
}

func (a Access) UserInsert(ctx context.Context, data *models.User) (model *models.User, err error) {
	// logger update ctx

	dataMap := map[string]interface{}{}
	dataStr, _ := json.Marshal(data)
	err = json.Unmarshal(dataStr, &dataMap)
	if err != nil {
		return
	}

	// set sql insert values
	sqlKeys := ""
	sqlValues := ""
	sqlArgs := []interface{}{}
	for key, value := range dataMap {
		if value == nil || slices.Contains([]string{"id"}, key) {
			continue
		}
		sqlArgs = append(sqlArgs, value)
		sqlValues += fmt.Sprintf(", $%v", len(sqlArgs))
		sqlKeys += fmt.Sprintf(", %v", key)
	}
	sqlKeys = strings.Trim(sqlKeys, ", ")
	sqlValues = strings.Trim(sqlValues, ", ")

	// run query
	err = a.runQuery(ctx, func(tx *pgxpool.Conn) (err error) {
		sqlTableParts := strings.Split(sqlUserTable, " ")
		sqlTable := sqlTableParts[0]
		sql := fmt.Sprintf(`insert into %v (%v) values (%v)`, sqlTable, sqlKeys, sqlValues)
		log.Println(sql)
		_, err = tx.Exec(ctx, sql, sqlArgs...)
		return
	})
	if err != nil {
		// describe err and log
	}
	return
}

func (a Access) UserUpdate(ctx context.Context, data *models.User) (model *models.User, err error) {
	// logger update ctx

	dataMap := map[string]interface{}{}
	dataStr, _ := json.Marshal(data)
	err = json.Unmarshal(dataStr, &dataMap)
	if err != nil {
		return
	}
	id := data.Id

	// set sql insert values
	sqlSets := ""
	sqlArgs := []interface{}{id}
	for key, value := range dataMap {
		if value == nil || slices.Contains([]string{"id"}, key) {
			continue
		}
		sqlArgs = append(sqlArgs, value)
		sqlSets += fmt.Sprintf("%v=$%v", key, len(sqlArgs))
	}

	// run query
	err = a.runQuery(ctx, func(tx *pgxpool.Conn) (err error) {
		sqlTableParts := strings.Split(sqlUserTable, " ")
		sqlTable := sqlTableParts[0]
		sql := fmt.Sprintf(`update %v set %v where id=$1`, sqlTable, sqlSets)
		_, err = tx.Exec(ctx, sql, sqlArgs...)
		return
	})
	if err != nil {
		// describe err and log
	}
	return
}

func (a Access) UserFindBy(ctx context.Context, opts map[string]interface{}) (list *models.Users, err error) {
	// logger update ctx
	list = &models.Users{
		Users: []*models.User{},
	}

	// set sql wheres
	sqlWheres := "1=1"
	sqlArgs := []interface{}{}
	sqlLimit := 12
	sqlOffset := 0
	if v, ok := opts["ids"]; ok {
		sqlArgs = append(sqlArgs, v)
		sqlWheres += fmt.Sprintf(` AND u.id=ANY($%v)`, len(sqlArgs))
	}
	if v, ok := opts["school_id"]; ok {
		sqlArgs = append(sqlArgs, v)
		sqlWheres += fmt.Sprintf(` AND u.school_id=$%v`, len(sqlArgs))
	}
	if v, ok := opts["search"]; ok {
		sqlArgs = append(sqlArgs, v)
		sqlWheres += fmt.Sprintf(` AND u.first_name like $%v`, len(sqlArgs))
	}
	if v, ok := opts["limit"].(int); ok {
		sqlLimit = v
	}
	if v, ok := opts["offset"].(int); ok {
		sqlOffset = v
	}

	// run query
	err = a.runQuery(ctx, func(tx *pgxpool.Conn) (err error) {
		sqlAppend := fmt.Sprintf(`%v LIMIT %v OFFSET %v`, sqlUserOrder, sqlLimit, sqlOffset)
		sql := fmt.Sprintf(`select %v, count(u.id) over() from %v where %v %v`, sqlUserKeys, sqlUserTable, sqlWheres, sqlAppend)
		rows, err := tx.Query(ctx, sql, sqlArgs...)
		if err != nil {
			return err
		}
		if rows.Next() {
			model := &models.User{}
			err = scanUser(rows, model, &list.Total)
			if err != nil {
				return err
			}
			list.Users = append(list.Users, model)
		}
		return
	})
	if err != nil {
		// describe err and log
	}
	return
}

func (a Access) UserDelete(ctx context.Context, ids []uint) (list *models.Users, err error) {
	// logger update ctx

	list, err = a.UserFindBy(ctx, map[string]interface{}{
		"ids": ids,
	})
	if err != nil {
		return nil, err
	}
	if len(list.Users) < 1 {
		return nil, utils.ErrNotfound
	}

	// run query
	err = a.runQuery(ctx, func(tx *pgxpool.Conn) (err error) {
		sqlTableParts := strings.Split(sqlUserTable, " ")
		sqlTable := sqlTableParts[0]
		sql := fmt.Sprintf(`delete from %v where id=ANY($1)`, sqlTable)
		_, err = tx.Exec(ctx, sql, ids)
		return
	})
	if err != nil {
		// describe err and log
	}
	return
}
