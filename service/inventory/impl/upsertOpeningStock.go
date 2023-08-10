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

func (a inventoryService) UpsertOpeningStock(w http.ResponseWriter, r *http.Request) error {

	var req model.UpsertOpeningStockRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.UpsertOpeningStockTrxParams{
		Id:              req.Id,
		TransactionDate: util.StringToDate(req.TransactionDate),
		CompanyId:       req.CompanyId,
		BranchId:        req.BranchId,
		WarehouseId:     req.WarehouseId,
		WarehouseRackId: req.WarehouseRackId,
		VariantId:       req.VariantId,
		ItemUnitId:      req.ItemUnitId,
		ItemUnitValue:   req.ItemUnitValue,
		Amount:          req.Amount,
		Price:           req.Price,
		Batch:           req.Batch,
		ExpiredDate:     req.ExpiredDate,
	}

	result, err := a.dbTrx.UpsertOpeningStockTrx(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpsertNewOpeningStockError, err.Error())
	}

	argWarehouse := client.GetWarehousesRequest{
		Id:       result.WarehouseId,
		BranchId: "1",
	}
	warehouse, err := client.GetWarehouses(argWarehouse)
	if err != nil {
		return errors.NewServerError(model.UpsertNewOpeningStockError, err.Error())
	}

	argWarehouseRack := client.GetWarehouseRacksRequest{
		Id:          result.WarehouseRackId,
		WarehouseId: "1",
	}
	warehouseRack, err := client.GetWarehouseRacks(argWarehouseRack)
	if err != nil {
		return errors.NewServerError(model.UpsertNewOpeningStockError, err.Error())
	}

	res := model.UpsertOpeningStockResponse{
		OpeningStock: model.OpeningStock{
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
			Amount:            result.Amount,
			Price:             result.Price,
			Batch:             result.Batch,
			ExpiredDate:       result.ExpiredDate,
		},
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
