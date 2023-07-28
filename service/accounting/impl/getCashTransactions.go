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

func (a accountingService) GetCashTransactions(w http.ResponseWriter, r *http.Request) error {

	var req model.GetCashTransactionsRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetCashTransactions(context.Background(), db.GetCashTransactionsParams{
		CompanyID: req.CompanyId,
		BranchID:  req.BranchId,
		Type:      util.WildCardString(req.Type),
		StartDate: util.StringToDate(req.StartDate),
		EndDate:   util.StringToDate(req.EndDate),
	})
	if err != nil {
		return errors.NewServerError(model.GetCashTransactionsError, err.Error())
	}

	var cash_transactions = make([]model.CashTransaction, 0)

	for _, d := range result {
		var cash_transaction = model.CashTransaction{
			CompanyId:       d.CompanyID,
			BranchId:        d.BranchID,
			TransactionId:   d.ID,
			TransactionDate: d.TransactionDate.Format(util.DateLayoutYMD),
			Type:            d.Type,
			MainChartOfAccount: model.ChartOfAccountIdName{
				ChartOfAccountId: d.MainChartOfAccountID,
				AccountName:      d.MainChartOfAccountName,
			},
			ContraChartOfAccount: model.ChartOfAccountIdName{
				ChartOfAccountId: d.ContraChartOfAccountID,
				AccountName:      d.ContraChartOfAccountName.String,
			},
			Amount:      strconv.FormatInt(d.Amount, 10),
			Description: d.Description,
		}
		cash_transactions = append(cash_transactions, cash_transaction)
	}

	res := cash_transactions
	httpHandler.WriteResponse(w, res)

	return nil
}
