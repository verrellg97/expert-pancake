package impl

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/inventory/model"
)

func (a inventoryService) GetVariantWarehouseStocks(w http.ResponseWriter, r *http.Request) error {

	var req model.GetVariantWarehouseStocksRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetVariantWarehouseStocks(context.Background(), req.WarehouseId)
	log.Println(err)
	if err != nil {
		return errors.NewServerError(model.GetVariantWarehouseStocksError, err.Error())
	}

	var datas = make([]model.VariantStock, 0)

	for _, d := range result {
		var data = model.VariantStock{
			ItemId:      d.ItemID,
			ItemName:    d.ItemName,
			VariantId:   d.VariantID,
			VariantName: d.VariantName,
			Stock:       strconv.FormatInt(d.Stock, 10),
		}
		datas = append(datas, data)
	}

	res := model.GetVariantWarehouseStocksResponse{
		VariantStocks: datas,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
