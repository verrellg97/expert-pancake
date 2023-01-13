package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

func AddDefaultContactBook(req AddDefaultContactBookRequest) error {

	reqJSON, err := json.Marshal(req)

	res, err := http.Post(BusinessRelationRootPath+AddDefaultContactBookPath, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		// You may read / inspect response body
		return errors.New(res.Status)
	}

	return nil
}
