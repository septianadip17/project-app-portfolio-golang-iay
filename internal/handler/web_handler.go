package handler

import (
	"html/template"
	"net/http"
	"path/filepath"
	"project-app-portfolio-golang-iay/internal/model"
	"project-app-portfolio-golang-iay/internal/service"

	"go.uber.org/zap"
)

type WebHandler struct {
	projectService    service.ProjectService
	experienceService service.ExperienceService
	contactService    service.ContactService
	logger            *zap.Logger
}

// Kita inject semua service yang dibutuhkan ke sini
func NewWebHandler(
	ps service.ProjectService,
	es service.ExperienceService,
	cs service.ContactService,
	log *zap.Logger,
) *WebHandler {
	return &WebHandler{
		projectService:    ps,
		experienceService: es,
		contactService:    cs,
		logger:            log,
	}
}

// Struct data untuk dikirim ke Template HTML (ViewModel)
type HomeData struct {
	Projects    []model.Project
	Experiences []model.Experience
}

// 1. GET / -> Halaman Utama
func (h *WebHandler) Home(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Ambil Data Projects
	projects, err := h.projectService.GetProjects(ctx)
	if err != nil {
		h.logger.Error("Gagal mengambil data project", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Ambil Data Experiences
	experiences, err := h.experienceService.GetExperiences(ctx)
	if err != nil {
		h.logger.Error("Gagal mengambil data experience", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Gabungkan data
	data := HomeData{
		Projects:    projects,
		Experiences: experiences,
	}

	// Render Template
	// Pastikan folder templates ada dan berisi file .html
	tmpl, err := template.ParseGlob(filepath.Join("templates", "*.html"))
	if err != nil {
		h.logger.Error("Gagal parsing template", zap.Error(err))
		http.Error(w, "Template Error", http.StatusInternalServerError)
		return
	}

	// Eksekusi template "layout" dengan data
	err = tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		h.logger.Error("Gagal render template", zap.Error(err))
	}
}

// 2. GET /project/new -> Tampilkan Form Tambah Project
func (h *WebHandler) NewProjectForm(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseGlob(filepath.Join("templates", "*.html"))
	if err != nil {
		h.logger.Error("Gagal parsing template", zap.Error(err))
		return
	}
	// Render file khusus form (misal namanya project_form.html di dalam layout)
	tmpl.ExecuteTemplate(w, "project_form_page", nil)
}

// 3. POST /project/create -> Proses Simpan Project
func (h *WebHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Ambil input dari form HTML
	p := model.Project{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		Link:        r.FormValue("link"),
		ImageURL:    r.FormValue("image_url"),
	}

	// Panggil Service
	if err := h.projectService.AddProject(r.Context(), p); err != nil {
		h.logger.Warn("Gagal tambah project (validasi/db)", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Redirect kembali ke home jika sukses
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// 4. POST /contact -> Proses Kirim Pesan
func (h *WebHandler) SubmitContact(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	c := model.Contact{
		Name:    r.FormValue("name"),
		Email:   r.FormValue("email"),
		Message: r.FormValue("message"),
	}

	if err := h.contactService.SubmitContact(r.Context(), c); err != nil {
		h.logger.Warn("Gagal submit kontak", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Redirect kembali ke home dengan query param sukses (opsional buat notif)
	http.Redirect(w, r, "/?contact=success", http.StatusSeeOther)
}
