package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HandlerFunc func(c *gin.Context) (interface{}, error)

func HandleRequest(handler HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := handler(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
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
