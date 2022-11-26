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
		StartMonth: int32(req.StartMonth),
		StartYear:  int32(req.StartYear),
		EndMonth:   int32(req.EndMonth),
		EndYear:    int32(req.EndYear),
	}

	result, err := a.dbTrx.UpsertCompanyFiscalYear(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpsertCompanyFiscalYearError, err.Error())
	}

	res := model.UpsertCompanyFiscalYearRequestResponse{
		CompanyId:  result.CompanyID,
		StartMonth: int(result.StartMonth),
		StartYear:  int(result.StartYear),
		EndMonth:   int(result.EndMonth),
		EndYear:    int(result.EndYear),
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
