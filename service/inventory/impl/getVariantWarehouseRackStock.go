package impl

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/inventory/db/sqlc"
	"github.com/expert-pancake/service/inventory/model"
	"github.com/expert-pancake/service/inventory/util"
)

func (a inventoryService) GetVariantWarehouseRackStock(w http.ResponseWriter, r *http.Request) error {

	var req model.GetVariantWarehouseRackStockRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetVariantWarehouseRackStock(context.Background(), db.GetVariantWarehouseRackStockParams{
		WarehouseRackID:   req.WarehouseRackId,
		VariantID:         req.VariantId,
		IsNullBatch:       !util.NewNullableString(req.Batch).Valid,
		Batch:             util.NewNullableString(req.Batch),
		IsNullExpiredDate: !util.NewNullableDate(util.StringToDate(req.ExpiredDate)).Valid,
		ExpiredDate:       util.NewNullableDate(util.StringToDate(req.ExpiredDate)),
	})
	log.Println(err)
	if err != nil {
		return errors.NewServerError(model.GetVariantWarehouseRackStockError, err.Error())
	}

	res := model.GetVariantWarehouseRackStockResponse{
		ItemBarcodeId: result.ID,
		Stock:         strconv.FormatInt(result.Stock, 10),
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
