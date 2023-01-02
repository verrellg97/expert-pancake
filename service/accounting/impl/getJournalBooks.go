package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/accounting/model"
	"github.com/expert-pancake/service/accounting/util"
)

func (a accountingService) GetJournalBooks(w http.ResponseWriter, r *http.Request) error {

	var req model.GetJournalBooksRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetJournalBooks(context.Background(), req.CompanyId)
	if err != nil {
		return errors.NewServerError(model.GetJournalBooksError, err.Error())
	}

	var journal_books = make([]model.JournalBook, 0)

	for _, d := range result {
		resultChartOfAccounts, err := a.dbTrx.GetJournalBookAccounts(context.Background(), d.ID)
		if err != nil {
			return errors.NewServerError(model.GetJournalBooksError, err.Error())
		}
		var journal_book = model.JournalBook{
			JournalBookId:   d.ID,
			CompanyId:       d.CompanyID,
			Name:            d.Name,
			StartPeriod:     d.StartPeriod.Format(util.DateLayoutYMD),
			EndPeriod:       d.EndPeriod.Format(util.DateLayoutYMD),
			IsClosed:        d.IsClosed,
			ChartOfAccounts: util.JournalBookAccountDbToApi(resultChartOfAccounts),
		}
		journal_books = append(journal_books, journal_book)
	}

	res := journal_books
	httpHandler.WriteResponse(w, res)

	return nil
}
