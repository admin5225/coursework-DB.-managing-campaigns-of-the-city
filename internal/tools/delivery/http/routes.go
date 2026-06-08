package http

import "github.com/gin-gonic/gin"

func RegisterToolRoutes(router *gin.RouterGroup, handler *Handler) {
	tool := router.Group("/tools")

	{
		tool.POST("/", handler.Create)
		tool.POST("/delete", handler.Delete)
		tool.POST("/update", handler.Update)
	}
}
