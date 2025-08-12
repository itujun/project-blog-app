package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User mewakili pengguna aplikasi.
type User struct {
	ID             uuid.UUID      `gorm:"type:binary(16);primaryKey" json:"id"`
	Email          string         `gorm:"size:191;uniqueIndex;not null" json:"email"`
	PasswordHash   string         `gorm:"size:255;not null" json:"-"`
	DisplayName    string         `gorm:"size:100;not null" json:"display_name"`
	AvatarFilename *string        `gorm:"size:191" json:"avatar_filename,omitempty"`
	IsActive       bool           `gorm:"not null;default:true" json:"is_active"`

	// Relasi (opsional untuk pemetaan)
	Roles []Role `gorm:"many2many:user_roles;joinForeignKey:UserID;joinReferences:RoleID" json:"roles,omitempty"`

	TimeFields
}

// BeforeCreate: isi UUID otomatis
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}

func (User) TableName() string {
	return "users"
}