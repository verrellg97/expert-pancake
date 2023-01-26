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

func (a inventoryService) AddItemUnit(w http.ResponseWriter, r *http.Request) error {

	var req model.AddItemUnitRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.InsertItemUnitParams{
		ID:        uuid.NewV4().String(),
		CompanyID: req.CompanyId,
		Name:      req.Name,
	}

	result, err := a.dbTrx.InsertItemUnit(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.AddNewItemUnitError, err.Error())
	}

	res := model.AddItemUnitResponse{
		Unit: model.Unit{
			ItemUnitId: result.ID,
			CompanyId:   result.CompanyID,
			Name:        result.Name,
		},
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
