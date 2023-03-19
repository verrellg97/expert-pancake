package impl

import (
	"context"
	"net/http"
	"strconv"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/business-relation/model"
)

func (a businessRelationService) GetCustomer(w http.ResponseWriter, r *http.Request) error {

	var req model.GetCustomerRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetCustomer(context.Background(), req.ContactBookId)
	if err != nil {
		return errors.NewServerError(model.GetCustomerError, err.Error())
	}

	res := model.GetCustomerResponse{
		CustomerInfo: model.CustomerInfo{
			ContactBookId:    result.ID,
			KonekinId:        result.KonekinID,
			CompanyId:        result.SecondaryCompanyID,
			ContactGroupName: result.ContactGroupName,
			Name:             result.Name,
			Email:            result.Email,
			Phone:            result.Phone,
			Mobile:           result.Mobile,
			Web:              result.Web,
			IsDefault:        result.IsDefault,
			IsTax:            result.IsTax,
			TaxId:            result.TaxID,
			Pic:              result.Pic,
			CreditLimit:      strconv.FormatInt(result.CreditLimit, 10),
			PaymentTerm:      strconv.Itoa(int(result.PaymentTerm)),
		},
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
