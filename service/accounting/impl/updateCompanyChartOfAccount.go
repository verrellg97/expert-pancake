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

func (a accountingService) UpdateCompanyChartOfAccount(w http.ResponseWriter, r *http.Request) error {

	var req model.UpdateCompanyChartOfAccountRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.UpdateCompanyChartOfAccountParams{
		ID:                req.ChartOfAccountId,
		AccountCode:       req.AccountCode,
		AccountName:       req.AccountName,
		AccountGroup:      req.AccountGroup,
		BankName:          req.BankName,
		BankAccountNumber: req.BankAccountNumber,
		BankCode:          req.BankCode,
		OpeningBalance:    util.StringToBigInt(req.OpeningBalance),
		IsDeleted:         req.IsDeleted,
	}

	result, err := a.dbTrx.UpdateCompanyChartOfAccount(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpdateCompanyChartOfAccountError, err.Error())
	}

	res := model.UpsertCompanyChartOfAccountResponse{
		ChartOfAccount: model.ChartOfAccount{
			ChartOfAccountId:  result.ID,
			CompanyId:         result.CompanyID,
			BranchId:          result.BranchID,
			AccountCode:       result.AccountCode,
			AccountName:       result.AccountName,
			AccountGroup:      result.AccountGroup,
			BankName:          result.BankName,
			BankAccountNumber: result.BankAccountNumber,
			BankCode:          result.BankCode,
			OpeningBalance:    util.BigIntToString(result.OpeningBalance),
			IsDeleted:         result.IsDeleted,
		},
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
