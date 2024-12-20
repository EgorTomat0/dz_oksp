package pgsql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"time"
)

type StorageCfg struct {
	Uname, Password, Host, Port, DbName string
}

type Conn interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}

func NewConn(ctx context.Context, sc StorageCfg) (*pgx.Conn, error) {
	dbs := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", sc.Uname, sc.Password, sc.Host, sc.Port, sc.DbName)
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	conn, err := pgx.Connect(ctx, dbs)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
