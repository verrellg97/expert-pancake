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

func (a inventoryService) GetItemReorders(w http.ResponseWriter, r *http.Request) error {

	var req model.GetItemReordersRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetItemReorders(context.Background(), db.GetItemReordersParams{
		WarehouseID: util.WildCardString(req.WarehouseId),
		ItemID:      util.WildCardString(req.ItemId),
	})
	if err != nil {
		return errors.NewServerError(model.GetItemReordersError, err.Error())
	}

	var itemReorders = make([]model.ItemReorder, 0)

	for _, d := range result {
		argGetWarehouse := client.GetWarehousesRequest{
			Id:       d.WarehouseID,
			BranchId: "1",
		}
		warehouse, err := client.GetWarehouses(argGetWarehouse)
		if err != nil {
			return errors.NewServerError(model.GetItemReordersError, err.Error())
		}

		var itemReorder = model.ItemReorder{
			Id:            d.ID,
			VariantId:     d.VariantID,
			VariantName:   d.VariantName,
			ItemUnitId:    d.ItemUnitID,
			ItemUnitName:  d.ItemUnitName,
			ItemId:        d.ItemID,
			ItemName:      d.ItemName,
			WarehouseId:   warehouse.Result.Warehouses[0].WarehouseId,
			WarehouseName: warehouse.Result.Warehouses[0].Name,
			MinimumStock:  strconv.FormatInt(d.MinimumStock, 10),
		}

		itemReorders = append(itemReorders, itemReorder)
	}

	res := itemReorders
	httpHandler.WriteResponse(w, res)

	return nil
}
