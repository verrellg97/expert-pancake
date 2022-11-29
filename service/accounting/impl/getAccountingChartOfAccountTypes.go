package impl

import (
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/accounting/model"
)

var defaultChartOfAccounts = []string{
	"BANK",
	"KAS",
	"DEPOSIT",
}

func (a accountingService) GetAccountingChartOfAccountTypes(w http.ResponseWriter, r *http.Request) error {

	res := model.GetAccountingChartOfAccountTypesResponse{ChartOfAccountTypes: defaultChartOfAccounts}
	httpHandler.WriteResponse(w, res)

	return nil
}
