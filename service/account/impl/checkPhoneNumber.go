package impl

import (
	"context"
	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/account/model"
	"net/http"
)

func (a accountService) CheckPhoneNumber(w http.ResponseWriter, r *http.Request) error {

	var req model.PhoneNumberRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetUserByPhoneNumber(context.Background(), req.PhoneNumber)
	if err != nil {
		return errors.NewServerError(model.GetUserByPhoneNumberError, err.Error())
	}

	res := model.PhoneNumberResponse{AccountId: result}
	httpHandler.WriteResponse(w, res)

	return nil
}
