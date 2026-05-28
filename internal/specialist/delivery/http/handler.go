package http

import (
	"log"
	"net/http"

	"github.com/admin5225/coursework-DB.-managing-campaigns-of-the-city/internal/specialist/usecase"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	createUC *usecase.CreateUseCase
	deleteUC *usecase.DeleteUseCase
}

func NewHandler(createUC *usecase.CreateUseCase, deleteUC *usecase.DeleteUseCase) *Handler {
	return &Handler{
		createUC: createUC,
		deleteUC: deleteUC,
	}
}

func (h *Handler) Create(c *gin.Context) {
	var dto CreateSpecialistDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "StatusBadRequest",
			},
		)

		return
	}

	err := h.createUC.Execute(
		c.Request.Context(),
		usecase.CreateInput{
			FullName:           dto.FullName,
			Position:           dto.Position,
			PhoneNumber:        dto.PhoneNumber,
			ManagingCamaiginID: dto.ManagingCamaiginID,
		},
	)

	if err != nil {
		log.Print(err)
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
			"message": "specialist created",
		},
	)
}

func (h *Handler) Delete(c *gin.Context) {
	var dto DeleteSpecialistDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	err := h.deleteUC.Execute(
		c.Request.Context(),
		usecase.DeleteInput{
			ID: dto.ID,
		},
	)

	if err != nil {
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
			"message":       "specialist deleted",
			"specialist_id": dto.ID,
		},
	)

}
