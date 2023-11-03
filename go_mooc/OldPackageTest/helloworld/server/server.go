package main

import (
	"net"
	"net/rpc"
)

type HelloServer struct {
}

func (s *HelloServer) Hello(request string, reply *string) error {
	*reply = "hello," + request
	return nil
}
func main() {
	//1 实例化一个server
	listerner, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		return
	}
	//2 注册一个处理逻辑handler
	//内部其实是Helloserver.Hello name 就是远程连接的id
	err = rpc.RegisterName("HelloServer", &HelloServer{})
	if err != nil {
		return
	}
	//3 启动服务
	conn, _ := listerner.Accept() //当一个新的连接进来的时候，
	rpc.ServeConn(conn)
	//一连串的代码大部分是net 好像和rpc没有关系
	//几个问题 1 call id 2序列化和反序列化 是rpc解决的
}
