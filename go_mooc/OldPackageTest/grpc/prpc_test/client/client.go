package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithInsecure())
	if err != nil {
		panic("conn to faild:" + err.Error())
	}
	defer conn.Close()

	c := helloworldProto.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &helloworldProto.HelloRequest{Name: "chenlin"})
	if err != nil {
		panic("faild" + err.Error())
	}

	fmt.Println(r.Message)

}
