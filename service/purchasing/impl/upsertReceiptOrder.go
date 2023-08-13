package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/purchasing/db/transaction"
	"github.com/expert-pancake/service/purchasing/model"
)

func (a purchasingService) UpsertReceiptOrder(w http.ResponseWriter, r *http.Request) error {

	var req model.UpsertReceiptOrderRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.UpsertReceiptOrderTrxParams{
		Id:                 req.Id,
		DeliveryOrderId:    req.DeliveryOrderId,
		WarehouseId:        req.WarehouseId,
		CompanyId:          req.CompanyId,
		BranchId:           req.BranchId,
		TransactionDate:    req.TransactionDate,
		ContactBookId:      req.ContactBookId,
		SecondaryCompanyId: req.SecondaryCompanyId,
		KonekinId:          req.KonekinId,
		TotalItems:         req.TotalItems,
		ReceiptOrderItems:  req.UpsertReceiptOrderItemsRequest,
	}

	result, err := a.dbTrx.UpsertReceiptOrderTrx(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpsertReceiptOrderError, err.Error())
	}

	res := model.UpsertReceiptOrderResponse{
		ReceiptOrder: model.ReceiptOrder{
			Id:                 result.TransactionId,
			DeliveryOrderId:    result.DeliveryOrderId,
			WarehouseId:        result.WarehouseId,
			CompanyId:          result.CompanyId,
			BranchId:           result.BranchId,
			FormNumber:         result.FormNumber,
			TransactionDate:    result.TransactionDate,
			ContactBookId:      result.ContactBookId,
			SecondaryCompanyId: result.SecondaryCompanyId,
			KonekinId:          result.KonekinId,
			TotalItems:         result.TotalItems,
			Status:             result.Status,
		},
		ReceiptOrderItems: result.ReceiptOrderItems,
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
