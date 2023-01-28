package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/warehouse/db/sqlc"
	"github.com/expert-pancake/service/warehouse/model"
	uuid "github.com/satori/go.uuid"
)

func (a warehouseService) UpsertRack(w http.ResponseWriter, r *http.Request) error {

	var req model.UpsertRackRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	var id = "";
	if req.Id == "" {
		id = uuid.NewV4().String()
	} else {
		id = req.Id
	}

	arg := db.UpsertRackParams{
		ID:       id,
		BranchID: req.BranchId,
		Name:     req.Name,
	}

	result, err := a.dbTrx.UpsertRack(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpsertRackError, err.Error())
	}

	res := model.UpsertRackResponse{
		Rack: model.Rack{
			RackId:    result.ID,
			BranchId:       result.BranchID,
			Name:           result.Name,
		},
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
