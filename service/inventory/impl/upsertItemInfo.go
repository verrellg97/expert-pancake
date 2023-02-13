package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/inventory/db/sqlc"
	"github.com/expert-pancake/service/inventory/impl/client"
	"github.com/expert-pancake/service/inventory/model"
)

func (a inventoryService) UpsertItemInfo(w http.ResponseWriter, r *http.Request) error {

	var req model.UpsertItemInfoRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.UpsertItemInfoParams{
		ItemID:                   req.ItemId,
		IsPurchase:               req.IsPurchase,
		IsSale:                   req.IsSale,
		IsRawMaterial:            req.IsRawMaterial,
		IsAsset:                  req.IsAsset,
		PurchaseChartOfAccountID: req.PurchaseChartOfAccountId,
		SaleChartOfAccountID:     req.SaleChartOfAccountId,
		PurchaseItemUnitID:       req.PurchaseItemUnitId,
	}

	err := a.dbTrx.UpsertItemInfo(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpsertItemInfoError, err.Error())
	}

	result, err := a.dbTrx.GetItemInfo(context.Background(), req.ItemId)
	if err != nil {
		return errors.NewServerError(model.UpsertItemInfoError, err.Error())
	}

	argPurchase := client.GetCompanyChartOfAccountsRequest{
		CompanyId: result.CompanyID,
		Id:        result.PurchaseChartOfAccountID,
	}
	purchaseCoa, err := client.GetCompanyChartOfAccounts(argPurchase)
	if err != nil {
		return errors.NewServerError(model.UpsertItemInfoError, err.Error())
	}

	argSale := client.GetCompanyChartOfAccountsRequest{
		CompanyId: result.CompanyID,
		Id:        result.SaleChartOfAccountID,
	}
	saleCoa, err := client.GetCompanyChartOfAccounts(argSale)
	if err != nil {
		return errors.NewServerError(model.UpsertItemInfoError, err.Error())
	}

	res := model.UpsertItemInfoResponse{
		ItemInfo: model.ItemInfo{
			ItemId:                     result.ItemID,
			IsPurchase:                 result.IsPurchase,
			IsSale:                     result.IsSale,
			IsRawMaterial:              result.IsRawMaterial,
			IsAsset:                    result.IsAsset,
			PurchaseChartOfAccountId:   result.PurchaseChartOfAccountID,
			PurchaseChartOfAccountName: purchaseCoa.Result[0].AccountName,
			SaleChartOfAccountId:       result.SaleChartOfAccountID,
			SaleChartOfAccountName:     saleCoa.Result[0].AccountName,
			PurchaseItemUnitId:         result.PurchaseItemUnitID,
			PurchaseItemUnitName:       result.PurchaseItemUnitName,
		},
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
