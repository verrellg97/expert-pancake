package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/warehouse/model"
)

func (a warehouseService) DeleteWarehouse(w http.ResponseWriter, r *http.Request) error {

	var req model.DeleteWarehouseRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	err := a.dbTrx.DeleteWarehouse(context.Background(), req.WarehouseId)
	if err != nil {
		return errors.NewServerError(model.DeleteWarehouseError, err.Error())
	}

	res := model.DeleteWarehouseResponse{
		Message: "OK",
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
