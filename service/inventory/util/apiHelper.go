package util

import (
	"strconv"

	db "github.com/expert-pancake/service/inventory/db/sqlc"
	"github.com/expert-pancake/service/inventory/impl/client"
	"github.com/expert-pancake/service/inventory/model"
)

func InternalStockTransferItemDbToApi(data []db.GetInternalStockTransferItemsRow) []model.InternalStockTransferItem {
	var items = make([]model.InternalStockTransferItem, 0)

	for _, d := range data {
		var batch, expired_date *string
		if d.Batch.Valid {
			batch = &d.Batch.String
		}
		if d.ExpiredDate.Valid {
			expired_date = new(string)
			*expired_date = d.ExpiredDate.Time.Format(DateLayoutYMD)
		}

		argWarehouseRack := client.GetWarehouseRacksRequest{
			Id:          d.WarehouseRackID,
			WarehouseId: "1",
		}
		warehouseRack, _ := client.GetWarehouseRacks(argWarehouseRack)

		items = append(items, model.InternalStockTransferItem{
			DetailId:          d.ID,
			WarehouseRackId:   d.WarehouseRackID,
			WarehouseRackName: warehouseRack.Result.WarehouseRacks[0].Name,
			ItemName:          d.ItemName,
			VariantId:         d.VariantID,
			VariantName:       d.VariantName,
			ItemUnitId:        d.ItemUnitID,
			ItemUnitName:      d.ItemUnitName,
			ItemUnitValue:     strconv.FormatInt(d.ItemUnitValue, 10),
			Amount:            strconv.FormatInt(d.Amount, 10),
			Batch:             batch,
			ExpiredDate:       expired_date,
		})
	}

	return items
}
