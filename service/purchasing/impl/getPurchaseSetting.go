package impl

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/purchasing/db/sqlc"
	"github.com/expert-pancake/service/purchasing/model"
)

func (a purchasingService) GetPurchaseSetting(w http.ResponseWriter, r *http.Request) error {

	var req model.GetPurchaseSettingRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetPurchaseSetting(context.Background(), req.CompanyId)
	if err == sql.ErrNoRows {
		arg := db.UpsertPurchaseSettingParams{
			CompanyID: req.CompanyId,
		}

		result, err = a.dbTrx.UpsertPurchaseSetting(context.Background(), arg)
		if err != nil {
			return errors.NewServerError(model.GetPurchaseSettingError, err.Error())
		}
	} else if err != nil {
		return errors.NewServerError(model.GetPurchaseSettingError, err.Error())
	}

	res := model.GetPurchaseSettingResponse{
		PurchaseSetting: model.PurchaseSetting{
			CompanyId:          result.CompanyID,
			IsAutoApproveOrder: result.IsAutoApproveOrder,
		},
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
