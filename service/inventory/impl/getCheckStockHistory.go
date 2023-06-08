package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/inventory/impl/client"
	"github.com/expert-pancake/service/inventory/model"
)

func (a inventoryService) GetCheckStockHistory(w http.ResponseWriter, r *http.Request) error {

	var req model.GetCheckStockHistoryRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	argWarehouseIds := client.GetWarehousesRequest{
		BranchIds: req.BranchIds,
		BranchId:  req.BranchIds[0],
	}
	warehouseIds, err := client.GetWarehouses(argWarehouseIds)
	if err != nil {
		return errors.NewServerError(model.GetCheckStockHistoryError, err.Error())
	}

	var warehouseIdsParams = make([]string, 0)
	for _, d := range warehouseIds.Result.Warehouses {
		warehouseIdsParams = append(warehouseIdsParams, d.WarehouseId)
	}

	result, err := a.dbTrx.GetCheckStockHistory(context.Background(), warehouseIdsParams)
	if err != nil {
		return errors.NewServerError(model.GetCheckStockHistoryError, err.Error())
	}

	var status = false
	if result > 0 {
		status = true
	}

	res := model.GetCheckStockHistoryResponse{
		Status: status,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
