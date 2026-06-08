package postgres

import (
	"context"

	"github.com/admin5225/coursework-DB.-managing-campaigns-of-the-city/internal/tools/domain"
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

func (r *Repository) Create(ctx context.Context, tool *domain.Tool) error {
	query := `
		CALL main.create_tool($1, $2, $3)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		tool.Name,
		tool.ManagingCampaiginID,
		tool.Quantity,
	)

	return err
}

func (r *Repository) Delete(ctx context.Context, id int) error {
	query := `
		CALL main.delete_tool($1)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		id,
	)

	return err
}

func (r *Repository) Update(ctx context.Context, id int, quantity int) error {
	query := `
		CALL main.update_tool_quantity($1, $2)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		id,
		quantity,
	)

	return err
}
