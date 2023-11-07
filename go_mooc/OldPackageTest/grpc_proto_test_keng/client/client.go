package main

import (
	"OldPackageTest/grpc_proto_test_keng/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:50053", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	//type GreeterClient interface {
	//	Sayhello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	//}
	c := proto.NewGreeterClient(conn)
	sayhello, err := c.Sayhello(context.Background(), &proto.HelloRequest{
		Data:    "chenlin",
		Url:     "www.baidu.com",
		G:       proto.Gender_FEMALE,
		Mp:      map[string]string{"name": "boby", "chen": "chenlinguo"},
		AddTime: timestamppb.New(time.Now()),
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(sayhello.Data)
}
