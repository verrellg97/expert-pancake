package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/sales/db/transaction"
	"github.com/expert-pancake/service/sales/model"
)

func (a salesService) UpdateDeliveryOrderStatus(w http.ResponseWriter, r *http.Request) error {

	var req model.UpdateDeliveryOrderStatusRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.UpdateDeliveryOrderStatusTrxParams{
		DeliveryOrderId: req.DeliveryOrderId,
		Status:          req.Status,
	}

	result, err := a.dbTrx.UpdateDeliveryOrderStatusTrx(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpdateDeliveryOrderStatusError, err.Error())
	}

	res := model.UpdateDeliveryOrderStatusResponse{
		Message: result.Message,
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
