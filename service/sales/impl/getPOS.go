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

		argCOA := client.GetCompanyChartOfAccountsRequest{
			CompanyId: d.CompanyID,
			Id:        d.ChartOfAccountID,
		}
		coa, err := client.GetCompanyChartOfAccounts(argCOA)
		if err != nil {
			return errors.NewServerError(model.GetPOSError, err.Error())
		}

		var posItems = make([]model.POSItem, 0)
		detailResult, err := a.dbTrx.GetPOSItemsByPOSId(context.Background(), d.ID)
		if err != nil {
			return errors.NewServerError(model.GetPOSError, err.Error())
		}

		for _, dd := range detailResult {
			argItemVariant := client.GetItemVariantsRequest{
				Id: dd.ItemVariantID,
			}
			itemVariant, err := client.GetItemVariants(argItemVariant)
			if err != nil {
				return errors.NewServerError(model.GetPOSError, err.Error())
			}

			argItemUnit := client.GetItemUnitsRequest{
				Id:     dd.ItemUnitID,
				ItemId: itemVariant.Result.ItemVariants[0].ItemId,
			}
			itemUnit, err := client.GetItemUnits(argItemUnit)
			if err != nil {
				return errors.NewServerError(model.GetPOSError, err.Error())
			}

			argWarehouseRack := client.GetWarehouseRacksRequest{
				Id:          dd.WarehouseRackID,
				WarehouseId: "1",
			}
			warehouseRack, err := client.GetWarehouseRacks(argWarehouseRack)
			if err != nil {
				return errors.NewServerError(model.GetPOSError, err.Error())
			}

			var batch, expiredDate *string
			if dd.Batch.Valid {
				batch = &dd.Batch.String
			}
			if dd.ExpiredDate.Valid {
				expiredDate = new(string)
				*expiredDate = dd.ExpiredDate.Time.Format(util.DateLayoutYMD)
			}

			var posItem = model.POSItem{
				DetailId:          dd.ID,
				POSId:             d.ID,
				WarehouseRackId:   dd.WarehouseRackID,
				WarehouseRackName: warehouseRack.Result.WarehouseRacks[0].Name,
				ItemVariantId:     dd.ItemVariantID,
				ItemVariantName:   itemVariant.Result.ItemVariants[0].VariantName,
				ItemUnitId:        dd.ItemUnitID,
				ItemUnitName:      itemUnit.Result.ItemUnits[0].UnitName,
				ItemUnitValue:     strconv.FormatInt(dd.ItemUnitValue, 10),
				ItemCode:          itemVariant.Result.ItemVariants[0].Code,
				ItemName:          itemVariant.Result.ItemVariants[0].Name,
				Batch:             batch,
				ExpiredDate:       expiredDate,
				ItemBarcodeId:     dd.ItemBarcodeID,
				Amount:            strconv.FormatInt(dd.Amount, 10),
				Price:             strconv.FormatInt(dd.Price, 10),
			}
			posItems = append(posItems, posItem)
		}

		var pos = model.POS{
			Id:                 d.ID,
			CompanyId:          d.CompanyID,
			BranchId:           d.BranchID,
			WarehouseId:        d.WarehouseID,
			WarehouseName:      warehouse.Result.Warehouses[0].Name,
			FormNumber:         d.FormNumber,
			TransactionDate:    d.TransactionDate.Format(util.DateLayoutYMD),
			ContactBookId:      d.ContactBookID,
			ContactBookName:    contactBook.Result[0].Name,
			SecondaryCompanyId: d.SecondaryCompanyID,
			KonekinId:          d.KonekinID,
			CurrencyCode:       d.CurrencyCode,
			ChartOfAccountId:   d.ChartOfAccountID,
			ChartOfAccountName: coa.Result[0].AccountName,
			TotalItems:         strconv.FormatInt(d.TotalItems, 10),
			Total:              strconv.FormatInt(d.Total, 10),
			PosItems:           posItems,
		}
		point_of_sales = append(point_of_sales, pos)
	}

	res := model.GetPOSResponse{
		POS: point_of_sales,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
