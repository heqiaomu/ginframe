package api

import (
	"github.com/Heqiaomu/ginframe/pkg/server"
	"github.com/gin-gonic/gin"
)

func DefaultRouter() []server.Route {
	return []server.Route{
		{Method: "GET", Path: "/api/v1/version", Handlers: []gin.HandlerFunc{HandleRequest(Version)}},
		{Method: "GET", Path: "/api/v1/healthcheck", Handlers: []gin.HandlerFunc{HandleRequest(HealthCheck)}},
	}
}
