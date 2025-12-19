package service

import (
	"context"
	"errors"
	"project-app-portfolio-golang-iay/internal/model"
	"project-app-portfolio-golang-iay/internal/repository"
	"strings"
)

type ContactService interface {
	SubmitContact(ctx context.Context, c model.Contact) error
}

type contactService struct {
	repo repository.ContactRepository
}

func NewContactService(repo repository.ContactRepository) ContactService {
	return &contactService{repo: repo}
}

func (s *contactService) SubmitContact(ctx context.Context, c model.Contact) error {
	// --- VALIDASI ---
	if c.Name == "" {
		return errors.New("nama harus diisi")
	}
	if c.Email == "" || !strings.Contains(c.Email, "@") {
		return errors.New("format email tidak valid")
	}
	if c.Message == "" {
		return errors.New("pesan tidak boleh kosong")
	}

	return s.repo.Create(ctx, c)
}
