package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/accounting/model"
)

func (a accountingService) GetChartOfAccountGroups(w http.ResponseWriter, r *http.Request) error {

	var req model.GetChartOfAccountGroupsRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetChartOfAccountGroups(context.Background(), req.CompanyId)
	if err != nil {
		return errors.NewServerError(model.GetChartOfAccountGroupsError, err.Error())
	}

	var chart_of_account_groups = make([]model.ChartOfAccountGroup, 0)

	for _, d := range result {
		var chart_of_account_group = model.ChartOfAccountGroup{
			ChartOfAccountGroupId: d.ID,
			CompanyId:             d.CompanyID,
			ReportType:            d.ReportType,
			AccountType:           d.AccountType,
			AccountGroupName:      d.AccountGroupName,
		}
		chart_of_account_groups = append(chart_of_account_groups, chart_of_account_group)
	}

	res := chart_of_account_groups
	httpHandler.WriteResponse(w, res)

	return nil
}
