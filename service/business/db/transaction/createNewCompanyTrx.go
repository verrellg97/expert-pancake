package db

import (
	"context"

	db "github.com/expert-pancake/service/business/db/sqlc"
	uuid "github.com/satori/go.uuid"
)

type CreateNewCompanyTrxResult struct {
	AccountId         string
	CompanyId         string
	Name              string
	InitialName       string
	Type              string
	ResponsiblePerson string
}

func (trx *Trx) CreateNewCompanyTrx(ctx context.Context, arg db.UpsertCompanyParams) (CreateNewCompanyTrxResult, error) {
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
		})
		if err != nil {
			return err
		}

		branchId := uuid.NewV4().String()

		err = q.InsertCompanyBranch(ctx, db.InsertCompanyBranchParams{
			ID:          branchId,
			UserID:      arg.UserID,
			CompanyID:   id,
			Name:        "Pusat",
			Address:     "",
			PhoneNumber: "",
			IsCentral:   true,
		})
		if err != nil {
			return err
		}

		result.CompanyId = id
		result.AccountId = companyRes.UserID
		result.Name = companyRes.Name
		result.InitialName = companyRes.InitialName
		result.Type = companyRes.Type
		result.ResponsiblePerson = companyRes.ResponsiblePerson

		return err
	})

	return result, err
}
