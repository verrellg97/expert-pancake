package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/inventory/db/sqlc"
	"github.com/expert-pancake/service/inventory/model"
	"github.com/expert-pancake/service/inventory/util"
)

func (a inventoryService) GetUnits(w http.ResponseWriter, r *http.Request) error {

	var req model.GetUnitsRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetUnits(context.Background(), db.GetUnitsParams{
		CompanyID:      req.CompanyId,
		UnitCategoryID: util.WildCardString(req.UnitCategoryId),
		Name:           util.WildCardString(req.Keyword),
	})
	if err != nil {
		return errors.NewServerError(model.GetUnitsError, err.Error())
	}

	var units = make([]model.Unit, 0)

	for _, d := range result {
		var unit = model.Unit{
			UnitId:         d.ID,
			CompanyId:      d.CompanyID,
			UnitCategoryId: d.UnitCategoryID,
			Name:           d.Name,
		}
		units = append(units, unit)
	}

	res := units
	httpHandler.WriteResponse(w, res)

	return nil
}
