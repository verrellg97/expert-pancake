package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/accounting/db/sqlc"
	"github.com/expert-pancake/service/accounting/model"
)

func (a accountingService) UpdateChartOfAccountGroup(w http.ResponseWriter, r *http.Request) error {

	var req model.UpdateChartOfAccounGroupRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.UpdateChartOfAccountGroupParams{
		ID:               req.ChartOfAccountGroupId,
		ReportType:       req.ReportType,
		AccountType:      req.AccountType,
		AccountGroupName: req.AccountGroupName,
	}

	result, err := a.dbTrx.UpdateChartOfAccountGroup(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpdateChartOfAccountGroupError, err.Error())
	}

	res := model.UpsertChartOfAccountGroupResponse{
		ChartOfAccountGroup: model.ChartOfAccountGroup{
			ChartOfAccountGroupId: result.ID,
			CompanyId:             result.CompanyID,
			ReportType:            result.ReportType,
			AccountType:           result.AccountType,
			AccountGroupName:      result.AccountGroupName,
		},
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
