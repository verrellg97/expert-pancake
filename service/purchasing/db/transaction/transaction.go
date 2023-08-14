package db

import (
	"context"
	"database/sql"
	"fmt"

	. "github.com/expert-pancake/service/purchasing/db/sqlc"
)

type PurchasingTrx interface {
	Querier
	UpsertPurchaseOrderTrx(ctx context.Context, arg UpsertPurchaseOrderTrxParams) (UpsertPurchaseOrderTrxResult, error)
	UpsertPurchaseOrderItemTrx(ctx context.Context, arg UpsertPurchaseOrderItemTrxParams) (UpsertPurchaseOrderItemTrxResult, error)
	UpdatePurchaseOrderItemsTrx(ctx context.Context, arg UpdatePurchaseOrderItemsTrxParams) (UpdatePurchaseOrderItemsTrxResult, error)
	UpdatePurchaseOrderStatusTrx(ctx context.Context, arg UpdatePurchaseOrderStatusTrxParams) (UpdatePurchaseOrderStatusTrxResult, error)
	UpdateReceiptOrderItemsTrx(ctx context.Context, arg UpdateReceiptOrderItemsTrxParams) error
	UpsertReceiptOrderTrx(ctx context.Context, arg UpsertReceiptOrderTrxParams) (UpsertReceiptOrderTrxResult, error)
	UpdateReceiptOrderStatusTrx(ctx context.Context, arg UpdateReceiptOrderStatusTrxParams) (UpdateReceiptOrderStatusTrxResult, error)
	UpsertPurchaseInvoiceTrx(ctx context.Context, arg UpsertPurchaseInvoiceTrxParams) error
}

// Trx provides all functions to execute SQL queries and transactions
type Trx struct {
	db *sql.DB
	*Queries
}

func NewPurchasingTrx(db *sql.DB) PurchasingTrx {
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
