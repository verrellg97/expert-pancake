package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/sales/db/transaction"
	"github.com/expert-pancake/service/sales/model"
)

func (a salesService) UpdatePOSCustomerSetting(w http.ResponseWriter, r *http.Request) error {

	var req model.UpdatePOSCustomerSettingRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.UpdatePOSCustomerSettingTrxParams{
		BranchId:  req.BranchId,
		Customers: req.Customers,
	}

	result, err := a.dbTrx.UpdatePOSCustomerSettingTrx(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpdatePOSCustomerSettingError, err.Error())
	}

	res := model.UpdatePOSCustomerSettingResponse{
		Message: result.Message,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
