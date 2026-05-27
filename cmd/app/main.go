package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	applicationHttp "github.com/admin5225/coursework-DB.-managing-campaigns-of-the-city/internal/application/delivery/http"
	applicationRepo "github.com/admin5225/coursework-DB.-managing-campaigns-of-the-city/internal/application/repository/postgres"
	applicationUseCase "github.com/admin5225/coursework-DB.-managing-campaigns-of-the-city/internal/application/usecase"
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

	// usecases
	createUC := applicationUseCase.NewCreateUseCase(
		applicationRepository,
	)

	// handlers
	requestHandler := applicationHttp.NewHandler(
		createUC,
	)

	router := gin.Default()

	api := router.Group("/api")

	applicationHttp.RegisterRoutes(
		api,
		requestHandler,
	)

	log.Fatal(router.Run(":8080"))
}
