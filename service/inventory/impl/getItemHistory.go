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

func (a inventoryService) GetItemHistory(w http.ResponseWriter, r *http.Request) error {

	var req model.GetItemHistoryRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetItemHistory(context.Background(), db.GetItemHistoryParams{
		BranchID:  req.BranchId,
		ID:        req.ItemId,
		VariantID: util.WildCardString(req.VariantId),
		StartDate: util.StringToDate(req.StartDate),
		EndDate:   util.StringToDate(req.EndDate),
	})
	if err != nil {
		return errors.NewServerError(model.GetItemHistoryError, err.Error())
	}

	var datas = make([]model.ItemHistory, 0)

	for _, d := range result {
		var data = model.ItemHistory{
			TransactionCode: d.TransactionCode,
			TransactionDate: d.TransactionDate.Format(util.DateLayoutYMD),
			ItemId:          d.ItemID,
			ItemCode:        d.ItemCode,
			ItemName:        d.ItemName,
			VariantId:       d.VariantID,
			VariantName:     d.VariantName,
			UnitId:          d.UnitID,
			UnitName:        d.UnitName,
			Amount:          strconv.FormatInt(d.Amount, 10),
		}
		datas = append(datas, data)
	}

	res := model.GetItemHistoryResponse{
		ItemHistories: datas,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
