package impl

import (
	"context"
	"net/http"
	"strconv"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/accounting/db/sqlc"
	"github.com/expert-pancake/service/accounting/model"
	"github.com/expert-pancake/service/accounting/util"
)

func (a accountingService) GetCashTransactionsGroupByDate(w http.ResponseWriter, r *http.Request) error {

	var req model.GetCashTransactionsGroupByDateRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetCashTransactionsGroupByDate(context.Background(), db.GetCashTransactionsGroupByDateParams{
		CompanyID: req.CompanyId,
		BranchID:  req.BranchId,
		Type:      util.WildCardString(req.Type),
		StartDate: util.StringToDate(req.StartDate),
		EndDate:   util.StringToDate(req.EndDate),
	})
	if err != nil {
		return errors.NewServerError(model.GetCashTransactionsGroupByDateError, err.Error())
	}

	var cash_transactions = make([]model.GetCashTransactionsGroupByDateResponse, 0)

	for _, d := range result {
		var cash_transaction = model.GetCashTransactionsGroupByDateResponse{
			TransactionDate: d.TransactionDate.Format(util.DateLayoutYMD),
			Amount: model.CashInCashOut{
				CashIn:  strconv.FormatInt(d.CashIn, 10),
				CashOut: strconv.FormatInt(d.CashOut, 10),
			},
		}
		cash_transactions = append(cash_transactions, cash_transaction)
	}

	res := cash_transactions
	httpHandler.WriteResponse(w, res)

	return nil
}
