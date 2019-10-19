package utils

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
)

func SendPostRequest(server *mux.Router, url string, params interface{}) *httptest.ResponseRecorder {
	data, _ := json.Marshal(params)
	r, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))

	w := httptest.NewRecorder()
	server.ServeHTTP(w, r)

	return w
}
