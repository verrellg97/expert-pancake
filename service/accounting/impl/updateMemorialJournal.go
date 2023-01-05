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

func (a accountingService) UpdateMemorialJournal(w http.ResponseWriter, r *http.Request) error {

	var req model.UpdateMemorialJournalRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.UpdateMemorialJournalTrxParams{
		Id:              req.MemorialJournalId,
		TransactionDate: util.StringToDate(req.TransactionDate),
		Description:     req.Description,
		ChartOfAccounts: req.ChartOfAccounts,
	}

	result, err := a.dbTrx.UpdateMemorialJournalTrx(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpdateMemorialJournalError, err.Error())
	}

	resultChartOfAccounts, err := a.dbTrx.GetMemorialJournalAccounts(context.Background(), result.MemorialJournalId)

	res := model.AddMemorialJournalResponse{
		MemorialJournal: model.MemorialJournal{
			MemorialJournalId: result.MemorialJournalId,
			CompanyId:         result.CompanyId,
			TransactionDate:   result.TransactionDate,
			Description:       result.Description,
			ChartOfAccounts:   util.MemorialJournalAccountDbToApi(resultChartOfAccounts),
		},
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
