package handlers

import "github.com/gin-gonic/gin"

func HealthCheck(c *gin.Context) (interface{}, int, error) {
	return "success", 200, nil
}
