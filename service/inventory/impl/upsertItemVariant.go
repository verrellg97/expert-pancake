package impl

import (
	"context"
	"net/http"
	"strconv"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/inventory/db/sqlc"
	"github.com/expert-pancake/service/inventory/model"
	uuid "github.com/satori/go.uuid"
)

func (a inventoryService) UpsertItemVariant(w http.ResponseWriter, r *http.Request) error {

	var req model.UpsertItemVariantRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	var id = ""
	if req.ItemVariantId == "" {
		id = uuid.NewV4().String()
	} else {
		id = req.ItemVariantId
	}

	price, _ := strconv.ParseInt(req.Price, 10, 64)
	stock, _ := strconv.ParseInt(req.Stock, 10, 64)
	arg := db.UpsertItemVariantParams{
		ID:       id,
		ItemID:   req.ItemId,
		ImageUrl: req.ImageUrl,
		Name:     req.Name,
		Price:    price,
		Stock:    stock,
	}

	err := a.dbTrx.UpsertItemVariant(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpsertItemVariantError, err.Error())
	}

	result, err := a.dbTrx.GetItemVariant(context.Background(), id)
	if err != nil {
		return errors.NewServerError(model.UpsertItemVariantError, err.Error())
	}

	res := model.UpsertItemVariantResponse{
		Item: model.Item{
			CompanyId:   result.CompanyID,
			ItemId:      result.ID,
			VariantId:   result.VariantID,
			ImageUrl:    result.ImageUrl,
			Code:        result.Code,
			Name:        result.Name,
			VariantName: result.VariantName,
			BrandId:     result.BrandID,
			BrandName:   result.BrandName,
			GroupId:     result.GroupID,
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
