package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/purchasing/db/transaction"
	"github.com/expert-pancake/service/purchasing/impl/client"
	"github.com/expert-pancake/service/purchasing/model"
)

func (a purchasingService) UpdatePurchaseOrderItems(w http.ResponseWriter, r *http.Request) error {

	var req model.UpdatePurchaseOrderItemsRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.UpdatePurchaseOrderItemsTrxParams{
		PurchaseOrderId:    req.PurchaseOrderId,
		PurchaseOrderItems: req.PurchaseOrderItems,
	}

	result, err := a.dbTrx.UpdatePurchaseOrderItemsTrx(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpdatePurchaseOrderItemsError, err.Error())
	}

	for idx, d := range result.PurchaseOrderItems {
		argItemVariant := client.GetItemVariantsRequest{
			Id: d.PrimaryItemVariantId,
		}
		itemVariant, err := client.GetItemVariants(argItemVariant)
		if err != nil {
			return errors.NewServerError(model.UpdatePurchaseOrderItemsError, err.Error())
		}

		result.PurchaseOrderItems[idx].ItemCode = itemVariant.Result.ItemVariants[0].Code
		result.PurchaseOrderItems[idx].ItemName = itemVariant.Result.ItemVariants[0].Name
		result.PurchaseOrderItems[idx].ItemVariantName = itemVariant.Result.ItemVariants[0].VariantName

		argItemUnit := client.GetItemUnitsRequest{
			Id:     d.PrimaryItemUnitId,
			ItemId: itemVariant.Result.ItemVariants[0].ItemId,
		}
		itemUnit, err := client.GetItemUnits(argItemUnit)
		if err != nil {
			return errors.NewServerError(model.UpdatePurchaseOrderItemsError, err.Error())
		}

		result.PurchaseOrderItems[idx].ItemUnitName = itemUnit.Result.ItemUnits[0].UnitName
	}

	res := model.UpdatePurchaseOrderItemsResponse{
		PurchaseOrderItems: result.PurchaseOrderItems,
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
