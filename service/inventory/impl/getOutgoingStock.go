package impl

import (
	"context"
	"net/http"
	"strconv"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/inventory/db/sqlc"
	"github.com/expert-pancake/service/inventory/model"
	"github.com/expert-pancake/service/inventory/util"
)

func (a inventoryService) GetOutgoingStock(w http.ResponseWriter, r *http.Request) error {

	var req model.GetOutgoingStockRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}


	result, err := a.dbTrx.GetOutgoingStock(context.Background(), db.GetOutgoingStockParams{
		CompanyID: req.CompanyId,
		BranchID:  req.BranchId,
		StartDate: util.StringToDate(req.StartDate),
		EndDate: util.StringToDate(req.EndDate),
	})
	if err != nil {
		return errors.NewServerError(model.GetOutgoingStockError, err.Error())
	}

	var responseData = make([]model.GetOutgoingStockResponseStruct, 0)

	for _, d := range result {
		var data = model.GetOutgoingStockResponseStruct{
			ItemId:      d.ItemID,
			ItemCode:    d.ItemCode,
			ItemName:    d.ItemName,
			VariantId:   d.VariantID,
			VariantName: d.VariantName,
			UnitId:      d.UnitID,
			UnitName:    d.UnitName,
			Amount:      strconv.FormatInt(int64(d.Amount), 10),
		}
		responseData = append(responseData, data)
	}

	res := model.GetOutgoingStockResponse{
		OutgoingStock: responseData,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
