package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/business/db/transaction"
	"github.com/expert-pancake/service/business/model"
)

func (a businessService) RegisterCompany(w http.ResponseWriter, r *http.Request) error {

	var req model.RegisterCompanyRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.CreateNewCompanyTrxParams{
		UserID:            req.AccountId,
		Name:              req.Name,
		InitialName:       req.InitialName,
		Type:              req.Type,
		ResponsiblePerson: req.ResponsiblePerson,
	}

	result, err := a.dbTrx.CreateNewCompanyTrx(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.CreateNewCompanyTransactionError, err.Error())
	}

	res := model.RegisterCompanyResponse{
		Company: model.Company{
			AccountId:         result.UserID,
			CompanyId:         result.CompanyID,
			Name:              result.Name,
			InitialName:       result.InitialName,
			Type:              result.Type,
			ResponsiblePerson: result.ResponsiblePerson,
		},
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
