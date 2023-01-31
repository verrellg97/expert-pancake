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

func (a inventoryService) AddUnit(w http.ResponseWriter, r *http.Request) error {

	var req model.AddUnitRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.InsertUnitParams{
		ID:        uuid.NewV4().String(),
		CompanyID: req.CompanyId,
		Name:      req.Name,
	}

	result, err := a.dbTrx.InsertUnit(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.AddNewUnitError, err.Error())
	}

	res := model.AddUnitResponse{
		Unit: model.Unit{
			UnitId:    result.ID,
			CompanyId: result.CompanyID,
			Name:      result.Name,
		},
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
