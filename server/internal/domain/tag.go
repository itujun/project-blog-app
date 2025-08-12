package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Tag untuk labeling post.
type Tag struct {
	ID   uuid.UUID `gorm:"type:binary(16);primaryKey" json:"id"`
	Name string    `gorm:"size:100;uniqueIndex;not null" json:"name"`

	TimeFields
}

func (t *Tag) BeforeCreate(tx *gorm.DB) (err error) {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}
	return nil
}

func (Tag) TableName() string { return "tags" }
