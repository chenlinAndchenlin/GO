package main

import (
	"2020_go/go_micro/day02/consul_test/pb"
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"strconv"
)

func main() {

	// 初始化 consul 配置
	consulConfig := api.DefaultConfig()
	consulConfig.Address = "192.168.150.132:8500"
	// 创建consul对象 -- (可以重新指定 consul 属性: IP/Port , 也可以使用默认)
	consulClient, err := api.NewClient(consulConfig)
	// 服务发现. 从consuL上, 获取健康的服务
	services, _, err := consulClient.Health().Service("grpc And Consul", "grcp", true, nil)
	fmt.Println("*****************************************")
	fmt.Println("services:", services[0].Service)
	fmt.Println("*****************************************")

	// 简单的负载均衡.
	addr := services[0].Service.Address + ":" + strconv.Itoa(services[0].Service.Port)
	/*
		一下是grpc远程调用实现
	*/
	//grpc 拨号
	grpcConn, err := grpc.Dial(addr, grpc.WithInsecure())
	//grpcConn, err := grpc.Dial("192.168.150.132:8800", grpc.WithInsecure())
	if err != nil {
		return
	}
	// pb包初始化grpc客户端
	grpcClient := pb.NewHelloClient(grpcConn)

	var p pb.Person
	p.Name = "itcast"
	p.Age = 18
	P, err := grpcClient.SayHello(context.TODO(), &p)
	if err != nil {
		return
	}

	fmt.Println(P, "**************", err)

}
