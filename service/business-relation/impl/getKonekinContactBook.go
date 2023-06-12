package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/business-relation/db/sqlc"
	"github.com/expert-pancake/service/business-relation/model"
)

func (a businessRelationService) GetKonekinContactBook(w http.ResponseWriter, r *http.Request) error {

	var req model.GetKonekinContactBookRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetKonekinContactBook(context.Background(), db.GetKonekinContactBookParams{
		PrimaryCompanyID:   req.PrimaryCompanyId,
		SecondaryCompanyID: req.SecondaryCompanyId,
	})
	if err != nil {
		return errors.NewServerError(model.GetKonekinContactBookError, err.Error())
	}

	res := model.KonekinContactBook{
		ContactBookId:    result.ID,
		KonekinId:        result.KonekinID,
		PrimaryCompanyId: result.PrimaryCompanyID,
		Name:             result.Name,
		Email:            result.Email,
		Phone:            result.Phone,
		Mobile:           result.Mobile,
		Web:              result.Web,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
