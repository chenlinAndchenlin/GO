package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	bj38 "tidy/go_micro/day03/web/proto"
)

func CallRemote(ctx *gin.Context) {
	// 初始化服务发现 consul
	consulReg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"192.168.150.132:8500",
		}
	})

	// 初始化micro服务对象, 指定consul 为服务发现
	service := micro.NewService(
		micro.Registry(consulReg),
	)
	microClient := bj38.NewBj38Service("bj38", service.Client())
	//microClient := bj38.NewBj38Service("bj38", client.DefaultClient)
	fmt.Println()
	resp, err := microClient.Call(context.TODO(), &bj38.CallRequest{
		Name: "xiaowang",
	})
	if err != nil {
		fmt.Println("call err:", err)
		return
	}
	// 为了方便查看, 在打印之前将结果返回给浏览器
	ctx.Writer.WriteString(resp.Msg)

	fmt.Println(resp, err)

}
func main() {

	router := gin.Default()
	/*router.GET("/", func(context *gin.Context) {
		context.Writer.WriteString("hello chenlin")
		fmt.Println("gin test")
	})*/
	router.GET("/", CallRemote)
	router.Run("192.168.150.1:9090")
}
