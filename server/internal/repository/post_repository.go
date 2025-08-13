package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/itujun/project-blog-app/server/internal/domain"
)

// ErrSlugExists dipakai jika slug sudah terpakai.
var ErrSlugExists = errors.New("slug exists")

// PostRepository mengelola akses data Post.
type PostRepository interface {
	Create(ctx context.Context, p *domain.Post) error
	GetByID(ctx context.Context, id string) (*domain.Post, error)
	ExistsSlug(ctx context.Context, slug string) (bool, error)
}

// postRepoGorm implementasi dengan GORM.
type postRepoGorm struct {
	db *gorm.DB
}

// NewPostRepository membuat repo post berbasis GORM.
func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepoGorm{db: db}
}

func (r *postRepoGorm) ExistsSlug(ctx context.Context, slug string) (bool, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&domain.Post{}).
		Where("slug = ?", slug).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *postRepoGorm) Create(ctx context.Context, p *domain.Post) error {
	// Cek unik slug secara aplikatif untuk error yang ramah FE.
	exists, err := r.ExistsSlug(ctx, p.Slug)
	if err != nil {
		return err
	}
	if exists {
		return ErrSlugExists
	}
	// Simpan
	if err := r.db.WithContext(ctx).Create(p).Error; err != nil {
		// Jika error karena constraint unik di DB, tetap kembalikan ErrSlugExists
		// agar FE mendapatkan kode error yang konsisten.
		return err
	}
	return nil
}

func (r *postRepoGorm) GetByID(ctx context.Context, id string) (*domain.Post, error) {
	var post domain.Post
	if err := r.db.WithContext(ctx).First(&post, "id = UNHEX(REPLACE(?, '-', ''))", id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}
