package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/inventory/db/sqlc"
	"github.com/expert-pancake/service/inventory/model"
	uuid "github.com/satori/go.uuid"
)

func (a inventoryService) AddItemGroup(w http.ResponseWriter, r *http.Request) error {

	var req model.AddItemGroupRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.InsertItemGroupParams{
		ID:        uuid.NewV4().String(),
		CompanyID: req.CompanyId,
		Name:      req.Name,
	}

	result, err := a.dbTrx.InsertItemGroup(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.AddNewItemGroupError, err.Error())
	}

	res := model.AddItemGroupResponse{
		Group: model.Group{
			ItemGroupId: result.ID,
			CompanyId:   result.CompanyID,
			Name:        result.Name,
		},
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
