package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/inventory/db/sqlc"
	"github.com/expert-pancake/service/inventory/model"
)

func (a inventoryService) DeleteStockMovement(w http.ResponseWriter, r *http.Request) error {

	var req model.DeleteStockMovementRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.DeleteStockMovementParams{
		TransactionID:        req.TransactionId,
		TransactionReference: req.TransactionReference,
	}

	a.dbTrx.DeleteStockMovement(context.Background(), arg)
	res := model.DeleteStockMovementResponse{
		Message: `OK`,
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
