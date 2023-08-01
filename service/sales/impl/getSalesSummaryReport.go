package impl

import (
	"context"
	"net/http"
	"strconv"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/sales/db/sqlc"
	"github.com/expert-pancake/service/sales/impl/client"
	"github.com/expert-pancake/service/sales/model"
	"github.com/expert-pancake/service/sales/util"
)

func (a salesService) GetSalesSummaryReport(w http.ResponseWriter, r *http.Request) error {

	var req model.GetSalesSummaryReportRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetSalesSummaryReport(context.Background(), db.GetSalesSummaryReportParams{
		CompanyID: req.CompanyId,
		BranchID:  util.WildCardString(req.BranchId),
		StartDate: util.StringToDate(req.StartDate),
		EndDate:   util.StringToDate(req.EndDate),
	})
	if err != nil {
		return errors.NewServerError(model.GetSalesSummaryReportError, err.Error())
	}

	branches, err := client.GetCompanyBranches(
		client.GetCompanyBranchesRequest{
			CompanyId: req.CompanyId,
		})
	if err != nil {
		return err
	}
	branchMap := util.BranchApiToMap(branches.Result)

	var salesSummaries = make([]model.SalesSummary, 0)

	for _, d := range result {
		argContactBook := client.GetContactBooksRequest{
			Id:        d.ContactBookID,
			CompanyId: d.SecondaryCompanyID,
		}
		contactBook, err := client.GetContactBooks(argContactBook)
		if err != nil {
			return errors.NewServerError(model.GetSalesSummaryReportError, err.Error())
		}
		customerName := ""
		if len(contactBook.Result) > 0 {
			customerName = contactBook.Result[0].Name
		}

		var salesSummary = model.SalesSummary{
			TransactionCode:    d.FormNumber,
			TransactionDate:    d.TransactionDate.Format(util.DateLayoutYMD),
			BranchId:           d.BranchID,
			BranchName:         branchMap[d.BranchID].Name,
			ContactBookId:      d.ContactBookID,
			SecondaryCompanyId: d.SecondaryCompanyID,
			KonekinId:          d.KonekinID,
			CustomerName:       customerName,
			TotalItems:         strconv.FormatInt(d.TotalItems, 10),
			CurrencyCode:       d.CurrencyCode,
			Total:              strconv.FormatInt(d.Total, 10),
		}
		salesSummaries = append(salesSummaries, salesSummary)
	}

	res := model.GetSalesSummaryReportResponse{
		SalesSummaries: salesSummaries,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
