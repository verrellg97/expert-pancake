package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/inventory/db/sqlc"
	"github.com/expert-pancake/service/inventory/model"
	uuid "github.com/satori/go.uuid"
)

func (a inventoryService) UpsertUnitCategory(w http.ResponseWriter, r *http.Request) error {

	var req model.UpsertUnitCategoryRequest

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

	arg := db.UpsertUnitCategoryParams{
		ID:        id,
		CompanyID: req.CompanyId,
		Name:      req.Name,
	}

	result, err := a.dbTrx.UpsertUnitCategory(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpsertUnitCategoryError, err.Error())
	}

	res := model.UpsertUnitCategoryResponse{
		UnitCategory: model.UnitCategory{
			Id:        result.ID,
			CompanyId: result.CompanyID,
			Name:      result.Name,
		},
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
