package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/purchasing/db/transaction"
	"github.com/expert-pancake/service/purchasing/model"
)

func (a purchasingService) UpdateReceiptOrderStatus(w http.ResponseWriter, r *http.Request) error {

	var req model.UpdateReceiptOrderStatusRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.UpdateReceiptOrderStatusTrxParams{
		ReceiptOrderId: req.ReceiptOrderId,
		Status:         req.Status,
	}

	result, err := a.dbTrx.UpdateReceiptOrderStatusTrx(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpdateReceiptOrderStatusError, err.Error())
	}

	res := model.UpdateReceiptOrderStatusResponse{
		Message: result.Message,
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
