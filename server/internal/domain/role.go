package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Role merepresentasikan peran user (SuperAdmin, Admin, Author, Reader)
type Role struct {
	ID 		uuid.UUID	`gorm:"type:binary(16);primaryKey" json:"id"`
	Name 	string		`gorm:"size:32;uniqueIndex;not null" json:"name"`

	TimeFields
}

// BeforeCreate: auto-generate UUID jika kosong.
func (r *Role) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == uuid.Nil {
		r.ID = uuid.New()
	}
	return nil
}

func (Role) TableName() string { return "roles" }