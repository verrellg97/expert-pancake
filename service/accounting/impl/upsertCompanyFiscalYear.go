package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/accounting/db/sqlc"
	"github.com/expert-pancake/service/accounting/model"
)

func (a accountingService) UpsertCompanyFiscalYear(w http.ResponseWriter, r *http.Request) error {

	var req model.UpsertCompanyFiscalYearRequestResponse

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.UpsertCompanyFiscalYearParams{
		CompanyID:  req.CompanyId,
		StartMonth: req.StartMonth,
		StartYear:  req.StartYear,
		EndMonth:   req.EndMonth,
		EndYear:    req.EndYear,
	}

	result, err := a.dbTrx.UpsertCompanyFiscalYear(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpsertCompanyFiscalYearError, err.Error())
	}

	res := model.UpsertCompanyFiscalYearRequestResponse{
		CompanyId:  result.CompanyID,
		StartMonth: result.StartMonth,
		StartYear:  result.StartYear,
		EndMonth:   result.EndMonth,
		EndYear:    result.EndYear,
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
