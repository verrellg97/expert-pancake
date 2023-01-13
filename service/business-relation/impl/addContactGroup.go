package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/business-relation/db/transaction"
	"github.com/expert-pancake/service/business-relation/model"
)

func (a businessRelationService) AddContactGroup(w http.ResponseWriter, r *http.Request) error {

	var req model.AddContactGroupRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.CreateNewContactGroupTrxParams{
		CompanyId:   req.CompanyId,
		ImageUrl:    req.ImageUrl,
		Name:        req.Name,
		Description: req.Description,
		Members:     req.Members,
	}

	result, err := a.dbTrx.CreateNewContactGroupTrx(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.AddNewContactGroupError, err.Error())
	}

	res := model.AddContactGroupResponse{
		ContactGroup: model.ContactGroup{
			GroupId:     result.ContactGroupId,
			CompanyId:   result.CompanyId,
			ImageUrl:    result.ImageUrl,
			Name:        result.Name,
			Description: result.Description,
		},
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
