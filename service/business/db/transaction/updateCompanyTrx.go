package db

import (
	"context"

	db "github.com/expert-pancake/service/business/db/sqlc"
)

type UpdateCompanyTrxParams struct {
	UserID            string
	CompanyID         string
	Name              string
	InitialName       string
	Type              string
	ResponsiblePerson string
	IsDeleted         int
}

type UpdateCompanyTrxResult struct {
	UserID            string
	CompanyID         string
	Name              string
	InitialName       string
	Type              string
	ResponsiblePerson string
}

func (trx *Trx) UpdateCompanyTrx(ctx context.Context, arg UpdateCompanyTrxParams) (UpdateCompanyTrxResult, error) {
	var result UpdateCompanyTrxResult

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		companyRes, err := q.UpsertCompany(ctx, db.UpsertCompanyParams{
			ID:                arg.CompanyID,
			UserID:            arg.UserID,
			Name:              arg.Name,
			InitialName:       arg.InitialName,
			Type:              arg.Type,
			ResponsiblePerson: arg.ResponsiblePerson,
			IsDeleted:         int32(arg.IsDeleted),
		})
		if err != nil {
			return err
		}

		result.UserID = companyRes.UserID
		result.CompanyID = companyRes.ID
		result.Name = companyRes.Name
		result.InitialName = companyRes.InitialName
		result.Type = companyRes.Type
		result.ResponsiblePerson = companyRes.ResponsiblePerson

		return err
	})

	return result, err
}
