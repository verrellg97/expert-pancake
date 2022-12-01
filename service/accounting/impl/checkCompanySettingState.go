package impl

import (
	"context"
	"net/http"
	"strconv"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
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
			BranchId:          resultBank.BranchID,
			AccountCode:       resultBank.AccountCode,
			AccountName:       resultBank.AccountName,
			AccountGroup:      resultBank.AccountGroup,
			BankName:          resultBank.BankName,
			BankAccountNumber: resultBank.BankAccountNumber,
			BankCode:          resultBank.BankCode,
			OpeningBalance:    strconv.FormatInt(resultBank.OpeningBalance, 10),
			IsDeleted:         resultBank.IsDeleted,
		}
	}

	resultCash, errCash := a.dbTrx.GetCompanySettingCash(context.Background(), req.CompanyId)
	if errCash == nil {
		res.CashAccount = &model.ChartOfAccount{
			ChartOfAccountId:  resultCash.ID,
			CompanyId:         resultCash.CompanyID,
			BranchId:          resultCash.BranchID,
			AccountCode:       resultCash.AccountCode,
			AccountName:       resultCash.AccountName,
			AccountGroup:      resultCash.AccountGroup,
			BankName:          resultCash.BankName,
			BankAccountNumber: resultCash.BankAccountNumber,
			BankCode:          resultCash.BankCode,
			OpeningBalance:    strconv.FormatInt(resultCash.OpeningBalance, 10),
			IsDeleted:         resultCash.IsDeleted,
		}
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
