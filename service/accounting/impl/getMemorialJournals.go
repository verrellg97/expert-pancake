package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/accounting/model"
	"github.com/expert-pancake/service/accounting/util"
)

func (a accountingService) GetMemorialJournals(w http.ResponseWriter, r *http.Request) error {

	var req model.GetMemorialJournalsRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetMemorialJournals(context.Background(), req.CompanyId)
	if err != nil {
		return errors.NewServerError(model.GetMemorialJournalsError, err.Error())
	}

	var journal_books = make([]model.MemorialJournal, 0)

	for _, d := range result {
		resultChartOfAccounts, err := a.dbTrx.GetMemorialJournalAccounts(context.Background(), d.ID)
		if err != nil {
			return errors.NewServerError(model.GetMemorialJournalsError, err.Error())
		}
		var journal_book = model.MemorialJournal{
			MemorialJournalId: d.ID,
			CompanyId:         d.CompanyID,
			TransactionDate:   d.TransactionDate.Format(util.DateLayoutYMD),
			Description:       d.Description,
			ChartOfAccounts:   util.MemorialJournalAccountDbToApi(resultChartOfAccounts),
		}
		journal_books = append(journal_books, journal_book)
	}

	res := journal_books
	httpHandler.WriteResponse(w, res)

	return nil
}
