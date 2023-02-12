package impl

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/inventory/impl/client"
	"github.com/expert-pancake/service/inventory/model"
)

func (a inventoryService) GetItemInfo(w http.ResponseWriter, r *http.Request) error {

	var req model.GetItemInfoRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	purchaseCoaName := ""
	saleCoaName := ""

	result, err := a.dbTrx.GetItemInfo(context.Background(), req.ItemId)
	if err != nil && err != sql.ErrNoRows {
		return errors.NewServerError(model.GetItemInfoError, err.Error())
	}

	if err == nil {
		argPurchase := client.GetCompanyChartOfAccountsRequest{
			CompanyId: result.CompanyID,
			Id:        result.PurchaseChartOfAccountID,
		}
		purchaseCoa, err := client.GetCompanyChartOfAccounts(argPurchase)
		if err != nil {
			return errors.NewServerError(model.GetItemInfoError, err.Error())
		}
		purchaseCoaName = purchaseCoa.Result[0].AccountName

		argSale := client.GetCompanyChartOfAccountsRequest{
			CompanyId: result.CompanyID,
			Id:        result.SaleChartOfAccountID,
		}
		saleCoa, err := client.GetCompanyChartOfAccounts(argSale)
		if err != nil {
			return errors.NewServerError(model.GetItemInfoError, err.Error())
		}
		saleCoaName = saleCoa.Result[0].AccountName
	}

	res := model.GetItemInfoResponse{
		ItemInfo: model.ItemInfo{
			ItemId:                     result.ItemID,
			IsPurchase:                 result.IsPurchase,
			IsSale:                     result.IsSale,
			IsRawMaterial:              result.IsRawMaterial,
			IsAsset:                    result.IsAsset,
			PurchaseChartOfAccountId:   result.PurchaseChartOfAccountID,
			PurchaseChartOfAccountName: purchaseCoaName,
			SaleChartOfAccountId:       result.SaleChartOfAccountID,
			SaleChartOfAccountName:     saleCoaName,
			PurchaseItemUnitId:         result.PurchaseItemUnitID,
			PurchaseItemUnitName:       result.PurchaseItemUnitName,
		},
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
