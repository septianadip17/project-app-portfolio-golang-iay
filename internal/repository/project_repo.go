package service

import (
	"context"
	"errors"
	"portfolio/internal/model"
	"portfolio/internal/repository"
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
	// Validasi Data (Ketentuan Utama)
	if p.Title == "" {
		return errors.New("title cannot be empty")
	}
	if len(p.Description) < 10 {
		return errors.New("description must be at least 10 characters")
	}
	return s.repo.Create(ctx, p)
}

func (s *projectService) GetProjects(ctx context.Context) ([]model.Project, error) {
	return s.repo.GetAll(ctx)
}
