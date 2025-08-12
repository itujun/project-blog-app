package http

// APIError merepresentasikan error standar API.
type APIError struct {
	Code string			`json:"code"`				// kode error ringkas (mis: VALIDATION_ERROR, NOT_FOUND)
	Details interface{} `json:"details,omitempty"`	// detail (bisa array field error, dsb.)
} 

// APIResponse adalah envelope response umum.
type APIResponse struct {
	Success	bool		`json:"success"`			// true jika sukses
	Message	string		`json:"message,omitempty"`	// pesan singkat
	Data	interface{}	`json:"data,omitempty"`		// payload (opsional)
	Error	*APIError	`json:"error,omitempty"`	// error (opsional)
}