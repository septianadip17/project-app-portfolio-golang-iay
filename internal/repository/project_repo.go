package repository

import (
	"context"
	"project-app-portfolio-golang-iay/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Interface agar code lebih modular & bisa di-test
type ProjectRepository interface {
	Create(ctx context.Context, p model.Project) error
	GetAll(ctx context.Context) ([]model.Project, error)
}

// Struct implementasi
type projectRepo struct {
	db *pgxpool.Pool
}

// Constructor
func NewProjectRepository(db *pgxpool.Pool) ProjectRepository {
	return &projectRepo{db: db}
}

// Query Insert ke Database
func (r *projectRepo) Create(ctx context.Context, p model.Project) error {
	query := `INSERT INTO projects (title, description, image_url, link) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(ctx, query, p.Title, p.Description, p.ImageURL, p.Link)
	return err
}

// Query Select Semua Data
func (r *projectRepo) GetAll(ctx context.Context) ([]model.Project, error) {
	query := `SELECT id, title, description, image_url, link, created_at FROM projects ORDER BY created_at DESC`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []model.Project
	for rows.Next() {
		var p model.Project
		err := rows.Scan(&p.ID, &p.Title, &p.Description, &p.ImageURL, &p.Link, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}
	return projects, nil
}
