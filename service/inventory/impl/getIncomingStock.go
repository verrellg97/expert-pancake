package impl

import (
	"context"
	"net/http"
	"strconv"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/inventory/model"
)

func (a inventoryService) GetIncomingStock(w http.ResponseWriter, r *http.Request) error {
	result, err := a.dbTrx.GetIncomingStock(context.Background())
	if err != nil {
		return errors.NewServerError(model.GetIncomingStockError, err.Error())
	}

	var responseData = make([]model.GetIncomingStockResponseStruct, 0)

	for _, d := range result {
		var data = model.GetIncomingStockResponseStruct{
			ItemId:       d.ItemID,
			ItemCode:     d.ItemCode,
			ItemName:     d.ItemName,
			VariantId:    d.VariantID,
			VariantName:  d.VariantName,
			UnitId:       d.UnitID,
			UnitName:     d.UnitName,
			Amount:       strconv.FormatInt(int64(d.Amount), 10),
		}
		responseData = append(responseData, data)
	}

	res := model.GetIncomingStockResponse{
		responseData,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
