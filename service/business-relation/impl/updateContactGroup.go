package impl

import (
	"context"
	"net/http"
	"strconv"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/business-relation/db/transaction"
	"github.com/expert-pancake/service/business-relation/model"
)

func (a businessRelationService) UpdateContactGroup(w http.ResponseWriter, r *http.Request) error {

	var req model.UpdateContactGroupRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.UpdateContactGroupTrxParams{
		Id:          req.GroupId,
		ImageUrl:    req.ImageUrl,
		Name:        req.Name,
		Description: req.Description,
		Members:     req.Members,
	}

	result, err := a.dbTrx.UpdateContactGroupTrx(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpdateContactGroupError, err.Error())
	}

	res := model.UpdateContactGroupResponse{
		ContactGroupWithMember: model.ContactGroupWithMember{
			GroupId:     result.ContactGroupId,
			CompanyId:   result.CompanyId,
			ImageUrl:    result.ImageUrl,
			Name:        result.Name,
			Description: result.Description,
			Member:      strconv.Itoa(int(len(result.Members))),
		},
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
