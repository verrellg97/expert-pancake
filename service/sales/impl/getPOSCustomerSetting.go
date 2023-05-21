package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/sales/impl/client"
	"github.com/expert-pancake/service/sales/model"
)

func (a salesService) GetPOSCustomerSetting(w http.ResponseWriter, r *http.Request) error {

	var req model.GetPOSCustomerSettingRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	var posCustomers = make([]model.POSCustomer, 0)
	detailResult, err := a.dbTrx.GetPOSCustomerSetting(context.Background(), req.BranchId)
	if err != nil {
		return errors.NewServerError(model.GetPOSCustomerSettingError, err.Error())
	}

	for _, d := range detailResult {
		argContactBook := client.GetContactBooksRequest{
			Id:        d.ContactBookID,
			CompanyId: "1",
		}
		contactBook, err := client.GetContactBooks(argContactBook)
		if err != nil {
			return errors.NewServerError(model.GetPOSCustomerSettingError, err.Error())
		}

		if len(contactBook.Result) > 0 {
			contactBookName := contactBook.Result[0].Name
			secondaryCompanyId := contactBook.Result[0].SecondaryCompanyId
			konekinId := contactBook.Result[0].KonekinId

			var posCustomer = model.POSCustomer{
				ContactBookId:      d.ContactBookID,
				ContactBookName:    contactBookName,
				SecondaryCompanyId: secondaryCompanyId,
				KonekinId:          konekinId,
			}
			posCustomers = append(posCustomers, posCustomer)
		}
	}

	res := model.GetPOSCustomerSettingResponse{
		POSCustomers: posCustomers,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
