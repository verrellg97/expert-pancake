package impl

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/warehouse/db/sqlc"
	"github.com/expert-pancake/service/warehouse/model"
	uuid "github.com/satori/go.uuid"
)

func (a warehouseService) UpsertWarehouse(w http.ResponseWriter, r *http.Request) error {

	var req model.UpsertWarehouseRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	var id = ""
	if req.WarehouseId == "" {
		id = uuid.NewV4().String()
	} else {
		id = req.WarehouseId
	}

	arg := db.UpsertWarehouseParams{
		ID:       id,
		BranchID: req.BranchId,
		Code:     "WH-" + fmt.Sprintf("%08d", rand.Intn(100000000)),
		Name:     req.Name,
		Address:  req.Address,
		Type:     req.Type,
	}

	result, err := a.dbTrx.UpsertWarehouse(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpsertWarehouseError, err.Error())
	}

	res := model.UpsertWarehouseResponse{
		Warehouse: model.Warehouse{
			WarehouseId: result.ID,
			BranchId:    result.BranchID,
			Code:        result.Code,
			Name:        result.Name,
			Address:     result.Address,
			Type:        result.Type,
		},
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
