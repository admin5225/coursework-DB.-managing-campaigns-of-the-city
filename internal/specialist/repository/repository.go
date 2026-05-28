package repository

import (
	"context"

	"github.com/admin5225/coursework-DB.-managing-campaigns-of-the-city/internal/specialist/domain"
)

type Repository interface {
	Create(
		ctx context.Context,
		specialist *domain.Specialist,
	) error

	Delete(
		ctx context.Context,
		id int,
	) error
}
