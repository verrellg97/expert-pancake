package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/inventory/db/sqlc"
	"github.com/expert-pancake/service/inventory/model"
)

func (a inventoryService) UpdateGroup(w http.ResponseWriter, r *http.Request) error {

	var req model.UpdateGroupRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.UpdateGroupParams{
		ID:   req.Id,
		Name: req.Name,
	}

	result, err := a.dbTrx.UpdateGroup(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpdateGroupError, err.Error())
	}

	res := model.UpdateGroupResponse{
		Group: model.Group{
			GroupId:   result.ID,
			CompanyId: result.CompanyID,
			Name:      result.Name,
		},
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
