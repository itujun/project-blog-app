package domain

import (
	"github.com/google/uuid"
)

// UserRole adalah tabel pivot many-to-many antara users dan roles.
type UserRole struct {
	UserID uuid.UUID `gorm:"type:binary(16);primaryKey" json:"user_id"`
	RoleID uuid.UUID `gorm:"type:binary(16);primaryKey"  json:"role_id"`
}

func (UserRole) TableName() string { return "user_roles" }
