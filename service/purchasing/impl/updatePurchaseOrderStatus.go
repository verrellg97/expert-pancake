package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/purchasing/db/transaction"
	"github.com/expert-pancake/service/purchasing/model"
)

func (a purchasingService) UpdatePurchaseOrderStatus(w http.ResponseWriter, r *http.Request) error {

	var req model.UpdatePurchaseOrderStatusRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.UpdatePurchaseOrderStatusTrxParams{
		PurchaseOrderId: req.PurchaseOrderId,
		Status:          req.Status,
	}

	result, err := a.dbTrx.UpdatePurchaseOrderStatusTrx(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpdatePurchaseOrderStatusError, err.Error())
	}

	res := model.UpdatePurchaseOrderStatusResponse{
		Message: result.Message,
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
