package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/business-relation/model"
)

func (a businessRelationService) AddCustomer(w http.ResponseWriter, r *http.Request) error {

	var req model.AddCustomerRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	err := a.dbTrx.AddCustomer(context.Background(), req.ContactBookIds)
	if err != nil {
		return errors.NewServerError(model.AddCustomerError, err.Error())
	}

	res := model.AddCustomerResponse{
		Message: "OK",
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
