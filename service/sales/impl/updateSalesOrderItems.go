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

func (a salesService) UpdateSalesOrderItems(w http.ResponseWriter, r *http.Request) error {

	var req model.UpdateSalesOrderItemsRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.UpdateSalesOrderItemsTrxParams{
		SalesOrderId:    req.SalesOrderId,
		SalesOrderItems: req.SalesOrderItems,
	}

	result, err := a.dbTrx.UpdateSalesOrderItemsTrx(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpdateSalesOrderItemsError, err.Error())
	}

	for idx, d := range result.SalesOrderItems {
		argItemVariant := client.GetItemVariantsRequest{
			Id: d.PrimaryItemVariantId,
		}
		itemVariant, err := client.GetItemVariants(argItemVariant)
		if err != nil {
			return errors.NewServerError(model.UpdateSalesOrderItemsError, err.Error())
		}

		result.SalesOrderItems[idx].ItemCode = itemVariant.Result.ItemVariants[0].Code
		result.SalesOrderItems[idx].ItemName = itemVariant.Result.ItemVariants[0].Name
		result.SalesOrderItems[idx].ItemVariantName = itemVariant.Result.ItemVariants[0].VariantName

		argItemUnit := client.GetItemUnitsRequest{
			Id:     d.PrimaryItemUnitId,
			ItemId: itemVariant.Result.ItemVariants[0].ItemId,
		}
		itemUnit, err := client.GetItemUnits(argItemUnit)
		if err != nil {
			return errors.NewServerError(model.UpdateSalesOrderItemsError, err.Error())
		}

		result.SalesOrderItems[idx].ItemUnitName = itemUnit.Result.ItemUnits[0].UnitName
	}

	res := model.UpdateSalesOrderItemsResponse{
		SalesOrderItems: result.SalesOrderItems,
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
