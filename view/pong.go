package view

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg":  "pong",
		"code": 1,
	})
}
