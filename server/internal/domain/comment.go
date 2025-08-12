package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Comment merepresentasikan komentar pada artikel.
type Comment struct {
	ID        uuid.UUID  `gorm:"type:binary(16);primaryKey" json:"id"`
	PostID    uuid.UUID  `gorm:"type:binary(16);not null;index" json:"post_id"`
	AuthorID  uuid.UUID  `gorm:"type:binary(16);not null;index" json:"author_id"` // Wajib login: NOT NULL
	Content   string     `gorm:"type:text;not null" json:"content"`
	IsApproved bool      `gorm:"not null;default:true" json:"is_approved"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relasi opsional
	Author *User `gorm:"foreignKey:AuthorID;references:ID" json:"author,omitempty"`
	Post   *Post `gorm:"foreignKey:PostID;references:ID" json:"post,omitempty"`
}

// BeforeCreate: auto UUID.
func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	return nil
}

func (Comment) TableName() string { return "comments" }
