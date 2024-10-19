package app

import (
	"fmt"
	"user_api/biz/domain/model"
	"user_api/biz/domain/service"
	"user_api/common/utils"

	"github.com/gin-gonic/gin"
)

type UserApp struct {
	userService *service.UserService
}

func NewUserApp(userService *service.UserService) *UserApp {
	return &UserApp{
		userService: userService,
	}
}

func (app *UserApp) GetUserInfoByUserID(c *gin.Context) {
	req, err := parseGetUserInfoByUserIDReq(c)
	if err != nil {
		fmt.Println(err)
		utils.MakeErrorResponse(c, utils.QueryFailedCode, err.Error())
		return
	}

	resp, err := app.userService.GetUserInfoByID(c, req)
	if err != nil {
		utils.MakeErrorResponse(c, utils.InternalFailed, err.Error())
		return
	}
	utils.MakeOKResponse(c, resp)
}

func parseGetUserInfoByUserIDReq(c *gin.Context) (*model.GetUserInfoByIDReq, error) {
	req := &model.GetUserInfoByIDReq{}

	userID := c.GetString("id")
	req.UserID = userID

	return req, nil
}
