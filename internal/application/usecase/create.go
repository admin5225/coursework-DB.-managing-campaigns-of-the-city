package usecase

import (
	"context"

	"github.com/admin5225/coursework-DB.-managing-campaigns-of-the-city/internal/application/domain"
	"github.com/admin5225/coursework-DB.-managing-campaigns-of-the-city/internal/application/repository"
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
	Description  string
	HouseID      int
	SpecialistID int
	WorkTypeID   int
	StatusID     int
}

func (u *CreateUseCase) Execute(ctx context.Context, input CreateInput) error {
	application := &domain.Application{
		Description:  input.Description,
		HouseID:      input.HouseID,
		SpecialistID: input.SpecialistID,
		WorkTypeID:   input.WorkTypeID,
		StatusID:     input.StatusID,
	}

	return u.repo.Create(ctx, application)
}
