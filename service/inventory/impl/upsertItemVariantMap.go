package impl

import (
	"context"
	"net/http"
	"strconv"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/inventory/db/sqlc"
	"github.com/expert-pancake/service/inventory/model"
	uuid "github.com/satori/go.uuid"
)

func (a inventoryService) UpsertItemVariantMap(w http.ResponseWriter, r *http.Request) error {

	var req model.UpsertItemVariantMapRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	var id = ""
	if req.ItemVariantMapId == "" {
		id = uuid.NewV4().String()
	} else {
		id = req.ItemVariantMapId
	}

	arg := db.UpsertItemVariantMapParams{
		ID:                     id,
		PrimaryCompanyID:       req.PrimaryCompanyId,
		PrimaryItemVariantID:   req.PrimaryItemVariantId,
		PrimaryItemUnitID:      req.PrimaryItemUnitId,
		SecondaryCompanyID:     req.SecondaryCompanyId,
		SecondaryItemVariantID: req.SecondaryItemVariantId,
		SecondaryItemUnitID:    req.SecondaryItemUnitId,
	}

	err := a.dbTrx.UpsertItemVariantMap(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpsertItemVariantMapError, err.Error())
	}

	result, err := a.dbTrx.GetItemVariantMap(context.Background(), id)
	if err != nil {
		return errors.NewServerError(model.UpsertItemVariantMapError, err.Error())
	}

	res := model.UpsertItemVariantMapResponse{
		ItemVariantMap: model.ItemVariantMap{
			ItemVariantMap:            result.ID,
			PrimaryItemId:             result.PrimaryItemID,
			PrimaryItemName:           result.PrimaryItemName,
			PrimaryItemVariantId:      result.PrimaryItemVariantID,
			PrimaryItemVariantName:    result.PrimaryItemVariantName,
			PrimaryItemVariantPrice:   strconv.FormatInt(result.PrimaryItemVariantPrice, 10),
			PrimaryItemUnitId:         result.PrimaryItemUnitID,
			PrimaryItemUnitName:       result.PrimaryItemUnitName,
			SecondaryItemId:           result.SecondaryItemID,
			SecondaryItemName:         result.SecondaryItemName,
			SecondaryItemVariantId:    result.SecondaryItemVariantID,
			SecondaryItemVariantName:  result.SecondaryItemVariantName,
			SecondaryItemVariantPrice: strconv.FormatInt(result.SecondaryItemVariantPrice, 10),
			SecondaryItemUnitId:       result.SecondaryItemUnitID,
			SecondaryItemUnitName:     result.SecondaryItemUnitName,
		},
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
