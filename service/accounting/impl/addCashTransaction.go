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
	uuid "github.com/satori/go.uuid"
)

func (a accountingService) AddCashTransaction(w http.ResponseWriter, r *http.Request) error {

	var req model.AddCashTransactionRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	amount, _ := strconv.ParseInt(req.Amount, 10, 64)
	arg := db.InsertCashTransactionParams{
		ID:                     uuid.NewV4().String(),
		CompanyID:              req.CompanyId,
		BranchID:               req.BranchId,
		TransactionDate:        util.StringToDate(req.TransactionDate),
		Type:                   req.Type,
		MainChartOfAccountID:   req.MainChartOfAccountId,
		ContraChartOfAccountID: req.ContraChartOfAccountId,
		Amount:                 amount,
		Description:            req.Description,
	}

	result, err := a.dbTrx.CreateNewCashTransactionTrx(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.CreateNewCashTransactionError, err.Error())
	}

	resultMain, _ := a.dbTrx.GetCompanyChartOfAccount(context.Background(), req.MainChartOfAccountId)
	resMainChartOfAccount := model.ChartOfAccountIdName{
		ChartOfAccountId: resultMain.ID,
		AccountName:      resultMain.AccountName,
	}

	resultContra, _ := a.dbTrx.GetCompanyChartOfAccount(context.Background(), req.ContraChartOfAccountId)
	resContraChartOfAccount := model.ChartOfAccountIdName{
		ChartOfAccountId: resultContra.ID,
		AccountName:      resultContra.AccountName,
	}

	res := model.AddCashTransactionResponse{
		CashTransaction: model.CashTransaction{
			CompanyId:            result.CompanyId,
			BranchId:             result.BranchId,
			TransactionId:        result.TransactionId,
			TransactionDate:      result.TransactionDate,
			Type:                 result.Type,
			MainChartOfAccount:   resMainChartOfAccount,
			ContraChartOfAccount: resContraChartOfAccount,
			Amount:               result.Amount,
			Description:          result.Description,
		},
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
