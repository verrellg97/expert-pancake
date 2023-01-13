package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/business-relation/db/transaction"
	"github.com/expert-pancake/service/business-relation/model"
)

func (a businessRelationService) AddDefaultContactBook(w http.ResponseWriter, r *http.Request) error {

	var req model.AddDefaultContactBookRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.AddDefaultContactBookTrxParams{
		CompanyId:   req.CompanyId,
		CompanyName: req.CompanyName,
	}

	err := a.dbTrx.AddDefaultContactBookTrx(context.Background(), arg)

	if err != nil {
		return errors.NewServerError(model.AddDefaultContactBookError, err.Error())
	}

	res := model.AddDefaultContactBookResponse{
		Message: "OK",
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
