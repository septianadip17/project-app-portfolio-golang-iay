package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"

	// Pastikan module path ini sesuai go.mod kamu
	"project-app-portfolio-golang-iay/internal/handler"
	"project-app-portfolio-golang-iay/internal/repository"
	"project-app-portfolio-golang-iay/internal/service"
	"project-app-portfolio-golang-iay/pkg/database"
	"project-app-portfolio-golang-iay/pkg/logger"

	"go.uber.org/zap"
)

func main() {
	// 1. Inisialisasi Logger
	// PERBAIKAN: Menangkap 2 return value (logger, err)
	logr, err := logger.NewLogger()
	if err != nil {
		// Jika logger gagal, pakai panic biasa karena kita butuh logger untuk start
		panic(err)
	}
	defer logr.Sync()
	logr.Info("Starting Portfolio App...")

	// 2. Inisialisasi Database
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:password@localhost:5432/portfolio_db" // Ganti password sesuai pgadmin mu
	}

	// Fungsi ini sekarang sudah ada di Langkah 1
	dbPool, err := database.NewPostgresConnection(dbURL)
	if err != nil {
		logr.Fatal("Gagal koneksi ke database", zap.Error(err))
	}
	defer dbPool.Close()
	logr.Info("Database connected successfully")

	// 3. Setup Template Caching
	templateCache, err := newTemplateCache("./templates")
	if err != nil {
		logr.Fatal("Gagal parsing templates", zap.Error(err))
	}

	// 4. Dependency Injection (WIRING)

	// PERBAIKAN: Hapus 'logr' karena Repository cuma minta dbPool (lihat error log tadi)
	projectRepo := repository.NewProjectRepository(dbPool)

	// PERBAIKAN: Hapus 'logr' karena Service cuma minta Repo
	projectService := service.NewProjectService(projectRepo)

	// Handler butuh Service, Logger, dan Template
	projectHandler := handler.NewProjectHandler(projectService, logr, templateCache)

	// 5. Routing
	mux := http.NewServeMux()

	mux.HandleFunc("/", projectHandler.Home)
	mux.HandleFunc("/project/add", projectHandler.AddProjectForm)
	mux.HandleFunc("/project/store", projectHandler.StoreProject)

	// 6. Jalankan Server
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logr.Info("Server berjalan", zap.String("addr", "http://localhost:8080"))

	err = srv.ListenAndServe()
	if err != nil {
		logr.Fatal("Server berhenti", zap.Error(err))
	}
}

// Helper Template Cache
func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	// Pastikan folder templates ada isinya (layout.html, index.html, dll)
	pages, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("gagal baca folder templates: %w", err)
	}

	for _, page := range pages {
		name := page.Name()
		// Skip file layout agar tidak diparse dobel, kita parse manual di bawah
		if name == "layout.html" {
			continue
		}

		// Pola: Parsing layout.html + file halaman spesifik
		files := []string{
			dir + "/layout.html",
			dir + "/" + name,
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}
	return cache, nil
}
