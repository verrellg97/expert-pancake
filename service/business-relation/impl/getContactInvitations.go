package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/business-relation/model"
)

func (a businessRelationService) GetContactInvitations(w http.ResponseWriter, r *http.Request) error {

	var req model.GetContactInvitationsRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetContactInvitations(context.Background(), req.CompanyId)
	if err != nil {
		return errors.NewServerError(model.GetContactInvitationsError, err.Error())
	}

	var my_contact_books = make([]model.MyContactBook, 0)

	for _, d := range result {
		var my_contact_book = model.MyContactBook{
			ContactBookId:    d.ID,
			KonekinId:        d.KonekinID,
			PrimaryCompanyId: d.PrimaryCompanyID,
			Name:             d.Name,
			Email:            d.Email,
			Phone:            d.Phone,
			Mobile:           d.Mobile,
			Web:              d.Web,
			AdditionalInfo: model.ContactBookAdditionaInfo{
				Nickname: d.Nickname,
				Tag:      d.Tag,
				Note:     d.Note,
			},
			MailingAddress: model.ContactBookAddress{
				Province:    d.MailingProvince,
				Regency:     d.MailingRegency,
				District:    d.MailingDistrict,
				PostalCode:  d.MailingPostalCode,
				FullAddress: d.MailingFullAddress,
			},
			ShippingAddress: model.ContactBookAddress{
				Province:    d.ShippingProvince,
				Regency:     d.ShippingRegency,
				District:    d.ShippingDistrict,
				PostalCode:  d.ShippingPostalCode,
				FullAddress: d.ShippingFullAddress,
			},
		}
		my_contact_books = append(my_contact_books, my_contact_book)
	}

	res := my_contact_books
	httpHandler.WriteResponse(w, res)

	return nil
}
