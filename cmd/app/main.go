package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	applicationHttp "github.com/admin5225/coursework-DB.-managing-campaigns-of-the-city/internal/application/delivery/http"
	applicationRepo "github.com/admin5225/coursework-DB.-managing-campaigns-of-the-city/internal/application/repository/postgres"
	applicationUseCase "github.com/admin5225/coursework-DB.-managing-campaigns-of-the-city/internal/application/usecase"

	specialistHttp "github.com/admin5225/coursework-DB.-managing-campaigns-of-the-city/internal/specialist/delivery/http"
	specialistRepo "github.com/admin5225/coursework-DB.-managing-campaigns-of-the-city/internal/specialist/repository/postgres"
	specialistUseCase "github.com/admin5225/coursework-DB.-managing-campaigns-of-the-city/internal/specialist/usecase"
)

func main() {

	ctx := context.Background()

	db, err := pgxpool.New(
		ctx,
		"postgres://my-user:123123@localhost:5432/managing_campaigns",
	)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// repository
	applicationRepository := applicationRepo.NewRepository(db)
	specialistRepository := specialistRepo.NewRepository(db)

	// application usecases
	applicationCreateUC := applicationUseCase.NewCreateUseCase(
		applicationRepository,
	)
	applicationDeleteUC := applicationUseCase.NewDeleteUseCase(
		applicationRepository,
	)
	applicationCloseUC := applicationUseCase.NewCloseUseCase(
		applicationRepository,
	)

	// specialist usecases
	specialistCreateUC := specialistUseCase.NewCreateUseCase(
		specialistRepository,
	)
	specialistDeleteUC := specialistUseCase.NewDeleteUseCase(
		specialistRepository,
	)

	// handlers
	applicationHandler := applicationHttp.NewHandler(
		applicationCreateUC,
		applicationDeleteUC,
		applicationCloseUC,
	)

	specialistHandler := specialistHttp.NewHandler(
		specialistCreateUC,
		specialistDeleteUC,
	)

	router := gin.Default()

	api := router.Group("/api")

	applicationHttp.RegisterRoutes(
		api,
		applicationHandler,
	)

	specialistHttp.RegisterSpecialistRoutes(
		api,
		specialistHandler,
	)

	log.Fatal(router.Run(":8080"))
}
