package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/inventory/db/sqlc"
	"github.com/expert-pancake/service/inventory/model"
	"github.com/expert-pancake/service/inventory/util"
	uuid "github.com/satori/go.uuid"
)

func (a inventoryService) UpsertPricelist(w http.ResponseWriter, r *http.Request) error {

	var req model.UpsertPricelistRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	var id = ""
	if req.PricelistId == "" {
		id = uuid.NewV4().String()
	} else {
		id = req.PricelistId
	}

	arg := db.UpsertPricelistParams{
		ID:        id,
		CompanyID: req.CompanyId,
		Name:      req.Name,
		StartDate: util.StringToDate(req.StartDate),
		EndDate:   util.NewNullableDate(util.StringToDate(req.EndDate)),
		IsDefault: req.IsDefault,
	}

	result, err := a.dbTrx.UpsertPricelist(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpsertPricelistError, err.Error())
	}

	var end_date *string

	if result.EndDate.Valid {
		end_date = new(string)
		*end_date = result.EndDate.Time.Format(util.DateLayoutYMD)
	}

	res := model.UpsertPricelistResponse{
		Pricelist: model.Pricelist{
			Id:        result.ID,
			Name:      result.Name,
			StartDate: result.StartDate.Format(util.DateLayoutYMD),
			EndDate:   end_date,
			IsDefault: result.IsDefault,
		},
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
