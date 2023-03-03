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

func (a inventoryService) GetStockHistory(w http.ResponseWriter, r *http.Request) error {

	var req model.GetStockHistoryRequest
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

	result, err := a.dbTrx.GetStockHistory(context.Background(), db.GetStockHistoryParams{
		StartDate:    util.StringToDate(req.StartDate),
		EndDate:      util.StringToDate(req.EndDate),
		ItemID:       util.WildCardString(req.ItemId),
		WarehouseIds: warehouseIdsParams,
	})
	if err != nil {
		return errors.NewServerError(model.GetStockHistoryError, err.Error())
	}

	var datas = make([]model.StockHistory, 0)

	for _, d := range result {
		var data = model.StockHistory{
			FormNumber:      d.FormNumber,
			TransactionDate: d.TransactionDate.Format(util.DateLayoutYMD),
			ItemId:          d.ItemID,
			ItemName:        d.ItemName,
			ItemImageUrl:    d.ImageUrl,
			VariantId:       d.VariantID,
			VariantName:     d.VariantName,
			Onhand:          strconv.FormatInt(d.Onhand, 10),
			Calculated:      strconv.FormatInt(d.Calculated, 10),
		}
		datas = append(datas, data)
	}

	res := model.GetStockHistoryResponse{
		StockHistories: datas,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
