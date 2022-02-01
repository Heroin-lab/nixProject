package services

import (
	"encoding/json"
	"net/http"
)

type RespondService struct {
}

func (res *RespondService) Error(w http.ResponseWriter, r *http.Request, code int, err error) {
	res.Respond(w, r, code, map[string]string{"Error": err.Error()})
}

func (*RespondService) Respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
