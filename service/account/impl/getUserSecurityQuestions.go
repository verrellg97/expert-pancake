package impl

import (
	"context"
	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/account/db/sqlc"
	"github.com/expert-pancake/service/account/model"
	"net/http"
)

func (a accountService) GetUserSecurityQuestion(w http.ResponseWriter, r *http.Request) error {

	var req model.GetUserSecurityQuestionRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetUserInfo(context.Background(), db.GetUserInfoParams{
		UserID: req.AccountId,
		Key:    model.UserSecurityQuestionKey,
	})
	if err != nil {
		return errors.NewServerError(model.GetUserSecurityQuestionError, err.Error())
	}

	res := model.GetUserSecurityQuestionResponse{Question: result.Value}
	httpHandler.WriteResponse(w, res)

	return nil
}
