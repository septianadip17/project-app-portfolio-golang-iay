package service

import (
	"context"
	"project-app-portfolio-golang-iay/internal/model"
	"project-app-portfolio-golang-iay/internal/repository"
)

type ExperienceService interface {
	GetExperiences(ctx context.Context) ([]model.Experience, error)
}

type experienceService struct {
	repo repository.ExperienceRepository
}

func NewExperienceService(repo repository.ExperienceRepository) ExperienceService {
	return &experienceService{repo: repo}
}

func (s *experienceService) GetExperiences(ctx context.Context) ([]model.Experience, error) {
	return s.repo.GetAll(ctx)
}
