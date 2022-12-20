package util

import (
	db "github.com/expert-pancake/service/accounting/db/sqlc"
)

func ChartOfAccountBranchDbToApi(data []db.GetChartOfAccountBranchesRow) []string {
	var chartOfAccountBranches = make([]string, 0)

	for _, d := range data {
		chartOfAccountBranches = append(chartOfAccountBranches, d.BranchID)
	}

	return chartOfAccountBranches
}
