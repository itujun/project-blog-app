package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PostStatus adalah enum status publikasi post.
type PostStatus string

const (
	PostDraft     PostStatus = "draft"
	PostPending   PostStatus = "pending"
	PostPublished PostStatus = "published"
	PostRejected  PostStatus = "rejected"
)

// Post merepresentasikan artikel/blog post.
type Post struct {
	ID            uuid.UUID  `gorm:"type:binary(16);primaryKey" json:"id"`
	AuthorID      uuid.UUID  `gorm:"type:binary(16);not null;index" json:"author_id"`
	Title         string     `gorm:"size:200;not null" json:"title"`
	Slug          string     `gorm:"size:200;uniqueIndex;not null" json:"slug"`
	Excerpt       *string    `gorm:"type:text" json:"excerpt,omitempty"`
	Content       string     `gorm:"type:longtext;not null" json:"content"`
	CoverFilename *string    `gorm:"size:191" json:"cover_filename,omitempty"`
	Status        PostStatus `gorm:"type:ENUM('draft','pending','published','rejected');default:'draft';not null" json:"status"`
	PublishedAt   *time.Time `json:"published_at,omitempty"`

	// Relasi (opsional): Author, Tags
	Author *User `gorm:"foreignKey:AuthorID;references:ID" json:"author,omitempty"`
	Tags   []Tag `gorm:"many2many:post_tags;joinForeignKey:PostID;JoinReferences:TagID" json:"tags,omitempty"`

	TimeFields
}

// BeforeCreate: generate UUID saat create.
func (p *Post) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return nil
}

func (Post) TableName() string { return "posts" }
