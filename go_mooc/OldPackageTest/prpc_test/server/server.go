package main

import (
	"OldPackageTest/prpc_test/proto"
	"context"
	"google.golang.org/grpc"
	"net"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, request *helloworldProto.HelloRequest) (*helloworldProto.HelloReply, error) {
	return &helloworldProto.HelloReply{
		Message: "hello," + request.Name,
	}, nil
}
func main() {
	g := grpc.NewServer()
	helloworldProto.RegisterGreeterServer(g, &Server{})
	lis, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	err = g.Serve(lis)
	if err != nil {
		panic("failed ti start grpc:" + err.Error())
	}
}
