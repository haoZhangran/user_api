package repo

import "context"

type IUserServiceRpcClient interface {
	GetUserInfoByID(ctx context.Context, id string) // TODO(@yua) 需要先建表、确定User结构体
}
