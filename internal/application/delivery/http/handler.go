package http

import (
	"net/http"
	"strconv"

	"github.com/admin5225/coursework-DB.-managing-campaigns-of-the-city/internal/application/usecase"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	createUC          *usecase.CreateUseCase
	deleteUC          *usecase.DeleteUseCase
	closeUC           *usecase.CloseUseCase
	getClosedUC       *usecase.GetClosedUseCase
	getStatsUC        *usecase.GetStatsUseCase
	getByHouseUC      *usecase.GetByHouseUseCase
	getBySpecialistUC *usecase.GetBySpecialistUseCase
}

func NewHandler(createUseCase *usecase.CreateUseCase, deleteUseCase *usecase.DeleteUseCase, closeUseCase *usecase.CloseUseCase, getClosed *usecase.GetClosedUseCase, getStats *usecase.GetStatsUseCase, getByHouse *usecase.GetByHouseUseCase, getBySpecialist *usecase.GetBySpecialistUseCase) *Handler {
	return &Handler{
		createUC:          createUseCase,
		deleteUC:          deleteUseCase,
		closeUC:           closeUseCase,
		getClosedUC:       getClosed,
		getStatsUC:        getStats,
		getByHouseUC:      getByHouse,
		getBySpecialistUC: getBySpecialist,
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

func (h *Handler) GetClosed(
	c *gin.Context,
) {

	applications, err := h.getClosedUC.Execute(
		c.Request.Context(),
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
		http.StatusOK,
		applications,
	)
}

func (h *Handler) GetStatistics(
	c *gin.Context,
) {

	stats, err := h.getStatsUC.Execute(
		c.Request.Context(),
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
		http.StatusOK,
		stats,
	)
}

func (h *Handler) GetByHouse(
	c *gin.Context,
) {
	id64, err := strconv.ParseInt(
		c.Param("house_id"),
		10,
		64,
	)
	houseID := int(id64)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "invalid id",
			},
		)

		return
	}

	applications, err := h.getByHouseUC.Execute(
		c.Request.Context(),
		houseID,
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
		http.StatusOK,
		applications,
	)
}

func (h *Handler) GetBySpecialist(
	c *gin.Context,
) {
	id64, err := strconv.ParseInt(
		c.Param("specialist_id"),
		10,
		64,
	)
	specialistID := int(id64)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "invalid id",
			},
		)

		return
	}

	applications, err := h.getBySpecialistUC.Execute(
		c.Request.Context(),
		specialistID,
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
		http.StatusOK,
		applications,
	)
}
