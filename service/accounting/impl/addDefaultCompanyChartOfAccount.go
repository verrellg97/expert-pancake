package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/accounting/db/transaction"
	"github.com/expert-pancake/service/accounting/model"
)

func (a accountingService) AddDefaultCompanyChartOfAccount(w http.ResponseWriter, r *http.Request) error {

	var req model.AddDefaultCompanyChartOfAccountRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.AddDefaultCompanyChartOfAccountTrxParams{
		CompanyId: req.CompanyId,
	}

	err := a.dbTrx.AddDefaultCompanyChartOfAccountTransactionTrx(context.Background(), arg)

	if err != nil {
		return errors.NewServerError(model.AddDefaultCompanyChartOfAccountError, err.Error())
	}

	res := model.AddDefaultDataResponse{
		Message: "OK",
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
