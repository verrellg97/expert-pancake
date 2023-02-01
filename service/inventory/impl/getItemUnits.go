package impl

import (
	"context"
	"net/http"
	"strconv"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/inventory/db/sqlc"
	"github.com/expert-pancake/service/inventory/model"
	"github.com/expert-pancake/service/inventory/util"
)

func (a inventoryService) GetItemUnits(w http.ResponseWriter, r *http.Request) error {

	var req model.GetItemUnitsRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetItemUnits(context.Background(), db.GetItemUnitsParams{
		ItemID: req.ItemId,
		Name:   util.WildCardString(req.Keyword),
	})
	if err != nil {
		return errors.NewServerError(model.GetItemUnitsError, err.Error())
	}

	var itemUnits = make([]model.ItemUnit, 0)

	for _, d := range result {
		var itemUnit = model.ItemUnit{
			ItemUnitId: d.ID,
			ItemId:     d.ItemID,
			UnitId:     d.UnitID,
			UnitName:   d.UnitName,
			Value:      strconv.FormatInt(d.Value, 10),
			IsDefault:  d.IsDefault,
		}
		itemUnits = append(itemUnits, itemUnit)
	}

	res := model.GetItemUnitsResponse{
		ItemUnits: itemUnits,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
