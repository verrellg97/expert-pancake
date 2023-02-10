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

func (a inventoryService) GetUnitCategories(w http.ResponseWriter, r *http.Request) error {

	var req model.GetUnitCategoriesRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetUnitCategories(context.Background(), db.GetUnitCategoriesParams{
		CompanyID: req.CompanyId,
		Name:      util.WildCardString(req.Keyword),
	})
	if err != nil {
		return errors.NewServerError(model.GetUnitCategoriesError, err.Error())
	}

	var unitCategories = make([]model.UnitCategory, 0)

	for _, d := range result {
		var unitCategory = model.UnitCategory{
			Id:        d.ID,
			CompanyId: d.CompanyID,
			Name:      d.Name,
		}
		unitCategories = append(unitCategories, unitCategory)
	}

	res := model.GetUnitCategoriesResponse{
		UnitCategories: unitCategories,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
