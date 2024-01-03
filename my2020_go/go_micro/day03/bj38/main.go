package main

import (
	"bj38/handler"
	pb "bj38/proto"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4/registry"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

var (
	service = "bj38"
	version = "latest"
)

func main() {

	// 初始化服务发现
	consulReg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"192.168.150.132:8500",
		}
	})

	// Create service
	srv := micro.NewService()
	srv.Init(
		micro.Name(service),
		micro.Registry(consulReg),
		micro.Version(version),
	)

	// Register handler注册服务
	if err := pb.RegisterBj38Handler(srv.Server(), new(handler.Bj38)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
