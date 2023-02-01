package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/warehouse/db/sqlc"
	"github.com/expert-pancake/service/warehouse/model"
	"github.com/expert-pancake/service/warehouse/util"
)

func (a warehouseService) GetWarehouseRacks(w http.ResponseWriter, r *http.Request) error {

	var req model.GetWarehouseRacksRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetWarehouseRacks(context.Background(), db.GetWarehouseRacksParams{
		WarehouseID: req.WarehouseId,
		Name:        util.WildCardString(req.Keyword),
	})
	if err != nil {
		return errors.NewServerError(model.GetWarehouseRacksError, err.Error())
	}

	var warehouseRacks = make([]model.WarehouseRack, 0)

	for _, d := range result {
		var warehouseRack = model.WarehouseRack{
			RackId:      d.ID,
			WarehouseId: d.WarehouseID,
			Name:        d.Name,
		}
		warehouseRacks = append(warehouseRacks, warehouseRack)
	}

	res := model.GetWarehouseRacksResponse{
		WarehouseRacks: warehouseRacks,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
