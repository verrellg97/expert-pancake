package impl

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/sales/db/sqlc"
	"github.com/expert-pancake/service/sales/impl/client"
	"github.com/expert-pancake/service/sales/model"
)

func (a salesService) GetPOSUserSetting(w http.ResponseWriter, r *http.Request) error {

	var req model.GetPOSUserSettingRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetPOSUserSetting(context.Background(), db.GetPOSUserSettingParams{
		UserID:   req.UserId,
		BranchID: req.BranchId,
	})
	if err == sql.ErrNoRows {
		arg := db.UpsertPOSUserSettingParams{
			UserID:          req.UserId,
			BranchID:        req.BranchId,
			WarehouseID:     "",
			WarehouseRackID: "",
		}

		result, err = a.dbTrx.UpsertPOSUserSetting(context.Background(), arg)
		if err != nil {
			return errors.NewServerError(model.GetPOSUserSettingError, err.Error())
		}
	} else if err != nil {
		return errors.NewServerError(model.GetPOSUserSettingError, err.Error())
	}

	argWarehouse := client.GetWarehousesRequest{
		Id:       result.WarehouseID,
		BranchId: "1",
	}
	warehouse, err := client.GetWarehouses(argWarehouse)
	if err != nil {
		return errors.NewServerError(model.GetPOSUserSettingError, err.Error())
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
		return errors.NewServerError(model.GetPOSUserSettingError, err.Error())
	}
	warehouseRackName := ""
	if len(warehouseRack.Result.WarehouseRacks) > 0 {
		warehouseRackName = warehouseRack.Result.WarehouseRacks[0].Name
	}

	res := model.GetPOSUserSettingResponse{
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
