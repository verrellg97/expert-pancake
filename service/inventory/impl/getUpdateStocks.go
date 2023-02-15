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

func (a inventoryService) GetUpdateStocks(w http.ResponseWriter, r *http.Request) error {

	var req model.GetUpdateStocksRequest
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
		return errors.NewServerError(model.GetUpdateStocksError, err.Error())
	}

	var warehouseIdsParams = make([]string, 0)
	for _, d := range warehouseIds.Result.Warehouses {
		warehouseIdsParams = append(warehouseIdsParams, d.WarehouseId)
	}

	result, err := a.dbTrx.GetUpdateStocks(context.Background(), db.GetUpdateStocksParams{
		StartDate:    util.StringToDate(req.StartDate),
		EndDate:      util.StringToDate(req.EndDate),
		WarehouseIds: warehouseIdsParams,
	})
	if err != nil {
		return errors.NewServerError(model.GetUpdateStocksError, err.Error())
	}

	var datas = make([]model.UpdateStock, 0)

	for _, d := range result {
		argWarehouse := client.GetWarehousesRequest{
			Id:       d.WarehouseID,
			BranchId: "1",
		}
		warehouse, err := client.GetWarehouses(argWarehouse)
		if err != nil {
			return errors.NewServerError(model.GetUpdateStocksError, err.Error())
		}

		argWarehouseRack := client.GetWarehouseRacksRequest{
			Id:          d.WarehouseRackID,
			WarehouseId: "1",
		}
		warehouseRack, err := client.GetWarehouseRacks(argWarehouseRack)
		if err != nil {
			return errors.NewServerError(model.GetUpdateStocksError, err.Error())
		}

		var batch, expired_date *string
		if d.Batch.Valid {
			batch = &d.Batch.String
		}
		if d.ExpiredDate.Valid {
			expired_date = new(string)
			*expired_date = d.ExpiredDate.Time.Format(util.DateLayoutYMD)
		}

		var data = model.UpdateStock{
			TransactionId:     d.ID,
			FormNumber:        d.FormNumber,
			TransactionDate:   d.TransactionDate.Format(util.DateLayoutYMD),
			WarehouseId:       d.WarehouseID,
			WarehouseName:     warehouse.Result.Warehouses[0].Name,
			WarehouseRackId:   d.WarehouseRackID,
			WarehouseRackName: warehouseRack.Result.WarehouseRacks[0].Name,
			ItemId:            d.ItemID,
			ItemName:          d.ItemName,
			VariantId:         d.VariantID,
			VariantName:       d.VariantName,
			ItemUnitId:        d.ItemUnitID,
			ItemUnitName:      d.ItemUnitName,
			ItemUnitValue:     strconv.FormatInt(d.ItemUnitValue, 10),
			BeginningStock:    strconv.FormatInt(d.BeginningStock, 10),
			EndingStock:       strconv.FormatInt(d.EndingStock, 10),
			Batch:             batch,
			ExpiredDate:       expired_date,
		}
		datas = append(datas, data)
	}

	res := model.GetUpdateStocksResponse{
		UpdateStocks: datas,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
