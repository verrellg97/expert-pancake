package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/inventory/db/sqlc"
	"github.com/expert-pancake/service/inventory/impl/client"
	"github.com/expert-pancake/service/inventory/model"
)

func (a inventoryService) GetVariantWarehouseRacks(w http.ResponseWriter, r *http.Request) error {

	var req model.GetVariantWarehouseRacksRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetVariantWarehouseRacks(context.Background(), db.GetVariantWarehouseRacksParams{
		WarehouseID: req.WarehouseId,
		VariantID:   req.VariantId,
	})
	if err != nil {
		return errors.NewServerError(model.GetVariantWarehouseRacksError, err.Error())
	}

	var datas = make([]model.WarehouseRack, 0)

	for _, d := range result {
		argWarehouseRack := client.GetWarehouseRacksRequest{
			Id:          d,
			WarehouseId: "1",
		}
		warehouseRack, err := client.GetWarehouseRacks(argWarehouseRack)
		if err != nil {
			return errors.NewServerError(model.GetVariantWarehouseRacksError, err.Error())
		}

		var data = model.WarehouseRack{
			RackId:      warehouseRack.Result.WarehouseRacks[0].RackId,
			WarehouseId: warehouseRack.Result.WarehouseRacks[0].WarehouseId,
			Name:        warehouseRack.Result.WarehouseRacks[0].Name,
		}
		datas = append(datas, data)
	}

	res := model.GetVariantWarehouseRacksResponse{
		WarehouseRacks: datas,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
