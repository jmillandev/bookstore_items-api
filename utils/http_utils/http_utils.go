package http_utils

import (
	"encoding/json"
	"net/http"

	"github.com/jmillandev/bookstore_utils-go/rest_errors"
)

func ResponseJson(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func ResponseJsonError(w http.ResponseWriter, err *rest_errors.RestErr) {
	ResponseJson(w, err.Status, err)
}
