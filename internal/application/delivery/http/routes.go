package http

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, handler *Handler) {
	applications := router.Group("/applications")

	{
		applications.POST("/", handler.Create)
		applications.POST("/delete", handler.Delete)
		applications.POST("/close", handler.Close)
		applications.GET("/closed", handler.GetClosed)
		applications.GET("/statistics", handler.GetStatistics)
		applications.GET("/house/:house_id", handler.GetByHouse)
		applications.GET("/specialist/:specialist_id", handler.GetBySpecialist)
	}
}
