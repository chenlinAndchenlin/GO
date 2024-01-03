package main

import (
	"2020_go/go_micro/day02/consul_test/pb"
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"net"
)

type Children struct {
}

//SayHello(context.Context, *Person) (*Person, error)

func (this *Children) SayHello(ctx context.Context, p *pb.Person) (*pb.Person, error) {
	p.Name = "chen " + p.Name
	p.Age = 1 + p.Age
	return p, nil
}
func main() {
	// 把grpc服务,注册到consul上.
	// 1. 初始化 consul 配置
	consulConfig := api.DefaultConfig()
	consulConfig.Address = "192.168.150.132:8500"

	// 2. 创建 consul 对象
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		fmt.Println("api.NewClient err:", err)
		return
	}
	// 3. 告诉consul, 即将注册的服务的配置信息
	reg := api.AgentServiceRegistration{
		ID:      "My_bj38",
		Tags:    []string{"grcp", "consul"},
		Name:    "grpc And Consul",
		Address: "192.168.150.1",
		Port:    9090,
		Check: &api.AgentServiceCheck{
			CheckID:  "consul grpc test",
			TCP:      "192.168.150.1:9090",
			Timeout:  "1s",
			Interval: "5s",
		},
	}
	// 4. 注册 grpc 服务到 consul 上
	err = consulClient.Agent().ServiceRegister(&reg)
	if err != nil {
		fmt.Println("注册错误，error:", err)
		return
	}

	/*
		一下是grpc远程调用实现
	*/
	//初始化grpc对象
	fmt.Println("开始初始化grpc")
	grpcServer := grpc.NewServer()
	//利用proto——register注册服务
	pb.RegisterHelloServer(grpcServer, new(Children))

	listen, err := net.Listen("tcp", "192.168.150.1:9090")
	if err != nil {
		println("Listen error:", err)
		return
	}

	defer listen.Close()

	//	启动服务

	grpcServer.Serve(listen)
	fmt.Println("结束grpc")
}
