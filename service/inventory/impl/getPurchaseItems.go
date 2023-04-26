package impl

import (
	"context"
	"log"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/inventory/db/sqlc"
	"github.com/expert-pancake/service/inventory/model"
	"github.com/expert-pancake/service/inventory/util"
)

func (a inventoryService) GetPurchaseItems(w http.ResponseWriter, r *http.Request) error {

	var req model.GetPurchaseItemsRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetPurchaseItems(context.Background(), db.GetPurchaseItemsParams{
		PrimaryCompanyID:   req.PrimaryCompanyId,
		SecondaryCompanyID: req.SecondaryCompanyId,
		Name:               util.WildCardString(req.Keyword),
	})
	log.Println(err)
	if err != nil {
		return errors.NewServerError(model.GetPurchaseItemsError, err.Error())
	}

	var datas = make([]model.PurchaseItem, 0)

	for _, d := range result {
		var data = model.PurchaseItem{
			PrimaryItemId:     d.PrimaryItemID,
			PrimaryItemCode:   d.PrimaryItemCode,
			PrimaryItemName:   d.PrimaryItemName,
			SecondaryItemId:   d.SecondaryItemID,
			SecondaryItemCode: d.SecondaryItemCode,
			SecondaryItemName: d.SecondaryItemName,
		}
		datas = append(datas, data)
	}

	res := model.GetPurchaseItemsResponse{
		PurchaseItems: datas,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
