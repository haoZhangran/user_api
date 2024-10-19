package service

import (
	"user_api/biz/domain/model"

	"github.com/gin-gonic/gin"
)

type UserService struct {
}

func (service *UserService) GetUserInfoByID(c *gin.Context, req *model.GetUserInfoByIDReq) (*model.GetUserInfoByIDResp, error) {
	panic("code here")
	return nil, nil
}
