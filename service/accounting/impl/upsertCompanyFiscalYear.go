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

func (a accountingService) UpsertCompanyFiscalYear(w http.ResponseWriter, r *http.Request) error {

	var req model.UpsertCompanyFiscalYearRequestResponse

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.UpsertCompanyFiscalYearParams{
		CompanyID:   req.CompanyId,
		StartPeriod: util.StringToDate(req.StartPeriod),
		EndPeriod:   util.StringToDate(req.EndPeriod),
	}

	result, err := a.dbTrx.UpsertCompanyFiscalYear(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpsertCompanyFiscalYearError, err.Error())
	}

	res := model.UpsertCompanyFiscalYearRequestResponse{
		CompanyId:   result.CompanyID,
		StartPeriod: result.StartPeriod.Format(util.DateLayoutYMD),
		EndPeriod:   result.EndPeriod.Format(util.DateLayoutYMD),
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
