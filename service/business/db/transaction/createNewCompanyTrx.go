package db

import (
	"context"
	"time"

	db "github.com/expert-pancake/service/business/db/sqlc"
	"github.com/expert-pancake/service/business/impl/client"
	"github.com/expert-pancake/service/business/model"
	"github.com/expert-pancake/service/business/util"
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

		err = client.AddDefaultContactBook(
			client.AddDefaultContactBookRequest{
				CompanyId:   id,
				CompanyName: arg.Name,
			})
		if err != nil {
			return err
		}

		categoryAmount, err := client.UpsertUnitCategory(
			client.UpsertUnitCategoryRequest{
				CompanyId: id,
				Name:      "Jumlah",
			})
		if err != nil {
			return err
		}

		err = client.AddUnit(
			client.AddUnitRequest{
				CompanyId:      id,
				UnitCategoryId: categoryAmount.Result.UnitCategory.Id,
				Name:           "Pieces",
			})
		if err != nil {
			return err
		}

		err = client.AddUnit(
			client.AddUnitRequest{
				CompanyId:      id,
				UnitCategoryId: categoryAmount.Result.UnitCategory.Id,
				Name:           "Dozen",
			})
		if err != nil {
			return err
		}

		err = client.AddUnit(
			client.AddUnitRequest{
				CompanyId:      id,
				UnitCategoryId: categoryAmount.Result.UnitCategory.Id,
				Name:           "Box",
			})
		if err != nil {
			return err
		}

		categoryWeight, err := client.UpsertUnitCategory(
			client.UpsertUnitCategoryRequest{
				CompanyId: id,
				Name:      "Berat",
			})
		if err != nil {
			return err
		}

		err = client.AddUnit(
			client.AddUnitRequest{
				CompanyId:      id,
				UnitCategoryId: categoryWeight.Result.UnitCategory.Id,
				Name:           "kilogram",
			})
		if err != nil {
			return err
		}

		err = client.AddUnit(
			client.AddUnitRequest{
				CompanyId:      id,
				UnitCategoryId: categoryWeight.Result.UnitCategory.Id,
				Name:           "gram",
			})
		if err != nil {
			return err
		}

		categoryLength, err := client.UpsertUnitCategory(
			client.UpsertUnitCategoryRequest{
				CompanyId: id,
				Name:      "Panjang",
			})
		if err != nil {
			return err
		}

		err = client.AddUnit(
			client.AddUnitRequest{
				CompanyId:      id,
				UnitCategoryId: categoryLength.Result.UnitCategory.Id,
				Name:           "m",
			})
		if err != nil {
			return err
		}

		err = client.AddUnit(
			client.AddUnitRequest{
				CompanyId:      id,
				UnitCategoryId: categoryLength.Result.UnitCategory.Id,
				Name:           "cm",
			})
		if err != nil {
			return err
		}

		err = client.UpsertPricelist(
			client.UpsertPricelistRequest{
				CompanyId: id,
				Name:      "Default",
				StartDate: time.Now().Format(util.DateLayoutYMD),
				IsDefault: true,
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
