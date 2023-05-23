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

func (a inventoryService) GetVariantWarehouseRacksByBranch(w http.ResponseWriter, r *http.Request) error {

	var req model.GetVariantWarehouseRacksByBranchRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	argWarehouseIds := client.GetWarehousesRequest{
		BranchId: req.BranchId,
	}
	warehouseIds, err := client.GetWarehouses(argWarehouseIds)
	if err != nil {
		return errors.NewServerError(model.GetVariantWarehouseRacksByBranchError, err.Error())
	}

	var warehouseIdsParams = make([]string, 0)
	var warehouseIdNames = make(map[string]string, 0)
	for _, d := range warehouseIds.Result.Warehouses {
		warehouseIdsParams = append(warehouseIdsParams, d.WarehouseId)
		warehouseIdNames[d.WarehouseId] = d.Name
	}

	result, err := a.dbTrx.GetVariantWarehouseRacksByBranch(context.Background(), db.GetVariantWarehouseRacksByBranchParams{
		VariantID:    req.VariantId,
		WarehouseIds: warehouseIdsParams,
	})
	if err != nil {
		return errors.NewServerError(model.GetVariantWarehouseRacksByBranchError, err.Error())
	}

	var datas = make([]model.WarehouseAndRack, 0)

	for _, d := range result {
		argWarehouseRack := client.GetWarehouseRacksRequest{
			Id:          d.WarehouseRackID,
			WarehouseId: "1",
		}
		warehouseRack, err := client.GetWarehouseRacks(argWarehouseRack)
		if err != nil {
			return errors.NewServerError(model.GetVariantWarehouseRacksByBranchError, err.Error())
		}

		var data = model.WarehouseAndRack{
			WarehouseId:   d.WarehouseID,
			WarehouseName: warehouseIdNames[d.WarehouseID],
			RackId:        d.WarehouseRackID,
			RackName:      warehouseRack.Result.WarehouseRacks[0].Name,
		}
		datas = append(datas, data)
	}

	res := model.GetVariantWarehouseRacksByBranchResponse{
		WarehouseAndRacks: datas,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
