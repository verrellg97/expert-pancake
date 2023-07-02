package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/purchasing/db/transaction"
	"github.com/expert-pancake/service/purchasing/model"
)

func (a purchasingService) UpsertPurchaseInvoice(w http.ResponseWriter, r *http.Request) error {

	var req model.UpsertPurchaseInvoiceRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}
	
	arg := db.UpsertPurchaseInvoiceTrxParams{
			Id: req.Id,
			SalesInvoiceId: req.SalesInvoiceId,
			ReceiptOrderId: req.ReceiptOrderId,
			CompanyId: req.CompanyId,
			BranchId: req.BranchId,
			TransactionDate: req.TransactionDate,
			ContactBookId: req.ContactBookId,
			SecondaryCompanyId: req.SecondaryCompanyId,
			KonekinId: req.KonekinId,
			CurrencyCode: req.CurrencyCode,
			TotalItems: req.TotalItems,
			Total: req.Total,
			PurchaseInvoiceItems: req.PurchaseInvoiceItems,
	}
	
	err := a.dbTrx.UpsertPurchaseInvoiceTrx(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpsertPurchaseInvoiceError, err.Error())
	}

	res := model.UpsertPurchaseInvoiceResponse{
		Message: `OK`,
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
