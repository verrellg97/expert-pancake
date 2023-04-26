package impl

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/inventory/db/sqlc"
	"github.com/expert-pancake/service/inventory/model"
	"github.com/expert-pancake/service/inventory/util"
)

func (a inventoryService) GetPurchaseItemVariantUnits(w http.ResponseWriter, r *http.Request) error {

	var req model.GetPurchaseItemVariantUnitsRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetPurchaseItemVariantUnits(context.Background(), db.GetPurchaseItemVariantUnitsParams{
		SecondaryCompanyID:   req.SecondaryCompanyId,
		PrimaryItemVariantID: req.PrimaryItemVariantId,
		Name:                 util.WildCardString(req.Keyword),
	})
	log.Println(err)
	if err != nil {
		return errors.NewServerError(model.GetPurchaseItemVariantUnitsError, err.Error())
	}

	var datas = make([]model.PurchaseItemVariantUnit, 0)

	for _, d := range result {
		var data = model.PurchaseItemVariantUnit{
			PrimaryItemUnitId:      d.PrimaryItemUnitID,
			PrimaryItemUnitName:    d.PrimaryItemUnitName,
			PrimaryItemUnitValue:   strconv.FormatInt(d.PrimaryItemUnitValue, 10),
			SecondaryItemUnitId:    d.SecondaryItemUnitID,
			SecondaryItemUnitName:  d.SecondaryItemUnitName,
			SecondaryItemUnitValue: strconv.FormatInt(d.SecondaryItemUnitValue, 10),
			Price:                  strconv.FormatInt(d.Price, 10),
		}
		datas = append(datas, data)
	}

	res := model.GetPurchaseItemVariantUnitsResponse{
		PurchaseItemVariantUnits: datas,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
