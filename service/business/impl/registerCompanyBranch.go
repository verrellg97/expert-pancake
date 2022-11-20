package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/business/db/transaction"
	"github.com/expert-pancake/service/business/model"
)

func (a businessService) RegisterCompanyBranch(w http.ResponseWriter, r *http.Request) error {

	var req model.RegisterCompanyBranchRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.CreateNewCompanyBranchTrxParams{
		UserID:      req.AccountId,
		CompanyID:   req.CompanyId,
		Name:        req.Name,
		Address:     req.Address,
		PhoneNumber: req.PhoneNumber,
	}

	result, err := a.dbTrx.CreateNewCompanyBranchTrx(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.CreateNewCompanyBranchTransactionError, err.Error())
	}

	res := model.RegisterCompanyBranchResponse{
		CompanyBranch: model.CompanyBranch{
			AccountId:   result.UserID,
			CompanyId:   result.CompanyID,
			BranchId:    result.BranchID,
			Name:        result.Name,
			Address:     result.Address,
			PhoneNumber: result.PhoneNumber,
		},
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
