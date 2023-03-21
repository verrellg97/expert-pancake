package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/inventory/model"
	"github.com/expert-pancake/service/inventory/util"
)

func (a inventoryService) GetPricelists(w http.ResponseWriter, r *http.Request) error {

	var req model.GetPricelistsRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetPricelists(context.Background(), req.CompanyId)
	if err != nil {
		return errors.NewServerError(model.GetPricelistsError, err.Error())
	}

	var pricelists = make([]model.Pricelist, 0)

	for _, d := range result {
		var end_date *string

		if d.EndDate.Valid {
			end_date = new(string)
			*end_date = d.EndDate.Time.Format(util.DateLayoutYMD)
		}

		var pricelist = model.Pricelist{
			Id:        d.ID,
			Name:      d.Name,
			StartDate: d.StartDate.Format(util.DateLayoutYMD),
			EndDate:   end_date,
			IsDefault: d.IsDefault,
		}
		pricelists = append(pricelists, pricelist)
	}

	res := model.GetPricelistsResponse{
		Pricelists: pricelists,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
