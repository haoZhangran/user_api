package repo

import (
	"context"

	"github.com/haoZhangran/common_sdk/sql/zproject/z_entity"
)

type IUserServiceRpcClient interface {
	GetUserInfoByID(ctx context.Context, id string) (*z_entity.User, error)
}
