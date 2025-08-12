package http

import (
	"encoding/json"
	"net/http"
)

// APIError memuat kode error ringkas & detail opsional.
type APIError struct {
	Code    string      `json:"code"`
	Details interface{} `json:"details,omitempty"`
}

// APIResponse adalah amplop respons baku.
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   *APIError   `json:"error,omitempty"`
}

// WriteJSON menulis JSON dengan status code & header tepat.
func WriteJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload) // abaikan error encode demi ringkas
}

// WriteSuccess menulis respons sukses standar.
func WriteSuccess(w http.ResponseWriter, status int, message string, data interface{}) {
	WriteJSON(w, status, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// WriteError menulis error dengan Code baku & detail.
func WriteError(w http.ResponseWriter, code string, message string, details interface{}) {
	status, ok := codeToHTTP(code)
	if !ok {
		status = http.StatusInternalServerError
		code = CodeInternal
	}
	WriteJSON(w, status, APIResponse{
		Success: false,
		Message: message,
		Error:   &APIError{Code: code, Details: details},
	})
}
