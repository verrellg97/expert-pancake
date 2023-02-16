package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/inventory/db/sqlc"
	"github.com/expert-pancake/service/inventory/model"
)

func (a inventoryService) GetVariantWarehouseRackBatches(w http.ResponseWriter, r *http.Request) error {

	var req model.GetVariantWarehouseRackBatchesRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetVariantWarehouseRackBatches(context.Background(), db.GetVariantWarehouseRackBatchesParams{
		WarehouseRackID: req.WarehouseRackId,
		VariantID:       req.VariantId,
	})
	if err != nil {
		return errors.NewServerError(model.GetVariantWarehouseRackBatchesError, err.Error())
	}

	var datas = make([]*string, 0)

	for _, d := range result {
		var batch *string
		if d.Valid {
			batch = new(string)
			*batch = d.String
		}
		datas = append(datas, batch)
	}

	res := model.GetVariantWarehouseRackBatchesResponse{
		Batches: datas,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
