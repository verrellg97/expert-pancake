package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/sales/impl/client"
	"github.com/expert-pancake/service/sales/model"
	"github.com/expert-pancake/service/sales/util"
)

func (a salesService) GetPOSPaymentMethod(w http.ResponseWriter, r *http.Request) error {

	var req model.GetPOSPaymentMethodRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetPOSPaymentMethod(context.Background(), util.WildCardString(req.Keyword))
	if err != nil {
		return errors.NewServerError(model.GetPOSPaymentMethodError, err.Error())
	}

	var payment_methods = make([]model.POSPaymentMethod, 0)
	for _, d := range result {
		argCOA := client.GetCompanyChartOfAccountsRequest{
			CompanyId: d.CompanyID,
			Id:        d.ChartOfAccountID,
		}
		coa, err := client.GetCompanyChartOfAccounts(argCOA)
		if err != nil {
			return errors.NewServerError(model.GetPOSPaymentMethodError, err.Error())
		}

		var payment_method = model.POSPaymentMethod{
			Id:                 d.ID,
			ChartOfAccountId:   d.ChartOfAccountID,
			ChartOfAccountName: coa.Result[0].AccountName,
			Name:               d.Name,
		}
		payment_methods = append(payment_methods, payment_method)
	}

	res := model.GetPOSPaymentMethodResponse{
		POSPaymentMethods: payment_methods,
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
