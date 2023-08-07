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

func (a businessService) GetPublicCompanies(w http.ResponseWriter, r *http.Request) error {

	var req model.GetPublicCompaniesRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetPublicCompaniesFilteredByName(context.Background(), db.GetPublicCompaniesFilteredByNameParams{
		UserID: req.AccountId,
		Name:   util.WildCardString(req.Keyword),
	})
	if err != nil {
		return errors.NewServerError(model.GetPublicCompaniesError, err.Error())
	}

	var companies = make([]model.Company, 0)

	for _, d := range result {
		resultBranches, err := a.dbTrx.GetUserCompanyBranches(context.Background(), db.GetUserCompanyBranchesParams{
			UserID:    d.UserID,
			CompanyID: d.ID,
		})
		if err != nil {
			return errors.NewServerError(model.GetPublicCompaniesError, err.Error())
		}

		var company = model.Company{
			AccountId:         d.UserID,
			CompanyId:         d.ID,
			Name:              d.Name,
			InitialName:       d.InitialName,
			Type:              d.Type,
			ResponsiblePerson: d.ResponsiblePerson,
			ImageUrl:          d.ImageUrl,
			Branches:          util.CompanyBranchDbToApi(resultBranches),
		}
		companies = append(companies, company)
	}

	res := model.GetPublicCompaniesResponse{
		Companies: companies,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
