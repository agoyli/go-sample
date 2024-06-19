package pgx

import (
	"context"
	"fmt"
	"log"
	"main/config"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Access struct {
	pool *pgxpool.Pool
}

func (d Access) Pool() *pgxpool.Pool {
	return d.pool
}

func (d *Access) Close() {
	d.pool.Close()
}

type pgxWithTx func(tx pgx.Tx) (rollback bool, err error)
type pgxQuery func(conn *pgxpool.Conn) (err error)

func (d *Access) runQuery(ctx context.Context, f pgxQuery) (err error) {
	err = d.Pool().AcquireFunc(ctx, f)
	if err != nil {
		return err
	}
	return
}

func Connect() *Access {
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable connect_timeout=5", config.App.DbUsername, config.App.DbDatabase, config.App.DbPassword, config.App.DbHost, config.App.DbPort)

	cfg, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		log.Fatal(err)
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), cfg)

	if err != nil {
		log.Fatal(err)
	}
	return &Access{pool: pool}
}
