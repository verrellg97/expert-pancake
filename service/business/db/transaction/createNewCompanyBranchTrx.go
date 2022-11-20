package db

import (
	"context"

	db "github.com/expert-pancake/service/business/db/sqlc"
	uuid "github.com/satori/go.uuid"
)

type CreateNewCompanyBranchTrxParams struct {
	UserID      string
	CompanyID   string
	Name        string
	Address     string
	PhoneNumber string
}

type CreateNewCompanyBranchTrxResult struct {
	UserID      string
	CompanyID   string
	BranchID    string
	Name        string
	Address     string
	PhoneNumber string
}

func (trx *Trx) CreateNewCompanyBranchTrx(ctx context.Context, arg CreateNewCompanyBranchTrxParams) (CreateNewCompanyBranchTrxResult, error) {
	var result CreateNewCompanyBranchTrxResult

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		id := uuid.NewV4().String()

		companyBranchRes, err := q.UpsertCompanyBranch(ctx, db.UpsertCompanyBranchParams{
			ID:          id,
			UserID:      arg.UserID,
			CompanyID:   arg.CompanyID,
			Name:        arg.Name,
			Address:     arg.Address,
			PhoneNumber: arg.PhoneNumber,
			IsDeleted:   0,
		})
		if err != nil {
			return err
		}

		result.UserID = companyBranchRes.UserID
		result.CompanyID = companyBranchRes.CompanyID
		result.BranchID = id
		result.Name = companyBranchRes.Name
		result.Address = companyBranchRes.Address
		result.PhoneNumber = companyBranchRes.PhoneNumber

		return err
	})

	return result, err
}
