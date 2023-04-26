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

func (a inventoryService) GetPurchaseItemVariants(w http.ResponseWriter, r *http.Request) error {

	var req model.GetPurchaseItemVariantsRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetPurchaseItemVariants(context.Background(), db.GetPurchaseItemVariantsParams{
		PrimaryCompanyID: req.PrimaryCompanyId,
		SecondaryItemID:  req.SecondaryItemId,
		Name:             util.WildCardString(req.Keyword),
	})
	log.Println(err)
	if err != nil {
		return errors.NewServerError(model.GetPurchaseItemVariantsError, err.Error())
	}

	var datas = make([]model.PurchaseItemVariant, 0)

	for _, d := range result {
		var data = model.PurchaseItemVariant{
			PrimaryItemVariantId:     d.PrimaryItemVariantID,
			PrimaryItemVariantName:   d.PrimaryItemVariantName,
			SecondaryItemVariantId:   d.SecondaryItemVariantID,
			SecondaryItemVariantName: d.SecondaryItemVariantName,
		}
		datas = append(datas, data)
	}

	res := model.GetPurchaseItemVariantsResponse{
		PurchaseItemVariants: datas,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
