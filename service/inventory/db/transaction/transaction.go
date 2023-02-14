package db

import (
	"context"
	"database/sql"
	"fmt"

	. "github.com/expert-pancake/service/inventory/db/sqlc"
)

type InventoryTrx interface {
	Querier
	AddItemTrx(ctx context.Context, arg AddItemTrxParams) (AddItemTrxResult, error)
	UpdateItemTrx(ctx context.Context, arg UpdateItemTrxParams) (UpdateItemTrxResult, error)
	AddInternalStockTransferTrx(ctx context.Context, arg AddInternalStockTransferTrxParams) (AddInternalStockTransferTrxResult, error)
	AddUpdateStockTrx(ctx context.Context, arg AddUpdateStockTrxParams) (AddUpdateStockTrxResult, error)
}

// Trx provides all functions to execute SQL queries and transactions
type Trx struct {
	db *sql.DB
	*Queries
}

func NewInventoryTrx(db *sql.DB) InventoryTrx {
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
