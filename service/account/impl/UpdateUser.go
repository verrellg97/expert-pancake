package impl

import (
	"context"
	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/account/db/transaction"
	"github.com/expert-pancake/service/account/model"
	"net/http"
)

func (a accountService) UpdateUser(w http.ResponseWriter, r *http.Request) error {

	var req model.UpdateUserRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	/*if req.NewPassword != "" && req.OldPassword != "" {
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
	}*/

	result, err := a.dbTrx.UpdateUserTrx(context.Background(), db.UpdateUserTrxParams{
		AccountId:   req.AccountId,
		FullName:    req.FullName,
		Nickname:    req.Nickname,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Location: model.Location{
			Province:    req.Location.Province,
			Regency:     req.Location.Regency,
			District:    req.Location.District,
			FullAddress: req.Location.FullAddress,
		},
	})
	if err != nil {
		return errors.NewServerError(model.UpdateUserError, err.Error())
	}

	res := model.UpdateUserResponse{
		AccountId:   result.AccountId,
		FullName:    result.FullName,
		Nickname:    result.Nickname,
		Email:       result.Email,
		PhoneNumber: result.PhoneNumber,
		Location:    result.Location,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
