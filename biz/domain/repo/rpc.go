package repo

import (
	"context"

	"github.com/haoZhangran/common_sdk/sql/zproject/z_entity"
)

type IUserServiceRpcClient interface {
	GetUserInfoByID(ctx context.Context, id string) (*z_entity.User, error) // TODO(@yua) 需要先建表、确定User结构体
}
