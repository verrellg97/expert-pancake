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

func (a inventoryService) GetSupplierCatalogs(w http.ResponseWriter, r *http.Request) error {

	var req model.GetSupplierCatalogsRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetSupplierCatalogs(context.Background(), db.GetSupplierCatalogsParams{
		CompanyID:          req.PrimaryCompanyId,
		SecondaryCompanyID: req.SecondaryCompanyId,
		Keyword:            util.WildCardString(req.Keyword),
	})
	if err != nil {
		return errors.NewServerError(model.GetSupplierCatalogsError, err.Error())
	}

	var items = make([]model.SupplierCatalog, 0)

	for _, d := range result {
		var item = model.SupplierCatalog{
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
			IsMapped:    d.IsMapped,
		}
		items = append(items, item)
	}

	res := model.GetSupplierCatalogsResponse{
		SupplierCatalogs: items,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
