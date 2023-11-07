package main

import (
	"OldPackageTest/metadata_grpc/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net"
)

type Server struct {
}

//	type GreeterServer interface {
//		Sayhello(context.Context, *HelloRequest) (*HelloReply, error)
//	}
func (s *Server) SayHello(ctx context.Context, request *helloworldProto.HelloRequest) (*helloworldProto.HelloReply, error) {
	fmt.Println(request.Name)
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		fmt.Println("get metadata ok")
	}
	for key, val := range md {
		fmt.Println(key, val)
	}
	if nameSlice, ok := md["name"]; ok {
		fmt.Println(nameSlice)
		for i, e := range nameSlice {
			fmt.Println(i, e)
		}
	}
	return &helloworldProto.HelloReply{
		Message: "hello," + request.Name,
	}, nil
}

func main() {
	g := grpc.NewServer()
	helloworldProto.RegisterGreeterServer(g, &Server{})

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
