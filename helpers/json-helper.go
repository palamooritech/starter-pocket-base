package helper

import (
	"encoding/json"
	"net/http"
)

// the standard json helper, which remains hidden but is used by every HTTP request
// to send it's data secretly.
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
