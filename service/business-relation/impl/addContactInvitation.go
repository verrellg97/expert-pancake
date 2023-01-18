package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/business-relation/db/sqlc"
	"github.com/expert-pancake/service/business-relation/model"
	uuid "github.com/satori/go.uuid"
)

func (a businessRelationService) AddContactInvitation(w http.ResponseWriter, r *http.Request) error {

	var req model.AddContactInvitationRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.InsertContactInvitationParams{
		ID:                     uuid.NewV4().String(),
		PrimaryContactBookID:   req.PrimaryContactBookId,
		SecondaryContactBookID: req.SecondaryContactBookId,
		PrimaryCompanyID:       req.PrimaryCompanyId,
		SecondaryCompanyID:     req.SecondaryCompanyId,
	}

	_, err := a.dbTrx.InsertContactInvitation(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.AddNewContactInvitationError, err.Error())
	}

	res := model.AddContactInvitationResponse{
		Message: "OK",
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
