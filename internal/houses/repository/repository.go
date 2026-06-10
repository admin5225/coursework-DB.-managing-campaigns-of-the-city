package repository

import (
	"context"

	"github.com/admin5225/coursework-DB.-managing-campaigns-of-the-city/internal/houses/domain"
)

type Repository interface {
	Create(
		ctx context.Context,
		house *domain.House,
	) error

	Delete(
		ctx context.Context,
		id int,
	) error
}
