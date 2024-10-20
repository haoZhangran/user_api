package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	Route(r)

	r.Run(":9480")
}

//func main() {
//	for {
//		time.Sleep(time.Second)
//		r := rpc.InitUserServiceRPCImpl()
//		userInfo, err := r.GetUserInfoByID(context.Background(), "356")
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		fmt.Println("time is ", time.Now().Format(time.DateTime))
//		fmt.Println(userInfo)
//	}
//}
