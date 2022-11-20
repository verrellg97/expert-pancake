package db

import (
	"context"

	db "github.com/expert-pancake/service/business/db/sqlc"
	uuid "github.com/satori/go.uuid"
)

type CreateNewCompanyTrxParams struct {
	UserID            string
	Name              string
	InitialName       string
	Type              string
	ResponsiblePerson string
}

type CreateNewCompanyTrxResult struct {
	UserID            string
	CompanyID         string
	Name              string
	InitialName       string
	Type              string
	ResponsiblePerson string
}

func (trx *Trx) CreateNewCompanyTrx(ctx context.Context, arg CreateNewCompanyTrxParams) (CreateNewCompanyTrxResult, error) {
	var result CreateNewCompanyTrxResult

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		id := uuid.NewV4().String()

		companyRes, err := q.UpsertCompany(ctx, db.UpsertCompanyParams{
			ID:                id,
			UserID:            arg.UserID,
			Name:              arg.Name,
			InitialName:       arg.InitialName,
			Type:              arg.Type,
			ResponsiblePerson: arg.ResponsiblePerson,
			IsDeleted:         0,
		})
		if err != nil {
			return err
		}

		result.UserID = companyRes.UserID
		result.CompanyID = id
		result.Name = companyRes.Name
		result.InitialName = companyRes.InitialName
		result.Type = companyRes.Type
		result.ResponsiblePerson = companyRes.ResponsiblePerson

		return err
	})

	return result, err
}
