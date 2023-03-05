package impl

import (
	"context"
	"net/http"
	"strconv"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/business-relation/model"
)

func (a businessRelationService) GetSuppliers(w http.ResponseWriter, r *http.Request) error {

	var req model.GetSuppliersRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetSuppliers(context.Background(), req.CompanyId)
	if err != nil {
		return errors.NewServerError(model.GetSuppliersError, err.Error())
	}

	var suppliers = make([]model.SupplierInfo, 0)

	for _, d := range result {
		var supplier = model.SupplierInfo{
			ContactBookId:    d.ID,
			KonekinId:        d.KonekinID,
			CompanyId:        d.SecondaryCompanyID,
			ContactGroupName: d.ContactGroupName,
			Name:             d.Name,
			Email:            d.Email,
			Phone:            d.Phone,
			Mobile:           d.Mobile,
			Web:              d.Web,
			IsDefault:        d.IsDefault,
			IsTax:            d.IsTax,
			TaxId:            d.TaxID,
			Pic:              d.Pic,
			CreditLimit:      strconv.FormatInt(d.CreditLimit, 10),
			PaymentTerm:      strconv.Itoa(int(d.PaymentTerm)),
		}
		suppliers = append(suppliers, supplier)
	}

	res := suppliers
	httpHandler.WriteResponse(w, res)

	return nil
}
