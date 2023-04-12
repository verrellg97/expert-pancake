package impl

import (
	"context"
	"net/http"
	"strconv"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/purchasing/db/sqlc"
	"github.com/expert-pancake/service/purchasing/model"
	"github.com/expert-pancake/service/purchasing/util"
	"github.com/expert-pancake/service/purchasing/impl/client"
)

func (a purchasingService) GetPurchaseOrders(w http.ResponseWriter, r *http.Request) error {

	var req model.GetPurchaseOrdersRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetPurchaseOrders(context.Background(), db.GetPurchaseOrdersParams{
		CompanyID: req.CompanyId,
		BranchID:  req.BranchId,
		StartDate: util.StringToDate(req.StartDate),
		EndDate:   util.StringToDate(req.EndDate),
	})
	if err != nil {
		return errors.NewServerError(model.GetPurchaseOrdersError, err.Error())
	}

	var purchaseOrders = make([]model.PurchaseOrder, 0)

	for _, d := range result {
		argContactBook := client.GetContactBooksRequest{
			Id:        d.ContactBookID,
			CompanyId: d.CompanyID,
		}
		contactBook, err := client.GetContactBooks(argContactBook)
		if err != nil {
			return errors.NewServerError(model.UpsertPurchaseOrderError, err.Error())
		}

		var purchaseOrder = model.PurchaseOrder{
			TransactionId:      d.ID,
			CompanyId:          d.CompanyID,
			BranchId:           d.BranchID,
			FormNumber:         d.FormNumber,
			TransactionDate:    d.TransactionDate.Format(util.DateLayoutYMD),
			ContactBookId:      d.ContactBookID,
			SecondaryCompanyId: d.SecondaryCompanyID,
			SupplierName:       contactBook.Result[0].Name,
			KonekinId:          d.KonekinID,
			CurrencyCode:       d.CurrencyCode,
			TotalItems:         strconv.FormatInt(d.TotalItems, 10),
			Total:              strconv.FormatInt(d.Total, 10),
			Status:             d.Status,
		}
		purchaseOrders = append(purchaseOrders, purchaseOrder)
	}

	res := model.GetPurchaseOrdersResponse{
		PurchaseOrders: purchaseOrders,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
