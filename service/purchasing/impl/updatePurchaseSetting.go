package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/purchasing/db/sqlc"
	"github.com/expert-pancake/service/purchasing/model"
)

func (a purchasingService) UpdatePurchaseSetting(w http.ResponseWriter, r *http.Request) error {

	var req model.UpdatePurchaseSettingRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.UpsertPurchaseSettingParams{
		CompanyID:          req.CompanyId,
		IsAutoApproveOrder: req.IsAutoApproveOrder,
	}

	result, err := a.dbTrx.UpsertPurchaseSetting(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpdatePurchaseSettingError, err.Error())
	}

	res := model.UpdatePurchaseSettingResponse{
		PurchaseSetting: model.PurchaseSetting{
			CompanyId:          result.CompanyID,
			IsAutoApproveOrder: result.IsAutoApproveOrder,
		},
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
