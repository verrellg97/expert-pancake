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

func (a purchasingService) GetReceiptOrders(w http.ResponseWriter, r *http.Request) error {

	var req model.GetReceiptOrdersRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetReceiptOrders(context.Background(), db.GetReceiptOrdersParams{
		CompanyID: req.CompanyId,
		BranchID:  req.BranchId,
		StartDate: util.StringToDate(req.StartDate),
		EndDate:   util.StringToDate(req.EndDate),
	})
	if err != nil {
		return errors.NewServerError(model.GetReceiptOrdersError, err.Error())
	}

	var receiptOrders = make([]model.ReceiptOrder, 0)

	for _, d := range result {
		argContactBook := client.GetContactBooksRequest{
			Id:        d.ContactBookID,
			CompanyId: d.CompanyID,
		}
		contactBook, err := client.GetContactBooks(argContactBook)
		if err != nil {
			return errors.NewServerError(model.GetReceiptOrdersError, err.Error())
		}

		var receiptOrder = model.ReceiptOrder{
			Id:                 d.ID,
			DeliveryOrderId:    d.DeliveryOrderID,
			CompanyId:          d.CompanyID,
			BranchId:           d.BranchID,
			FormNumber:         d.FormNumber,
			TransactionDate:    d.TransactionDate.Format(util.DateLayoutYMD),
			ContactBookId:      d.ContactBookID,
			SupplierName:       contactBook.Result[0].Name,
			SecondaryCompanyId: d.SecondaryCompanyID,
			KonekinId:          d.KonekinID,
			TotalItems:         strconv.FormatInt(d.TotalItems, 10),
			Status:             d.Status,
		}
		receiptOrders = append(receiptOrders, receiptOrder)
	}

	res := model.GetReceiptOrdersResponse{
		ReceiptOrders: receiptOrders,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
