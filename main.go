package main

import (
	"user_api/view"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello", view.Hello)

	r.GET("/ping", view.Pong)

	r.Run(":9480")
}
