package repository

import (
	"context"

	"github.com/admin5225/coursework-DB.-managing-campaigns-of-the-city/internal/application/domain"
)

type Repository interface {
	Create(
		ctx context.Context,
		application *domain.Application,
	) error
	Update(
		ctx context.Context,
		application *domain.Application,
	)
	Delete(
		ctx context.Context,
		id int,
	)
	Close(
		ctx context.Context,
		id int,
	)
	GetClosed(
		ctx context.Context,
	) ([]domain.Application, error)
}
