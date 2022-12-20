package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/accounting/db/transaction"
	"github.com/expert-pancake/service/accounting/model"
	uuid "github.com/satori/go.uuid"
)

func (a accountingService) AddCompanyChartOfAccount(w http.ResponseWriter, r *http.Request) error {

	var req model.AddCompanyChartOfAccountRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	var branches = make([]string, 0)
	if req.Branches != nil {
		branches = req.Branches
	}

	arg := db.CreateNewChartOfAccountTrxParams{
		ID:                uuid.NewV4().String(),
		CompanyID:         req.CompanyId,
		CurrencyCode:      req.CurrencyCode,
		ReportType:        req.ReportType,
		AccountType:       req.AccountType,
		AccountGroup:      req.AccountGroup,
		AccountCode:       req.AccountCode,
		AccountName:       req.AccountName,
		BankName:          req.BankName,
		BankAccountNumber: req.BankAccountNumber,
		BankCode:          req.BankCode,
		IsAllBranches:     req.IsAllBranches,
		Branches:          branches,
	}

	result, err := a.dbTrx.CreateNewChartOfAccountTrx(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.AddCompanyChartOfAccountError, err.Error())
	}

	res := model.UpsertCompanyChartOfAccountResponse{
		ChartOfAccount: model.ChartOfAccount{
			ChartOfAccountId:  result.ChartOfAccountId,
			CompanyId:         result.CompanyId,
			CurrencyCode:      result.CurrencyCode,
			ReportType:        result.ReportType,
			AccountType:       result.AccountType,
			AccountGroup:      result.AccountGroup,
			AccountCode:       result.AccountCode,
			AccountName:       result.AccountName,
			BankName:          result.BankName,
			BankAccountNumber: result.BankAccountNumber,
			BankCode:          result.BankCode,
			IsAllBranches:     result.IsAllBranches,
			Branches:          result.Branches,
			IsDeleted:         result.IsDeleted,
		},
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
