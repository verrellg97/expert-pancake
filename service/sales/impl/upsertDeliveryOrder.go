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

func (a salesService) UpsertDeliveryOrder(w http.ResponseWriter, r *http.Request) error {

	var req model.UpsertDeliveryOrderRequest

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

	arg := db.UpsertDeliveryOrderTrxParams{
		Id:                 id,
		CompanyId:          req.CompanyId,
		BranchId:           req.BranchId,
		TransactionDate:    util.StringToDate(req.TransactionDate),
		ContactBookId:      req.ContactBookId,
		SecondaryCompanyId: req.SecondaryCompanyId,
		KonekinId:          req.KonekinId,
		SecondaryBranchId:  req.SecondaryBranchId,
	}

	result, err := a.dbTrx.UpsertDeliveryOrderTrx(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpsertDeliveryOrderError, err.Error())
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

	branchName := ""
	branches, err := client.GetCompanyBranches(
		client.GetCompanyBranchesRequest{
			CompanyId: result.SecondaryCompanyId,
		})
	if err != nil {
		return err
	}
	for _, d := range branches.Result {
		if d.BranchId == result.SecondaryBranchId {
			branchName = d.Name
			break
		}
	}

	res := model.UpsertDeliveryOrderResponse{
		DeliveryOrder: model.DeliveryOrder{
			TransactionId:       result.TransactionId,
			CompanyId:           result.CompanyId,
			BranchId:            result.BranchId,
			FormNumber:          result.FormNumber,
			TransactionDate:     result.TransactionDate,
			ContactBookId:       result.ContactBookId,
			SecondaryCompanyId:  result.SecondaryCompanyId,
			CustomerName:        customerName,
			KonekinId:           result.KonekinId,
			SecondaryBranchId:   result.SecondaryBranchId,
			SecondaryBranchName: branchName,
			TotalItems:          result.TotalItems,
			Status:              result.Status,
		},
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
