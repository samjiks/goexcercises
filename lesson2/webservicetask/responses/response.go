package responses

import (
	"net/http"
	"encoding/json"
)

type Reponse struct {
	Code   int    `json:"code"`
	Result string `json:"result"`
}


func SendReponse(
	res string,
	w http.ResponseWriter,
	code int,
) {
	if code == 500 {
		res = "Bad request"
	}
	r := Reponse{
		Code:   code,
		Result: res,
	}
	js, err := json.Marshal(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	w.Write(js)
}

