package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/business/db/sqlc"
	"github.com/expert-pancake/service/business/model"
	"github.com/expert-pancake/service/business/util"
)

func (a businessService) GetUserCompanies(w http.ResponseWriter, r *http.Request) error {

	var req model.UserCompaniesRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetUserCompaniesFilteredByName(context.Background(), db.GetUserCompaniesFilteredByNameParams{
		UserID: req.AccountId,
		Name:   util.WildCardString(req.Keyword),
	})
	if err != nil {
		return errors.NewServerError(model.GetUserCompaniesError, err.Error())
	}

	var companies = make([]model.Company, 0)

	for _, d := range result {
		var companyBranches = make([]model.CompanyBranch, 0)
		resultBranches, err := a.dbTrx.GetUserCompanyBranches(context.Background(), db.GetUserCompanyBranchesParams{
			UserID:    req.AccountId,
			CompanyID: d.ID,
		})
		if err != nil {
			return errors.NewServerError(model.GetUserCompanyBranchesError, err.Error())
		}

		for _, e := range resultBranches {
			var companyBranch = model.CompanyBranch{
				AccountId:   e.UserID,
				CompanyId:   e.CompanyID,
				BranchId:    e.ID,
				Name:        e.Name,
				Address:     e.Address,
				PhoneNumber: e.PhoneNumber,
				IsCentral:   e.IsCentral,
			}
			companyBranches = append(companyBranches, companyBranch)
		}
		var company = model.Company{
			AccountId:         d.UserID,
			CompanyId:         d.ID,
			Name:              d.Name,
			InitialName:       d.InitialName,
			Type:              d.Type,
			ResponsiblePerson: d.ResponsiblePerson,
			Branches:          companyBranches,
		}
		companies = append(companies, company)
	}

	res := companies
	httpHandler.WriteResponse(w, res)

	return nil
}
