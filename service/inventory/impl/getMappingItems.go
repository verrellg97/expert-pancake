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

func (a inventoryService) GetMappingItems(w http.ResponseWriter, r *http.Request) error {

	var req model.GetMappingItemsRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetMappingItems(context.Background(), db.GetMappingItemsParams{
		ItemID:             req.ItemId,
		PrimaryCompanyID:   req.PrimaryCompanyId,
		SecondaryCompanyID: req.SecondaryCompanyId,
		Name:               util.WildCardString(req.Keyword),
	})
	if err != nil {
		return errors.NewServerError(model.GetMappingItemsError, err.Error())
	}

	var mappingItems = make([]model.MappingItem, 0)

	for _, d := range result {
		var mappingItem = model.MappingItem{
			ItemId: d.ID,
			Code:   d.Code,
			Name:   d.Name,
		}
		mappingItems = append(mappingItems, mappingItem)
	}

	res := model.GetMappingItemsResponse{
		Items: mappingItems,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
