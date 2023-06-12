package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/business/model"
)

func (a businessService) GetCompanyBranches(w http.ResponseWriter, r *http.Request) error {

	var req model.CompanyBranchesRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetCompanyBranches(context.Background(), req.CompanyId)
	if err != nil {
		return errors.NewServerError(model.GetCompanyBranchesError, err.Error())
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
