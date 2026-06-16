package usecase

import (
	"context"

	"github.com/admin5225/coursework-DB.-managing-campaigns-of-the-city/internal/application/domain"
	"github.com/admin5225/coursework-DB.-managing-campaigns-of-the-city/internal/application/repository"
)

type GetByHouseUseCase struct {
	repo repository.Repository
}

func NewGetByHouseUseCase(repo repository.Repository) *GetByHouseUseCase {
	return &GetByHouseUseCase{
		repo: repo,
	}
}

func (u *GetByHouseUseCase) Execute(
	ctx context.Context,
	houseID int,
) ([]domain.FullApplication, error) {

	return u.repo.GetHouseApplications(
		ctx,
		houseID,
	)
}
