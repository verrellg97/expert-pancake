package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/sales/model"
)

func (a salesService) DeletePOS(w http.ResponseWriter, r *http.Request) error {

	var req model.DeletePOSRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	err := a.dbTrx.DeletePOS(context.Background(), req.Id)
	if err != nil {
		return errors.NewServerError(model.DeletePOSError, err.Error())
	}

	res := model.DeletePOSResponse{
		Message: "OK",
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
