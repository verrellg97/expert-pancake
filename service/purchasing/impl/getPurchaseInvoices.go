package impl

import (
	"context"
	"net/http"
	"strconv"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/purchasing/db/sqlc"
	"github.com/expert-pancake/service/purchasing/impl/client"
	"github.com/expert-pancake/service/purchasing/model"
	"github.com/expert-pancake/service/purchasing/util"
)

func (a purchasingService) GetPurchaseInvoices(w http.ResponseWriter, r *http.Request) error {

	var req model.GetPurchaseInvoicesRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetPurchaseInvoices(context.Background(), db.GetPurchaseInvoicesParams{
		CompanyID: req.CompanyId,
		BranchID:  req.BranchId,
		StartDate: util.StringToDate(req.StartDate),
		EndDate:   util.StringToDate(req.EndDate),
	})
	if err != nil {
		return errors.NewServerError(model.GetPurchaseInvoicesError, err.Error())
	}

	var purchaseInvoices = make([]model.PurchaseInvoice, 0)

	for _, d := range result {
		argContactBook := client.GetContactBooksRequest{
			Id:        d.ContactBookID,
			CompanyId: d.CompanyID,
		}
		contactBook, err := client.GetContactBooks(argContactBook)
		if err != nil {
			return errors.NewServerError(model.GetPurchaseInvoicesError, err.Error())
		}

		argGetWarehouse := client.GetWarehousesRequest{
			Id:       d.WarehouseID,
			BranchId: "1",
		}
		warehouse, err := client.GetWarehouses(argGetWarehouse)
		if err != nil {
			return errors.NewServerError(model.GetPurchaseInvoicesError, err.Error())
		}

		var purchaseInvoice = model.PurchaseInvoice{
			Id:                          d.ID,
			FormNumber:                  d.FormNumber,
			SalesInvoiceId:              d.SalesInvoiceID,
			ReceiptOrderId:              d.ReceiptOrderID,
			ReceiptOrderFormNumber:      d.ReceiptOrderFormNumber,
			ReceiptOrderTransactionDate: d.ReceiptOrderTransactionDate.Format(util.DateLayoutYMD),
			WarehouseId:                 d.WarehouseID,
			WarehouseName:               warehouse.Result.Warehouses[0].Name,
			CompanyId:                   d.CompanyID,
			BranchId:                    d.BranchID,
			TransactionDate:             d.TransactionDate.Format(util.DateLayoutYMD),
			ContactBookId:               d.ContactBookID,
			SecondaryCompanyId:          d.SecondaryCompanyID,
			KonekinId:                   d.KonekinID,
			SupplierName:                contactBook.Result[0].Name,
			CurrencyCode:                d.CurrencyCode,
			Total:                       strconv.FormatInt(d.Total, 10),
			TotalItems:                  strconv.FormatInt(d.TotalItems, 10),
			Status:                      d.Status,
		}
		purchaseInvoices = append(purchaseInvoices, purchaseInvoice)
	}

	res := model.GetPurchaseInvoicesResponse{
		PurchaseInvoices: purchaseInvoices,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
