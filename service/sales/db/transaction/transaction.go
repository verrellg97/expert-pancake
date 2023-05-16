package db

import (
	"context"
	"database/sql"
	"fmt"

	. "github.com/expert-pancake/service/sales/db/sqlc"
)

type SalesTrx interface {
	Querier
	UpserPOSTrx(ctx context.Context, arg UpsertPOSTrxParams) (UpsertPOSTrxResult, error)
}

// Trx provides all functions to execute SQL queries and transactions
type Trx struct {
	db *sql.DB
	*Queries
}

func NewSalesTrx(db *sql.DB) SalesTrx {
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
