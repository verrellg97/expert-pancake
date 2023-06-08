package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/sales/model"
)

func (a salesService) GetCheckPOS(w http.ResponseWriter, r *http.Request) error {

	var req model.GetCheckPOSRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetCheckPOS(context.Background(), req.CompanyId)
	if err != nil {
		return errors.NewServerError(model.GetCheckPOSError, err.Error())
	}

	var status = false
	if result > 0 {
		status = true
	}

	res := model.GetCheckPOSResponse{
		Status: status,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
