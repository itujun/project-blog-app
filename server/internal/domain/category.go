package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Category untuk pengelompokan artikel.
// (Catatan: pada migrasi awal kita BELUM mapping posts<->categories.
//  Kalau nanti butuh, kita bisa tambah kolom posts.category_id atau tabel pivot post_categories.)
type Category struct {
	ID   uuid.UUID `gorm:"type:binary(16);primaryKey" json:"id"`
	Name string    `gorm:"size:100;uniqueIndex;not null" json:"name"`

	TimeFields
}

func (c *Category) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	return nil
}

func (Category) TableName() string { return "categories" }
