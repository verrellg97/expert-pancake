package db

import (
	"context"

	db "github.com/expert-pancake/service/inventory/db/sqlc"
)

type UpdateItemTrxParams struct {
	ItemId      string
	ImageUrl    string
	Name        string
	BrandId     string
	GroupId     string
	Tag         string
	Description string
}

type UpdateItemTrxResult struct {
	CompanyId   string
	ItemId      string
	VariantId   string
	ImageUrl    string
	Code        string
	Name        string
	VariantName string
	BrandId     string
	BrandName   string
	GroupId     string
	GroupName   string
	Tag         string
	Description string
	IsDefault   bool
	Price       int64
	Stock       int64
}

func (trx *Trx) UpdateItemTrx(ctx context.Context, arg UpdateItemTrxParams) (UpdateItemTrxResult, error) {
	var result UpdateItemTrxResult

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		itemRes, err := q.UpdateItem(ctx, db.UpdateItemParams{
			ID:          arg.ItemId,
			ImageUrl:    arg.ImageUrl,
			Name:        arg.Name,
			BrandID:     arg.BrandId,
			GroupID:     arg.GroupId,
			Tag:         arg.Tag,
			Description: arg.Description,
		})
		if err != nil {
			return err
		}

		brandRes, err := q.GetBrandById(ctx, arg.BrandId)
		if err != nil {
			return err
		}

		groupRes, err := q.GetGroupById(ctx, arg.GroupId)
		if err != nil {
			return err
		}

		itemVariantRes, err := q.UpdateItemVariantDefault(ctx, db.UpdateItemVariantDefaultParams{
			ItemID:   arg.ItemId,
			ImageUrl: arg.ImageUrl,
		})
		if err != nil {
			return err
		}

		result.CompanyId = itemRes.CompanyID
		result.ItemId = arg.ItemId
		result.VariantId = itemVariantRes.ID
		result.ImageUrl = arg.ImageUrl
		result.Code = itemRes.Code
		result.Name = arg.Name
		result.VariantName = itemVariantRes.Name
		result.BrandId = arg.BrandId
		result.BrandName = brandRes.Name
		result.GroupId = arg.GroupId
		result.GroupName = groupRes.Name
		result.Tag = arg.Tag
		result.Description = arg.Description
		result.IsDefault = itemVariantRes.IsDefault
		result.Price = itemVariantRes.Price
		result.Stock = itemVariantRes.Stock

		return err
	})

	return result, err
}
