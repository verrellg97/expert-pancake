package db

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"

	db "github.com/expert-pancake/service/inventory/db/sqlc"
	"github.com/expert-pancake/service/inventory/util"
	uuid "github.com/satori/go.uuid"
)

type AddItemTrxParams struct {
	CompanyId   string
	ImageUrl    string
	Code        string
	Name        string
	Barcode     string
	BrandId     string
	GroupIds    []string
	Tag         string
	Price       string
	Description string
}

type AddItemTrxResult struct {
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

func (trx *Trx) AddItemTrx(ctx context.Context, arg AddItemTrxParams) (AddItemTrxResult, error) {
	var result AddItemTrxResult

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		id := uuid.NewV4().String()

		if arg.Code == "" {
			arg.Code = "BRG-" + fmt.Sprintf("%08d", rand.Intn(100000000))
		}

		itemRes, err := q.InsertItem(ctx, db.InsertItemParams{
			ID:          id,
			CompanyID:   arg.CompanyId,
			ImageUrl:    arg.ImageUrl,
			Code:        arg.Code,
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
		itemVariantRes, err := q.InsertItemVariant(ctx, db.InsertItemVariantParams{
			ID:        uuid.NewV4().String(),
			ItemID:    id,
			ImageUrl:  arg.ImageUrl,
			Barcode:   arg.Code,
			Price:     price,
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
