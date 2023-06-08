package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/purchasing/model"
)

func (a purchasingService) GetCheckPurchaseOrders(w http.ResponseWriter, r *http.Request) error {

	var req model.GetCheckPurchaseOrdersRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetCheckPurchaseOrders(context.Background(), req.CompanyId)
	if err != nil {
		return errors.NewServerError(model.GetCheckPurchaseOrdersError, err.Error())
	}

	var status = false
	if result > 0 {
		status = true
	}

	res := model.GetCheckPurchaseOrdersResponse{
		Status: status,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
