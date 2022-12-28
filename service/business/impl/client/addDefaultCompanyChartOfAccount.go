package client

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func AddDefaultCompanyChartOfAccount(req AddDefaultCompanyChartOfAccountRequest) error {

	reqJSON, err := json.Marshal(req)

	_, err = http.Post(RootPath+AddDefaultCompanyChartOfAccountPath, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		return err
	}

	return nil
}
