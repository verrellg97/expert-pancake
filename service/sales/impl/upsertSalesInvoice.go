package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/sales/db/transaction"
	"github.com/expert-pancake/service/sales/model"
)

func (a salesService) UpsertSalesInvoice(w http.ResponseWriter, r *http.Request) error {

	var req model.UpsertSalesInvoiceRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.UpsertSalesInvoiceTrxParams{
		Id:                 req.Id,
		SalesOrderId:       req.SalesOrderId,
		CompanyId:          req.CompanyId,
		BranchId:           req.BranchId,
		TransactionDate:    req.TransactionDate,
		ContactBookId:      req.ContactBookId,
		SecondaryCompanyId: req.SecondaryCompanyId,
		KonekinId:          req.KonekinId,
		CurrencyCode:       req.CurrencyCode,
		SalesInvoiceItems:  req.SalesInvoiceItems,
	}

	err := a.dbTrx.UpsertSalesInvoiceTrx(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpsertSalesInvoiceError, err.Error())
	}

	res := model.UpsertSalesInvoiceResponse{
		Message: `OK`,
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
