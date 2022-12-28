package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/accounting/db/sqlc"
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
		BankAccount: nil,
		CashAccount: nil,
	}

	resultBank, errBank := a.dbTrx.GetCompanySettingChartOfAccount(context.Background(),
		db.GetCompanySettingChartOfAccountParams{
			CompanyID:        req.CompanyId,
			AccountGroupName: "BANK",
		})
	if errBank == nil {
		resultBranches, err := a.dbTrx.GetChartOfAccountBranches(context.Background(), resultBank.ID)
		if err != nil {
			return errors.NewServerError(model.GetCompanyChartOfAccountsError, err.Error())
		}
		res.BankAccount = &model.ChartOfAccount{
			ChartOfAccountId:      resultBank.ID,
			CompanyId:             resultBank.CompanyID,
			CurrencyCode:          resultBank.CurrencyCode,
			ChartOfAccountGroupId: resultBank.ChartOfAccountGroupID,
			ReportType:            resultBank.ReportType,
			AccountType:           resultBank.AccountType,
			AccountGroup:          resultBank.AccountGroupName,
			AccountCode:           resultBank.AccountCode,
			AccountName:           resultBank.AccountName,
			BankName:              resultBank.BankName,
			BankAccountNumber:     resultBank.BankAccountNumber,
			BankCode:              resultBank.BankCode,
			IsAllBranches:         resultBank.IsAllBranches,
			Branches:              util.ChartOfAccountBranchDbToApi(resultBranches),
			IsDeleted:             resultBank.IsDeleted,
		}
	}

	resultCash, errCash := a.dbTrx.GetCompanySettingChartOfAccount(context.Background(),
		db.GetCompanySettingChartOfAccountParams{
			CompanyID:        req.CompanyId,
			AccountGroupName: "KAS",
		})
	if errCash == nil {
		resultBranches, err := a.dbTrx.GetChartOfAccountBranches(context.Background(), resultCash.ID)
		if err != nil {
			return errors.NewServerError(model.GetCompanyChartOfAccountsError, err.Error())
		}
		res.CashAccount = &model.ChartOfAccount{
			ChartOfAccountId:      resultCash.ID,
			CompanyId:             resultCash.CompanyID,
			CurrencyCode:          resultCash.CurrencyCode,
			ChartOfAccountGroupId: resultCash.ChartOfAccountGroupID,
			ReportType:            resultCash.ReportType,
			AccountType:           resultCash.AccountType,
			AccountGroup:          resultCash.AccountGroupName,
			AccountCode:           resultCash.AccountCode,
			AccountName:           resultCash.AccountName,
			BankName:              resultCash.BankName,
			BankAccountNumber:     resultCash.BankAccountNumber,
			BankCode:              resultCash.BankCode,
			IsAllBranches:         resultCash.IsAllBranches,
			Branches:              util.ChartOfAccountBranchDbToApi(resultBranches),
			IsDeleted:             resultCash.IsDeleted,
		}
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
