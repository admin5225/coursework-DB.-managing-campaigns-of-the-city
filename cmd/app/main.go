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

	toolHttp "github.com/admin5225/coursework-DB.-managing-campaigns-of-the-city/internal/tools/delivery/http"
	toolRepo "github.com/admin5225/coursework-DB.-managing-campaigns-of-the-city/internal/tools/repository/postgres"
	toolUseCase "github.com/admin5225/coursework-DB.-managing-campaigns-of-the-city/internal/tools/usecase"

	houseHttp "github.com/admin5225/coursework-DB.-managing-campaigns-of-the-city/internal/houses/delivery/http"
	houseRepo "github.com/admin5225/coursework-DB.-managing-campaigns-of-the-city/internal/houses/repository/postgres"
	houseUseCase "github.com/admin5225/coursework-DB.-managing-campaigns-of-the-city/internal/houses/usecase"
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
	toolRepository := toolRepo.NewRepository(db)
	houseRepository := houseRepo.NewRepository(db)

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
	applicationGetClosedUC := applicationUseCase.NewGetClosedUseCase(
		applicationRepository,
	)
	applicationGetStatisticsUC := applicationUseCase.NewGetStatsUseCase(
		applicationRepository,
	)
	applicationGetByHouseUC := applicationUseCase.NewGetByHouseUseCase(
		applicationRepository,
	)
	applicationGetBySpecialistUC := applicationUseCase.NewGetBySpecialistUseCase(
		applicationRepository,
	)

	// specialist usecases
	specialistCreateUC := specialistUseCase.NewCreateUseCase(
		specialistRepository,
	)
	specialistDeleteUC := specialistUseCase.NewDeleteUseCase(
		specialistRepository,
	)

	// tool usecases
	toolCreateUC := toolUseCase.NewCreateUseCase(
		toolRepository,
	)
	toolDeleteUC := toolUseCase.NewDeleteUseCase(
		toolRepository,
	)
	toolUpdateUC := toolUseCase.NewUpdateUseCase(
		toolRepository,
	)

	// house usecases
	houseCreateUC := houseUseCase.NewCreateUseCase(
		houseRepository,
	)
	houseDeleteUC := houseUseCase.NewDeleteUseCase(
		houseRepository,
	)

	// handlers
	applicationHandler := applicationHttp.NewHandler(
		applicationCreateUC,
		applicationDeleteUC,
		applicationCloseUC,
		applicationGetClosedUC,
		applicationGetStatisticsUC,
		applicationGetByHouseUC,
		applicationGetBySpecialistUC,
	)

	specialistHandler := specialistHttp.NewHandler(
		specialistCreateUC,
		specialistDeleteUC,
	)

	toolHandler := toolHttp.NewHandler(
		toolCreateUC,
		toolDeleteUC,
		toolUpdateUC,
	)

	houseHandler := houseHttp.NewHandler(
		houseCreateUC,
		houseDeleteUC,
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

	toolHttp.RegisterToolRoutes(
		api,
		toolHandler,
	)

	houseHttp.RegisterHouseRoutes(
		api,
		houseHandler,
	)

	log.Fatal(router.Run(":8080"))
}
