package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/sales/db/transaction"
	"github.com/expert-pancake/service/sales/model"
)

func (a salesService) UpsertPOS(w http.ResponseWriter, r *http.Request) error {

	var req model.UpsertPOSRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.UpsertPOSTrxParams{
		Id:                 req.Id,
		CompanyId:          req.CompanyId,
		BranchId:           req.BranchId,
		WarehouseId:        req.WarehouseId,
		FormNumber:         req.FormNumber,
		TransactionDate:    req.TransactionDate,
		ContactBookId:      req.ContactBookId,
		SecondaryCompanyId: req.SecondaryCompanyId,
		KonekinId:          req.KonekinId,
		CurrencyCode:       req.CurrencyCode,
		POSPaymentMethodId: req.POSPaymentMethodId,
		TotalItems:         req.TotalItems,
		Total:              req.Total,
		POSItems:           req.POSItems,
	}

	result, err := a.dbTrx.UpsertPOSTrx(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpsertPOSError, err.Error())
	}

	res := model.UpsertPOSResponse{
		Message: result.Message,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
