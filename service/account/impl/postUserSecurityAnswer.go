package impl

import (
	"context"
	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/account/db/sqlc"
	"github.com/expert-pancake/service/account/model"
	"net/http"
)

func (a accountService) PostUserSecurityAnswer(w http.ResponseWriter, r *http.Request) error {

	var req model.PostUserSecurityAnswerRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	questionResult, err := a.dbTrx.GetUserInfo(context.Background(), db.GetUserInfoParams{
		UserID: req.AccountId,
		Key:    model.UserSecurityQuestionKey,
	})
	if err != nil {
		return errors.NewServerError(model.GetUserSecurityQuestionError, err.Error())
	}

	answerResult, err := a.dbTrx.GetUserInfo(context.Background(), db.GetUserInfoParams{
		UserID: req.AccountId,
		Key:    model.UserSecurityAnswerKey,
	})

	if questionResult.Value != req.Question || answerResult.Value != req.Answer {
		return errors.NewServerError(model.PostUserSecurityAnswerError, "security question or answer doesn't match our record")
	}

	res := model.PostUserSecurityAnswerResponse{Message: "OK"}
	httpHandler.WriteResponse(w, res)

	return nil
}
