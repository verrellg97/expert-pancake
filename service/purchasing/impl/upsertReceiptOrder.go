package impl

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/purchasing/db/sqlc"
	"github.com/expert-pancake/service/purchasing/model"
	"github.com/expert-pancake/service/purchasing/util"
	uuid "github.com/satori/go.uuid"
)

func (a purchasingService) UpsertReceiptOrder(w http.ResponseWriter, r *http.Request) error {

	var req model.UpsertReceiptOrderRequest
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
	
	totalItmes, _ := strconv.ParseInt(req.TotalItems, 10, 64)
	
	arg := db.UpsertReceiptOrderParams{
		ID:                 id,
		DeliveryOrderID:    req.DeliveryOrderId,
		FormNumber:         "RCO-" + fmt.Sprintf("%08d", rand.Intn(100000000)),
		CompanyID:          req.CompanyId,
		BranchID:           req.BranchId,
		TransactionDate:    util.StringToDate(req.TransactionDate),
		ContactBookID:      req.ContactBookId,
		SecondaryCompanyID: req.SecondaryCompanyId,
		KonekinID:          req.KonekinId,
		TotalItems:         totalItmes,
	}
	
	err := a.dbTrx.UpsertReceiptOrder(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpsertReceiptOrderError, err.Error())
	}

	res := model.UpsertReceiptOrderResponse{
		Message: `OK`,
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
