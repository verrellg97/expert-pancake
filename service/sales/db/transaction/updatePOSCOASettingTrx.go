package db

import (
	"context"

	db "github.com/expert-pancake/service/sales/db/sqlc"
)

type UpdatePOSCOASettingTrxParams struct {
	BranchId        string
	ChartOfAccounts []string
}

type UpdatePOSCOASettingTrxResult struct {
	Message string
}

func (trx *Trx) UpdatePOSCOASettingTrx(ctx context.Context, arg UpdatePOSCOASettingTrxParams) (UpdatePOSCOASettingTrxResult, error) {
	var result UpdatePOSCOASettingTrxResult

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		err = q.DeletePOSCOASetting(ctx, arg.BranchId)
		if err != nil {
			return err
		}

		for _, d := range arg.ChartOfAccounts {
			_, err := q.InsertPOSCOASetting(ctx, db.InsertPOSCOASettingParams{
				BranchID:         arg.BranchId,
				ChartOfAccountID: d,
			})
			if err != nil {
				return err
			}
		}

		result.Message = "OK"

		return err
	})

	return result, err
}
