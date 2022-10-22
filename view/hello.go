package view

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/haoranz527/gocommon/packer"
)

func Hello(c *gin.Context) {
	packer.MakeOkResponse(c, "hello")
}
