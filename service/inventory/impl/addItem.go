package impl

import (
	"context"
	"net/http"
	"strconv"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/inventory/db/transaction"
	"github.com/expert-pancake/service/inventory/model"
)

func (a inventoryService) AddItem(w http.ResponseWriter, r *http.Request) error {

	var req model.AddItemRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.AddItemTrxParams{
		CompanyId:   req.CompanyId,
		ImageUrl:    req.ImageUrl,
		Name:        req.Name,
		BrandId:     req.BrandId,
		GroupId:     req.GroupId,
		Tag:         req.Tag,
		Description: req.Description,
	}

	result, err := a.dbTrx.AddItemTrx(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.AddNewItemError, err.Error())
	}

	res := model.AddItemResponse{
		Item: model.Item{
			CompanyId:   result.CompanyId,
			ItemId:      result.ItemId,
			VariantId:   result.VariantId,
			ImageUrl:    result.ImageUrl,
			Code:        result.Code,
			Name:        result.Name,
			BrandId:     result.BrandId,
			BrandName:   result.BrandName,
			GroupId:     result.GroupId,
			GroupName:   result.GroupName,
			Tag:         result.Tag,
			Description: result.Description,
			IsDefault:   result.IsDefault,
			Price:       strconv.FormatInt(result.Price, 10),
			Stock:       strconv.FormatInt(result.Stock, 10),
		},
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
