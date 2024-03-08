package routes

import (
	"github.com/Heqiaomu/ginframe/api/handlers"
	"github.com/gin-gonic/gin"
)

func Router(rootGroup *gin.RouterGroup) *gin.RouterGroup {
	apiV1 := rootGroup.Group("api/v1")
	apiV1.GET("/healthcheck", handlers.HandleRequest(handlers.HealthCheck))

	return rootGroup
}
