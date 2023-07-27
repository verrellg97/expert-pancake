package impl

import (
	"context"
	"net/http"
	"strconv"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/inventory/model"
)

func (a inventoryService) GetOutgoingStock(w http.ResponseWriter, r *http.Request) error {
	result, err := a.dbTrx.GetOutgoingStock(context.Background())
	if err != nil {
		return errors.NewServerError(model.GetOutgoingStockError, err.Error())
	}

	var responseData = make([]model.GetOutgoingStockResponseStruct, 0)

	for _, d := range result {
		var data = model.GetOutgoingStockResponseStruct{
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

	res := model.GetOutgoingStockResponse{
		responseData,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
