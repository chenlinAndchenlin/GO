package main

import (
	"flag"
	"fmt"
	"mxshop_chen/user_srv/handler"
	"mxshop_chen/user_srv/proto"
	"net"

	"google.golang.org/grpc"
)

// type UserServer struct {
// }
func main() {
	//两个变量 地址和端口号 用户启动的时候 进行输入
	IP := flag.String("ip", "127.0.0.1", "输入ip地址")
	Port := flag.Int("port", 50051, "输入端口号")
	flag.Parse()
	fmt.Println("ip:", *IP)
	fmt.Println("port:", *Port)
	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic("failed to listen:" + err.Error())
	}

	err = server.Serve(listen)
	if err != nil {
		panic("failed to start server:" + err.Error())
	}
}
