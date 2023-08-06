package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/business/db/transaction"
	"github.com/expert-pancake/service/business/model"
)

func (a businessService) UpdateMemberRequest(w http.ResponseWriter, r *http.Request) error {

	var req model.UpdateMemberRequestRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.UpdateMemberRequestTrxParams{
		Id:     req.Id,
		Status: req.Status,
	}

	result, err := a.dbTrx.UpdateMemberRequestTrx(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpdateMemberRequestError, err.Error())
	}

	res := model.UpdateMemberRequestResponse{
		Message: result.Message,
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
