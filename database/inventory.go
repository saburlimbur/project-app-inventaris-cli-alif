package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type PgxIface interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
}

func InitDB() (*pgx.Conn, error) {
	connStr := "user=postgres password=secret host=localhost port=5432 dbname=inventory sslmode=disable"

	conn, err := pgx.Connect(context.Background(), connStr)
	return conn, err
}
