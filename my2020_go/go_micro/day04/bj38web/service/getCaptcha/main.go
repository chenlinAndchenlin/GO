package main

import (
	"getCaptcha/handler"
	pb "getCaptcha/proto"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4/registry"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

var (
	service = "getcaptcha"
	version = "latest"
)

func main() {
	// 初始化consul
	consulReg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"192.168.150.132:8500",
		}
	})
	// Create service
	srv := micro.NewService()
	srv.Init(
		//micro.Address("192.168.150.132:8880"), // 防止随机生成 port
		micro.Registry(consulReg), // 添加注册
		micro.Name(service),
		micro.Version(version),
	)

	// Register handler
	if err := pb.RegisterGetCaptchaHandler(srv.Server(), new(handler.GetCaptcha)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
