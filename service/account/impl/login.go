package impl

import (
	"context"
	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/account/model"
	"github.com/expert-pancake/service/account/util"
	"net/http"
)

func (a accountService) Login(w http.ResponseWriter, r *http.Request) error {
	var req model.LoginRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetUserPassword(context.Background(), req.AccountId)
	if err != nil {
		return errors.NewServerError(model.GetUserPasswordError, err.Error())
	}

	err = util.CheckPassword(req.Password, result.Password)
	if err != nil {
		return errors.NewClientError().With(errors.ClientErrorData{
			Field:     "Password",
			ErrorType: "Password doesn't match our record",
		})
	}

	account, err := a.dbTrx.GetUser(context.Background(), req.AccountId)
	if err != nil {
		return errors.NewServerError(model.GetUserError, err.Error())
	}

	accountAddress, _ := a.dbTrx.GetUserAddress(context.Background(), req.AccountId)

	accessToken, accessPayload, err := a.tokenMaker.CreateToken(
		result.UserID,
		a.config.Token.AccessTokenDuration,
	)
	if err != nil {
		return errors.NewServerError(model.CreateAccessTokenError, err.Error())
	}

	refreshToken, refreshPayload, err := a.tokenMaker.CreateToken(
		result.UserID,
		a.config.Token.AccessTokenDuration,
	)
	if err != nil {
		return errors.NewServerError(model.CreateRefreshTokenError, err.Error())
	}

	res := model.LoginResponse{
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
		User: model.LoginUserResponse{
			User: model.User{
				Id:          account.ID,
				FullName:    account.Fullname,
				Nickname:    account.Nickname,
				Email:       account.Email.String,
				PhoneNumber: account.PhoneNumber,
			},
			Location: model.Location{
				Province:    accountAddress.Province,
				Regency:     accountAddress.Regency,
				District:    accountAddress.District,
				FullAddress: accountAddress.FullAddress,
			},
		},
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
