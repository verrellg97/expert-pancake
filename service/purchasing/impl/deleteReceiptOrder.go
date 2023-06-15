package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/purchasing/model"
)

func (a purchasingService) DeleteReceiptOrder(w http.ResponseWriter, r *http.Request) error {

	var req model.DeleteReceiptOrderRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	err := a.dbTrx.DeleteReceiptOrder(context.Background(), req.ReceiptOrderId)
	if err != nil {
		return errors.NewServerError(model.DeleteReceiptOrderError, err.Error())
	}

	res := model.DeleteReceiptOrderResponse{
		Message: `OK`,
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
