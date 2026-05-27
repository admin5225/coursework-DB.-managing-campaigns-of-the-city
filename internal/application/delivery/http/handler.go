package http

import (
	"net/http"

	"github.com/admin5225/coursework-DB.-managing-campaigns-of-the-city/internal/application/usecase"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	createUC *usecase.CreateUseCase
	deleteUC *usecase.DeleteUseCase
	closeUC  *usecase.CloseUseCase
}

func NewHandler(createUseCase *usecase.CreateUseCase, deleteUseCase *usecase.DeleteUseCase, closeUseCase *usecase.CloseUseCase) *Handler {
	return &Handler{
		createUC: createUseCase,
		deleteUC: deleteUseCase,
		closeUC:  closeUseCase,
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

func (h *Handler) Delete(c *gin.Context) {
	var dto DeleteApplicationDTO

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
			"message":        "application deleted",
			"application_id": dto.ID,
		},
	)

}

func (h *Handler) Close(c *gin.Context) {
	var dto CloseApplicationDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	err := h.closeUC.Execute(
		c.Request.Context(),
		usecase.CloseInput{
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
			"message":        "application closed",
			"application_id": dto.ID,
		},
	)

}
