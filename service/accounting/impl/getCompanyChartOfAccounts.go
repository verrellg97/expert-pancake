package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/accounting/db/sqlc"
	"github.com/expert-pancake/service/accounting/model"
	"github.com/expert-pancake/service/accounting/util"
)

func (a accountingService) GetCompanyChartOfAccounts(w http.ResponseWriter, r *http.Request) error {

	var req model.GetCompanyChartOfAccountsRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	var isDeletedFilter = false
	if req.IsDeletedFilter != nil {
		isDeletedFilter = true
	} else {
		req.IsDeletedFilter = &isDeletedFilter
	}

	var isFilterJournalType = false
	var accountTypes []string
	if req.TypeFilter != nil {
		isFilterJournalType = true
		switch *req.TypeFilter {
		case "DEBET":
			accountTypes = append(accountTypes, "CASH & EQUALS")
		case "KREDIT":
			accountTypes = append(accountTypes, "MAIN COST OF REVENUE", "OTHER COST OF REVENUE", "EXPENSE")
		case "BUKU":
			accountTypes = append(accountTypes, "CASH & EQUALS", "INVENTORY")
		case "":
			isFilterJournalType = false
		default:
			accountTypes = append(accountTypes, *req.TypeFilter)
		}
	}

	var isFilterId = false
	if req.Id != "" {
		isFilterId = true
	}

	result, err := a.dbTrx.GetCompanyChartOfAccounts(context.Background(), db.GetCompanyChartOfAccountsParams{
		CompanyID:           req.CompanyId,
		AccountName:         util.WildCardString(req.Keyword),
		IsDeletedFilter:     isDeletedFilter,
		IsDeleted:           *req.IsDeletedFilter,
		IsFilterJournalType: isFilterJournalType,
		AccountTypes:        accountTypes,
		IsFilterID:          isFilterId,
		ID:                  req.Id,
	})
	if err != nil {
		return errors.NewServerError(model.GetCompanyChartOfAccountsError, err.Error())
	}

	var chart_of_accounts = make([]model.ChartOfAccount, 0)

	for _, d := range result {
		resultBranches, err := a.dbTrx.GetChartOfAccountBranches(context.Background(), d.ID)
		if err != nil {
			return errors.NewServerError(model.GetCompanyChartOfAccountsError, err.Error())
		}
		var chart_of_account = model.ChartOfAccount{
			ChartOfAccountId:      d.ID,
			CompanyId:             d.CompanyID,
			CurrencyCode:          d.CurrencyCode,
			ChartOfAccountGroupId: d.ChartOfAccountGroupID,
			ReportType:            d.ReportType,
			AccountType:           d.AccountType,
			AccountGroup:          d.AccountGroupName,
			AccountCode:           d.AccountCode,
			AccountName:           d.AccountName,
			BankName:              d.BankName,
			BankAccountNumber:     d.BankAccountNumber,
			BankCode:              d.BankCode,
			IsAllBranches:         d.IsAllBranches,
			Branches:              util.ChartOfAccountBranchDbToApi(resultBranches),
			IsDeleted:             d.IsDeleted,
		}
		chart_of_accounts = append(chart_of_accounts, chart_of_account)
	}

	res := chart_of_accounts
	httpHandler.WriteResponse(w, res)

	return nil
}
