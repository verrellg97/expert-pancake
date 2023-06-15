package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/sales/db/sqlc"
	"github.com/expert-pancake/service/sales/model"
)

func (a salesService) UpdateSalesOrderStatus(w http.ResponseWriter, r *http.Request) error {

	var req model.UpdateSalesOrderStatusRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetSalesOrder(context.Background(), req.SalesOrderId)
	if err != nil {
		return errors.NewServerError(model.UpdateSalesOrderStatusError, err.Error())
	}

	var message = "OK"
	if !result.IsDeleted {
		if result.Status != req.Status && result.Status == "created" {
			arg := db.UpdateSalesOrderStatusParams{
				ID:       req.SalesOrderId,
				BranchID: req.BranchId,
				Status:   req.Status,
			}

			err := a.dbTrx.UpdateSalesOrderStatus(context.Background(), arg)
			if err != nil {
				return errors.NewServerError(model.UpdateSalesOrderStatusError, err.Error())
			}
		} else {
			message = "No data updated"
		}
	} else {
		message = "Transaction has been deleted"
	}

	res := model.UpdateSalesOrderStatusResponse{
		Message: message,
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
