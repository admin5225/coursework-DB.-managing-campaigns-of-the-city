package usecase

import (
	"context"

	"github.com/admin5225/coursework-DB.-managing-campaigns-of-the-city/internal/specialist/domain"
	"github.com/admin5225/coursework-DB.-managing-campaigns-of-the-city/internal/specialist/repository"
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
	FullName           string
	Position           string
	PhoneNumber        string
	ManagingCamaiginID int
}

func (u *CreateUseCase) Execute(ctx context.Context, input CreateInput) error {
	specialist := &domain.Specialist{
		FullName:           input.FullName,
		Position:           input.Position,
		PhoneNumber:        input.PhoneNumber,
		ManagingCamaiginID: input.ManagingCamaiginID,
	}

	return u.repo.Create(ctx, specialist)
}
