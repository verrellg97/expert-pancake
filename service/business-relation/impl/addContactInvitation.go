package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/business-relation/db/transaction"
	"github.com/expert-pancake/service/business-relation/model"
)

func (a businessRelationService) AddContactInvitation(w http.ResponseWriter, r *http.Request) error {

	var req model.AddContactInvitationRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.AddContactInvitationTrxParams{
		PrimaryContactBookId: req.PrimaryContactBookId,
		PrimaryCompanyId:     req.PrimaryCompanyId,
		SecondaryCompanyId:   req.SecondaryCompanyId,
	}

	result, err := a.dbTrx.AddContactInvitationTrx(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.AddNewContactInvitationError, err.Error())
	}

	res := model.AddContactInvitationResponse{
		Message: result.Message,
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
