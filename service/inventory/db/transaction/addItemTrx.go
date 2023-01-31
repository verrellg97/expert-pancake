package db

import (
	"context"
	"fmt"
	"math/rand"

	db "github.com/expert-pancake/service/inventory/db/sqlc"
	"github.com/expert-pancake/service/inventory/model"
	uuid "github.com/satori/go.uuid"
)

type AddItemTrxResult struct {
	CompanyId   string
	ItemId      string
	VariantId   string
	ImageUrl    string
	Code        string
	Name        string
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

func (trx *Trx) AddItemTrx(ctx context.Context, arg model.AddItemRequest) (AddItemTrxResult, error) {
	var result AddItemTrxResult

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		id := uuid.NewV4().String()

		itemRes, err := q.InsertItem(ctx, db.InsertItemParams{
			ID:          id,
			CompanyID:   arg.CompanyId,
			ImageUrl:    arg.ImageUrl,
			Code:        "BRG-" + fmt.Sprintf("%08d", rand.Intn(100000000)),
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

		itemVariantRes, err := q.InsertItemVariant(ctx, db.InsertItemVariantParams{
			ID:        uuid.NewV4().String(),
			ItemID:    id,
			ImageUrl:  arg.ImageUrl,
			Name:      arg.Name,
			IsDefault: true,
		})
		if err != nil {
			return err
		}

		result.CompanyId = arg.CompanyId
		result.ItemId = id
		result.VariantId = itemVariantRes.ID
		result.ImageUrl = arg.ImageUrl
		result.Code = itemRes.Code
		result.Name = arg.Name
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
