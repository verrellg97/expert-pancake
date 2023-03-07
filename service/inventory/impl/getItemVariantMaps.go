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
)

func (a inventoryService) GetItemVariantMaps(w http.ResponseWriter, r *http.Request) error {

	var req model.GetItemVariantMapsRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetItemVariantMaps(context.Background(), db.GetItemVariantMapsParams{
		CompanyID:     req.CompanyId,
		PrimaryItemID: req.ItemId,
	})
	log.Println(err)
	if err != nil {
		return errors.NewServerError(model.GetItemVariantMapsError, err.Error())
	}

	var datas = make([]model.ItemVariantMap, 0)

	for _, d := range result {
		var data = model.ItemVariantMap{
			ItemVariantMap:            d.ID,
			PrimaryItemId:             d.PrimaryItemID,
			PrimaryItemName:           d.PrimaryItemName,
			PrimaryItemVariantId:      d.PrimaryItemVariantID,
			PrimaryItemVariantName:    d.PrimaryItemVariantName,
			PrimaryItemVariantPrice:   strconv.FormatInt(d.PrimaryItemVariantPrice, 10),
			PrimaryItemUnitId:         d.PrimaryItemUnitID,
			PrimaryItemUnitName:       d.PrimaryItemUnitName,
			SecondaryItemId:           d.SecondaryItemID,
			SecondaryItemName:         d.SecondaryItemName,
			SecondaryItemVariantId:    d.SecondaryItemVariantID,
			SecondaryItemVariantName:  d.SecondaryItemVariantName,
			SecondaryItemVariantPrice: strconv.FormatInt(d.SecondaryItemVariantPrice, 10),
			SecondaryItemUnitId:       d.SecondaryItemUnitID,
			SecondaryItemUnitName:     d.SecondaryItemUnitName,
		}
		datas = append(datas, data)
	}

	res := model.GetItemVariantMapsResponse{
		ItemVariantMaps: datas,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
