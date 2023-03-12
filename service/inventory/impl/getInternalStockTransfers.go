package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/inventory/db/sqlc"
	"github.com/expert-pancake/service/inventory/impl/client"
	"github.com/expert-pancake/service/inventory/model"
	"github.com/expert-pancake/service/inventory/util"
)

func (a inventoryService) GetInternalStockTransfers(w http.ResponseWriter, r *http.Request) error {

	var req model.GetInternalStockTransfersRequest
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
		return errors.NewServerError(model.GetInternalStockTransfersError, err.Error())
	}

	var warehouseIdsParams = make([]string, 0)
	for _, d := range warehouseIds.Result.Warehouses {
		warehouseIdsParams = append(warehouseIdsParams, d.WarehouseId)
	}

	result, err := a.dbTrx.GetInternalStockTransfers(context.Background(), db.GetInternalStockTransfersParams{
		StartDate:    util.StringToDate(req.StartDate),
		EndDate:      util.StringToDate(req.EndDate),
		WarehouseIds: warehouseIdsParams,
	})
	if err != nil {
		return errors.NewServerError(model.GetInternalStockTransfersError, err.Error())
	}

	warehouseMap := util.WarehouseApiToMap(warehouseIds.Result.Warehouses)

	var datas = make([]model.InternalStockTransfer, 0)

	for _, d := range result {

		resultItems, err := a.dbTrx.GetInternalStockTransferItems(context.Background(), d.ID)
		if err != nil {
			return errors.NewServerError(model.GetInternalStockTransfersError, err.Error())
		}

		var data = model.InternalStockTransfer{
			TransactionId:            d.ID,
			SourceWarehouseId:        d.SourceWarehouseID,
			SourceWarehouseName:      warehouseMap[d.SourceWarehouseID],
			DestinationWarehouseId:   d.DestinationWarehouseID,
			DestinationWarehouseName: warehouseMap[d.DestinationWarehouseID],
			FormNumber:               d.FormNumber,
			TransactionDate:          d.TransactionDate.Format(util.DateLayoutYMD),
			Items:                    util.InternalStockTransferItemDbToApi(resultItems),
		}
		datas = append(datas, data)
	}

	res := model.GetInternalStockTransfersResponse{
		InternalStockTransfers: datas,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
