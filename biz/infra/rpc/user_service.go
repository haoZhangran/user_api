package rpc

import (
	"context"
	"fmt"
	"strconv"
	"user_api/common/consts"

	"google.golang.org/grpc/credentials/insecure"

	"github.com/haoZhangran/common_sdk/sql/zproject/z_entity"

	"github.com/haoZhangran/common_sdk/protoc_gen_code/zproject/user_service"
	"google.golang.org/grpc"
)

type UserServiceRPCImpl struct {
	client user_service.UserServiceClient
}

func InitUserServiceRPCImpl() *UserServiceRPCImpl {
	conn, err := grpc.NewClient(consts.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
		return nil
	}
	c := user_service.NewUserServiceClient(conn)
	return &UserServiceRPCImpl{
		client: c,
	}
}

func (u UserServiceRPCImpl) GetUserInfoByID(ctx context.Context, id string) (*z_entity.User, error) {
	req := &user_service.GetUserInfoByIDRequest{
		ID: id,
	}
	resp, err := u.client.GetUserInfoByID(ctx, req)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	userInfo := &z_entity.User{
		Name: resp.GetName(),
		Id: func() int {
			id, err := strconv.ParseInt(resp.GetID(), 10, 64)
			if err != nil {
				return 0
			}
			return int(id)
		}(),
	}
	return userInfo, nil
}
