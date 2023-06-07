package impl

import (
	"context"
	"net/http"
	"strconv"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/sales/db/sqlc"
	"github.com/expert-pancake/service/sales/impl/client"
	"github.com/expert-pancake/service/sales/model"
	"github.com/expert-pancake/service/sales/util"
)

func (a salesService) GetPOS(w http.ResponseWriter, r *http.Request) error {

	var req model.GetPOSRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	headerResult, err := a.dbTrx.GetPOS(context.Background(), db.GetPOSParams{
		CompanyID: req.CompanyId,
		BranchID:  req.BranchId,
		StartDate: util.StringToDate(req.StartDate),
		EndDate:   util.StringToDate(req.EndDate),
	})
	if err != nil {
		return errors.NewServerError(model.GetPOSError, err.Error())
	}

	var point_of_sales = make([]model.POS, 0)
	for _, d := range headerResult {
		argContactBook := client.GetContactBooksRequest{
			Id:        d.ContactBookID,
			CompanyId: d.CompanyID,
		}
		contactBook, err := client.GetContactBooks(argContactBook)
		if err != nil {
			return errors.NewServerError(model.GetPOSError, err.Error())
		}

		argWarehouse := client.GetWarehousesRequest{
			Id:       d.WarehouseID,
			BranchId: "1",
		}
		warehouse, err := client.GetWarehouses(argWarehouse)
		if err != nil {
			return errors.NewServerError(model.GetPOSError, err.Error())
		}

		var pos = model.POS{
			Id:                   d.ID,
			CompanyId:            d.CompanyID,
			BranchId:             d.BranchID,
			WarehouseId:          d.WarehouseID,
			WarehouseName:        warehouse.Result.Warehouses[0].Name,
			FormNumber:           d.FormNumber,
			TransactionDate:      d.TransactionDate.Format(util.DateLayoutYMD),
			ContactBookId:        d.ContactBookID,
			ContactBookName:      contactBook.Result[0].Name,
			SecondaryCompanyId:   d.SecondaryCompanyID,
			KonekinId:            d.KonekinID,
			CurrencyCode:         d.CurrencyCode,
			POSPaymentMethodId:   d.PosPaymentMethodID,
			POSPaymentMethodName: d.PosPaymentMethodName,
			TotalItems:           strconv.FormatInt(d.TotalItems, 10),
			Total:                strconv.FormatInt(d.Total, 10),
		}
		point_of_sales = append(point_of_sales, pos)
	}

	res := model.GetPOSResponse{
		POS: point_of_sales,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
