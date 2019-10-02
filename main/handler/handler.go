package handler

import (
	"encoding/json"
	"net/http"
)

func Welcome(w http.ResponseWriter, req *http.Request) {
	data := map[string]string{
		"somedata": "valuedata",
	}
	json.NewEncoder(w).Encode(data)
}
