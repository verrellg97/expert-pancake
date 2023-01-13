package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/business-relation/model"
)

func (a businessRelationService) GetMyContactBook(w http.ResponseWriter, r *http.Request) error {

	var req model.GetMyContactBookRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetMyContactBook(context.Background(), req.CompanyId)
	if err != nil {
		return errors.NewServerError(model.GetMyContactBookError, err.Error())
	}

	res := model.MyContactBook{
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
