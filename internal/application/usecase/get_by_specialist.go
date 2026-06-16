package usecase

import (
	"context"

	"github.com/admin5225/coursework-DB.-managing-campaigns-of-the-city/internal/application/domain"
	"github.com/admin5225/coursework-DB.-managing-campaigns-of-the-city/internal/application/repository"
)

type GetBySpecialistUseCase struct {
	repo repository.Repository
}

func NewGetBySpecialistUseCase(repo repository.Repository) *GetBySpecialistUseCase {
	return &GetBySpecialistUseCase{
		repo: repo,
	}
}

func (u *GetBySpecialistUseCase) Execute(
	ctx context.Context,
	specialistID int,
) ([]domain.FullApplication, error) {

	return u.repo.GetSpecialistApplications(
		ctx,
		specialistID,
	)
}
