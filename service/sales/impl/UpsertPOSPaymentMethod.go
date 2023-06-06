package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/sales/db/sqlc"
	"github.com/expert-pancake/service/sales/model"
	uuid "github.com/satori/go.uuid"
)

func (a salesService) UpsertPOSPaymentMethod(w http.ResponseWriter, r *http.Request) error {

	var req model.UpsertPOSPaymentMethodRequest

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

	arg := db.UpsertPOSPaymentMethodParams{
		ID:               id,
		CompanyID:        req.CompanyId,
		ChartOfAccountID: req.ChartOfAccountId,
		Name:             req.Name,
	}

	err := a.dbTrx.UpsertPOSPaymentMethod(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpsertPOSPaymentMethodError, err.Error())
	}

	res := model.UpsertPOSPaymentMethodResponse{
		Message: `OK`,
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
