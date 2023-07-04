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

func (a salesService) GetSalesInvoiceItems(w http.ResponseWriter, r *http.Request) error {

	var req model.GetSalesInvoiceItemsRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetSalesInvoiceItems(context.Background(), req.SalesInvoiceId)
	if err != nil {
		return errors.NewServerError(model.GetSalesInvoiceItemsError, err.Error())
	}

	var salesInvoiceitems = make([]model.SalesInvoiceItem, 0)

	for _, d := range result {
		argItemVariant := client.GetItemVariantsRequest{
			Id: d.PrimaryItemVariantID,
		}
		itemVariant, err := client.GetItemVariants(argItemVariant)
		if err != nil {
			return errors.NewServerError(model.GetSalesInvoiceItemsError, err.Error())
		}

		argItemUnit := client.GetItemUnitsRequest{
			Id:     d.PrimaryItemUnitID,
			ItemId: itemVariant.Result.ItemVariants[0].ItemId,
		}
		itemUnit, err := client.GetItemUnits(argItemUnit)
		if err != nil {
			return errors.NewServerError(model.GetSalesInvoiceItemsError, err.Error())
		}

		var salesInvoiceItem = model.SalesInvoiceItem{
			Id:                     d.ID,
			PurchaseOrderItemId:    d.PurchaseOrderItemID,
			SalesOrderItemId:       d.SalesOrderItemID,
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
		salesInvoiceitems = append(salesInvoiceitems, salesInvoiceItem)
	}

	res := model.GetSalesInvoiceItemsResponse{
		SalesInvoiceItems: salesInvoiceitems,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
