package helpers

import (
	"encoding/json"
	"net/http"
)

func ResponseJson(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Failed to marshal response"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func ResponseError(w http.ResponseWriter, code int, msg string) {
	ResponseJson(w, code, map[string]string{"error": msg})
}
