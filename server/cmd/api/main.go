package main

import (
	"fmt"
	"net/http"

	"github.com/itujun/project-blog-app/server/config"
	httpRouter "github.com/itujun/project-blog-app/server/internal/http"
	"github.com/itujun/project-blog-app/server/internal/logger"
)

// main adalah entrypoint service API.
// - Inisialisasi config & logger (singleton)
// - Siapkan router dan jalankan HTTP server
func main() {
	cfg := config.Load()	// muat konffigurasi (ENV/file)
	logger.Init()			// inisialisasi zap

	port := cfg.GetInt("app.port")	// baca port dari config
	r := httpRouter.NewRouter()			// router chi dengan /health

	addr := fmt.Sprintf(":%d", port)	// alamat server HTTP
	logger.L().Infow("starting server", "addr", addr, "env", cfg.GetString("app.env")) // log structured
	
	if err := http.ListenAndServe(addr, r); err != nil {
		logger.L().Fatalw("server stopped with error", "error", err) // log fatal jika server gagal berjalan
	} 
	}