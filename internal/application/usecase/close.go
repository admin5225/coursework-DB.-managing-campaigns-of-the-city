package usecase

import (
	"context"

	"github.com/admin5225/coursework-DB.-managing-campaigns-of-the-city/internal/application/repository"
)

type CloseUseCase struct {
	repo repository.Repository
}

func NewCloseUseCase(repo repository.Repository) *CloseUseCase {
	return &CloseUseCase{
		repo: repo,
	}
}

type CloseInput struct {
	ID int
}

func (u *CloseUseCase) Execute(ctx context.Context, input CloseInput) error {

	return u.repo.Close(ctx, input.ID)
}
