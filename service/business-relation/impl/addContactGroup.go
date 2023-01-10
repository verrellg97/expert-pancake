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

func (a businessRelationService) AddContactGroup(w http.ResponseWriter, r *http.Request) error {

	var req model.AddContactGroupRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.InsertContactGroupParams{
		ID:          uuid.NewV4().String(),
		CompanyID:   req.CompanyId,
		ImageUrl:    req.ImageUrl,
		Name:        req.Name,
		Description: req.Description,
	}

	result, err := a.dbTrx.InsertContactGroup(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.AddNewContactGroupError, err.Error())
	}

	res := model.AddContactGroupResponse{
		ContactGroup: model.ContactGroup{
			GroupId:     result.ID,
			CompanyId:   result.CompanyID,
			ImageUrl:    result.ImageUrl,
			Name:        result.Name,
			Description: result.Description,
		},
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
