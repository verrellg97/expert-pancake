package impl

import (
	"context"
	"net/http"
	"strconv"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/sales/db/sqlc"
	"github.com/expert-pancake/service/sales/impl/client"
	"github.com/expert-pancake/service/sales/model"
	"github.com/expert-pancake/service/sales/util"
)

func (a salesService) GetSalesInvoices(w http.ResponseWriter, r *http.Request) error {

	var req model.GetSalesInvoicesRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetSalesInvoices(context.Background(), db.GetSalesInvoicesParams{
		CompanyID: req.CompanyId,
		BranchID:  req.BranchId,
		StartDate: util.StringToDate(req.StartDate),
		EndDate:   util.StringToDate(req.EndDate),
	})
	if err != nil {
		return errors.NewServerError(model.GetSalesInvoicesError, err.Error())
	}

	var salesInvoices = make([]model.SalesInvoice, 0)

	for _, d := range result {
		argContactBook := client.GetContactBooksRequest{
			Id:        d.ContactBookID,
			CompanyId: d.CompanyID,
		}
		contactBook, err := client.GetContactBooks(argContactBook)
		if err != nil {
			return errors.NewServerError(model.GetSalesInvoicesError, err.Error())
		}

		var salesInvoice = model.SalesInvoice{
			Id:                   d.ID,
			FormNumber:           d.FormNumber,
			SalesOrderId:         d.SalesOrderID,
			SalesOrderFormNumber: d.SalesOrderFormNumber,
			CompanyId:            d.CompanyID,
			BranchId:             d.BranchID,
			TransactionDate:      d.TransactionDate.Format(util.DateLayoutYMD),
			ContactBookId:        d.ContactBookID,
			SecondaryCompanyId:   d.SecondaryCompanyID,
			KonekinId:            d.KonekinID,
			CustomerName:         contactBook.Result[0].Name,
			CurrencyCode:         d.CurrencyCode,
			Total:                strconv.FormatInt(d.Total, 10),
			TotalItems:           strconv.FormatInt(d.TotalItems, 10),
			Status:               d.Status,
		}
		salesInvoices = append(salesInvoices, salesInvoice)
	}

	res := model.GetSalesInvoicesResponse{
		SalesInvoices: salesInvoices,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
