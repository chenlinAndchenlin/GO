package main

import (
	"OldPackageTest/grpc/grpc_proto_test_keng/proto"
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
}

//	type GreeterServer interface {
//		Sayhello(context.Context, *HelloRequest) (*HelloReply, error)
//	}
func (s *Server) Sayhello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	fmt.Println(request.Data)
	return &proto.HelloReply{
		Data: "hello" + request.Data,
	}, nil
}

func main() {
	g := grpc.NewServer()
	proto.RegisterGreeterServer(g, &Server{})

	lis, err := net.Listen("tcp", "127.0.0.1:50053")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = g.Serve(lis)
	if err != nil {
		fmt.Println(err)
		return
	}
}
