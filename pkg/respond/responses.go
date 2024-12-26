package respond

import (
	"encoding/json"
	"net/http"
)

type (
	ErrorResp struct {
		Error string `json:"error"`
	}
)

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func Ok(w http.ResponseWriter, payload any) {
	writeJSON(w, http.StatusOK, payload)
}

func InternalErr(w http.ResponseWriter, err string) {
	writeJSON(w, http.StatusInternalServerError, err)
}

func BadRequest(w http.ResponseWriter, err string) {
	writeJSON(w, http.StatusBadRequest, err)
}
