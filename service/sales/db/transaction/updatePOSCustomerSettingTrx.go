package db

import (
	"context"

	db "github.com/expert-pancake/service/sales/db/sqlc"
)

type UpdatePOSCustomerSettingTrxParams struct {
	BranchId  string
	Customers []string
}

type UpdatePOSCustomerSettingTrxResult struct {
	Message string
}

func (trx *Trx) UpdatePOSCustomerSettingTrx(ctx context.Context, arg UpdatePOSCustomerSettingTrxParams) (UpdatePOSCustomerSettingTrxResult, error) {
	var result UpdatePOSCustomerSettingTrxResult

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		err = q.DeletePOSCustomerSetting(ctx, arg.BranchId)
		if err != nil {
			return err
		}

		for _, d := range arg.Customers {
			_, err := q.InsertPOSCustomerSetting(ctx, db.InsertPOSCustomerSettingParams{
				BranchID:      arg.BranchId,
				ContactBookID: d,
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
