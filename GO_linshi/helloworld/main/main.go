package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"hellowrold/proto"
)

func main() {
	fmt.Println("chen")
	req := myproto.HelloRequest{
		Name: "chenlin",
	}
	rsq, _ := proto.Marshal(&req)
	fmt.Println(string(rsq))
}
