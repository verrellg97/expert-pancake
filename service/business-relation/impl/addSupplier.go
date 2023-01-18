package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/business-relation/model"
)

func (a businessRelationService) AddSupplier(w http.ResponseWriter, r *http.Request) error {

	var req model.AddSupplierRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	err := a.dbTrx.AddSupplier(context.Background(), req.ContactBookIds)
	if err != nil {
		return errors.NewServerError(model.AddSupplierError, err.Error())
	}

	res := model.AddSupplierResponse{
		Message: "OK",
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
