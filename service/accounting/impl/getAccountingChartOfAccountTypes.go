package impl

import (
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/accounting/model"
)

var defaultChartOfAccounts = []model.ChartOfAccountType{
	{ReportType: "NERACA", AccountType: "CASH & EQUALS"},
	{ReportType: "NERACA", AccountType: "RECEIVABLE"},
	{ReportType: "NERACA", AccountType: "INVENTORY"},
	{ReportType: "NERACA", AccountType: "CURRENT ASSET"},
	{ReportType: "NERACA", AccountType: "NON CURRENT ASSET"},
	{ReportType: "NERACA", AccountType: "PAYABLE"},
	{ReportType: "NERACA", AccountType: "CURRENT LIABILITY"},
	{ReportType: "NERACA", AccountType: "NON CURRENT LIABILITY"},
	{ReportType: "NERACA", AccountType: "EQUITY"},
	{ReportType: "NERACA", AccountType: "CURRENT YEAR EARNING"},
	{ReportType: "LABA RUGI", AccountType: "MAIN REVENUE"},
	{ReportType: "LABA RUGI", AccountType: "OTHER REVENUE"},
	{ReportType: "LABA RUGI", AccountType: "MAIN COST OF REVENUE"},
	{ReportType: "LABA RUGI", AccountType: "OTHER COST OF REVENUE"},
	{ReportType: "LABA RUGI", AccountType: "EXPENSE"},
	{ReportType: "LABA RUGI", AccountType: "INCOME"},
}

func (a accountingService) GetAccountingChartOfAccountTypes(w http.ResponseWriter, r *http.Request) error {

	res := model.GetAccountingChartOfAccountTypesResponse{ChartOfAccountTypes: defaultChartOfAccounts}
	httpHandler.WriteResponse(w, res)

	return nil
}
