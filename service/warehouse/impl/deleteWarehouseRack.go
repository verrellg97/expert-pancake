package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/warehouse/model"
)

func (a warehouseService) DeleteWarehouseRack(w http.ResponseWriter, r *http.Request) error {

	var req model.DeleteWarehouseRackRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	err := a.dbTrx.DeleteWarehouseRack(context.Background(), req.Id)
	if err != nil {
		return errors.NewServerError(model.DeleteWarehouseRackError, err.Error())
	}

	res := model.DeleteWarehouseRackResponse{
		Message: "OK",
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
