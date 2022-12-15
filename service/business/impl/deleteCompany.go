package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/business/model"
)

func (a businessService) DeleteCompany(w http.ResponseWriter, r *http.Request) error {

	var req model.DeleteCompanyRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	err := a.dbTrx.DeleteCompany(context.Background(), req.CompanyId)
	if err != nil {
		return errors.NewServerError(model.DeleteCompanyError, err.Error())
	}

	errCompanyBranch := a.dbTrx.DeleteCompanyBranchesByCompanyId(context.Background(), req.CompanyId)
	if errCompanyBranch != nil {
		return errors.NewServerError(model.DeleteCompanyError, err.Error())
	}

	res := model.DeleteDataResponse{
		Message: "OK",
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
