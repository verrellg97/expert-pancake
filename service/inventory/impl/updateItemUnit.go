package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/inventory/db/sqlc"
	"github.com/expert-pancake/service/inventory/model"
)

func (a inventoryService) UpdateItemUnit(w http.ResponseWriter, r *http.Request) error {

	var req model.UpdateItemUnitRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.UpdateItemUnitParams{
		ID:        req.Id,
		Name:      req.Name,
	}

	result, err := a.dbTrx.UpdateItemUnit(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpdateItemUnitError, err.Error())
	}

	res := model.UpdateItemUnitResponse{
		Unit: model.Unit{
			ItemUnitId: result.ID,
			CompanyId:   result.CompanyID,
			Name:        result.Name,
		},
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
