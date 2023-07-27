package impl

import (
	"context"
	"net/http"
	"strconv"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/inventory/model"
)

func (a inventoryService) GetUnderMinimumOrder(w http.ResponseWriter, r *http.Request) error {
	result, err := a.dbTrx.GetUnderMinimumOrder(context.Background())
	if err != nil {
		return errors.NewServerError(model.GetUnderMinimumOrderError, err.Error())
	}

	var responseData = make([]model.GetUnderMinimumOrderResponseStruct, 0)

	for _, d := range result {
		var data = model.GetUnderMinimumOrderResponseStruct{
			ItemId:       d.ItemID,
			ItemCode:     d.ItemCode,
			ItemName:     d.ItemName,
			VariantId:    d.VariantID,
			VariantName:  d.VariantName,
			UnitId:       d.UnitID,
			UnitName:     d.UnitName,
			MinimumStock: strconv.FormatInt(d.MinimumStock, 10),
			Amount:       strconv.FormatInt(d.Amount, 10),
		}
		responseData = append(responseData, data)
	}

	res := model.GetUnderMinimumOrderResponse{
		UnderMinimumOrder: responseData,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
