package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/accounting/db/sqlc"
	"github.com/expert-pancake/service/accounting/model"
)

func (a accountingService) GetCompanyChartOfAccounts(w http.ResponseWriter, r *http.Request) error {

	var req model.CompanyChartOfAccountsRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetCompanyChartOfAccounts(context.Background(), db.GetCompanyChartOfAccountsParams{
		CompanyID:   req.CompanyId,
		AccountName: wildCardString(req.Keyword),
	})
	if err != nil {
		return errors.NewServerError(model.GetCompanyChartOfAccountsError, err.Error())
	}

	var chart_of_accounts = make([]model.ChartOfAccount, 0)

	for _, d := range result {
		var chart_of_account = model.ChartOfAccount{
			ChartOfAccountId:  d.ID,
			CompanyId:         d.CompanyID,
			BranchId:          d.BranchID,
			AccountCode:       d.AccountCode,
			AccountName:       d.AccountName,
			AccountGroup:      d.AccountGroup,
			BankName:          d.BankName,
			BankAccountNumber: d.BankAccountNumber,
			BankCode:          d.BankCode,
			OpeningBalance:    d.OpeningBalance,
			IsDeleted:         d.IsDeleted,
		}
		chart_of_accounts = append(chart_of_accounts, chart_of_account)
	}

	res := chart_of_accounts
	httpHandler.WriteResponse(w, res)

	return nil
}
