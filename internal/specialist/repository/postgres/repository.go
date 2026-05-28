package postgres

import (
	"context"

	"github.com/admin5225/coursework-DB.-managing-campaigns-of-the-city/internal/specialist/domain"
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

func (r *Repository) Create(ctx context.Context, specialist *domain.Specialist) error {
	query := `
		CALL main.create_specialist($1, $2, $3, $4)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		specialist.FullName,
		specialist.Position,
		specialist.PhoneNumber,
		specialist.ManagingCamaiginID,
	)

	return err
}

func (r *Repository) Delete(ctx context.Context, id int) error {
	query := `
		CALL main.delete_specialist($1)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		id,
	)

	return err
}
