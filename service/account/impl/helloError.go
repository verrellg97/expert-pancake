package impl

import (
	"github.com/calvinkmts/expert-pancake/engine/errors"
	"net/http"
)

func (a accountService) HelloError(w http.ResponseWriter, r *http.Request) error {

	return errors.NewServerError(100001, "INI ERROR")
}
