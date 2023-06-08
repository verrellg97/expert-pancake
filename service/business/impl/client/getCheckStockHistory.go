package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func GetCheckStockHistory(req GetCheckStockHistoryRequest) (GetCheckStockHistoryResponse, error) {

	var result GetCheckStockHistoryResponse

	reqJSON, err := json.Marshal(req)

	res, err := http.Post(InventoryRootPath+GetCheckStockHistoryPath, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		return result, err
	}
	if res.StatusCode != http.StatusOK {
		// You may read / inspect response body
		return result, errors.New(res.Status)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err := json.Unmarshal(body, &result); err != nil {
		return result, err
	}

	return result, err
}
