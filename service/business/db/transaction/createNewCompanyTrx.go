package db

import (
	"context"

	db "github.com/expert-pancake/service/business/db/sqlc"
	"github.com/expert-pancake/service/business/impl/client"
	"github.com/expert-pancake/service/business/model"
	uuid "github.com/satori/go.uuid"
)

type CreateNewCompanyTrxResult struct {
	AccountId         string
	CompanyId         string
	Name              string
	InitialName       string
	Type              string
	ResponsiblePerson string
	Branches          []model.CompanyBranch
}

func (trx *Trx) CreateNewCompanyTrx(ctx context.Context, arg db.InsertCompanyParams) (CreateNewCompanyTrxResult, error) {
	var result CreateNewCompanyTrxResult

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		id := uuid.NewV4().String()

		companyRes, err := q.InsertCompany(ctx, db.InsertCompanyParams{
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

		_, err = q.InsertCompanyBranch(ctx, db.InsertCompanyBranchParams{
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

		err = client.AddDefaultCompanyChartOfAccount(
			client.AddDefaultCompanyChartOfAccountRequest{
				CompanyId: id,
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
		result.Branches = []model.CompanyBranch{
			{
				AccountId:   arg.ID,
				CompanyId:   id,
				BranchId:    branchId,
				Name:        "Pusat",
				Address:     "",
				PhoneNumber: "",
				IsCentral:   true,
			},
		}

		return err
	})

	return result, err
}
