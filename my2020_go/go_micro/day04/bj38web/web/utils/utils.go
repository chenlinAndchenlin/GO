package utils

import (
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
)

// 初始化micro

func InitMicro() micro.Service {
	// 指定 consul 服务发现
	consulReg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"192.168.150.132:8500",
		}
	})
	//consulReg := consul.NewRegistry()
	return micro.NewService(
		micro.Registry(consulReg),
	)
}
