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

	Delete(
		ctx context.Context,
		id int,
	) error

	Close(
		ctx context.Context,
		id int,
	) error
}
