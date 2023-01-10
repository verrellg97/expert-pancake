package impl

import (
	"context"
	"net/http"
	"strconv"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/business-relation/model"
)

func (a businessRelationService) GetCustomers(w http.ResponseWriter, r *http.Request) error {

	var req model.GetCustomersRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetCustomers(context.Background(), req.CompanyId)
	if err != nil {
		return errors.NewServerError(model.GetCustomersError, err.Error())
	}

	var customers = make([]model.CustomerInfo, 0)

	for _, d := range result {
		var customer = model.CustomerInfo{
			ContactBookId:    d.ID,
			ContactGroupName: d.ContactGroupName,
			Name:             d.Name,
			Email:            d.Email,
			Phone:            d.Phone,
			Mobile:           d.Mobile,
			Web:              d.Web,
			IsTax:            d.IsTax,
			TaxId:            d.TaxID,
			Pic:              d.Pic,
			CreditLimit:      strconv.FormatInt(d.CreditLimit, 10),
			PaymentTerm:      strconv.Itoa(int(d.PaymentTerm)),
		}
		customers = append(customers, customer)
	}

	res := customers
	httpHandler.WriteResponse(w, res)

	return nil
}
