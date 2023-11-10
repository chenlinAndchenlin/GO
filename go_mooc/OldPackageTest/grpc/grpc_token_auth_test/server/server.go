package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Server struct {
}

//	type GreeterServer interface {
//		Sayhello(context.Context, *HelloRequest) (*HelloReply, error)
//	}
func (s *Server) SayHello(ctx context.Context, request *helloworldProto.helloworldProto) (*helloworldProto.HelloReply, error) {
	fmt.Println(request.Name)

	return &helloworldProto.HelloReply{
		Message: "hello," + request.Name,
	}, nil
}

func main() {
	interceptor := func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		fmt.Println("接受一个新的请求")
		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			fmt.Println("get metadata ok")
		} else {
			fmt.Println("get metadata faild")
			return resp, status.Error(codes.Unauthenticated, "无token！！！")

		}
		var (
			appid  string
			appkey string
		)
		//for key, val := range md {
		//	fmt.Println(key, val)
		//}
		if appid_, ok := md["appid"]; ok {
			//fmt.Println(appid_)
			appid = appid_[0]

		}
		if appkey_, ok := md["appkey"]; ok {
			//fmt.Println(appkey_)
			appkey = appkey_[0]
		}
		if appid != "110" || appkey != "1234" {
			return resp, status.Error(codes.Unauthenticated, "无token！！！")
		}
		res, err := handler(ctx, req)
		fmt.Println("请求已经完成")
		return res, nil
	}
	//type UnaryServerInterceptor func(ctx context.Context, req any, info *UnaryServerInfo, handler UnaryHandler) (resp any, err error)
	opt := grpc.UnaryInterceptor(interceptor)
	g := grpc.NewServer(opt)
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
