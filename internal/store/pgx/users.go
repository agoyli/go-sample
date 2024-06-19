package pgx

import (
	"context"
	"main/internal/models"
	"strconv"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

const sqlUserTable = `users u`
const sqlUserKeys = `u.id, u.first_name, u.last_name, u.middle_name, u.username, u.password, u.status, u.phone, u.phone_verified_at, u.email, u.email_verified_at, u.last_active, u.created_at, u.updated_at`
const sqlUserOrder = `order by u.id`

func scanUser(rows pgx.Row, model *models.User, addColumns ...interface{}) error {
	return rows.Scan(parseColumnsForScan(model, addColumns...)...)
}

func (a Access) UserFindById(ctx context.Context, id uint) (model *models.User, err error) {
	// logger update ctx
	err = a.runQuery(ctx, func(tx *pgxpool.Conn) (err error) {
		err = scanUser(tx.QueryRow(ctx,
			`select `+sqlUserKeys+` from `+sqlUserTable+` where id=$1`, id),
			model)
		return nil
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
	sqlWheres := "id=id"
	sqlArgs := []interface{}{}
	if v, ok := opts["school_id"]; ok {
		sqlArgs = append(sqlArgs, v)
		sqlWheres += `u.school_id=$` + strconv.Itoa(len(sqlArgs))
	}
	if v, ok := opts["search"]; ok {
		sqlArgs = append(sqlArgs, v)
		sqlWheres += `u.first_name like $` + strconv.Itoa(len(sqlArgs))
	}

	err = a.runQuery(ctx, func(tx *pgxpool.Conn) (err error) {
		sql := `select ` + sqlUserKeys + `, count(u.id) over() from ` + sqlUserTable + ` where ` + sqlWheres + ` ` + sqlUserOrder
		rows, err := tx.Query(ctx, sql, sqlArgs...)
		if err != nil {
			return err
		}
		if rows.Next() {
			model := &models.User{}
			err = scanUser(rows, model, list.Total)
			if err != nil {
				return err
			}
			list.Users = append(list.Users, model)
		}
		return nil
	})
	if err != nil {
		// describe err and log
	}
	return
}
