package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/business-relation/db/sqlc"
	"github.com/expert-pancake/service/business-relation/model"
)

func (a businessRelationService) UpdateContactGroup(w http.ResponseWriter, r *http.Request) error {

	var req model.UpdateContactGroupRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.UpdateContactGroup(context.Background(), db.UpdateContactGroupParams{
		ID:   req.GroupId,
		Name: req.Name,
	})
	if err != nil {
		return errors.NewServerError(model.UpdateContactGroupError, err.Error())
	}

	res := model.UpdateContactGroupResponse{
		ContactGroup: model.ContactGroup{
			GroupId:   result.ID,
			CompanyId: result.CompanyID,
			Name:      result.Name,
		},
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
