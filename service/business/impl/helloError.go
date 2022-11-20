package impl

import (
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
)

func (a businessService) HelloError(w http.ResponseWriter, r *http.Request) error {

	return errors.NewServerError(100001, "INI ERROR")
}
