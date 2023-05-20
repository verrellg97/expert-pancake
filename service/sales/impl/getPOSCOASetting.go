package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/sales/impl/client"
	"github.com/expert-pancake/service/sales/model"
)

func (a salesService) GetPOSCOASetting(w http.ResponseWriter, r *http.Request) error {

	var req model.GetPOSCOASettingRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	var posCOAs = make([]model.POSCOA, 0)
	detailResult, err := a.dbTrx.GetPOSCOASetting(context.Background(), req.BranchId)
	if err != nil {
		return errors.NewServerError(model.GetPOSCOASettingError, err.Error())
	}

	for _, d := range detailResult {
		argCOA := client.GetCompanyChartOfAccountsRequest{
			CompanyId: "1",
			Id:        d.ChartOfAccountID,
		}
		coa, err := client.GetCompanyChartOfAccounts(argCOA)
		if err != nil {
			return errors.NewServerError(model.GetPOSCOASettingError, err.Error())
		}

		if len(coa.Result) > 0 {
			coaName := coa.Result[0].AccountName
			var posCOA = model.POSCOA{
				ChartOfAccountId:   d.ChartOfAccountID,
				ChartOfAccountName: coaName,
			}
			posCOAs = append(posCOAs, posCOA)
		}
	}

	res := model.GetPOSCOASettingResponse{
		POSCOAs: posCOAs,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
