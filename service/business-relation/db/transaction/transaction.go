package db

import (
	"context"
	"database/sql"
	"fmt"

	. "github.com/expert-pancake/service/business-relation/db/sqlc"
)

type BusinessRelationTrx interface {
	Querier
	CreateNewContactGroupTrx(ctx context.Context, arg CreateNewContactGroupTrxParams) (CreateNewContactGroupTrxResult, error)
	UpdateContactGroupTrx(ctx context.Context, arg UpdateContactGroupTrxParams) (UpdateContactGroupTrxResult, error)
	AddDefaultContactBookTrx(ctx context.Context, arg AddDefaultContactBookTrxParams) error
	CreateNewContactBookTrx(ctx context.Context, arg CreateNewContactBookTrxParams) (CreateNewContactBookTrxResult, error)
	UpdateContactBookTrx(ctx context.Context, arg UpdateContactBookTrxParams) (UpdateContactBookTrxResult, error)
	UpdateCustomerTrx(ctx context.Context, arg UpdateCustomerTrxParams) (UpdateCustomerTrxResult, error)
	UpdateSupplierTrx(ctx context.Context, arg UpdateSupplierTrxParams) (UpdateSupplierTrxResult, error)
}

// Trx provides all functions to execute SQL queries and transactions
type Trx struct {
	db *sql.DB
	*Queries
}

func NewBusinessRelationTrx(db *sql.DB) BusinessRelationTrx {
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
