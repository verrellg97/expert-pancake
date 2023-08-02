package impl

import (
	"context"
	"net/http"
	"strconv"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/inventory/db/transaction"
	"github.com/expert-pancake/service/inventory/model"
	"github.com/expert-pancake/service/inventory/util"
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
		Code:        req.Code,
		Name:        req.Name,
		Barcode:     req.Barcode,
		BrandId:     req.BrandId,
		GroupIds:    req.GroupIds,
		Tag:         util.ArrayToString(req.Tag),
		Price:       req.Price,
		Description: req.Description,
	}

	result, err := a.dbTrx.AddItemTrx(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.AddNewItemError, err.Error())
	}

	var amount_in_stock *string
	amount_in_stock = new (string)
	*amount_in_stock = strconv.FormatInt(0, 10)
	
	res := model.AddItemResponse{
		Item: model.Item{
			CompanyId:   result.CompanyId,
			ItemId:      result.ItemId,
			VariantId:   result.VariantId,
			ImageUrl:    result.ImageUrl,
			Code:        result.Code,
			Barcode:     result.Barcode,
			Name:        result.Name,
			VariantName: result.VariantName,
			BrandId:     result.BrandId,
			BrandName:   result.BrandName,
			Groups:      util.StringToArrayOfGroup(result.Groups, result.CompanyId),
			Tag:         util.StringToArray(result.Tag),
			Description: result.Description,
			IsDefault:   result.IsDefault,
			Price:       strconv.FormatInt(result.Price, 10),
			AmountInStock: amount_in_stock,
		},
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
