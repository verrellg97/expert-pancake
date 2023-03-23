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

func (a inventoryService) GetMappingItemVariants(w http.ResponseWriter, r *http.Request) error {

	var req model.GetMappingItemVariantsRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetMappingItemVariants(context.Background(), db.GetMappingItemVariantsParams{
		ItemID:             req.ItemId,
		PrimaryCompanyID:   req.PrimaryCompanyId,
		SecondaryCompanyID: req.SecondaryCompanyId,
		Name:               util.WildCardString(req.Keyword),
	})
	if err != nil {
		return errors.NewServerError(model.GetMappingItemVariantsError, err.Error())
	}

	var mappingItemVariants = make([]model.MappingItemVariant, 0)

	for _, d := range result {
		var mappingItemVariant = model.MappingItemVariant{
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
			IsDefault:   d.IsDefault,
			Price:       strconv.FormatInt(d.Price, 10),
		}
		mappingItemVariants = append(mappingItemVariants, mappingItemVariant)
	}

	res := model.GetMappingItemVariantsResponse{
		ItemVariants: mappingItemVariants,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
