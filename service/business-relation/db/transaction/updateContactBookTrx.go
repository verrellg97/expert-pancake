package db

import (
	"context"

	db "github.com/expert-pancake/service/business-relation/db/sqlc"
	"github.com/expert-pancake/service/business-relation/model"
)

type UpdateContactBookTrxParams struct {
	Id              string
	ContactGroupId  string
	Name            string
	Email           string
	Phone           string
	Mobile          string
	Web             string
	AdditionalInfo  model.ContactBookAdditionaInfo
	MailingAddress  model.ContactBookAddress
	ShippingAddress model.ContactBookAddress
	IsAllBranches   bool
	Branches        []string
	IsCustomer      bool
	IsSupplier      bool
}

type UpdateContactBookTrxResult struct {
	ContactBookId      string
	KonekinId          string
	PrimaryCompanyId   string
	SecondaryCompanyId string
	ContactGroupId     string
	ContactGroupName   string
	Name               string
	Email              string
	Phone              string
	Mobile             string
	Web                string
	AdditionalInfo     model.ContactBookAdditionaInfo
	MailingAddress     model.ContactBookAddress
	ShippingAddress    model.ContactBookAddress
	IsAllBranches      bool
	Branches           []string
	IsCustomer         bool
	IsSupplier         bool
}

func (trx *Trx) UpdateContactBookTrx(ctx context.Context, arg UpdateContactBookTrxParams) (UpdateContactBookTrxResult, error) {
	var result UpdateContactBookTrxResult

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		_, err = q.UpdateContactBook(ctx, db.UpdateContactBookParams{
			ID:             arg.Id,
			ContactGroupID: arg.ContactGroupId,
			Name:           arg.Name,
			Email:          arg.Email,
			Phone:          arg.Phone,
			Mobile:         arg.Mobile,
			Web:            arg.Web,
			IsAllBranches:  arg.IsAllBranches,
			IsCustomer:     arg.IsCustomer,
			IsSupplier:     arg.IsSupplier,
		})
		if err != nil {
			return err
		}

		contactBookRes, err := q.GetContactBookById(ctx, arg.Id)
		if err != nil {
			return err
		}

		err = q.UpdateContactBookAdditionalInfo(ctx, db.UpdateContactBookAdditionalInfoParams{
			ContactBookID: arg.Id,
			Nickname:      arg.AdditionalInfo.Nickname,
			Tag:           arg.AdditionalInfo.Tag,
			Note:          arg.AdditionalInfo.Note,
		})
		if err != nil {
			return err
		}

		err = q.UpdateContactBookMailingAddress(ctx, db.UpdateContactBookMailingAddressParams{
			ContactBookID: arg.Id,
			Province:      arg.MailingAddress.Province,
			Regency:       arg.MailingAddress.Regency,
			District:      arg.MailingAddress.District,
			PostalCode:    arg.MailingAddress.PostalCode,
			FullAddress:   arg.MailingAddress.FullAddress,
		})
		if err != nil {
			return err
		}

		err = q.UpdateContactBookShippingAddress(ctx, db.UpdateContactBookShippingAddressParams{
			ContactBookID: arg.Id,
			Province:      arg.ShippingAddress.Province,
			Regency:       arg.ShippingAddress.Regency,
			District:      arg.ShippingAddress.District,
			PostalCode:    arg.ShippingAddress.PostalCode,
			FullAddress:   arg.ShippingAddress.FullAddress,
		})
		if err != nil {
			return err
		}

		err = q.DeleteContactBookBranches(ctx, arg.Id)
		if err != nil {
			return err
		}

		if !arg.IsAllBranches {
			for _, d := range arg.Branches {
				err = q.InsertContactBookBranch(ctx, db.InsertContactBookBranchParams{
					ContactBookID:   arg.Id,
					CompanyBranchID: d,
				})
				if err != nil {
					return err
				}
			}
		}

		result.ContactBookId = contactBookRes.ID
		result.KonekinId = contactBookRes.KonekinID
		result.PrimaryCompanyId = contactBookRes.PrimaryCompanyID
		result.SecondaryCompanyId = contactBookRes.SecondaryCompanyID
		result.ContactGroupId = contactBookRes.ContactGroupID
		result.ContactGroupName = contactBookRes.ContactGroupName
		result.Name = contactBookRes.Name
		result.Email = contactBookRes.Email
		result.Phone = contactBookRes.Phone
		result.Mobile = contactBookRes.Mobile
		result.Web = contactBookRes.Web
		result.AdditionalInfo = arg.AdditionalInfo
		result.MailingAddress = arg.MailingAddress
		result.ShippingAddress = arg.ShippingAddress
		result.IsAllBranches = contactBookRes.IsAllBranches
		result.Branches = arg.Branches
		result.IsCustomer = contactBookRes.IsCustomer
		result.IsSupplier = contactBookRes.IsSupplier

		return err
	})

	return result, err
}
