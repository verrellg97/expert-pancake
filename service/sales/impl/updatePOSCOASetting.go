package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/sales/db/transaction"
	"github.com/expert-pancake/service/sales/model"
)

func (a salesService) UpdatePOSCOASetting(w http.ResponseWriter, r *http.Request) error {

	var req model.UpdatePOSCOASettingRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.UpdatePOSCOASettingTrxParams{
		BranchId:        req.BranchId,
		ChartOfAccounts: req.ChartOfAccounts,
	}

	result, err := a.dbTrx.UpdatePOSCOASettingTrx(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpdatePOSCOASettingError, err.Error())
	}

	res := model.UpdatePOSCOASettingResponse{
		Message: result.Message,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
