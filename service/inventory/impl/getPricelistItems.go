package impl

import (
	"context"
	"net/http"
	"strconv"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/inventory/db/sqlc"
	"github.com/expert-pancake/service/inventory/model"
	"github.com/expert-pancake/service/inventory/util"
)

func (a inventoryService) GetPricelistItems(w http.ResponseWriter, r *http.Request) error {

	var req model.GetPricelistItemsRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetPricelistItems(context.Background(), db.GetPricelistItemsParams{
		CompanyID:   req.CompanyId,
		PricelistID: req.PricelistId,
		Name:        util.WildCardString(req.Keyword),
	})
	if err != nil {
		return errors.NewServerError(model.GetPricelistItemsError, err.Error())
	}

	var pricelistItems = make([]model.PricelistItem, 0)

	for _, d := range result {
		var pricelistItem = model.PricelistItem{
			ItemId:      d.ItemID,
			ItemName:    d.ItemName,
			VariantId:   d.VariantID,
			VariantName: d.VariantName,
			ItemUnitId:  d.ItemUnitID,
			UnitName:    d.UnitName,
			Price:       strconv.FormatInt(d.Price, 10),
		}
		pricelistItems = append(pricelistItems, pricelistItem)
	}

	res := model.GetPricelistItemsResponse{
		PricelistItems: pricelistItems,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
