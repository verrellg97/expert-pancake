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

func (a businessService) GetUserCompanyBranches(w http.ResponseWriter, r *http.Request) error {

	var req model.UserCompanyBranchesRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetUserCompanyBranchesFilteredByName(context.Background(), db.GetUserCompanyBranchesFilteredByNameParams{
		UserID:    req.AccountId,
		CompanyID: req.CompanyId,
		Name:      util.WildCardString(req.Keyword),
	})
	if err != nil {
		return errors.NewServerError(model.GetUserCompanyBranchesError, err.Error())
	}

	var companyBranches = make([]model.CompanyBranch, 0)

	for _, d := range result {
		var companyBranch = model.CompanyBranch{
			AccountId:   d.UserID,
			CompanyId:   d.CompanyID,
			BranchId:    d.ID,
			Name:        d.Name,
			Address:     d.Address,
			PhoneNumber: d.PhoneNumber,
			IsCentral:   d.IsCentral,
		}
		companyBranches = append(companyBranches, companyBranch)
	}

	res := companyBranches
	httpHandler.WriteResponse(w, res)

	return nil
}
