package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerFunc func(c *gin.Context) (interface{}, int, error)

func HandleRequest(handler HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, code, err := handler(c)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": err.Error(),
				"data":    nil,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "OK",
			"data":    data,
		})
	}
}
