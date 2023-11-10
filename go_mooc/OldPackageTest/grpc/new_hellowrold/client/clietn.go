package main

import (
	"OldPackageTest/grpc/new_hellowrold/client_proxy"
	"fmt"
)

//	func main() {
//		//1 建立连接 拨号
//		client, err := rpc.Dial("tcp", "127.0.0.1:9090")
//		if err != nil {
//			println(err)
//			panic("连接失败")
//		}
//		//new 分配空间 只有地址的 可以直接赋值
//		//
//		var reply *string = new(string) //reply是nil 是没有任何地址的 因此不能直接赋值
//		//var reply *string //reply是nil 是没有任何地址的 因此不能直接赋值
//		err = client.Call("HelloServer.Hello", "chenlin", reply)
//		//但是我想使用client.Hello这种使用
//		//可以夸语言调用 核心的观点就是序列化的协议
//		// go语言的rpc序列化和反序列化协议是Gob 能否替换为Json
//		if err != nil {
//			fmt.Println(err)
//			panic("数据返回失败")
//		}
//		fmt.Println(*reply)
//	}
func main() {
	//1 建立连接 拨号
	client := client_proxy.NewHelloServiceClient("tcp", "127.0.0.1:9090")

	//new 分配空间 只有地址的 可以直接赋值
	//
	var reply *string = new(string) //reply是nil 是没有任何地址的 因此不能直接赋值
	//var reply *string //reply是nil 是没有任何地址的 因此不能直接赋值
	err := client.Hello("boddy", reply)
	//err = client.Call("HelloServer.Hello", "chenlin", reply)
	//但是我想使用client.Hello这种使用
	//可以夸语言调用 核心的观点就是序列化的协议
	// go语言的rpc序列化和反序列化协议是Gob 能否替换为Json
	if err != nil {
		fmt.Println(err)
		panic("数据返回失败")
	}
	fmt.Println(*reply)
}

//func main() {
//	//1 建立连接 拨号
//	conn, err := net.Dial("tcp", "127.0.0.1:9090")
//	if err != nil {
//		println(err)
//		panic("连接失败")
//	}
//	//new 分配空间 只有地址的 可以直接赋值
//	//
//	var reply *string = new(string) //reply是nil 是没有任何地址的 因此不能直接赋值
//	//var reply *string //reply是nil 是没有任何地址的 因此不能直接赋值
//	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
//	err = client.Call(handle.HelloServiceName+".Hello", "chenlinlin", reply)
//	//err = client.Call("HelloServer.Hello", "chenlin", reply)
//	//但是我想使用client.Hello这种使用
//	//可以夸语言调用 核心的观点就是序列化的协议
//	// go语言的rpc序列化和反序列化协议是Gob 能否替换为Json
//	if err != nil {
//		fmt.Println(err)
//		panic("数据返回失败")
//	}
//	fmt.Println(*reply)
//}
