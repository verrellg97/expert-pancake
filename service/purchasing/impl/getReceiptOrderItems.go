package impl

import (
	"context"
	"net/http"
	"strconv"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/purchasing/impl/client"
	"github.com/expert-pancake/service/purchasing/model"
	"github.com/expert-pancake/service/purchasing/util"
)

func (a purchasingService) GetReceiptOrderItems(w http.ResponseWriter, r *http.Request) error {

	var req model.GetReceiptOrderItemsRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetReceiptOrderItems(context.Background(), req.ReceiptOrderId)
	if err != nil {
		return errors.NewServerError(model.GetReceiptOrderItemsError, err.Error())
	}

	var receiptOrderitems = make([]model.ReceiptOrderItem, 0)

	for _, d := range result {
		argItemVariant := client.GetItemVariantsRequest{
			Id: d.PrimaryItemVariantID,
		}
		itemVariant, err := client.GetItemVariants(argItemVariant)
		if err != nil {
			return errors.NewServerError(model.GetReceiptOrderItemsError, err.Error())
		}

		argItemUnit := client.GetItemUnitsRequest{
			Id:     d.PrimaryItemUnitID,
			ItemId: itemVariant.Result.ItemVariants[0].ItemId,
		}
		itemUnit, err := client.GetItemUnits(argItemUnit)
		if err != nil {
			return errors.NewServerError(model.GetReceiptOrderItemsError, err.Error())
		}

		argWarehouseRack := client.GetWarehouseRacksRequest{
			Id:          d.WarehouseRackID,
			WarehouseId: "1",
		}
		warehouseRack, err := client.GetWarehouseRacks(argWarehouseRack)
		if err != nil {
			return errors.NewServerError(model.GetReceiptOrderItemsError, err.Error())
		}

		var batch, expiredDate *string
		if d.Batch.Valid {
			batch = &d.Batch.String
		}
		if d.ExpiredDate.Valid {
			expiredDate = new(string)
			*expiredDate = d.ExpiredDate.Time.Format(util.DateLayoutYMD)
		}
		var receiptOrderItem = model.ReceiptOrderItem{
			DetailId:                 d.ID,
			PurchaseOrderItemId:      d.PurchaseOrderItemID,
			SalesOrderItemId:         d.SalesOrderItemID,
			DeliveryOrderItemId:      d.DeliveryOrderItemID,
			ReceiptOrderId:           d.ReceiptOrderID,
			PrimaryItemVariantId:     d.PrimaryItemVariantID,
			ItemCode:                 itemVariant.Result.ItemVariants[0].Code,
			ItemName:                 itemVariant.Result.ItemVariants[0].Name,
			ItemVariantName:          itemVariant.Result.ItemVariants[0].VariantName,
			WarehouseRackId:          d.WarehouseRackID,
			WarehouseRackName:        warehouseRack.Result.WarehouseRacks[0].Name,
			Batch:                    batch,
			ExpiredDate:              expiredDate,
			ItemBarcodeId:            d.ItemBarcodeID,
			SecondaryItemVariantId:   d.SecondaryItemVariantID,
			PrimaryItemUnitId:        d.PrimaryItemUnitID,
			ItemUnitName:             itemUnit.Result.ItemUnits[0].UnitName,
			SecondaryItemUnitId:      d.SecondaryItemUnitID,
			PrimaryItemUnitValue:     strconv.FormatInt(d.PrimaryItemUnitValue, 10),
			SecondaryItemUnitValue:   strconv.FormatInt(d.SecondaryItemUnitValue, 10),
			AmountDelivered:          strconv.FormatInt(d.AmountDelivered, 10),
			Amount:                   strconv.FormatInt(d.Amount, 10),
		}
		receiptOrderitems = append(receiptOrderitems, receiptOrderItem)
	}

	res := model.GetReceiptOrderItemsResponse{
		ReceiptOrderItems: receiptOrderitems,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
