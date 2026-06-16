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

	GetCloseApplications(
		ctx context.Context,
	) ([]domain.FullApplication, error)

	GetApplicationsStats(
		ctx context.Context,
	) (*domain.Statistics, error)

	GetHouseApplications(
		ctx context.Context,
		housID int,
	) ([]domain.FullApplication, error)

	GetSpecialistApplications(
		ctx context.Context,
		specialistID int,
	) ([]domain.FullApplication, error)
}
