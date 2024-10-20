package main

import (
	"user_api/biz/app"

	"github.com/gin-gonic/gin"
)

var (
	userAPP *app.UserApp
)

// TODO(@yua) 初始化app

func Route(r *gin.Engine) {
	r.GET("/z100/user/info/query", userAPP.GetUserInfoByUserID)
}
