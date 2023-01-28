package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/warehouse/db/sqlc"
	"github.com/expert-pancake/service/warehouse/model"
	"github.com/expert-pancake/service/warehouse/util"
)

func (a warehouseService) GetRacks(w http.ResponseWriter, r *http.Request) error {

	var req model.GetRacksRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}
	
	var is_get_all = false
	if req.Mode == "all" {
		is_get_all = true
	} 
	
	arg := db.GetRacksParams{
		BranchID: req.BranchId,
		Name:     util.WildCardString(req.Keyword),
		IsGetAll: is_get_all,
	}
	result, err := a.dbTrx.GetRacks(context.Background(), arg)
	
	if err != nil {
		return errors.NewServerError(model.GetRacksError, err.Error())
	}

	var racks = make([]model.Rack, 0)

	for _, d := range result {
		var rack = model.Rack{
			RackId:   d.ID,
			BranchId: d.BranchID,
			Name:      d.Name,
		}
		racks = append(racks, rack)
	}

	res := racks
	httpHandler.WriteResponse(w, res)

	return nil
}
