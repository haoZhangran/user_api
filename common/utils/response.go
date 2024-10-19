package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	OkCode          = 1
	QueryFailedCode = 1001
	InternalFailed  = 1002
)

func MakeErrorResponse(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    msg,
	})
}

func MakeOKResponse(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{
		"status": OkCode,
		"data":   data,
	})
}
