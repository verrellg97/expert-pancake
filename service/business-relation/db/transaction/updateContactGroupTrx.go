package db

import (
	"context"

	db "github.com/expert-pancake/service/business-relation/db/sqlc"
)

type UpdateContactGroupTrxParams struct {
	Id          string
	ImageUrl    string
	Name        string
	Description string
	Members     []string
}

type UpdateContactGroupTrxResult struct {
	ContactGroupId string
	CompanyId      string
	ImageUrl       string
	Name           string
	Description    string
	Members        []string
}

func (trx *Trx) UpdateContactGroupTrx(ctx context.Context, arg UpdateContactGroupTrxParams) (UpdateContactGroupTrxResult, error) {
	var result UpdateContactGroupTrxResult

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		contactGroupRes, err := q.UpdateContactGroup(ctx, db.UpdateContactGroupParams{
			ID:          arg.Id,
			ImageUrl:    arg.ImageUrl,
			Name:        arg.Name,
			Description: arg.Description,
		})
		if err != nil {
			return err
		}

		err = q.UpdateContactBookGroupIdByGroupId(ctx, db.UpdateContactBookGroupIdByGroupIdParams{
			ContactGroupID:    arg.Id,
			NewContactGroupID: "",
		})
		if err != nil {
			return err
		}

		for _, d := range arg.Members {
			err = q.UpdateContactBookGroupId(ctx, db.UpdateContactBookGroupIdParams{
				ID:             d,
				ContactGroupID: arg.Id,
			})
			if err != nil {
				return err
			}
		}

		result.ContactGroupId = arg.Id
		result.CompanyId = contactGroupRes.CompanyID
		result.ImageUrl = contactGroupRes.ImageUrl
		result.Name = contactGroupRes.Name
		result.Description = contactGroupRes.Description
		result.Members = arg.Members

		return err
	})

	return result, err
}
