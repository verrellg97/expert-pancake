package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/warehouse/db/sqlc"
	"github.com/expert-pancake/service/warehouse/model"
	uuid "github.com/satori/go.uuid"
)

func (a warehouseService) UpsertWarehouseRack(w http.ResponseWriter, r *http.Request) error {

	var req model.UpsertWarehouseRackRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	var id = ""
	if req.RackId == "" {
		id = uuid.NewV4().String()
	} else {
		id = req.RackId
	}

	arg := db.UpsertWarehouseRackParams{
		ID:          id,
		WarehouseID: req.WarehouseId,
		Name:        req.Name,
	}

	result, err := a.dbTrx.UpsertWarehouseRack(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpsertWarehouseRackError, err.Error())
	}

	res := model.UpsertWarehouseRackResponse{
		WarehouseRack: model.WarehouseRack{
			RackId:      result.ID,
			WarehouseId: result.WarehouseID,
			Name:        result.Name,
		},
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
