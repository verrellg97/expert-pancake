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

func (a salesService) GetSalesOrders(w http.ResponseWriter, r *http.Request) error {

	var req model.GetSalesOrdersRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetSalesOrders(context.Background(), db.GetSalesOrdersParams{
		CompanyID: req.CompanyId,
		BranchID:  req.BranchId,
		StartDate: util.StringToDate(req.StartDate),
		EndDate:   util.StringToDate(req.EndDate),
	})
	if err != nil {
		return errors.NewServerError(model.GetSalesOrdersError, err.Error())
	}

	var salesOrders = make([]model.SalesOrder, 0)

	for _, d := range result {
		argContactBook := client.GetContactBooksRequest{
			Id:        d.ContactBookID,
			CompanyId: d.CompanyID,
		}
		contactBook, err := client.GetContactBooks(argContactBook)
		if err != nil {
			return errors.NewServerError(model.GetSalesOrdersError, err.Error())
		}
		customerName := ""
		if len(contactBook.Result) > 0 {
			customerName = contactBook.Result[0].Name
		}

		var salesOrder = model.SalesOrder{
			TransactionId:      d.ID,
			CompanyId:          d.CompanyID,
			BranchId:           d.BranchID,
			FormNumber:         d.FormNumber,
			TransactionDate:    d.TransactionDate.Format(util.DateLayoutYMD),
			ContactBookId:      d.ContactBookID,
			SecondaryCompanyId: d.SecondaryCompanyID,
			CustomerName:       customerName,
			KonekinId:          d.KonekinID,
			CurrencyCode:       d.CurrencyCode,
			TotalItems:         strconv.FormatInt(d.TotalItems, 10),
			Total:              strconv.FormatInt(d.Total, 10),
			Status:             d.Status,
		}
		salesOrders = append(salesOrders, salesOrder)
	}

	res := model.GetSalesOrdersResponse{
		SalesOrders: salesOrders,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
