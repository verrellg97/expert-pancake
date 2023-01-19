package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/business-relation/db/sqlc"
	"github.com/expert-pancake/service/business-relation/model"
	"github.com/expert-pancake/service/business-relation/util"
)

func (a businessRelationService) GetContactBooks(w http.ResponseWriter, r *http.Request) error {

	var req model.GetContactBooksRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	var isFilterGroupId = false
	if req.ContactGroupId != "" {
		isFilterGroupId = true
	}

	var isCustomerApplicant = false
	var isSupplierrApplicant = false
	if req.Applicant == "Customer" {
		isCustomerApplicant = true
	} else if req.Applicant == "Supplier" {
		isSupplierrApplicant = true
	}

	result, err := a.dbTrx.GetContactBooks(context.Background(), db.GetContactBooksParams{
		PrimaryCompanyID:    req.CompanyId,
		IsFilterGroupID:     isFilterGroupId,
		ContactGroupID:      req.ContactGroupId,
		IsCustomerApplicant: isCustomerApplicant,
		IsSupplierApplicant: isSupplierrApplicant,
	})
	if err != nil {
		return errors.NewServerError(model.GetContactBooksError, err.Error())
	}

	var contact_books = make([]model.ContactBook, 0)

	for _, d := range result {
		resultBranches, err := a.dbTrx.GetContactBookBranches(context.Background(), d.ID)
		if err != nil {
			return errors.NewServerError(model.GetContactBooksError, err.Error())
		}
		var contact_book = model.ContactBook{
			ContactBookId:      d.ID,
			KonekinId:          d.KonekinID,
			PrimaryCompanyId:   d.PrimaryCompanyID,
			SecondaryCompanyId: d.SecondaryCompanyID,
			ContactGroupId:     d.ContactGroupID,
			ContactGroupName:   d.ContactGroupName,
			Name:               d.Name,
			Email:              d.Email,
			Phone:              d.Phone,
			Mobile:             d.Mobile,
			Web:                d.Web,
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
			IsAllBranches: d.IsAllBranches,
			Branches:      util.ContactBookBranchDbToApi(resultBranches),
			IsCustomer:    d.IsCustomer,
			IsSupplier:    d.IsSupplier,
		}
		contact_books = append(contact_books, contact_book)
	}

	res := contact_books
	httpHandler.WriteResponse(w, res)

	return nil
}
