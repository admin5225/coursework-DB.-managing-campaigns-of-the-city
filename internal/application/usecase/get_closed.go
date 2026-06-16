package usecase

import (
	"context"

	"github.com/admin5225/coursework-DB.-managing-campaigns-of-the-city/internal/application/domain"
	"github.com/admin5225/coursework-DB.-managing-campaigns-of-the-city/internal/application/repository"
)

type GetClosedUseCase struct {
	repo repository.Repository
}

func NewGetClosedUseCase(repo repository.Repository) *GetClosedUseCase {
	return &GetClosedUseCase{
		repo: repo,
	}
}

func (u *GetClosedUseCase) Execute(
	ctx context.Context,
) ([]domain.FullApplication, error) {

	return u.repo.GetCloseApplications(
		ctx,
	)
}
