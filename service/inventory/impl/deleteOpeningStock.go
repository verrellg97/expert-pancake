package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
		db "github.com/expert-pancake/service/inventory/db/transaction"
	"github.com/expert-pancake/service/inventory/model"
)

func (a inventoryService) DeleteOpeningStock(w http.ResponseWriter, r *http.Request) error {

	var req model.DeleteOpeningStockRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.DeleteOpeningStockTrxParams{
		Id:              req.Id,
	}

	err := a.dbTrx.DeleteOpeningStockTrx(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.DeleteOpeningStockError, err.Error())
	}

	

	res := model.DeleteOpeningStockResponse{
		Message: `OK`,
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
