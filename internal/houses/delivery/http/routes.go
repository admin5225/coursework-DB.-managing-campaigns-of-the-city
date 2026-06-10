package http

import "github.com/gin-gonic/gin"

func RegisterHouseRoutes(router *gin.RouterGroup, handler *Handler) {
	tool := router.Group("/houses")

	{
		tool.POST("/", handler.Create)
		tool.POST("/delete", handler.Delete)
	}
}
