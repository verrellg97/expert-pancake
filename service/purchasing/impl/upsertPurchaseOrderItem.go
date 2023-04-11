package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/purchasing/db/transaction"
	"github.com/expert-pancake/service/purchasing/impl/client"
	"github.com/expert-pancake/service/purchasing/model"
	uuid "github.com/satori/go.uuid"
)

func (a purchasingService) UpsertPurchaseOrderItem(w http.ResponseWriter, r *http.Request) error {

	var req model.UpsertPurchaseOrderItemRequest

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

	arg := db.UpsertPurchaseOrderItemTrxParams{
		Id:                     id,
		PurchaseOrderId:        req.PurchaseOrderId,
		PrimaryItemVariantId:   req.PrimaryItemVariantId,
		SecondaryItemVariantId: req.SecondaryItemVariantId,
		PrimaryItemUnitId:      req.PrimaryItemUnitId,
		SecondaryItemUnitId:    req.SecondaryItemUnitId,
		PrimaryItemUnitValue:   req.PrimaryItemUnitValue,
		SecondaryItemUnitValue: req.SecondaryItemUnitValue,
		Amount:                 req.Amount,
		Price:                  req.Price,
	}

	result, err := a.dbTrx.UpsertPurchaseOrderItemTrx(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpsertPurchaseOrderItemError, err.Error())
	}

	argItemVariant := client.GetItemVariantsRequest{
		Id: result.PrimaryItemVariantId,
	}
	itemVariant, err := client.GetItemVariants(argItemVariant)
	if err != nil {
		return errors.NewServerError(model.UpsertPurchaseOrderItemError, err.Error())
	}

	argItemUnit := client.GetItemUnitsRequest{
		Id:     result.PrimaryItemUnitId,
		ItemId: itemVariant.Result.ItemVariants[0].ItemId,
	}
	itemUnit, err := client.GetItemUnits(argItemUnit)
	if err != nil {
		return errors.NewServerError(model.UpsertPurchaseOrderItemError, err.Error())
	}

	res := model.UpsertPurchaseOrderItemResponse{
		PurchaseOrderItem: model.PurchaseOrderItem{
			DetailId:               result.DetailId,
			PurchaseOrderId:        result.PurchaseOrderId,
			PrimaryItemVariantId:   result.PrimaryItemVariantId,
			ItemCode:               itemVariant.Result.ItemVariants[0].Code,
			ItemName:               itemVariant.Result.ItemVariants[0].Name,
			ItemVariantName:        itemVariant.Result.ItemVariants[0].VariantName,
			SecondaryItemVariantId: result.SecondaryItemVariantId,
			PrimaryItemUnitId:      result.PrimaryItemUnitId,
			ItemUnitName:           itemUnit.Result.ItemUnits[0].UnitName,
			SecondaryItemUnitId:    result.SecondaryItemUnitId,
			PrimaryItemUnitValue:   result.PrimaryItemUnitValue,
			SecondaryItemUnitValue: result.SecondaryItemUnitValue,
			Amount:                 result.Amount,
			Price:                  result.Price,
		},
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
