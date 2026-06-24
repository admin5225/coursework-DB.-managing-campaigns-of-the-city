package http

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/admin5225/coursework-DB.-managing-campaigns-of-the-city/internal/tools/usecase"
)

type Handler struct {
	createUC *usecase.CreateUseCase
	deleteUC *usecase.DeleteUseCase
	updateUC *usecase.UpdateUseCase
}

func NewHandler(createUC *usecase.CreateUseCase, deleteUC *usecase.DeleteUseCase, updateUC *usecase.UpdateUseCase) *Handler {
	return &Handler{
		createUC: createUC,
		deleteUC: deleteUC,
		updateUC: updateUC,
	}
}

func (h *Handler) Create(c *gin.Context) {
	var dto CreateToolDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	err := h.createUC.Execute(
		c.Request.Context(),
		usecase.CreateInput{
			Name:                dto.Name,
			ManagingCampaiginID: dto.ManagingCampaiginId,
			Quantity:            dto.Quantity,
		},
	)

	if err != nil {
		log.Print(err)
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(
		http.StatusCreated,
		gin.H{
			"message": "tool created",
		},
	)
}

func (h *Handler) Delete(c *gin.Context) {
	var dto DeleteToolDTO

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
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(
		http.StatusCreated,
		gin.H{
			"message": "tool deleted",
			"tool_id": dto.ID,
		},
	)

}

func (h *Handler) Update(c *gin.Context) {
	var dto UpdateToolDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	err := h.updateUC.Execute(
		c.Request.Context(),
		usecase.UpdateInput{
			ID:       dto.ID,
			Quantity: dto.Quantity,
		},
	)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(
		http.StatusCreated,
		gin.H{
			"message":  "tool updeted",
			"tool_id":  dto.ID,
			"quantity": dto.Quantity,
		},
	)

}
