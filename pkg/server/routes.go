package server

import (
	"github.com/Heqiaomu/ginframe/pkg/version"
	"github.com/gin-gonic/gin"
)

func DefaultRouter() []Route {
	return []Route{
		{Method: "GET", Path: "/api/v1/version", Handlers: []gin.HandlerFunc{HandleRequest(version.Handler)}},
	}
}
