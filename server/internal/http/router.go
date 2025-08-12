package http

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

// NewRouter membuat router chi & memasang rute dasar + versi API.
// Pola: /api/v1/** dan grup admin di /api/v1/admin/**
func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	// Middleware CORS: izinkan akses dari frontend dev (http://localhost:3000)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:*", "https://localhost:*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:    []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Middleware basic logging (bisa kita ganti ke Zap HTTP middleware nanti)
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, req)
			// Di tahap awal, cukup catat method+path+durasi; nanti kita pakai zap untuk struktur log
			_ = start // TODO: pasang zap logger untuk catat durasi, status, dsb.
		})
	})

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		WriteSuccess(w, http.StatusOK, "ok", map[string]string{"status": "ok"})
	})

	// ---- API v1 ----
	r.Route("/api/v1", func(api chi.Router) {

		// Public/semi-public endpoints (contoh)
		api.Get("/posts", handleListPosts)          // GET /api/v1/posts
		api.Get("/posts/{id}", HandleGetPost)       // GET /api/v1/posts/:id
		api.Get("/posts/{id}/comments", HandleListComments)

		// Protected user endpoints (nanti diberi JWT middleware)
		api.Group(func(g chi.Router) {
			// TODO: g.Use(AuthMiddleware()) // setelah Fase 3
			g.Post("/posts", HandleCreatePost)                               // Author/Admin
			g.Post("/posts/{id}/comments", HandleCreateComment)              // Reader+
			g.Patch("/users/me", HandleUpdateMe)                             // User login
		})

		// Admin group: konsisten di /api/v1/admin/**
		api.Route("/admin", func(ad chi.Router) {
			// TODO: ad.Use(AuthMiddleware(), RBACMiddleware("admin")) // setelah Fase 3 (Casbin)
			ad.Post("/posts/{id}/publish", HandlePublishPost)          // Admin/SuperAdmin
			ad.Post("/categories", HandleCreateCategory)
			ad.Post("/tags", HandleCreateTag)
			ad.Get("/users", HandleListUsers)
		})
	})

	return r
}
// Handler di atas placeholder (belum diisi), tujuannya membentuk pola URL konsisten dulu. Nanti saat Fase 3/4 kita lengkapi dengan JWT + Casbin.