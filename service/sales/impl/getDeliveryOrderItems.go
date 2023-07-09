package impl

import (
	"context"
	"net/http"
	"strconv"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/sales/impl/client"
	"github.com/expert-pancake/service/sales/model"
	"github.com/expert-pancake/service/sales/util"
)

func (a salesService) GetDeliveryOrderItems(w http.ResponseWriter, r *http.Request) error {

	var req model.GetDeliveryOrderItemsRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetDeliveryOrderItems(context.Background(), req.DeliveryOrderId)
	if err != nil {
		return errors.NewServerError(model.GetDeliveryOrderItemsError, err.Error())
	}

	var deliveryOrderItems = make([]model.DeliveryOrderItem, 0)

	for _, d := range result {
		warehouseRackName := ""
		argWarehouseRack := client.GetWarehouseRacksRequest{
			Id:          d.WarehouseRackID,
			WarehouseId: "1",
		}
		warehouseRack, err := client.GetWarehouseRacks(argWarehouseRack)
		if err != nil {
			return errors.NewServerError(model.GetDeliveryOrderItemsError, err.Error())
		}

		if len(warehouseRack.Result.WarehouseRacks) > 0 {
			warehouseRackName = warehouseRack.Result.WarehouseRacks[0].Name
		}

		argItemVariant := client.GetItemVariantsRequest{
			Id: d.PrimaryItemVariantID,
		}
		itemVariant, err := client.GetItemVariants(argItemVariant)
		if err != nil {
			return errors.NewServerError(model.GetDeliveryOrderItemsError, err.Error())
		}

		itemCode := itemVariant.Result.ItemVariants[0].Code
		itemName := itemVariant.Result.ItemVariants[0].Name
		itemVariantName := itemVariant.Result.ItemVariants[0].VariantName

		argItemUnit := client.GetItemUnitsRequest{
			Id:     d.PrimaryItemUnitID,
			ItemId: itemVariant.Result.ItemVariants[0].ItemId,
		}
		itemUnit, err := client.GetItemUnits(argItemUnit)
		if err != nil {
			return errors.NewServerError(model.GetDeliveryOrderItemsError, err.Error())
		}

		itemUnitName := itemUnit.Result.ItemUnits[0].UnitName

		var batch, expired_date string
		if d.Batch.Valid {
			batch = d.Batch.String
		}
		if d.ExpiredDate.Valid {
			expired_date = d.ExpiredDate.Time.Format(util.DateLayoutYMD)
		}

		var deliveryOrderItem = model.DeliveryOrderItem{
			DetailId:               d.ID,
			PurchaseOrderItemId:    d.PurchaseOrderItemID,
			SalesOrderItemId:       d.SalesOrderItemID,
			DeliveryOrderId:        d.DeliveryOrderID,
			PrimaryItemVariantId:   d.PrimaryItemVariantID,
			ItemCode:               itemCode,
			ItemName:               itemName,
			ItemVariantName:        itemVariantName,
			WarehouseRackId:        d.WarehouseRackID,
			WarehouseRackName:      warehouseRackName,
			Batch:                  batch,
			ExpiredDate:            expired_date,
			ItemBarcodeId:          d.ItemBarcodeID,
			SecondaryItemVariantId: d.SecondaryItemVariantID,
			PrimaryItemUnitId:      d.PrimaryItemUnitID,
			ItemUnitName:           itemUnitName,
			SecondaryItemUnitId:    d.SecondaryItemUnitID,
			PrimaryItemUnitValue:   strconv.FormatInt(d.PrimaryItemUnitValue, 10),
			SecondaryItemUnitValue: strconv.FormatInt(d.SecondaryItemUnitValue, 10),
			Amount:                 strconv.FormatInt(d.Amount, 10),
		}

		deliveryOrderItems = append(deliveryOrderItems, deliveryOrderItem)
	}

	res := model.UpdateDeliveryOrderItemsResponse{
		DeliveryOrderItems: deliveryOrderItems,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
