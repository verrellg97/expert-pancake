package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/business-relation/db/transaction"
	"github.com/expert-pancake/service/business-relation/model"
)

func (a businessRelationService) AddContactBook(w http.ResponseWriter, r *http.Request) error {

	var req model.AddContactBookRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	var branches = make([]string, 0)
	if req.Branches != nil {
		branches = req.Branches
	}

	arg := db.CreateNewContactBookTrxParams{
		PrimaryCompanyId: req.PrimaryCompanyId,
		ContactGroupId:   req.ContactGroupId,
		Name:             req.Name,
		Email:            req.Email,
		Phone:            req.Phone,
		Mobile:           req.Mobile,
		Web:              req.Web,
		AdditionalInfo:   req.AdditionalInfo,
		MailingAddress:   req.MailingAddress,
		ShippingAddress:  req.ShippingAddress,
		IsAllBranches:    req.IsAllBranches,
		Branches:         branches,
		IsCustomer:       req.IsCustomer,
		IsSupplier:       req.IsSupplier,
	}

	result, err := a.dbTrx.CreateNewContactBookTrx(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.AddNewContactBookError, err.Error())
	}

	res := model.AddContactBookResponse{
		ContactBook: model.ContactBook{
			ContactBookId:      result.ContactBookId,
			KonekinId:          result.KonekinId,
			PrimaryCompanyId:   result.PrimaryCompanyId,
			SecondaryCompanyId: result.SecondaryCompanyId,
			ContactGroupId:     result.ContactGroupId,
			ContactGroupName:   result.ContactGroupName,
			Name:               result.Name,
			Email:              result.Email,
			Phone:              result.Phone,
			Mobile:             result.Mobile,
			Web:                result.Web,
			AdditionalInfo:     result.AdditionalInfo,
			MailingAddress:     result.MailingAddress,
			ShippingAddress:    result.ShippingAddress,
			IsAllBranches:      result.IsAllBranches,
			Branches:           result.Branches,
			IsCustomer:         result.IsCustomer,
			IsSupplier:         result.IsSupplier,
		},
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
