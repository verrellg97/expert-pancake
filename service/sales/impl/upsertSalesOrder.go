package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/sales/db/transaction"
	"github.com/expert-pancake/service/sales/impl/client"
	"github.com/expert-pancake/service/sales/model"
	"github.com/expert-pancake/service/sales/util"
	uuid "github.com/satori/go.uuid"
)

func (a salesService) UpsertSalesOrder(w http.ResponseWriter, r *http.Request) error {

	var req model.UpsertSalesOrderRequest

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

	arg := db.UpsertSalesOrderTrxParams{
		Id:                    id,
		PurchaseOrderId:       req.PurchaseOrderId,
		PurchaseOrderBranchId: req.PurchaseOrderBranchId,
		CompanyId:             req.CompanyId,
		BranchId:              req.BranchId,
		TransactionDate:       util.StringToDate(req.TransactionDate),
		ContactBookId:         req.ContactBookId,
		SecondaryCompanyId:    req.SecondaryCompanyId,
		KonekinId:             req.KonekinId,
		CurrencyCode:          req.CurrencyCode,
		IsAllBranches:         req.IsAllBranches,
		Branches:              req.Branches,
	}

	result, err := a.dbTrx.UpsertSalesOrderTrx(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpsertSalesOrderError, err.Error())
	}

	argContactBook := client.GetContactBooksRequest{
		Id:        result.ContactBookId,
		CompanyId: result.CompanyId,
	}
	contactBook, err := client.GetContactBooks(argContactBook)
	if err != nil {
		return errors.NewServerError(model.UpsertSalesOrderError, err.Error())
	}
	customerName := ""
	if len(contactBook.Result) > 0 {
		customerName = contactBook.Result[0].Name
	}

	res := model.UpsertSalesOrderResponse{
		SalesOrder: model.SalesOrder{
			TransactionId:      result.TransactionId,
			CompanyId:          result.CompanyId,
			BranchId:           result.BranchId,
			FormNumber:         result.FormNumber,
			TransactionDate:    result.TransactionDate,
			ContactBookId:      result.ContactBookId,
			SecondaryCompanyId: result.SecondaryCompanyId,
			CustomerName:       customerName,
			KonekinId:          result.KonekinId,
			SecondaryBranchId:  result.PurchaseOrderBranchId,
			CurrencyCode:       result.CurrencyCode,
			TotalItems:         result.TotalItems,
			Total:              result.Total,
			Status:             result.Status,
		},
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
