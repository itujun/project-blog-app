package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/itujun/project-blog-app/server/internal/domain"
	"github.com/itujun/project-blog-app/server/internal/repository"
	util "github.com/itujun/project-blog-app/server/internal/utils"
	"github.com/itujun/project-blog-app/server/internal/validation"
)

// PostCreateRequest adalah payload pembuatan post.
type PostCreateRequest struct {
	Title   string  `json:"title" validate:"required,min=3"`
	Slug    *string `json:"slug" validate:"omitempty,slug"` // jika tidak dikirim, dibuat dari Title
	Excerpt *string `json:"excerpt"`
	Content string  `json:"content" validate:"required,min=10"`
}

// HandleCreatePost membuat artikel baru (Author/Admin).
// @Summary  Buat post baru
// @Tags     Posts
// @Accept   json
// @Produce  json
// @Security BearerAuth
// @Param    payload  body  PostCreateRequest  true  "data post"
// @Success  201 {object} APIResponse{data=domain.Post}
// @Failure  400 {object} APIResponse "VALIDATION_ERROR"
// @Failure  409 {object} APIResponse "CONFLICT (slug sudah digunakan)"
// @Router   /posts [post]
func HandleCreatePost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req PostCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteError(w, CodeBadRequest, "payload tidak valid", nil)
		return
	}

	// Validasi payload dasar
	v := validation.New()
	if err := v.Struct(req); err != nil {
		WriteError(w, CodeValidationError, "validasi gagal", err)
		return
	}

	// Normalisasi / generate slug
	var slug string
	if req.Slug != nil && *req.Slug != "" {
		slug = util.NormalizeSlug(*req.Slug) // pastikan konsisten
	} else {
		slug = util.NormalizeSlug(req.Title) // buat dari judul
	}
	if !util.ValidateSlug(slug, 200) {
		WriteError(w, CodeValidationError, "slug tidak valid", map[string]string{
			"pattern": "^[a-z0-9]+(?:-[a-z0-9]+)*$",
		})
		return
	}

	// TODO: ambil user dari context (setelah JWT di Fase 3)
	authorID := uuid.New() // placeholder; ganti dengan user.ID dari token

	post := &domain.Post{
		AuthorID:    authorID,
		Title:       req.Title,
		Slug:        slug,
		Excerpt:     req.Excerpt,
		Content:     req.Content,
		Status:      domain.PostDraft, // default draft
		PublishedAt: nil,
	}

	// Simpan via repository
	repo := r.Context().Value("postRepo").(repository.PostRepository) // contoh: injeksi via context/DI
	if err := repo.Create(ctx, post); err != nil {
		if err == repository.ErrSlugExists {
			WriteError(w, CodeConflict, "slug sudah digunakan", map[string]string{"slug": slug})
			return
		}
		WriteError(w, CodeInternal, "gagal simpan post", nil)
		return
	}

	WriteSuccess(w, http.StatusCreated, "berhasil membuat post", post)
	_ = time.Now() // placeholder untuk ilustrasi; hapus bila tidak dipakai
}
