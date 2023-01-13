package db

import (
	"context"

	db "github.com/expert-pancake/service/business-relation/db/sqlc"
	uuid "github.com/satori/go.uuid"
)

type CreateNewContactGroupTrxParams struct {
	CompanyId   string
	ImageUrl    string
	Name        string
	Description string
	Members     []string
}

type CreateNewContactGroupTrxResult struct {
	ContactGroupId string
	CompanyId      string
	ImageUrl       string
	Name           string
	Description    string
	Members        []string
}

func (trx *Trx) CreateNewContactGroupTrx(ctx context.Context, arg CreateNewContactGroupTrxParams) (CreateNewContactGroupTrxResult, error) {
	var result CreateNewContactGroupTrxResult

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		id := uuid.NewV4().String()

		contactGroupRes, err := q.InsertContactGroup(ctx, db.InsertContactGroupParams{
			ID:          id,
			CompanyID:   arg.CompanyId,
			ImageUrl:    arg.ImageUrl,
			Name:        arg.Name,
			Description: arg.Description,
		})
		if err != nil {
			return err
		}

		for _, d := range arg.Members {
			err = q.UpdateContactBookGroupId(ctx, db.UpdateContactBookGroupIdParams{
				ID:             d,
				ContactGroupID: id,
			})
			if err != nil {
				return err
			}
		}

		result.ContactGroupId = id
		result.CompanyId = contactGroupRes.CompanyID
		result.ImageUrl = contactGroupRes.ImageUrl
		result.Name = contactGroupRes.Name
		result.Description = contactGroupRes.Description
		result.Members = arg.Members

		return err
	})

	return result, err
}
