package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	db "github.com/expert-pancake/service/purchasing/db/transaction"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/purchasing/model"
)

func (a purchasingService) UpdateReceiptOrderItems(w http.ResponseWriter, r *http.Request) error {

	var req model.UpdateReceiptOrderItemsRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.UpdateReceiptOrderItemsTrxParams{
		ReceiptOrderItems: req.ReceiptOrderItems,
		ReceiptOrderId:    req.ReceiptOrderId,
	}
	
	err := a.dbTrx.UpdateReceiptOrderItemsTrx(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpdateReceiptOrderItemsError, err.Error())
	}

	res := model.UpdateReceiptOrderItemsResponse{
		Message: `OK`,
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
