package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/inventory/db/sqlc"
	"github.com/expert-pancake/service/inventory/model"
	uuid "github.com/satori/go.uuid"
)

func (a inventoryService) AddItemBrand(w http.ResponseWriter, r *http.Request) error {

	var req model.AddItemBrandRequest


	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.InsertItemBrandParams{
		ID:        uuid.NewV4().String(),
		CompanyID: req.CompanyId,
		Name:      req.Name,
	}

	result, err := a.dbTrx.InsertItemBrand(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.AddNewItemBrandError, err.Error())
	}

	res := model.AddItemBrandResponse{
		Brand: model.Brand{
			ItemBrandId: result.ID,
			CompanyId:   result.CompanyID,
			Name:        result.Name,
		},
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
