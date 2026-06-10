package usecase

import (
	"context"

	"github.com/admin5225/coursework-DB.-managing-campaigns-of-the-city/internal/houses/domain"
	"github.com/admin5225/coursework-DB.-managing-campaigns-of-the-city/internal/houses/repository"
)

type CreateUseCase struct {
	repo repository.Repository
}

func NewCreateUseCase(repo repository.Repository) *CreateUseCase {
	return &CreateUseCase{
		repo: repo,
	}
}

type CreateInput struct {
	Street              string
	HouseNumber         int
	EntrancesNumber     int
	ManagingCampaiginId int
}

func (u *CreateUseCase) Execute(ctx context.Context, input CreateInput) error {
	house := &domain.House{
		Street:              input.Street,
		HouseNumber:         input.HouseNumber,
		EntrancesNumber:     input.EntrancesNumber,
		ManagingCampaiginId: input.ManagingCampaiginId,
	}

	return u.repo.Create(ctx, house)
}
