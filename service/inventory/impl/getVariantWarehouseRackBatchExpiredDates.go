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

func (a inventoryService) GetVariantWarehouseRackBatchExpiredDates(w http.ResponseWriter, r *http.Request) error {

	var req model.GetVariantWarehouseRackBatchExpiredDatesRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetVariantWarehouseRackBatchExpiredDates(context.Background(), db.GetVariantWarehouseRackBatchExpiredDatesParams{
		WarehouseRackID: req.WarehouseRackId,
		VariantID:       req.VariantId,
		IsNullBatch:     !util.NewNullableString(req.Batch).Valid,
		Batch:           util.NewNullableString(req.Batch),
	})
	if err != nil {
		return errors.NewServerError(model.GetVariantWarehouseRackBatchExpiredDatesError, err.Error())
	}

	var datas = make([]*string, 0)

	for _, d := range result {
		var expired_date *string
		if d.Valid {
			expired_date = new(string)
			*expired_date = d.Time.Format(util.DateLayoutYMD)
		}
		datas = append(datas, expired_date)
	}

	res := model.GetVariantWarehouseRackBatchExpiredDatesResponse{
		ExpiredDates: datas,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
