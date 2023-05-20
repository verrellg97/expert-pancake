package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/sales/db/sqlc"
	"github.com/expert-pancake/service/sales/impl/client"
	"github.com/expert-pancake/service/sales/model"
)

func (a salesService) UpdatePOSUserSetting(w http.ResponseWriter, r *http.Request) error {

	var req model.UpdatePOSUserSettingRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.UpsertPOSUserSettingParams{
		UserID:          req.UserId,
		BranchID:        req.BranchId,
		WarehouseID:     req.WarehouseId,
		WarehouseRackID: req.WarehouseRackId,
	}

	result, err := a.dbTrx.UpsertPOSUserSetting(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpdatePOSUserSettingError, err.Error())
	}

	argWarehouse := client.GetWarehousesRequest{
		Id:       result.WarehouseID,
		BranchId: "1",
	}
	warehouse, err := client.GetWarehouses(argWarehouse)
	if err != nil {
		return errors.NewServerError(model.UpdatePOSUserSettingError, err.Error())
	}
	warehouseName := ""
	if len(warehouse.Result.Warehouses) > 0 {
		warehouseName = warehouse.Result.Warehouses[0].Name
	}

	argWarehouseRack := client.GetWarehouseRacksRequest{
		Id:          result.WarehouseRackID,
		WarehouseId: "1",
	}
	warehouseRack, err := client.GetWarehouseRacks(argWarehouseRack)
	if err != nil {
		return errors.NewServerError(model.UpdatePOSUserSettingError, err.Error())
	}
	warehouseRackName := ""
	if len(warehouseRack.Result.WarehouseRacks) > 0 {
		warehouseRackName = warehouseRack.Result.WarehouseRacks[0].Name
	}

	res := model.UpdatePOSUserSettingResponse{
		POSUserSetting: model.POSUserSetting{
			UserId:            result.UserID,
			BranchId:          result.BranchID,
			WarehouseId:       result.WarehouseID,
			WarehouseName:     warehouseName,
			WarehouseRackId:   result.WarehouseRackID,
			WarehouseRackName: warehouseRackName,
		},
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
