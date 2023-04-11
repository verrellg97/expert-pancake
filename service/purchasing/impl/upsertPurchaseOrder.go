package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/purchasing/db/transaction"
	"github.com/expert-pancake/service/purchasing/model"
	"github.com/expert-pancake/service/purchasing/util"
	uuid "github.com/satori/go.uuid"
)

func (a purchasingService) UpsertPurchaseOrder(w http.ResponseWriter, r *http.Request) error {

	var req model.UpsertPurchaseOrderRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	var id = ""
	if req.Id == "" {
		id = uuid.NewV4().String()
	} else {
		id = req.Id
	}

	arg := db.UpsertPurchaseOrderTrxParams{
		Id:              id,
		CompanyId:       req.CompanyId,
		BranchId:        req.BranchId,
		TransactionDate: util.StringToDate(req.TransactionDate),
		ContactBookId:   req.ContactBookId,
		KonekinId:       req.KonekinId,
		CurrencyCode:    req.CurrencyCode,
	}

	result, err := a.dbTrx.UpsertPurchaseOrderTrx(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpsertPurchaseOrderError, err.Error())
	}

	res := model.UpsertPurchaseOrderResponse{
		PurchaseOrder: model.PurchaseOrder{
			TransactionId:   result.TransactionId,
			CompanyId:       result.CompanyId,
			BranchId:        result.BranchId,
			FormNumber:      result.FormNumber,
			TransactionDate: result.TransactionDate,
			ContactBookId:   result.ContactBookId,
			KonekinId:       result.KonekinId,
			CurrencyCode:    result.CurrencyCode,
			TotalItems:      result.TotalItems,
			Total:           result.Total,
			Status:          result.Status,
		},
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
