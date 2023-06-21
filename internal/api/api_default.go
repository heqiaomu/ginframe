package api

import "github.com/gin-gonic/gin"

func Version(c *gin.Context) (interface{}, error) {
	return "1.0", nil
}

func HealthCheck(c *gin.Context) (interface{}, error) {
	return "success", nil
}
