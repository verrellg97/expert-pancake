package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/inventory/db/sqlc"
	"github.com/expert-pancake/service/inventory/model"
	uuid "github.com/satori/go.uuid"
)

func (a inventoryService) UpsertItemReorder(w http.ResponseWriter, r *http.Request) error {

	var req model.UpsertItemReorderRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	var id = ""
	if req.Id == "" {
		id = uuid.NewV4().String()
	} else {
		id = req.Id
	}

	arg := db.UpsertItemReorderParams{
		ID:           id,
		VariantID:    req.VariantId,
		WarehouseID:  req.WarehouseId,
		MinimumStock: req.MinimumStock,
	}

	err := a.dbTrx.UpsertItemReorder(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpsertItemReorderError, err.Error())
	}

	res := model.UpsertItemReorderResponse{
		Message: "OK",
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
