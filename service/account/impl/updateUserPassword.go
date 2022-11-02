package impl

import (
	"context"
	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/account/db/sqlc"
	"github.com/expert-pancake/service/account/model"
	"github.com/expert-pancake/service/account/util"
	"net/http"
)

func (a accountService) UpdateUserPassword(w http.ResponseWriter, r *http.Request) error {

	var req model.UpdateUserPasswordRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetUserPassword(context.Background(), req.AccountId)
	if err != nil {
		return errors.NewServerError(model.GetUserPasswordError, err.Error())
	}

	err = util.CheckPassword(req.OldPassword, result.Password)
	if err != nil {
		return errors.NewClientError().With(errors.ClientErrorData{
			Field:     "Password",
			ErrorType: "Password doesn't match our record",
		})
	}

	hashedPassword, _ := util.HashPassword(req.NewPassword)

	err = a.dbTrx.UpsertUserPassword(context.Background(), db.UpsertUserPasswordParams{
		UserID:   req.AccountId,
		Password: hashedPassword,
	})
	if err != nil {
		return errors.NewServerError(model.UpdateUserPasswordError, err.Error())
	}
	
	res := model.UpdateUserPasswordResponse{Message: "OK"}
	httpHandler.WriteResponse(w, res)

	return nil
}
