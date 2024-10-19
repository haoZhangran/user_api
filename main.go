package main

import (
	"context"
	"fmt"
	"time"
	"user_api/biz/infra/rpc"
)

//func main() {
//	r := gin.Default()
//	r.GET("/hello", view.Hello)
//
//	r.GET("/ping", view.Pong)
//
//	r.Run(":9480")
//}

func main() {
	for {
		time.Sleep(time.Second)
		r := rpc.InitUserServiceRPCImpl()
		userInfo, err := r.GetUserInfoByID(context.Background(), "1")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("time is ", time.Now().Format(time.DateTime))
		fmt.Println(userInfo)
	}

}
