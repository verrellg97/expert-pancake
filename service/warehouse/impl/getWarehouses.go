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

func (a warehouseService) GetWarehouses(w http.ResponseWriter, r *http.Request) error {

	var req model.GetWarehousesRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.GetWarehousesParams{
		BranchID: req.BranchId,
		Name:     util.WildCardString(req.Keyword),
		Type:     util.WildCardString(req.Type),
	}
	result, err := a.dbTrx.GetWarehouses(context.Background(), arg)

	if err != nil {
		return errors.NewServerError(model.GetWarehousesError, err.Error())
	}

	var warehouses = make([]model.Warehouse, 0)

	for _, d := range result {
		var warehouse = model.Warehouse{
			WarehouseId: d.ID,
			BranchId:    d.BranchID,
			Code:        d.Code,
			Name:        d.Name,
			Type:        d.Type,
			Address:     d.Address,
		}
		warehouses = append(warehouses, warehouse)
	}

	res := warehouses
	httpHandler.WriteResponse(w, res)

	return nil
}
