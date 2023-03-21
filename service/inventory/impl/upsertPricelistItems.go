package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/inventory/db/transaction"
	"github.com/expert-pancake/service/inventory/model"
)

func (a inventoryService) UpsertPricelistItems(w http.ResponseWriter, r *http.Request) error {

	var req model.UpsertPricelistItemsRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.UpsertPricelistItemsTrxParams{
		PricelistId:    req.PricelistId,
		PricelistItems: req.PricelistItems,
	}

	result, err := a.dbTrx.UpsertPricelistItemsTrx(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpsertPricelistItemsError, err.Error())
	}

	res := model.UpsertPricelistItemsResponse{
		Message: result.Message,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
