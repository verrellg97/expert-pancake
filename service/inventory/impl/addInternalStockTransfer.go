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

func (a inventoryService) AddInternalStockTransfer(w http.ResponseWriter, r *http.Request) error {

	var req model.AddInternalStockTransferRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.AddInternalStockTransferTrxParams{
		CompanyId:              req.CompanyId,
		BranchId:               req.BranchId,
		SourceWarehouseId:      req.SourceWarehouseId,
		DestinationWarehouseId: req.DestinationWarehouseId,
		TransactionDate:        util.StringToDate(req.TransactionDate),
		Items:                  req.Items,
	}

	result, err := a.dbTrx.AddInternalStockTransferTrx(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.AddNewInternalStockTransferError, err.Error())
	}

	argSourceWarehouse := client.GetWarehousesRequest{
		Id:       result.SourceWarehouseId,
		BranchId: "1",
	}
	sourceWarehouse, err := client.GetWarehouses(argSourceWarehouse)
	if err != nil {
		return errors.NewServerError(model.AddNewInternalStockTransferError, err.Error())
	}

	argDestinationWarehouse := client.GetWarehousesRequest{
		Id:       result.DestinationWarehouseId,
		BranchId: "1",
	}
	destinationWarehouse, err := client.GetWarehouses(argDestinationWarehouse)
	if err != nil {
		return errors.NewServerError(model.AddNewInternalStockTransferError, err.Error())
	}

	resultItems, err := a.dbTrx.GetInternalStockTransferItems(context.Background(), result.TransactionId)

	res := model.AddInternalStockTransferResponse{
		InternalStockTransfer: model.InternalStockTransfer{
			TransactionId:            result.TransactionId,
			SourceWarehouseId:        result.SourceWarehouseId,
			SourceWarehouseName:      sourceWarehouse.Result.Warehouses[0].Name,
			DestinationWarehouseId:   result.DestinationWarehouseId,
			DestinationWarehouseName: destinationWarehouse.Result.Warehouses[0].Name,
			FormNumber:               result.FormNumber,
			TransactionDate:          result.TransactionDate,
			Items:                    util.InternalStockTransferItemDbToApi(resultItems),
		},
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
