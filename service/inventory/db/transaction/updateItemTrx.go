package db

import (
	"context"
	"strconv"

	db "github.com/expert-pancake/service/inventory/db/sqlc"
	"github.com/expert-pancake/service/inventory/util"
)

type UpdateItemTrxParams struct {
	ItemId      string
	ImageUrl    string
	Barcode     string
	Name        string
	BrandId     string
	GroupIds    []string
	Tag         string
	Price       string
	Description string
}

type UpdateItemTrxResult struct {
	CompanyId   string
	ItemId      string
	VariantId   string
	ImageUrl    string
	Code        string
	Barcode     string
	Name        string
	VariantName string
	BrandId     string
	BrandName   string
	Groups      string
	Tag         string
	Description string
	IsDefault   bool
	Price       int64
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
			GroupID:     util.ArrayToString(arg.GroupIds),
			Tag:         arg.Tag,
			Description: arg.Description,
		})
		if err != nil {
			return err
		}

		if arg.BrandId != "" {
			brandRes, err := q.GetBrandById(ctx, arg.BrandId)
			if err != nil {
				return err
			}
			result.BrandName = brandRes.Name
		}

		groupRes, err := q.GetItemGroups(ctx, arg.GroupIds)
		if err != nil {
			return err
		}

		price, _ := strconv.ParseInt(arg.Price, 10, 64)
		itemVariantRes, err := q.UpdateItemVariantDefault(ctx, db.UpdateItemVariantDefaultParams{
			ItemID:   arg.ItemId,
			ImageUrl: arg.ImageUrl,
			Barcode:  arg.Barcode,
			Price:    price,
		})
		if err != nil {
			return err
		}

		result.CompanyId = itemRes.CompanyID
		result.ItemId = arg.ItemId
		result.VariantId = itemVariantRes.ID
		result.ImageUrl = arg.ImageUrl
		result.Code = itemRes.Code
		result.Barcode = itemVariantRes.Barcode
		result.Name = arg.Name
		result.VariantName = itemVariantRes.Name
		result.BrandId = arg.BrandId
		result.Groups = groupRes
		result.Tag = arg.Tag
		result.Description = arg.Description
		result.IsDefault = itemVariantRes.IsDefault
		result.Price = itemVariantRes.Price

		return err
	})

	return result, err
}
