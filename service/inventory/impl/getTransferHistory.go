package impl

import (
	"context"
	"net/http"
	"strconv"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/inventory/db/sqlc"
	"github.com/expert-pancake/service/inventory/impl/client"
	"github.com/expert-pancake/service/inventory/model"
	"github.com/expert-pancake/service/inventory/util"
)

func (a inventoryService) GetTransferHistory(w http.ResponseWriter, r *http.Request) error {

	var req model.GetTransferHistoryRequest
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
		return errors.NewServerError(model.GetTransferHistoryError, err.Error())
	}

	var warehouseIdsParams = make([]string, 0)
	for _, d := range warehouseIds.Result.Warehouses {
		warehouseIdsParams = append(warehouseIdsParams, d.WarehouseId)
	}

	var isReceivedFilter = false
	if req.IsReceivedFilter != nil {
		isReceivedFilter = true
	} else {
		req.IsReceivedFilter = &isReceivedFilter
	}

	result, err := a.dbTrx.GetTransferHistory(context.Background(), db.GetTransferHistoryParams{
		StartDate:              util.StringToDate(req.StartDate),
		EndDate:                util.StringToDate(req.EndDate),
		ItemID:                 util.WildCardString(req.ItemId),
		IsReceivedFilter:       isReceivedFilter,
		IsReceived:             *req.IsReceivedFilter,
		WarehouseIds:           warehouseIdsParams,
		SourceWarehouseID:      util.WildCardString(req.WarehouseId),
		DestinationWarehouseID: util.WildCardString(req.WarehouseId),
	})
	if err != nil {
		return errors.NewServerError(model.GetTransferHistoryError, err.Error())
	}

	var datas = make([]model.TransferHistory, 0)

	for _, d := range result {
		argSourceWarehouse := client.GetWarehousesRequest{
			Id:       d.SourceWarehouseID,
			BranchId: "1",
		}
		sourceWarehouse, err := client.GetWarehouses(argSourceWarehouse)
		if err != nil {
			return errors.NewServerError(model.GetTransferHistoryError, err.Error())
		}

		argDestinationWarehouse := client.GetWarehousesRequest{
			Id:       d.DestinationWarehouseID,
			BranchId: "1",
		}
		destinationWarehouse, err := client.GetWarehouses(argDestinationWarehouse)
		if err != nil {
			return errors.NewServerError(model.GetTransferHistoryError, err.Error())
		}

		var data = model.TransferHistory{
			SourceWarehouseId:        d.SourceWarehouseID,
			SourceWarehouseName:      sourceWarehouse.Result.Warehouses[0].Name,
			DestinationWarehouseId:   d.DestinationWarehouseID,
			DestinationWarehouseName: destinationWarehouse.Result.Warehouses[0].Name,
			FormNumber:               d.FormNumber,
			TransactionDate:          d.TransactionDate.Format(util.DateLayoutYMD),
			ItemId:                   d.ItemID,
			ItemName:                 d.ItemName,
			ItemImageUrl:             d.ImageUrl,
			VariantId:                d.VariantID,
			VariantName:              d.VariantName,
			Amount:                   strconv.FormatInt(d.Amount, 10),
		}
		datas = append(datas, data)
	}

	res := model.GetTransferHistoryResponse{
		TransferHistories: datas,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
