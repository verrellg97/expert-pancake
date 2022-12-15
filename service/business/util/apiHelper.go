package util

import (
	db "github.com/expert-pancake/service/business/db/sqlc"
	"github.com/expert-pancake/service/business/model"
)

func CompanyBranchDbToApi(data []db.GetUserCompanyBranchesRow) []model.CompanyBranch {
	var companyBranches = make([]model.CompanyBranch, 0)

	for _, d := range data {
		var companyBranch = model.CompanyBranch{
			AccountId:   d.UserID,
			CompanyId:   d.CompanyID,
			BranchId:    d.ID,
			Name:        d.Name,
			Address:     d.Address,
			PhoneNumber: d.PhoneNumber,
			IsCentral:   d.IsCentral,
		}
		companyBranches = append(companyBranches, companyBranch)
	}

	return companyBranches
}
