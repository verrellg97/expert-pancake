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

func (a inventoryService) UpsertItemUnit(w http.ResponseWriter, r *http.Request) error {

	var req model.UpsertItemUnitRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	var id = ""
	if req.ItemUnitId == "" {
		id = uuid.NewV4().String()
	} else {
		id = req.ItemUnitId
	}

	value, _ := strconv.ParseInt(req.Value, 10, 64)
	arg := db.UpsertItemUnitParams{
		ID:        id,
		ItemID:    req.ItemId,
		UnitID:    req.UnitId,
		Value:     value,
		IsDefault: req.IsDefault,
	}

	result, err := a.dbTrx.UpsertItemUnit(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpsertItemUnitError, err.Error())
	}

	unitRes, err := a.dbTrx.GetUnit(context.Background(), result.UnitID)
	if err != nil {
		return errors.NewServerError(model.UpsertItemUnitError, err.Error())
	}

	res := model.UpsertItemUnitResponse{
		ItemUnit: model.ItemUnit{
			ItemUnitId: result.ID,
			ItemId:     result.ItemID,
			UnitId:     result.UnitID,
			UnitName:   unitRes.Name,
			Value:      strconv.FormatInt(result.Value, 10),
			IsDefault:  result.IsDefault,
		},
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
