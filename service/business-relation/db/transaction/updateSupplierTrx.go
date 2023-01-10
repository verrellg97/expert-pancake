package db

import (
	"context"
	"strconv"

	db "github.com/expert-pancake/service/business-relation/db/sqlc"
)

type UpdateSupplierTrxParams struct {
	Id          string
	IsTax       bool
	TaxId       string
	Pic         string
	CreditLimit int64
	PaymentTerm int32
}

type UpdateSupplierTrxResult struct {
	ContactBookId    string
	ContactGroupName string
	Name             string
	Email            string
	Phone            string
	Mobile           string
	Web              string
	IsTax            bool
	TaxId            string
	Pic              string
	CreditLimit      string
	PaymentTerm      string
}

func (trx *Trx) UpdateSupplierTrx(ctx context.Context, arg UpdateSupplierTrxParams) (UpdateSupplierTrxResult, error) {
	var result UpdateSupplierTrxResult

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		err = q.UpdateContactBookTaxInfo(ctx, db.UpdateContactBookTaxInfoParams{
			ID:    arg.Id,
			IsTax: arg.IsTax,
			TaxID: arg.TaxId,
		})
		if err != nil {
			return err
		}

		err = q.UpsertSupplierInfo(ctx, db.UpsertSupplierInfoParams{
			ContactBookID: arg.Id,
			Pic:           arg.Pic,
			CreditLimit:   arg.CreditLimit,
			PaymentTerm:   arg.PaymentTerm,
		})
		if err != nil {
			return err
		}

		contactBookRes, err := q.GetContactBookById(ctx, arg.Id)
		if err != nil {
			return err
		}

		result.ContactBookId = contactBookRes.ID
		result.ContactGroupName = contactBookRes.ContactGroupName
		result.Name = contactBookRes.Name
		result.Email = contactBookRes.Email
		result.Phone = contactBookRes.Phone
		result.Mobile = contactBookRes.Mobile
		result.Web = contactBookRes.Web
		result.IsTax = contactBookRes.IsTax
		result.TaxId = contactBookRes.TaxID
		result.Pic = arg.Pic
		result.CreditLimit = strconv.FormatInt(arg.CreditLimit, 10)
		result.PaymentTerm = strconv.Itoa(int(arg.PaymentTerm))

		return err
	})

	return result, err
}
