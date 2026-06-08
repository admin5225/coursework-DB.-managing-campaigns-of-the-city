package repository

import (
	"context"

	"github.com/admin5225/coursework-DB.-managing-campaigns-of-the-city/internal/tools/domain"
)

type Repository interface {
	Create(
		ctx context.Context,
		specialist *domain.Tool,
	) error

	Delete(
		ctx context.Context,
		id int,
	) error

	Update(
		ctx context.Context,
		id int,
		quantity int,
	) error
}
