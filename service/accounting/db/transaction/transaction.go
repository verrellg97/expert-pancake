package db

import (
	"context"
	"database/sql"
	"fmt"

	. "github.com/expert-pancake/service/accounting/db/sqlc"
	db "github.com/expert-pancake/service/accounting/db/sqlc"
)

type AccountingTrx interface {
	Querier
	CreateNewCashTransactionTrx(ctx context.Context, arg db.InsertCashTransactionParams) (CreateNewCashTransactionTrxResult, error)
	AddDefaultCompanyChartOfAccountTransactionTrx(ctx context.Context, arg AddDefaultCompanyChartOfAccountTrxParams) error
	CreateNewChartOfAccountTrx(ctx context.Context, arg CreateNewChartOfAccountTrxParams) (CreateNewChartOfAccountTrxResult, error)
	UpdateChartOfAccountTrx(ctx context.Context, arg UpdateChartOfAccountTrxParams) (CreateNewChartOfAccountTrxResult, error)
	CreateNewJournalBookTrx(ctx context.Context, arg CreateNewJournalBookTrxParams) (CreateNewJournalBookTrxResult, error)
	UpdateJournalBookTrx(ctx context.Context, arg UpdateJournalBookTrxParams) (UpdateJournalBookTrxResult, error)
	CloseJournalBookTrx(ctx context.Context, arg string) (CloseJournalBookTrxResult, error)
	CreateNewMemorialJournalTrx(ctx context.Context, arg CreateNewMemorialJournalTrxParams) (CreateNewMemorialJournalTrxResult, error)
	UpdateMemorialJournalTrx(ctx context.Context, arg UpdateMemorialJournalTrxParams) (UpdateMemorialJournalTrxResult, error)
}

// Trx provides all functions to execute SQL queries and transactions
type Trx struct {
	db *sql.DB
	*Queries
}

func NewAccountingTrx(db *sql.DB) AccountingTrx {
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
