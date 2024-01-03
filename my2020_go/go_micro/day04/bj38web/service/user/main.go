package main

import (
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4/registry"
	"user/handler"
	"user/model"
	pb "user/proto"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

var (
	service = "user"
	version = "latest"
)

func main() {

	// 初始化 MySQL 连接池
	model.InitDb()
	// 初始化 redis 连接池
	model.InitRedis()
	// 初始化consul
	consulReg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"192.168.150.132:8500",
		}
	})
	// Create service
	srv := micro.NewService()
	srv.Init(
		micro.Registry(consulReg), // 添加注册
		micro.Name(service),
		micro.Version(version),
	)

	// Register handler
	if err := pb.RegisterUserHandler(srv.Server(), new(handler.User)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
