package helper

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TxFunc func(tx pgx.Tx) error

type Uow struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Uow {
	return &Uow{db: db}
}

func (u *Uow) WithTx(ctx context.Context, fn TxFunc) error {
	tx, err := u.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			_ = tx.Rollback(ctx)
			panic(r)
		}
	}()

	if err := fn(tx); err != nil {
		_ = tx.Rollback(ctx)
		return err
	}

	return tx.Commit(ctx)
}
