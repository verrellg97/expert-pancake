package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/sales/db/transaction"
	"github.com/expert-pancake/service/sales/impl/client"
	"github.com/expert-pancake/service/sales/model"
)

func (a salesService) UpdateDeliveryOrderItems(w http.ResponseWriter, r *http.Request) error {

	var req model.UpdateDeliveryOrderItemsRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.UpdateDeliveryOrderItemsTrxParams{
		DeliveryOrderId:    req.DeliveryOrderId,
		DeliveryOrderItems: req.DeliveryOrderItems,
	}

	result, err := a.dbTrx.UpdateDeliveryOrderItemsTrx(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpdateDeliveryOrderItemsError, err.Error())
	}

	for idx, d := range result.DeliveryOrderItems {
		argWarehouseRack := client.GetWarehouseRacksRequest{
			Id:          d.WarehouseRackId,
			WarehouseId: "1",
		}
		warehouseRack, err := client.GetWarehouseRacks(argWarehouseRack)
		if err != nil {
			return errors.NewServerError(model.UpdateDeliveryOrderItemsError, err.Error())
		}

		result.DeliveryOrderItems[idx].WarehouseRackName = warehouseRack.Result.WarehouseRacks[0].Name

		argItemVariant := client.GetItemVariantsRequest{
			Id: d.PrimaryItemVariantId,
		}
		itemVariant, err := client.GetItemVariants(argItemVariant)
		if err != nil {
			return errors.NewServerError(model.UpdateDeliveryOrderItemsError, err.Error())
		}

		result.DeliveryOrderItems[idx].ItemCode = itemVariant.Result.ItemVariants[0].Code
		result.DeliveryOrderItems[idx].ItemName = itemVariant.Result.ItemVariants[0].Name
		result.DeliveryOrderItems[idx].ItemVariantName = itemVariant.Result.ItemVariants[0].VariantName

		argItemUnit := client.GetItemUnitsRequest{
			Id:     d.PrimaryItemUnitId,
			ItemId: itemVariant.Result.ItemVariants[0].ItemId,
		}
		itemUnit, err := client.GetItemUnits(argItemUnit)
		if err != nil {
			return errors.NewServerError(model.UpdateDeliveryOrderItemsError, err.Error())
		}

		result.DeliveryOrderItems[idx].ItemUnitName = itemUnit.Result.ItemUnits[0].UnitName
	}

	res := model.UpdateDeliveryOrderItemsResponse{
		DeliveryOrderItems: result.DeliveryOrderItems,
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
