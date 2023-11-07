package main

import (
	"OldPackageTest/metadata_grpc/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"
)

const (
	timestampformat = time.StampNano
)

func main() {
	Interceptor := func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		if err != nil {
			return err
		}
		fmt.Printf("时间：%s\n", time.Since(start))
		return err
	}

	opt := grpc.WithUnaryInterceptor(Interceptor)
	conn, err := grpc.Dial("127.0.0.1:50053", grpc.WithInsecure(), opt)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	//type GreeterClient interface {
	//	Sayhello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	//}
	//md := metadata.Pairs("timestamp", time.Now().Format(timestampformat))

	md := metadata.New(map[string]string{
		"name":     "bobby",
		"pawwword": "imooc",
	})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	c := helloworldProto.NewGreeterClient(conn)

	sayhello, err := c.SayHello(ctx, &helloworldProto.HelloRequest{
		Name: "chenlin",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(sayhello.Message)
}
