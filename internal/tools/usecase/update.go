package usecase

import (
	"context"

	"github.com/admin5225/coursework-DB.-managing-campaigns-of-the-city/internal/tools/repository"
)

type UpdateUseCase struct {
	repo repository.Repository
}

func NewUpdateUseCase(repo repository.Repository) *UpdateUseCase {
	return &UpdateUseCase{
		repo: repo,
	}
}

type UpdateInput struct {
	ID       int
	Quantity int
}

func (u *UpdateUseCase) Execute(ctx context.Context, input UpdateInput) error {
	id := input.ID
	quantity := input.Quantity

	return u.repo.Update(ctx, id, quantity)
}
