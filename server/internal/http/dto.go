package http

import "github.com/google/uuid"

// RegisterRequest adalah payload untuk /auth/register
type RegisterRequest struct {
	Email    string `json:"email" example:"user@example.com"`
	Password string `json:"password" example:"StrongP@ssw0rd"`
}

// RegisteredUser adalah respons ringkas saat user berhasil daftar
type RegisteredUser struct {
	ID          uuid.UUID `json:"id" example:"3fa85f64-5717-4562-b3fc-2c963f66afa6"`
	Email       string    `json:"email" example:"user@example.com"`
	DisplayName string    `json:"display_name" example:"User Baru"`
}
