package http

import "net/http"

// Kumpulan error code standar agar konsisten antara BE-FE.
// FE bisa switch-case berdasarkan Code berikut (bukan pesan panjang).
const (
	CodeValidationError = "VALIDATION_ERROR"
	CodeUnauthorized    = "UNAUTHORIZED"
	CodeForbidden       = "FORBIDDEN"
	CodeNotFound        = "NOT_FOUND"
	CodeConflict        = "CONFLICT"
	CodeInternal        = "INTERNAL_ERROR"
	CodeBadRequest      = "BAD_REQUEST"
)

// Map ke HTTP status agar seragam pemakaian di handler.
var codeToHTTP = map[string]int{
	CodeValidationError: http.StatusBadRequest,
	CodeUnauthorized:    http.StatusUnauthorized,
	CodeForbidden:       http.StatusForbidden,
	CodeNotFound:        http.StatusNotFound,
	CodeConflict:        http.StatusConflict,
	CodeInternal:        http.StatusInternalServerError,
	CodeBadRequest:      http.StatusBadRequest,
}
