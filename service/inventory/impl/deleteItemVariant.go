package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/inventory/model"
)

func (a inventoryService) DeleteItemVariant(w http.ResponseWriter, r *http.Request) error {

	var req model.DeleteItemVariantRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	err := a.dbTrx.DeleteItemVariant(context.Background(), req.Id)
	if err != nil {
		return errors.NewServerError(model.DeleteItemVariantError, err.Error())
	}

	res := model.DeleteItemVariantResponse{
		Message: "OK",
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
