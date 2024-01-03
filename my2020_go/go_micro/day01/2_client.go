package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	conn, err := rpc.Dial("tcp", ":8800")
	if err != nil {
		fmt.Println("Dial err:", err)
		return
	}
	defer conn.Close()

	var reply string

	err = conn.Call("hello.HelloWorld", "chenlin", &reply)
	if err != nil {
		fmt.Println("Call err:", err)
		return
	}

	fmt.Println(reply)
}
