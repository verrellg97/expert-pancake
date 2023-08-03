package impl

import (
	"context"
	"net/http"
	"strconv"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/sales/db/sqlc"
	"github.com/expert-pancake/service/sales/model"
	"github.com/expert-pancake/service/sales/util"
)

func (a salesService) GetMonthlyGrossSales(w http.ResponseWriter, r *http.Request) error {

	var req model.GetMonthlyGrossSalesRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetMonthlyGrossSales(context.Background(), db.GetMonthlyGrossSalesParams{
		CompanyID: req.CompanyId,
		BranchID:  util.WildCardString(req.BranchId),
		StartDate: util.StringToDate(req.StartDate),
		EndDate:   util.StringToDate(req.EndDate),
	})
	if err != nil {
		return errors.NewServerError(model.GetMonthlyGrossSalesError, err.Error())
	}

	var monthlyGrossSales = make([]model.MonthlyGrossSale, 0)

	for _, d := range result {

		var monthlyGrossSale = model.MonthlyGrossSale{
			Month: strconv.FormatInt(d.MonthNumber, 10),
			Year:  strconv.FormatInt(d.YearNumber, 10),
			Total: strconv.FormatInt(d.Total, 10),
		}
		monthlyGrossSales = append(monthlyGrossSales, monthlyGrossSale)
	}

	res := model.GetMonthlyGrossSalesResponse{
		MonthlyGrossSales: monthlyGrossSales,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
