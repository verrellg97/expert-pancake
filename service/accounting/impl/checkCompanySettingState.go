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

func (a accountingService) CheckCompanySettingState(w http.ResponseWriter, r *http.Request) error {

	var req model.CheckCompanySettingStateRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	_, errCOA := a.dbTrx.GetCompanyChartOfAccount(context.Background(), req.CompanyId)
	if errCOA != nil {
		arg := db.AddDefaultCompanyChartOfAccountTrxParams{
			CompanyId: req.CompanyId,
		}
		_ = a.dbTrx.AddDefaultCompanyChartOfAccountTransactionTrx(context.Background(), arg)
	}

	res := model.CheckCompanySettingStateResponse{
		FiscalYear:  nil,
		BankAccount: nil,
		CashAccount: nil,
	}

	resultFiscal, errFiscal := a.dbTrx.GetCompanySettingFiscalYear(context.Background(), req.CompanyId)
	if errFiscal == nil {
		res.FiscalYear = &model.FiscalYear{
			CompanyId:   resultFiscal.CompanyID,
			StartPeriod: resultFiscal.StartPeriod.Format(util.DateLayoutYMD),
			EndPeriod:   resultFiscal.EndPeriod.Format(util.DateLayoutYMD),
		}
	}

	resultBank, errBank := a.dbTrx.GetCompanySettingBank(context.Background(), req.CompanyId)
	if errBank == nil {
		res.BankAccount = &model.ChartOfAccount{
			ChartOfAccountId:  resultBank.ID,
			CompanyId:         resultBank.CompanyID,
			AccountCode:       resultBank.AccountCode,
			AccountName:       resultBank.AccountName,
			BankName:          resultBank.BankName,
			BankAccountNumber: resultBank.BankAccountNumber,
			BankCode:          resultBank.BankCode,
			IsDeleted:         resultBank.IsDeleted,
		}
	}

	resultCash, errCash := a.dbTrx.GetCompanySettingCash(context.Background(), req.CompanyId)
	if errCash == nil {
		res.CashAccount = &model.ChartOfAccount{
			ChartOfAccountId:  resultCash.ID,
			CompanyId:         resultCash.CompanyID,
			AccountCode:       resultCash.AccountCode,
			AccountName:       resultCash.AccountName,
			BankName:          resultCash.BankName,
			BankAccountNumber: resultCash.BankAccountNumber,
			BankCode:          resultCash.BankCode,
			IsDeleted:         resultCash.IsDeleted,
		}
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
