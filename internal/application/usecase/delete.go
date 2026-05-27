package usecase

import (
	"context"

	"github.com/admin5225/coursework-DB.-managing-campaigns-of-the-city/internal/application/repository"
)

type DeleteUseCase struct {
	repo repository.Repository
}

func NewDeleteUseCase(repo repository.Repository) *DeleteUseCase {
	return &DeleteUseCase{
		repo: repo,
	}
}

type DeleteInput struct {
	ID int
}

func (u *DeleteUseCase) Execute(ctx context.Context, input DeleteInput) error {
	return u.repo.Delete(ctx, input.ID)
}
