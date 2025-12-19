package service

import (
	"context"
	"errors"
	"project-app-portfolio-golang-iay/internal/model"
	"project-app-portfolio-golang-iay/internal/repository"
)

type ProjectService interface {
	AddProject(ctx context.Context, p model.Project) error
	GetProjects(ctx context.Context) ([]model.Project, error)
}

type projectService struct {
	repo repository.ProjectRepository
}

func NewProjectService(repo repository.ProjectRepository) ProjectService {
	return &projectService{repo: repo}
}

func (s *projectService) AddProject(ctx context.Context, p model.Project) error {
	// --- VALIDASI DATA (Business Logic) ---
	// 1. Cek Judul Kosong
	if p.Title == "" {
		return errors.New("judul project tidak boleh kosong")
	}
	// 2. Cek Panjang Deskripsi (Operator Perbandingan)
	if len(p.Description) < 10 {
		return errors.New("deskripsi terlalu pendek (minimal 10 karakter)")
	}

	// Jika lolos validasi, simpan ke repo
	return s.repo.Create(ctx, p)
}

func (s *projectService) GetProjects(ctx context.Context) ([]model.Project, error) {
	return s.repo.GetAll(ctx)
}
