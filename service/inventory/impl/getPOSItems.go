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

func (a inventoryService) GetPOSItems(w http.ResponseWriter, r *http.Request) error {

	var req model.GetPOSItemsRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetPOSItems(context.Background(), db.GetPOSItemsParams{
		CompanyID: req.CompanyId,
		Name:      util.WildCardString(req.Keyword),
	})
	log.Println(err)
	if err != nil {
		return errors.NewServerError(model.GetPOSItemsError, err.Error())
	}

	var datas = make([]model.POSItem, 0)

	for _, d := range result {
		var data = model.POSItem{
			ItemId:        d.ItemID,
			ItemName:      d.ItemName,
			VariantId:     d.VariantID,
			VariantName:   d.VariantName,
			ItemUnitId:    d.ItemUnitID,
			UnitName:      d.UnitName,
			ItemUnitValue: strconv.FormatInt(d.ItemUnitValue, 10),
			Price:         strconv.FormatInt(d.Price, 10),
		}
		datas = append(datas, data)
	}

	res := model.GetPOSItemsResponse{
		POSItems: datas,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
