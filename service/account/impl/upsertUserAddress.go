package impl

import (
	"context"
	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/account/db/sqlc"
	"github.com/expert-pancake/service/account/model"
	"net/http"
)

func (a accountService) UpsertUserAddress(w http.ResponseWriter, r *http.Request) error {

	var req model.UpsertUserAddressRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.UpsertUserAddressesParams{
		UserID:      req.AccountId,
		Country:     "INDONESIA",
		Province:    req.Province,
		Regency:     req.Regency,
		District:    req.District,
		FullAddress: req.FullAddress,
	}

	_, err := a.dbTrx.UpsertUserAddresses(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpsertUserAddressError, err.Error())
	}

	res := model.UpsertUserAddressResponse{Message: "OK"}
	httpHandler.WriteResponse(w, res)

	return nil
}
