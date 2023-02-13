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

func (a inventoryService) GetItems(w http.ResponseWriter, r *http.Request) error {

	var req model.GetItemsRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetItems(context.Background(), db.GetItemsParams{
		CompanyID: req.CompanyId,
		Name:      util.WildCardString(req.Keyword),
	})
	if err != nil {
		return errors.NewServerError(model.GetItemsError, err.Error())
	}

	var items = make([]model.Item, 0)

	for _, d := range result {
		var item = model.Item{
			CompanyId:   d.CompanyID,
			ItemId:      d.ID,
			VariantId:   d.VariantID,
			ImageUrl:    d.ImageUrl,
			Code:        d.Code,
			Barcode:     d.Barcode,
			Name:        d.Name,
			VariantName: d.VariantName,
			BrandId:     d.BrandID,
			BrandName:   d.BrandName,
			GroupId:     d.GroupID,
			GroupName:   d.GroupName,
			Tag:         d.Tag,
			Description: d.Description,
			IsDefault:   d.IsDefault,
			Price:       strconv.FormatInt(d.Price, 10),
		}
		items = append(items, item)
	}

	res := model.GetItemsResponse{
		Items: items,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
