package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/accounting/model"
)

func (a accountingService) CloseJournalBook(w http.ResponseWriter, r *http.Request) error {

	var req model.CloseJournalBookRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	res, err := a.dbTrx.CloseJournalBookTrx(context.Background(), req.JournalBookId)
	if err != nil {
		return errors.NewServerError(model.CloseJournalBookError, err.Error())
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
