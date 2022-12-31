package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/accounting/db/transaction"
	"github.com/expert-pancake/service/accounting/model"
	"github.com/expert-pancake/service/accounting/util"
)

func (a accountingService) AddJournalBook(w http.ResponseWriter, r *http.Request) error {

	var req model.AddJournalBookRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.CreateNewJournalBookTrxParams{
		CompanyId:       req.CompanyId,
		Name:            req.Name,
		StartPeriod:     util.StringToDate(req.StartPeriod),
		EndPeriod:       util.StringToDate(req.EndPeriod),
		ChartOfAccounts: req.ChartOfAccounts,
	}

	result, err := a.dbTrx.CreateNewJournalBookTrx(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.AddJournalBookError, err.Error())
	}

	resultChartOfAccounts, err := a.dbTrx.GetJournalBookAccounts(context.Background(), result.JournalBookId)

	res := model.AddJournalBookResponse{
		JournalBook: model.JournalBook{
			JournalBookId:   result.JournalBookId,
			CompanyId:       result.CompanyId,
			Name:            result.Name,
			StartPeriod:     result.StartPeriod,
			EndPeriod:       result.EndPeriod,
			IsClosed:        result.IsClosed,
			ChartOfAccounts: util.JournalBookAccountDbToApi(resultChartOfAccounts),
		},
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
