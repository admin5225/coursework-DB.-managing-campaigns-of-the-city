package http

import "github.com/gin-gonic/gin"

func RegisterSpecialistRoutes(router *gin.RouterGroup, handler *Handler) {
	specialist := router.Group("/specialists")

	{
		specialist.POST("/", handler.Create)
		specialist.POST("/delete", handler.Delete)
	}
}
