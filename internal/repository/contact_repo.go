package repository

import (
	"context"
	"project-app-portfolio-golang-iay/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ContactRepository interface {
	Create(ctx context.Context, c model.Contact) error
}

type contactRepo struct {
	db *pgxpool.Pool
}

func NewContactRepository(db *pgxpool.Pool) ContactRepository {
	return &contactRepo{db: db}
}

func (r *contactRepo) Create(ctx context.Context, c model.Contact) error {
	query := `INSERT INTO contacts (name, email, message) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(ctx, query, c.Name, c.Email, c.Message)
	return err
}
