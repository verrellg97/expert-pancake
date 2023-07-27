package impl

import (
	"context"
	"net/http"
	"strconv"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/inventory/db/sqlc"
	"github.com/expert-pancake/service/inventory/model"
	"github.com/expert-pancake/service/inventory/util"
	uuid "github.com/satori/go.uuid"
)

func (a inventoryService) InsertStockMovement(w http.ResponseWriter, r *http.Request) error {

	var req model.InsertStockMovementRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	amount, _ := strconv.ParseInt(req.Amount, 10, 64)
	arg := db.InsertStockMovementParams{
		ID:                   uuid.NewV4().String(),
		TransactionID:        req.TransactionId,
		TransactionCode:      req.TransactionCode,
		TransactionDate:      util.StringToDate(req.TransactionDate),
		TransactionReference: req.TransactionReference,
		DetailTransactionID:  req.DetailTransactionId,
		WarehouseID:          req.WarehouseId,
		WarehouseRackID:      req.WarehouseRackId,
		VariantID:            req.VariantId,
		ItemBarcodeID:        req.ItemBarcodeId,
		Amount:               amount,
	}

	a.dbTrx.InsertStockMovement(context.Background(), arg)
	res := model.InsertStockMovementResponse{
		Message: `OK`,
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
