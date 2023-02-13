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
	uuid "github.com/satori/go.uuid"
)

func (a inventoryService) UpsertItemReorder(w http.ResponseWriter, r *http.Request) error {

	var req model.UpsertItemReorderRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	var id = ""
	if req.Id == "" {
		id = uuid.NewV4().String()
	} else {
		id = req.Id
	}

	minimumStock, _ := strconv.ParseInt(req.MinimumStock, 10, 64)
	arg := db.UpsertItemReorderParams{
		ID:           id,
		VariantID:    req.VariantId,
		ItemUnitID:   req.ItemUnitId,
		WarehouseID:  req.WarehouseId,
		MinimumStock: minimumStock,
	}

	result, err := a.dbTrx.UpsertItemReorder(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpsertItemReorderError, err.Error())
	}

	itemReorder, err := a.dbTrx.GetItemReorder(context.Background(), result.ID)
	if err != nil {
		return errors.NewServerError(model.UpsertItemReorderError, err.Error())
	}

	argGetWarehouse := client.GetWarehousesRequest{
		Id:       itemReorder.WarehouseID,
		BranchId: "1",
	}
	warehouse, err := client.GetWarehouses(argGetWarehouse)
	if err != nil {
		return errors.NewServerError(model.UpsertItemReorderError, err.Error())
	}

	res := model.UpsertItemReorderResponse{
		ItemReorder: model.ItemReorder{
			Id:            itemReorder.ID,
			VariantId:     itemReorder.VariantID,
			VariantName:   itemReorder.VariantName,
			ItemUnitId:    itemReorder.ItemUnitID,
			ItemUnitName:  itemReorder.ItemUnitName,
			ItemId:        itemReorder.ItemID,
			ItemName:      itemReorder.ItemName,
			WarehouseId:   warehouse.Result.Warehouses[0].WarehouseId,
			WarehouseName: warehouse.Result.Warehouses[0].Name,
			MinimumStock:  strconv.FormatInt(itemReorder.MinimumStock, 10),
		},
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
