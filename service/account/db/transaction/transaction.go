package db

import (
	"context"
	"database/sql"
	"fmt"
	. "github.com/expert-pancake/service/account/db/sqlc"
)

type AccountTrx interface {
	Querier
	CreateNewUserTrx(ctx context.Context, arg CreateNewUserTrxParams) (CreateNewUserTrxResult, error)
	UpdateUserTrx(ctx context.Context, arg UpdateUserTrxParams) (UpdateUserTrxResult, error)
}

// Trx provides all functions to execute SQL queries and transactions
type Trx struct {
	db *sql.DB
	*Queries
}

func NewAccountTrx(db *sql.DB) AccountTrx {
	return &Trx{
		db:      db,
		Queries: New(db),
	}
}

// ExecTx executes a function within a database transaction
func (trx *Trx) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := trx.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
