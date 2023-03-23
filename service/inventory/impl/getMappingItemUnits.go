package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/inventory/db/sqlc"
	"github.com/expert-pancake/service/inventory/model"
)

func (a inventoryService) GetMappingItemUnits(w http.ResponseWriter, r *http.Request) error {

	var req model.GetMappingItemUnitsRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetMappingItemUnits(context.Background(), db.GetMappingItemUnitsParams{
		ID:                 req.VariantId,
		PrimaryCompanyID:   req.PrimaryCompanyId,
		SecondaryCompanyID: req.SecondaryCompanyId,
		IsSupplier:         req.IsSupplier,
	})
	if err != nil {
		return errors.NewServerError(model.GetMappingItemUnitsError, err.Error())
	}

	var mappingItemUnits = make([]model.MappingItemUnit, 0)

	for _, d := range result {
		var mappingItemUnit = model.MappingItemUnit{
			ItemUnitId: d.ItemUnitID,
			UnitName:   d.UnitName,
		}
		mappingItemUnits = append(mappingItemUnits, mappingItemUnit)
	}

	res := model.GetMappingItemUnitsResponse{
		ItemUnits: mappingItemUnits,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
