package impl

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/inventory/db/sqlc"
	"github.com/expert-pancake/service/inventory/model"
)

func (a inventoryService) GetItemReorderNotifications(w http.ResponseWriter, r *http.Request) error {

	var req model.GetItemReorderNotificationsRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetItemReorderNotifications(context.Background(), db.GetItemReorderNotificationsParams{
		WarehouseID:    req.WarehouseId,
		ItemVariantIds: req.VariantIds,
	})
	log.Println(err)
	if err != nil {
		return errors.NewServerError(model.GetItemReorderNotificationsError, err.Error())
	}

	var datas = make([]model.ItemReorderNotification, 0)

	for _, d := range result {
		var minimum_stock *string
		if d.ItemReorderID.Valid {
			minimum_stock = new(string)
			*minimum_stock = strconv.FormatInt(d.MinimumStock, 10)
		}
		var data = model.ItemReorderNotification{
			ItemId:       d.ItemID,
			ItemName:     d.ItemName,
			VariantId:    d.VariantID,
			VariantName:  d.VariantName,
			CurrentStock: strconv.FormatInt(d.CurrentStock, 10),
			MinimumStock: minimum_stock,
		}
		datas = append(datas, data)
	}

	res := model.GetItemReorderNotificationsResponse{
		ItemVariants: datas,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
