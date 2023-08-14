package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/sales/db/transaction"
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
		WarehouseId:        req.WarehouseId,
		TransactionDate:    util.StringToDate(req.TransactionDate),
		ContactBookId:      req.ContactBookId,
		SecondaryCompanyId: req.SecondaryCompanyId,
		KonekinId:          req.KonekinId,
		SalesOrderId:       req.SalesOrderId,
		Items:              req.Items,
	}

	result, err := a.dbTrx.UpsertDeliveryOrderTrx(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpsertDeliveryOrderError, err.Error())
	}

	res := model.UpsertDeliveryOrderResponse{
		Message: result.Message,
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
