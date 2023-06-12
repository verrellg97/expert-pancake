package impl

import (
	"context"
	"net/http"
	"strconv"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/sales/impl/client"
	"github.com/expert-pancake/service/sales/model"
)

func (a salesService) GetSalesOrderItems(w http.ResponseWriter, r *http.Request) error {

	var req model.GetSalesOrderItemsRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetSalesOrderItems(context.Background(), req.SalesOrderId)
	if err != nil {
		return errors.NewServerError(model.GetSalesOrderItemsError, err.Error())
	}

	var salesOrderItems = make([]model.SalesOrderItem, 0)

	for _, d := range result {
		argItemVariant := client.GetItemVariantsRequest{
			Id: d.PrimaryItemVariantID,
		}
		itemVariant, err := client.GetItemVariants(argItemVariant)
		if err != nil {
			return errors.NewServerError(model.GetSalesOrderItemsError, err.Error())
		}

		argItemUnit := client.GetItemUnitsRequest{
			Id:     d.PrimaryItemUnitID,
			ItemId: itemVariant.Result.ItemVariants[0].ItemId,
		}
		itemUnit, err := client.GetItemUnits(argItemUnit)
		if err != nil {
			return errors.NewServerError(model.GetSalesOrderItemsError, err.Error())
		}
		var salesOrderItem = model.SalesOrderItem{
			DetailId:               d.ID,
			PurchaseOrderItemId:    d.PurchaseOrderItemID,
			SalesOrderId:           d.SalesOrderID,
			PrimaryItemVariantId:   d.PrimaryItemVariantID,
			ItemCode:               itemVariant.Result.ItemVariants[0].Code,
			ItemName:               itemVariant.Result.ItemVariants[0].Name,
			ItemVariantName:        itemVariant.Result.ItemVariants[0].VariantName,
			SecondaryItemVariantId: d.SecondaryItemVariantID,
			PrimaryItemUnitId:      d.PrimaryItemUnitID,
			ItemUnitName:           itemUnit.Result.ItemUnits[0].UnitName,
			SecondaryItemUnitId:    d.SecondaryItemUnitID,
			PrimaryItemUnitValue:   strconv.FormatInt(d.PrimaryItemUnitValue, 10),
			SecondaryItemUnitValue: strconv.FormatInt(d.SecondaryItemUnitValue, 10),
			Amount:                 strconv.FormatInt(d.Amount, 10),
			Price:                  strconv.FormatInt(d.Price, 10),
		}
		salesOrderItems = append(salesOrderItems, salesOrderItem)
	}

	res := model.GetSalesOrderItemsResponse{
		SalesOrderItems: salesOrderItems,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
