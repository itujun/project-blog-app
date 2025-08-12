package http

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

// NewRouter membuat router chi dengan middleware CORS & request logging sederhana
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
			_ = start
		})
	})

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"ok"}`)) // respon sederhana untuk cek konektivitas
	})

	return r
}