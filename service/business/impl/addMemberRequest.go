package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/business/db/sqlc"
	"github.com/expert-pancake/service/business/model"
	uuid "github.com/satori/go.uuid"
)

func (a businessService) AddMemberRequest(w http.ResponseWriter, r *http.Request) error {

	var req model.AddMemberRequestRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.InsertMemberRequestParams{
		ID:        uuid.NewV4().String(),
		UserID:    req.UserId,
		CompanyID: req.CompanyId,
	}

	err := a.dbTrx.InsertMemberRequest(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.AddMemberRequestError, err.Error())
	}

	res := model.AddMemberRequestResponse{
		Message: "OK",
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
