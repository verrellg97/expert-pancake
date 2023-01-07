package util

import (
	db "github.com/expert-pancake/service/business-relation/db/sqlc"
)

func ContactBookBranchDbToApi(data []db.GetContactBookBranchesRow) []string {
	var contactBookBranches = make([]string, 0)

	for _, d := range data {
		contactBookBranches = append(contactBookBranches, d.CompanyBranchID)
	}

	return contactBookBranches
}
