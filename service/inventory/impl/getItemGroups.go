package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/inventory/db/sqlc"
	"github.com/expert-pancake/service/inventory/model"
	"github.com/expert-pancake/service/inventory/util"
)

func (a inventoryService) GetItemGroups(w http.ResponseWriter, r *http.Request) error {

	var req model.GetItemGroupsRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetItemGroups(context.Background(), db.GetItemGroupsParams{
		CompanyID: req.CompanyId,
		Name:   util.WildCardString(req.Keyword),
	})
	if err != nil {
		return errors.NewServerError(model.GetItemGroupsError, err.Error())
	}

	var groups = make([]model.Group, 0)

	for _, d := range result {
		var group = model.Group{
			ItemGroupId:   d.ID,
			CompanyId: d.CompanyID,
			Name:      d.Name,
		}
		groups = append(groups, group)
	}


	res := groups
	httpHandler.WriteResponse(w, res)

	return nil
}
