package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/business-relation/db/transaction"
	"github.com/expert-pancake/service/business-relation/model"
)

func (a businessRelationService) UpdateContactInvitation(w http.ResponseWriter, r *http.Request) error {

	var req model.UpdateContactInvitationRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.UpdateContactInvitationTrxParams{
		Id:                     req.InvitationId,
		SecondaryContactBookId: req.SecondaryContactBookId,
		Status:                 req.Status,
	}

	result, err := a.dbTrx.UpdateContactInvitationTrx(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpdateContactInvitationError, err.Error())
	}

	res := model.UpdateContactInvitationResponse{
		Message: result.Message,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
