package main

import (
	"fmt"
	"net/http"

	"github.com/itujun/project-blog-app/server/config"
	_ "github.com/itujun/project-blog-app/server/internal/docs" // <-- penting: import paket docs hasil generate
	httpRouter "github.com/itujun/project-blog-app/server/internal/http"
	"github.com/itujun/project-blog-app/server/internal/logger"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// @title			BlogApp API
// @version 		1.0
// @description		REST API untuk BlogApp (Go + Next.js)
// @termsOfService	http://swagger.io/terms/

// @contact.name	API Support
// @contact.email	support@example.com

// @license.name	MIT

// @BasePath					/api/v1
// @schemes						http
// @securityDefinitions.apikey 	BearerAuth
// @in 							header
// @name 						Authorization
func main() {
	cfg := config.Load()	// muat konffigurasi (ENV/file)
	logger.Init()			// inisialisasi zap

	port := cfg.GetInt("app.port")	// baca port dari config
	r := httpRouter.NewRouter()			// router chi dengan /health

	// Endpoint Swagger UI: http://localhost:8080/swagger/index.html
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", port)), // URL definisi swagger
	))

	addr := fmt.Sprintf(":%d", port)	// alamat server HTTP
	logger.L().Infow("starting server", "addr", addr, "env", cfg.GetString("app.env")) // log structured
	
	if err := http.ListenAndServe(addr, r); err != nil {
		logger.L().Fatalw("server stopped with error", "error", err) // log fatal jika server gagal berjalan
	} 
}
// main adalah entrypoint service API.
// - Inisialisasi config & logger (singleton)
// - Siapkan router dan jalankan HTTP server