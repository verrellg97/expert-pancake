package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/inventory/db/transaction"
	"github.com/expert-pancake/service/inventory/impl/client"
	"github.com/expert-pancake/service/inventory/model"
	"github.com/expert-pancake/service/inventory/util"
)

func (a inventoryService) AddUpdateStock(w http.ResponseWriter, r *http.Request) error {

	var req model.AddUpdateStockRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.AddUpdateStockTrxParams{
		TransactionDate: util.StringToDate(req.TransactionDate),
		WarehouseId:     req.WarehouseId,
		WarehouseRackId: req.WarehouseRackId,
		VariantId:       req.VariantId,
		ItemUnitId:      req.ItemUnitId,
		ItemUnitValue:   req.ItemUnitValue,
		BeginningStock:  req.BeginningStock,
		EndingStock:     req.EndingStock,
		Batch:           req.Batch,
		ExpiredDate:     req.ExpiredDate,
	}

	result, err := a.dbTrx.AddUpdateStockTrx(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.AddNewUpdateStockError, err.Error())
	}

	argWarehouse := client.GetWarehousesRequest{
		Id:       result.WarehouseId,
		BranchId: "1",
	}
	warehouse, err := client.GetWarehouses(argWarehouse)
	if err != nil {
		return errors.NewServerError(model.AddNewUpdateStockError, err.Error())
	}

	argWarehouseRack := client.GetWarehouseRacksRequest{
		Id:          result.WarehouseRackId,
		WarehouseId: "1",
	}
	warehouseRack, err := client.GetWarehouseRacks(argWarehouseRack)
	if err != nil {
		return errors.NewServerError(model.AddNewUpdateStockError, err.Error())
	}

	res := model.AddUpdateStockResponse{
		UpdateStock: model.UpdateStock{
			TransactionId:     result.TransactionId,
			FormNumber:        result.FormNumber,
			TransactionDate:   result.TransactionDate,
			WarehouseId:       result.WarehouseId,
			WarehouseName:     warehouse.Result.Warehouses[0].Name,
			WarehouseRackId:   result.WarehouseRackId,
			WarehouseRackName: warehouseRack.Result.WarehouseRacks[0].Name,
			ItemId:            result.ItemId,
			ItemName:          result.ItemName,
			VariantId:         result.VariantId,
			VariantName:       result.VariantName,
			ItemUnitId:        result.ItemUnitId,
			ItemUnitName:      result.ItemUnitName,
			ItemUnitValue:     result.ItemUnitValue,
			BeginningStock:    result.BeginningStock,
			EndingStock:       result.EndingStock,
			Batch:             result.Batch,
			ExpiredDate:       result.ExpiredDate,
		},
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
