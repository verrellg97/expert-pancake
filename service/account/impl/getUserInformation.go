package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/account/model"
)

func (a accountService) GetUserInformation(w http.ResponseWriter, r *http.Request) error {

	var req model.GetUserInformationRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetUser(context.Background(), req.AccountId)
	if err != nil {
		return errors.NewServerError(model.GetUserInformationError, err.Error())
	}

	res := model.GetUserInformationResponse{
		AccountId:   result.ID,
		ImageUrl:    result.ImageUrl,
		FullName:    result.Fullname,
		Nickname:    result.Nickname,
		Email:       result.Email.String,
		PhoneNumber: result.PhoneNumber,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
