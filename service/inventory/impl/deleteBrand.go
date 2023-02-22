package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/inventory/model"
)

func (a inventoryService) DeleteBrand(w http.ResponseWriter, r *http.Request) error {

	var req model.DeleteBrandRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	err := a.dbTrx.DeleteBrand(context.Background(), req.Id)
	if err != nil {
		return errors.NewServerError(model.DeleteBrandError, err.Error())
	}

	res := model.DeleteBrandResponse{
		Message: "OK",
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
