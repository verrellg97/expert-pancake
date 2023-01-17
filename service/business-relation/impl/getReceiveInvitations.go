package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/business-relation/model"
)

func (a businessRelationService) GetReceiveInvitations(w http.ResponseWriter, r *http.Request) error {

	var req model.GetReceiveInvitationsRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetReceiveInvitations(context.Background(), req.CompanyId)
	if err != nil {
		return errors.NewServerError(model.GetReceiveInvitationsError, err.Error())
	}

	var invitation_contact_books = make([]model.InvitationContactBook, 0)

	for _, d := range result {
		var invitation_contact_book = model.InvitationContactBook{
			InvitationId:     d.ID,
			ContactBookId:    d.ContactBookID,
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
			Status: d.Status,
		}
		invitation_contact_books = append(invitation_contact_books, invitation_contact_book)
	}

	res := invitation_contact_books
	httpHandler.WriteResponse(w, res)

	return nil
}
