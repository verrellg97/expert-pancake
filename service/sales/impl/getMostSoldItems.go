package impl

import (
	"context"
	"net/http"
	"strconv"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/sales/db/sqlc"
	"github.com/expert-pancake/service/sales/impl/client"
	"github.com/expert-pancake/service/sales/model"
	"github.com/expert-pancake/service/sales/util"
)

func (a salesService) GetMostSoldItems(w http.ResponseWriter, r *http.Request) error {

	var req model.GetMostSoldItemsRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetMostSoldItems(context.Background(), db.GetMostSoldItemsParams{
		CompanyID: req.CompanyId,
		BranchID:  util.WildCardString(req.BranchId),
		StartDate: util.StringToDate(req.StartDate),
		EndDate:   util.StringToDate(req.EndDate),
	})
	if err != nil {
		return errors.NewServerError(model.GetMostSoldItemsError, err.Error())
	}

	var mostSoldItems = make([]model.MostSoldItem, 0)

	for _, d := range result {
		argItemVariant := client.GetItemVariantsRequest{
			Id: d.ItemVariantID,
		}
		itemVariant, err := client.GetItemVariants(argItemVariant)
		if err != nil {
			return errors.NewServerError(model.GetPOSItemsError, err.Error())
		}
		itemVariantName := ""
		itemCode := ""
		itemName := ""
		if len(itemVariant.Result.ItemVariants) > 0 {
			itemVariantName = itemVariant.Result.ItemVariants[0].VariantName
			itemCode = itemVariant.Result.ItemVariants[0].Code
			itemName = itemVariant.Result.ItemVariants[0].Name
		}

		var mostSoldItem = model.MostSoldItem{
			ItemVariantId:   d.ItemVariantID,
			ItemVariantName: itemVariantName,
			ItemCode:        itemCode,
			ItemName:        itemName,
			Total:           strconv.FormatInt(d.Total, 10),
		}
		mostSoldItems = append(mostSoldItems, mostSoldItem)
	}

	res := model.GetMostSoldItemsResponse{
		MostSoldItems: mostSoldItems,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
