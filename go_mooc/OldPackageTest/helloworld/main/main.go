package main

import (
	helloworld "OldPackageTest/helloworld/protobuf"
	"fmt"

	"github.com/golang/protobuf/proto"
)

func main() {
	req := helloworld.HelloRequest{
		Name: "bobby",
	}
	rsp, _ := proto.Marshal(&req)
	fmt.Println(string(rsp))
}
