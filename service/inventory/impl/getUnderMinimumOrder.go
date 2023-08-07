package impl

import (
	"context"
	"net/http"
	"strconv"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/inventory/db/sqlc"
	"github.com/expert-pancake/service/inventory/impl/client"
	"github.com/expert-pancake/service/inventory/model"
)

func (a inventoryService) GetUnderMinimumOrder(w http.ResponseWriter, r *http.Request) error {

	var req model.GetUnderMinimumOrderRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetUnderMinimumOrder(context.Background(), db.GetUnderMinimumOrderParams{
		CompanyID: req.CompanyId,
		BranchID:  req.BranchId,
	})
	if err != nil {
		return errors.NewServerError(model.GetUnderMinimumOrderError, err.Error())
	}

	var responseData = make([]model.GetUnderMinimumOrderResponseStruct, 0)

	for _, d := range result {
		argWarehouse := client.GetWarehousesRequest{
			Id:       d.WarehouseID,
			BranchId: "1",
		}
		warehouse, err := client.GetWarehouses(argWarehouse)
		if err != nil {
			return errors.NewServerError(model.GetUnderMinimumOrderError, err.Error())
		}

		var data = model.GetUnderMinimumOrderResponseStruct{
			WarehouseName: warehouse.Result.Warehouses[0].Name,
			ItemId:        d.ItemID,
			ItemCode:      d.ItemCode,
			ItemName:      d.ItemName,
			VariantId:     d.VariantID,
			VariantName:   d.VariantName,
			UnitId:        d.UnitID,
			UnitName:      d.UnitName,
			MinimumStock:  strconv.FormatInt(d.MinimumStock, 10),
			Amount:        strconv.FormatInt(d.Amount, 10),
		}
		responseData = append(responseData, data)
	}

	res := model.GetUnderMinimumOrderResponse{
		UnderMinimumOrder: responseData,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
