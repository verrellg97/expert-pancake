package util

import "github.com/expert-pancake/service/sales/impl/client"

func BranchApiToMap(data []client.CompanyBranch) map[string]client.CompanyBranch {
	dataMap := map[string]client.CompanyBranch{}
	for _, v := range data {
		dataMap[v.BranchId] = v
	}

	return dataMap
}
