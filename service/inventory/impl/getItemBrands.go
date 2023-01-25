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

func (a inventoryService) GetItemBrands(w http.ResponseWriter, r *http.Request) error {

	var req model.GetItemBrandsRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetItemBrands(context.Background(), db.GetItemBrandsParams{
		CompanyID: req.CompanyId,
		Name:   util.WildCardString(req.Keyword),
	})
	if err != nil {
		return errors.NewServerError(model.GetItemBrandsError, err.Error())
	}

	var brands = make([]model.Brand, 0)

	for _, d := range result {
		var brand = model.Brand{
			ItemBrandId:   d.ID,
			CompanyId: d.CompanyID,
			Name:      d.Name,
		}
		brands = append(brands, brand)
	}


	res := brands
	httpHandler.WriteResponse(w, res)

	return nil
}
