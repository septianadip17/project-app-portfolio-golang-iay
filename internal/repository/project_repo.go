package repository

import (
	"context"
	"portfolio/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Interface untuk Abstraction
type ProjectRepository interface {
	Create(ctx context.Context, project model.Project) error
	GetAll(ctx context.Context) ([]model.Project, error)
}

type projectRepo struct {
	db *pgxpool.Pool
}

func NewProjectRepository(db *pgxpool.Pool) ProjectRepository {
	return &projectRepo{db: db}
}

func (r *projectRepo) Create(ctx context.Context, p model.Project) error {
	query := `INSERT INTO projects (title, description, image_url, link) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(ctx, query, p.Title, p.Description, p.ImageURL, p.Link)
	return err
}

func (r *projectRepo) GetAll(ctx context.Context) ([]model.Project, error) {
	rows, err := r.db.Query(ctx, "SELECT id, title, description, image_url, link FROM projects")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []model.Project
	for rows.Next() {
		var p model.Project
		if err := rows.Scan(&p.ID, &p.Title, &p.Description, &p.ImageURL, &p.Link); err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}
	return projects, nil
}
