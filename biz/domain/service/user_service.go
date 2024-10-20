package service

import (
	"fmt"
	"user_api/biz/domain/model"
	"user_api/biz/domain/repo"

	"github.com/gin-gonic/gin"
)

type UserService struct {
	rpc repo.IUserServiceRpcClient
}

func NewUserService(rpc repo.IUserServiceRpcClient) *UserService {
	return &UserService{
		rpc: rpc,
	}
}

func (service *UserService) GetUserInfoByID(c *gin.Context, req *model.GetUserInfoByIDReq) (*model.GetUserInfoByIDResp, error) {
	info, err := service.rpc.GetUserInfoByID(c, req.UserID)
	if err != nil {
		return nil, err
	}
	return &model.GetUserInfoByIDResp{
		UserID:   fmt.Sprintf("%d", info.Id),
		UserName: info.Name,
	}, nil
}
