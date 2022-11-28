package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/business/db/sqlc"
	"github.com/expert-pancake/service/business/model"
)

func (a businessService) UpdateCompanyBranch(w http.ResponseWriter, r *http.Request) error {

	var req model.UpdateCompanyBranchRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.UpdateCompanyBranch(context.Background(), db.UpdateCompanyBranchParams{
		ID:          req.BranchId,
		Name:        req.Name,
		Address:     req.Address,
		PhoneNumber: req.PhoneNumber,
		IsDeleted:   req.IsDeleted,
	})
	if err != nil {
		return errors.NewServerError(model.UpdateCompanyBranchError, err.Error())
	}

	res := model.RegisterCompanyBranchResponse{
		CompanyBranch: model.CompanyBranch{
			AccountId:   result.UserID,
			CompanyId:   result.CompanyID,
			BranchId:    result.ID,
			Name:        result.Name,
			Address:     result.Address,
			PhoneNumber: result.PhoneNumber,
			IsCentral:   result.IsCentral,
		},
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
