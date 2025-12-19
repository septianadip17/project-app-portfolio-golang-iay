package handler

import (
	"html/template"
	"net/http"
	"path/filepath"
	"portfolio/internal/model"
	"portfolio/internal/service"

	"go.uber.org/zap"
)

type WebHandler struct {
	service service.ProjectService
	logger  *zap.Logger
}

func NewWebHandler(s service.ProjectService, l *zap.Logger) *WebHandler {
	return &WebHandler{service: s, logger: l}
}

// 1. Menampilkan Halaman Home + List Project
func (h *WebHandler) Home(w http.ResponseWriter, r *http.Request) {
	projects, err := h.service.GetProjects(r.Context())
	if err != nil {
		h.logger.Error("Failed to fetch projects", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Render Template
	tmpl, err := template.ParseGlob(filepath.Join("templates", "*.html"))
	if err != nil {
		h.logger.Error("Template parsing error", zap.Error(err))
		return
	}

	// Kirim data projects ke frontend
	tmpl.ExecuteTemplate(w, "layout", projects)
}

// 2. Handle Submit Form Project
func (h *WebHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Binding Data Form
	p := model.Project{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		Link:        r.FormValue("link"),
	}

	// Panggil Service
	if err := h.service.AddProject(r.Context(), p); err != nil {
		h.logger.Warn("Validation or DB error", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
