package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/inventory/model"
)

func (a inventoryService) DeleteItemUnit(w http.ResponseWriter, r *http.Request) error {

	var req model.DeleteItemUnitRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	err := a.dbTrx.DeleteItemUnit(context.Background(), req.ItemUnitId)
	if err != nil {
		return errors.NewServerError(model.DeleteItemUnitError, err.Error())
	}

	res := model.DeleteItemUnitResponse{
		Message: "OK",
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
