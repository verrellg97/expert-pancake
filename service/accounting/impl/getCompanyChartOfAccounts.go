package impl

import (
	"context"
	"net/http"
	"strconv"

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

	var isFilterJournalGroup = false
	var accountGroups []string
	if req.GroupFilter != nil {
		isFilterJournalGroup = true
		switch *req.GroupFilter {
		case "DEBET":
			accountGroups = append(accountGroups, "BANK", "KAS")
		case "KREDIT":
			accountGroups = append(accountGroups, "BEBAN USAHA", "BEBAN LAIN-LAIN")
		default:
			accountGroups = append(accountGroups, *req.GroupFilter)
		}
	}

	result, err := a.dbTrx.GetCompanyChartOfAccounts(context.Background(), db.GetCompanyChartOfAccountsParams{
		CompanyID:            req.CompanyId,
		AccountName:          util.WildCardString(req.Keyword),
		IsDeletedFilter:      isDeletedFilter,
		IsDeleted:            *req.IsDeletedFilter,
		IsFilterJournalGroup: isFilterJournalGroup,
		AccountGroups:        accountGroups,
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
			OpeningBalance:    strconv.FormatInt(d.OpeningBalance, 10),
			IsDeleted:         d.IsDeleted,
		}
		chart_of_accounts = append(chart_of_accounts, chart_of_account)
	}

	res := chart_of_accounts
	httpHandler.WriteResponse(w, res)

	return nil
}
