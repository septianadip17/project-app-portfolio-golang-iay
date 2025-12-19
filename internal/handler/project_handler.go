package handler

import (
	"html/template"
	"net/http"
	"project-app-portfolio-golang-iay/internal/service"

	"go.uber.org/zap"
)

// ProjectHandler struct
type ProjectHandler struct {
	Service   service.ProjectService
	Log       *zap.Logger
	Templates map[string]*template.Template
}

// --- FUNGSI INI YANG HILANG TADI ---
// NewProjectHandler adalah constructor untuk membuat instance ProjectHandler baru
func NewProjectHandler(s service.ProjectService, log *zap.Logger, tmpl map[string]*template.Template) *ProjectHandler {
	return &ProjectHandler{
		Service:   s,
		Log:       log,
		Templates: tmpl,
	}
}

// -----------------------------------

// Home menangani halaman utama
func (h *ProjectHandler) Home(w http.ResponseWriter, r *http.Request) {
	// Panggil service untuk ambil data project
	projects, err := h.Service.GetAllProjects(r.Context())
	if err != nil {
		h.Log.Error("Gagal mengambil data project", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Render template index.html
	if tmpl, ok := h.Templates["index.html"]; ok {
		// Kirim data projects ke template
		data := map[string]interface{}{
			"Title":    "Portfolio Septian",
			"Projects": projects,
		}
		tmpl.ExecuteTemplate(w, "layout", data)
	} else {
		http.Error(w, "Template not found", http.StatusInternalServerError)
	}
}

// AddProjectForm menampilkan form tambah project
func (h *ProjectHandler) AddProjectForm(w http.ResponseWriter, r *http.Request) {
	if tmpl, ok := h.Templates["project_form.html"]; ok {
		data := map[string]interface{}{
			"Title": "Tambah Project Baru",
		}
		tmpl.ExecuteTemplate(w, "layout", data)
	}
}

// StoreProject memproses submit form
func (h *ProjectHandler) StoreProject(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Ambil data dari form
	title := r.FormValue("title")
	description := r.FormValue("description")
	imageURL := r.FormValue("image_url")
	gitURL := r.FormValue("git_url")

	// Panggil service untuk simpan
	err := h.Service.CreateProject(r.Context(), title, description, imageURL, gitURL)
	if err != nil {
		h.Log.Error("Gagal menyimpan project", zap.Error(err))
		http.Error(w, "Gagal menyimpan data", http.StatusInternalServerError)
		return
	}

	// Redirect kembali ke home
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
