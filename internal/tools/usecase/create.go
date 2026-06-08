package usecase

import (
	"context"

	"github.com/admin5225/coursework-DB.-managing-campaigns-of-the-city/internal/tools/domain"
	"github.com/admin5225/coursework-DB.-managing-campaigns-of-the-city/internal/tools/repository"
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
	Name                string
	ManagingCampaiginID int
	Quantity            int
}

func (u *CreateUseCase) Execute(ctx context.Context, input CreateInput) error {
	tool := &domain.Tool{
		Name:                input.Name,
		ManagingCampaiginID: input.ManagingCampaiginID,
		Quantity:            input.Quantity,
	}

	return u.repo.Create(ctx, tool)
}
