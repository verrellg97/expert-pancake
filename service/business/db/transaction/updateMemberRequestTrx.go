package db

import (
	"context"

	db "github.com/expert-pancake/service/business/db/sqlc"
	uuid "github.com/satori/go.uuid"
)

type UpdateMemberRequestTrxParams struct {
	Id     string
	Status string
}

type UpdateMemberRequestTrxResult struct {
	Message string
}

func (trx *Trx) UpdateMemberRequestTrx(ctx context.Context, arg UpdateMemberRequestTrxParams) (UpdateMemberRequestTrxResult, error) {
	var result UpdateMemberRequestTrxResult

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		memberRequest, err := q.UpdateMemberRequest(ctx, db.UpdateMemberRequestParams{
			ID:     arg.Id,
			Status: arg.Status,
		})
		if err != nil {
			return err
		}

		if arg.Status == "accepted" {
			err = q.InsertCompanyMember(ctx, db.InsertCompanyMemberParams{
				ID:        uuid.NewV4().String(),
				UserID:    memberRequest.UserID,
				CompanyID: memberRequest.CompanyID,
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
