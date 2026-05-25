package postgres

import (
	"context"

	"github.com/admin5225/coursework-DB.-managing-campaigns-of-the-city/internal/application/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(ctx context.Context, application *domain.Application) error {
	query := `
		CALL create_request($1, $2, $3, $4)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		application.Description,
		application.CreatedAt,
		application.HouseID,
		application.SpecialistID,
		application.WorkTypeID,
		application.StatusID,
	)

	return err
}
