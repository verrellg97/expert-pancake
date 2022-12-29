package util

import (
	"strconv"

	db "github.com/expert-pancake/service/accounting/db/sqlc"
	"github.com/expert-pancake/service/accounting/model"
)

func ChartOfAccountBranchDbToApi(data []db.GetChartOfAccountBranchesRow) []string {
	var chartOfAccountBranches = make([]string, 0)

	for _, d := range data {
		chartOfAccountBranches = append(chartOfAccountBranches, d.BranchID)
	}

	return chartOfAccountBranches
}

func JournalGroupAccountDbToApi(data []db.GetJournalBookAccountsRow) []model.JournalBookAccount {
	var journalBookAccounts = make([]model.JournalBookAccount, 0)

	for _, d := range data {
		journalBookAccounts = append(journalBookAccounts, model.JournalBookAccount{
			ChartOfAccountId: d.ChartOfAccountID,
			AccountType:      d.AccountType,
			AccountGroup:     d.AccountGroupName,
			AccountName:      d.AccountName,
			Amount:           strconv.FormatInt(d.Amount, 10),
		})
	}

	return journalBookAccounts
}
