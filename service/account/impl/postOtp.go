package impl

import (
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/account/model"
	"net/http"
)

func (a accountService) PostOtp(w http.ResponseWriter, r *http.Request) error {

	res := model.PostOtpResponse{Message: "OK"}

	httpHandler.WriteResponse(w, res)

	return nil
}
