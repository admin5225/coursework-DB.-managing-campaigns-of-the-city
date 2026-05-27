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
		CALL main.add_application($1, $2, $3, $4, $5)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		application.Description,
		application.HouseID,
		application.SpecialistID,
		application.WorkTypeID,
		application.StatusID,
	)

	return err
}

func (r *Repository) Delete(ctx context.Context, id int) error {
	query := `
		CALL main.delete_application($1)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		id,
	)

	return err
}

func (r *Repository) Close(ctx context.Context, id int) error {
	query := `
		CALL main.close_application($1)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		id,
	)

	return err
}
