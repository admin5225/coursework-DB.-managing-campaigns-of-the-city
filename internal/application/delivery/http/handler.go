package http

import (
	"log"
	"net/http"

	"github.com/admin5225/coursework-DB.-managing-campaigns-of-the-city/internal/application/usecase"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	createUC *usecase.CreateUseCase
}

func NewHandler(createUC *usecase.CreateUseCase) *Handler {
	return &Handler{
		createUC: createUC,
	}
}

func (h *Handler) Create(c *gin.Context) {
	var dto CreateApplicationDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "error",
			},
		)

		return
	}

	err := h.createUC.Execute(
		c.Request.Context(),
		usecase.CreateInput{
			Description:  dto.Description,
			HouseID:      dto.HouseID,
			SpecialistID: dto.SpecialistID,
			WorkTypeID:   dto.WorkTypeID,
			StatusID:     dto.StatusID,
		},
	)

	if err != nil {
		log.Printf("Error in CreateUseCase.Execute: %v", err)
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": "StatusInternalServerError",
			},
		)

		return
	}

	c.JSON(
		http.StatusCreated,
		gin.H{
			"message": "request creared",
		},
	)
}
