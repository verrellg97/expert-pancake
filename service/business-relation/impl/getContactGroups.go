package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/business-relation/model"
)

func (a businessRelationService) GetContactGroups(w http.ResponseWriter, r *http.Request) error {

	var req model.GetContactGroupsRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetContactGroups(context.Background(), req.CompanyId)
	if err != nil {
		return errors.NewServerError(model.GetContactGroupsError, err.Error())
	}

	var groups = make([]model.ContactGroup, 0)

	for _, d := range result {
		var group = model.ContactGroup{
			GroupId:   d.ID,
			CompanyId: d.CompanyID,
			Name:      d.Name,
		}
		groups = append(groups, group)
	}

	res := groups
	httpHandler.WriteResponse(w, res)

	return nil
}
