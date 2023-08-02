package impl

import (
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/account/model"
	"net/http"
)

var defaultSecurityQuestions = []string{
	"Di kota mana anda bertemu pasangan/mitra pertama Anda?",
	"Apa nama tengah ibu Anda?",
	"Apa nama sekolah pertama yang Anda hadiri?",
	"Apa nama panggilan anak Anda?",
}

func (a accountService) GetDefaultSecurityQuestions(w http.ResponseWriter, r *http.Request) error {

	res := model.GetDefaultSecurityQuestionsResponse{Questions: defaultSecurityQuestions}
	httpHandler.WriteResponse(w, res)

	return nil
}
