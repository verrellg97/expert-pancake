package impl

import (
	"context"
	"net/http"
	"strconv"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/sales/impl/client"
	"github.com/expert-pancake/service/sales/model"
	"github.com/expert-pancake/service/sales/util"
)

func (a salesService) GetPOSItems(w http.ResponseWriter, r *http.Request) error {

	var req model.GetPOSItemsRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	var posItems = make([]model.POSItem, 0)
	detailResult, err := a.dbTrx.GetPOSItemsByPOSId(context.Background(), req.POSId)
	if err != nil {
		return errors.NewServerError(model.GetPOSError, err.Error())
	}

	for _, d := range detailResult {
		argItemVariant := client.GetItemVariantsRequest{
			Id: d.ItemVariantID,
		}
		itemVariant, err := client.GetItemVariants(argItemVariant)
		if err != nil {
			return errors.NewServerError(model.GetPOSError, err.Error())
		}

		argItemUnit := client.GetItemUnitsRequest{
			Id:     d.ItemUnitID,
			ItemId: itemVariant.Result.ItemVariants[0].ItemId,
		}
		itemUnit, err := client.GetItemUnits(argItemUnit)
		if err != nil {
			return errors.NewServerError(model.GetPOSError, err.Error())
		}

		argWarehouseRack := client.GetWarehouseRacksRequest{
			Id:          d.WarehouseRackID,
			WarehouseId: "1",
		}
		warehouseRack, err := client.GetWarehouseRacks(argWarehouseRack)
		if err != nil {
			return errors.NewServerError(model.GetPOSError, err.Error())
		}

		var batch, expiredDate *string
		if d.Batch.Valid {
			batch = &d.Batch.String
		}
		if d.ExpiredDate.Valid {
			expiredDate = new(string)
			*expiredDate = d.ExpiredDate.Time.Format(util.DateLayoutYMD)
		}

		var posItem = model.POSItem{
			DetailId:          d.ID,
			POSId:             d.ID,
			WarehouseRackId:   d.WarehouseRackID,
			WarehouseRackName: warehouseRack.Result.WarehouseRacks[0].Name,
			ItemVariantId:     d.ItemVariantID,
			ItemVariantName:   itemVariant.Result.ItemVariants[0].VariantName,
			ItemUnitId:        d.ItemUnitID,
			ItemUnitName:      itemUnit.Result.ItemUnits[0].UnitName,
			ItemUnitValue:     strconv.FormatInt(d.ItemUnitValue, 10),
			ItemCode:          itemVariant.Result.ItemVariants[0].Code,
			ItemName:          itemVariant.Result.ItemVariants[0].Name,
			Batch:             batch,
			ExpiredDate:       expiredDate,
			ItemBarcodeId:     d.ItemBarcodeID,
			Amount:            strconv.FormatInt(d.Amount, 10),
			Price:             strconv.FormatInt(d.Price, 10),
		}
		posItems = append(posItems, posItem)
	}

	res := model.GetPOSItemsResponse{
		POSItems: posItems,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
