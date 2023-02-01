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

func (a inventoryService) UpdateItem(w http.ResponseWriter, r *http.Request) error {

	var req model.UpdateItemRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.UpdateItemTrxParams{
		ItemId:      req.ItemId,
		ImageUrl:    req.ImageUrl,
		Name:        req.Name,
		BrandId:     req.BrandId,
		GroupId:     req.GroupId,
		Tag:         req.Tag,
		Description: req.Description,
	}

	result, err := a.dbTrx.UpdateItemTrx(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpdateItemError, err.Error())
	}

	res := model.UpdateItemResponse{
		Item: model.Item{
			CompanyId:   result.CompanyId,
			ItemId:      result.ItemId,
			VariantId:   result.VariantId,
			ImageUrl:    result.ImageUrl,
			Code:        result.Code,
			Name:        result.Name,
			VariantName: result.VariantName,
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
