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

	result, err := a.dbTrx.GetInternalStockTransfers(context.Background(), db.GetInternalStockTransfersParams{
		StartDate: util.StringToDate(req.StartDate),
		EndDate:   util.StringToDate(req.EndDate),
	})
	if err != nil {
		return errors.NewServerError(model.GetInternalStockTransfersError, err.Error())
	}

	var datas = make([]model.InternalStockTransfer, 0)

	for _, d := range result {
		argSourceWarehouse := client.GetWarehousesRequest{
			Id:       d.SourceWarehouseID,
			BranchId: "1",
		}
		sourceWarehouse, err := client.GetWarehouses(argSourceWarehouse)
		if err != nil {
			return errors.NewServerError(model.GetInternalStockTransfersError, err.Error())
		}

		argDestinationWarehouse := client.GetWarehousesRequest{
			Id:       d.DestinationWarehouseID,
			BranchId: "1",
		}
		destinationWarehouse, err := client.GetWarehouses(argDestinationWarehouse)
		if err != nil {
			return errors.NewServerError(model.GetInternalStockTransfersError, err.Error())
		}

		resultItems, err := a.dbTrx.GetInternalStockTransferItems(context.Background(), d.ID)

		var data = model.InternalStockTransfer{
			TransactionId:            d.ID,
			SourceWarehouseId:        d.SourceWarehouseID,
			SourceWarehouseName:      sourceWarehouse.Result.Warehouses[0].Name,
			DestinationWarehouseId:   d.DestinationWarehouseID,
			DestinationWarehouseName: destinationWarehouse.Result.Warehouses[0].Name,
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
