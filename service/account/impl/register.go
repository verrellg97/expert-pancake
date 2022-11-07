package impl

import (
	"context"
	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/account/db/transaction"
	"github.com/expert-pancake/service/account/model"
	"net/http"
)

func (a accountService) Register(w http.ResponseWriter, r *http.Request) error {

	var req model.RegisterRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.CreateNewUserTrxParams{
		FullName:         req.FullName,
		Nickname:         req.Nickname,
		Email:            req.Email,
		PhoneNumber:      req.PhoneNumber,
		Password:         req.Password,
		SecurityQuestion: req.SecurityQuestion,
		SecurityAnswer:   req.SecurityAnswer,
	}

	result, err := a.dbTrx.CreateNewUserTrx(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.CreateNewUserTransactionError, err.Error())
	}

	res := model.RegisterResponse{
		User: model.User{
			AccountId:   result.Id,
			FullName:    result.FullName,
			Nickname:    result.Nickname,
			Email:       result.Email,
			PhoneNumber: result.PhoneNumber,
		},
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
