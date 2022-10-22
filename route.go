package main

import (
	"user_api/view"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) {
	r.GET("/hello", view.Hello)
}
