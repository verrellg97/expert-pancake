package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/business/db/sqlc"
	"github.com/expert-pancake/service/business/impl/client"
	"github.com/expert-pancake/service/business/model"
	"github.com/expert-pancake/service/business/util"
)

func (a businessService) UpdateCompany(w http.ResponseWriter, r *http.Request) error {

	var req model.UpdateCompanyRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	resultId, err := a.dbTrx.GetCompanyByName(context.Background(), req.Name)
	if err == nil && resultId != req.CompanyId {
		var errRes = errors.NewClientError()
		errRes.Code = model.CompanyUniqueNameError
		errRes.Message = model.CompanyUniqueNameErrorMessage
		return errRes
	}

	companyData, err := a.dbTrx.GetCompanyById(context.Background(), req.CompanyId)

	branchIds, err := a.dbTrx.GetCompanyBranchesByCompany(context.Background(), req.CompanyId)
	if err != nil {
		return errors.NewServerError(model.UpdateCompanyError, err.Error())
	}
	var branchIdsParams = make([]string, 0)
	for _, d := range branchIds {
		branchIdsParams = append(branchIdsParams, d.ID)
	}
	checkStockHistory, err := client.GetCheckStockHistory(
		client.GetCheckStockHistoryRequest{
			BranchIds: branchIdsParams,
		})
	if err != nil {
		return err
	}

	checkPOS, err := client.GetCheckPOS(
		client.GetCheckPOSRequest{
			CompanyId: req.CompanyId,
		})
	if err != nil {
		return err
	}

	checkPurchaseOrders, err := client.GetCheckPurchaseOrders(
		client.GetCheckPurchaseOrdersRequest{
			CompanyId: req.CompanyId,
		})
	if err != nil {
		return err
	}

	if companyData.Type != req.Type && (checkStockHistory.Result.Status || checkPOS.Result.Status || checkPurchaseOrders.Result.Status) {
		var errRes = errors.NewClientError()
		errRes.Code = model.CompanyUpdateTypeError
		errRes.Message = model.CompanyUpdateTypeErrorMessage
		return errRes
	}

	result, err := a.dbTrx.UpdateCompany(context.Background(), db.UpdateCompanyParams{
		ID:                req.CompanyId,
		Name:              req.Name,
		InitialName:       req.InitialName,
		Type:              req.Type,
		ResponsiblePerson: req.ResponsiblePerson,
		ImageUrl:          req.ImageUrl,
	})
	if err != nil {
		return errors.NewServerError(model.UpdateCompanyError, err.Error())
	}

	resultBranches, err := a.dbTrx.GetUserCompanyBranches(context.Background(), db.GetUserCompanyBranchesParams{
		UserID:    req.AccountId,
		CompanyID: req.CompanyId,
	})
	if err != nil {
		return errors.NewServerError(model.GetUserCompanyBranchesError, err.Error())
	}

	res := model.RegisterCompanyResponse{
		Company: model.Company{
			AccountId:         result.UserID,
			CompanyId:         result.ID,
			Name:              result.Name,
			InitialName:       result.InitialName,
			Type:              result.Type,
			ResponsiblePerson: result.ResponsiblePerson,
			ImageUrl:          result.ImageUrl,
			Branches:          util.CompanyBranchDbToApi(resultBranches),
		},
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
