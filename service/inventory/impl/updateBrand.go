package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/inventory/db/sqlc"
	"github.com/expert-pancake/service/inventory/model"
)

func (a inventoryService) UpdateBrand(w http.ResponseWriter, r *http.Request) error {

	var req model.UpdateBrandRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.UpdateBrandParams{
		ID:   req.Id,
		Name: req.Name,
	}

	result, err := a.dbTrx.UpdateBrand(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpdateBrandError, err.Error())
	}

	res := model.UpdateBrandResponse{
		Brand: model.Brand{
			BrandId:   result.ID,
			CompanyId: result.CompanyID,
			Name:      result.Name,
		},
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
