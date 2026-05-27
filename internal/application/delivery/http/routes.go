package http

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, handler *Handler) {
	applications := router.Group("/applications")

	{
		applications.POST("/", handler.Create)
		applications.POST("/delete", handler.Delete)
		applications.POST("/close", handler.Close)
	}
}
