package repository

import (
	"context"
	"project-app-portfolio-golang-iay/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ExperienceRepository interface {
	GetAll(ctx context.Context) ([]model.Experience, error)
}

type experienceRepo struct {
	db *pgxpool.Pool
}

func NewExperienceRepository(db *pgxpool.Pool) ExperienceRepository {
	return &experienceRepo{db: db}
}

func (r *experienceRepo) GetAll(ctx context.Context) ([]model.Experience, error) {
	query := `SELECT id, role, company, start_date, end_date, description FROM experiences ORDER BY start_date DESC`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var experiences []model.Experience
	for rows.Next() {
		var e model.Experience
		// Hati-hati, end_date bisa NULL, jadi pakai pointer
		err := rows.Scan(&e.ID, &e.Role, &e.Company, &e.StartDate, &e.EndDate, &e.Description)
		if err != nil {
			return nil, err
		}
		experiences = append(experiences, e)
	}
	return experiences, nil
}
