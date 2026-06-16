package usecase

import (
	"context"

	"github.com/admin5225/coursework-DB.-managing-campaigns-of-the-city/internal/application/domain"
	"github.com/admin5225/coursework-DB.-managing-campaigns-of-the-city/internal/application/repository"
)

type GetStatsUseCase struct {
	repo repository.Repository
}

func NewGetStatsUseCase(repo repository.Repository) *GetStatsUseCase {
	return &GetStatsUseCase{
		repo: repo,
	}
}

func (u *GetStatsUseCase) Execute(
	ctx context.Context,
) (*domain.Statistics, error) {

	return u.repo.GetApplicationsStats(
		ctx,
	)
}
