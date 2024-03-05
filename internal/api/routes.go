package api

import "github.com/gin-gonic/gin"

func Router(group *gin.RouterGroup) *gin.RouterGroup {
	apiV1 := group.Group("api/v1")
	apiV1.GET("/healthcheck", HandleRequest(HealthCheck))

	return group
}
