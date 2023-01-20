package db

import (
	"context"

	db "github.com/expert-pancake/service/business-relation/db/sqlc"
	uuid "github.com/satori/go.uuid"
)

type UpdateContactInvitationTrxParams struct {
	Id                     string
	SecondaryContactBookId string
	Status                 string
}

type UpdateContactInvitationTrxResult struct {
	Message string
}

func (trx *Trx) UpdateContactInvitationTrx(ctx context.Context, arg UpdateContactInvitationTrxParams) (UpdateContactInvitationTrxResult, error) {
	var result UpdateContactInvitationTrxResult

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		contactInvitation, err := q.UpdateContactInvitation(ctx, db.UpdateContactInvitationParams{
			ID:                     arg.Id,
			SecondaryContactBookID: arg.SecondaryContactBookId,
			Status:                 arg.Status,
		})
		if err != nil {
			return err
		}

		if arg.Status == "accepted" {
			primaryContactBook, err := q.GetMyContactBook(ctx, contactInvitation.PrimaryCompanyID)
			if err != nil {
				return err
			}
			secondaryContactBook, err := q.GetMyContactBook(ctx, contactInvitation.SecondaryCompanyID)
			if err != nil {
				return err
			}

			if contactInvitation.PrimaryContactBookID == "" {
				secondaryContactBookAdditionalInfo, err := q.GetContactBookAdditionalInfo(ctx, secondaryContactBook.ID)
				if err != nil {
					return err
				}
				secondaryContactBookMailingAddress, err := q.GetContactBookMailingAddress(ctx, secondaryContactBook.ID)
				if err != nil {
					return err
				}
				secondaryContactBookShippingAddress, err := q.GetContactBookShippingAddress(ctx, secondaryContactBook.ID)
				if err != nil {
					return err
				}

				id := uuid.NewV4().String()

				_, err = q.InsertContactBook(ctx, db.InsertContactBookParams{
					ID:                 id,
					KonekinID:          secondaryContactBook.KonekinID,
					PrimaryCompanyID:   contactInvitation.PrimaryCompanyID,
					SecondaryCompanyID: contactInvitation.SecondaryCompanyID,
					ContactGroupID:     "",
					Name:               secondaryContactBook.Name,
					Email:              secondaryContactBook.Email,
					Phone:              secondaryContactBook.Phone,
					Mobile:             secondaryContactBook.Mobile,
					Web:                secondaryContactBook.Web,
				})
				if err != nil {
					return err
				}

				err = q.InsertContactBookAdditionalInfo(ctx, db.InsertContactBookAdditionalInfoParams{
					ContactBookID: id,
					Nickname:      secondaryContactBookAdditionalInfo.Nickname,
					Tag:           secondaryContactBookAdditionalInfo.Tag,
					Note:          secondaryContactBookAdditionalInfo.Note,
				})
				if err != nil {
					return err
				}

				err = q.InsertContactBookMailingAddress(ctx, db.InsertContactBookMailingAddressParams{
					ContactBookID: id,
					Province:      secondaryContactBookMailingAddress.Province,
					Regency:       secondaryContactBookMailingAddress.Regency,
					District:      secondaryContactBookMailingAddress.District,
					PostalCode:    secondaryContactBookMailingAddress.PostalCode,
					FullAddress:   secondaryContactBookMailingAddress.FullAddress,
				})
				if err != nil {
					return err
				}

				err = q.InsertContactBookShippingAddress(ctx, db.InsertContactBookShippingAddressParams{
					ContactBookID: id,
					Province:      secondaryContactBookShippingAddress.Province,
					Regency:       secondaryContactBookShippingAddress.Regency,
					District:      secondaryContactBookShippingAddress.District,
					PostalCode:    secondaryContactBookShippingAddress.PostalCode,
					FullAddress:   secondaryContactBookShippingAddress.FullAddress,
				})
				if err != nil {
					return err
				}
			} else {
				err = q.UpdateContactBookRelation(ctx, db.UpdateContactBookRelationParams{
					ID:                 contactInvitation.PrimaryContactBookID,
					KonekinID:          secondaryContactBook.KonekinID,
					SecondaryCompanyID: contactInvitation.SecondaryCompanyID,
				})
				if err != nil {
					return err
				}
			}

			if contactInvitation.SecondaryContactBookID == "" {
				primaryContactBookAdditionalInfo, err := q.GetContactBookAdditionalInfo(ctx, primaryContactBook.ID)
				if err != nil {
					return err
				}
				primaryContactBookMailingAddress, err := q.GetContactBookMailingAddress(ctx, primaryContactBook.ID)
				if err != nil {
					return err
				}
				primaryContactBookShippingAddress, err := q.GetContactBookShippingAddress(ctx, primaryContactBook.ID)
				if err != nil {
					return err
				}

				id := uuid.NewV4().String()

				_, err = q.InsertContactBook(ctx, db.InsertContactBookParams{
					ID:                 id,
					KonekinID:          primaryContactBook.KonekinID,
					PrimaryCompanyID:   contactInvitation.SecondaryCompanyID,
					SecondaryCompanyID: contactInvitation.PrimaryCompanyID,
					ContactGroupID:     "",
					Name:               primaryContactBook.Name,
					Email:              primaryContactBook.Email,
					Phone:              primaryContactBook.Phone,
					Mobile:             primaryContactBook.Mobile,
					Web:                primaryContactBook.Web,
				})
				if err != nil {
					return err
				}

				err = q.InsertContactBookAdditionalInfo(ctx, db.InsertContactBookAdditionalInfoParams{
					ContactBookID: id,
					Nickname:      primaryContactBookAdditionalInfo.Nickname,
					Tag:           primaryContactBookAdditionalInfo.Tag,
					Note:          primaryContactBookAdditionalInfo.Note,
				})
				if err != nil {
					return err
				}

				err = q.InsertContactBookMailingAddress(ctx, db.InsertContactBookMailingAddressParams{
					ContactBookID: id,
					Province:      primaryContactBookMailingAddress.Province,
					Regency:       primaryContactBookMailingAddress.Regency,
					District:      primaryContactBookMailingAddress.District,
					PostalCode:    primaryContactBookMailingAddress.PostalCode,
					FullAddress:   primaryContactBookMailingAddress.FullAddress,
				})
				if err != nil {
					return err
				}

				err = q.InsertContactBookShippingAddress(ctx, db.InsertContactBookShippingAddressParams{
					ContactBookID: id,
					Province:      primaryContactBookShippingAddress.Province,
					Regency:       primaryContactBookShippingAddress.Regency,
					District:      primaryContactBookShippingAddress.District,
					PostalCode:    primaryContactBookShippingAddress.PostalCode,
					FullAddress:   primaryContactBookShippingAddress.FullAddress,
				})
				if err != nil {
					return err
				}
			} else {
				err = q.UpdateContactBookRelation(ctx, db.UpdateContactBookRelationParams{
					ID:                 contactInvitation.SecondaryContactBookID,
					KonekinID:          primaryContactBook.KonekinID,
					SecondaryCompanyID: contactInvitation.PrimaryCompanyID,
				})
				if err != nil {
					return err
				}
			}
		}

		result.Message = "OK"
		return err
	})

	return result, err
}
