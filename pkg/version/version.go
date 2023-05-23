package version

import "github.com/gin-gonic/gin"

func Handler(c *gin.Context) (interface{}, error) {
	return "1.0", nil
}
