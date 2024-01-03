package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	ip := "127.0.0.1"
	port := 8848
	address := fmt.Sprintf("%s:%d", ip, port)
	listen, err := net.Listen("tcp", address)

	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	fmt.Println("监听中...")

	accept, err := listen.Accept()
	if err != nil {
		fmt.Println("listener.Accept err:", err)
		return
	}
	fmt.Println("连接建立成功!")
	buf := make([]byte, 1024)

	read, err := accept.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err:", err)

		return
	}
	fmt.Println("Client =====> Server, 长度:", read, "，数据:", string(buf[0:read]))

	//服务器对客户端请求进行响应 ,将数据转成大写 "hello" ==> HELLO
	//func ToUpper(s string) string {
	upperData := strings.ToUpper(string(buf[0:read]))
	write, err := accept.Write([]byte(upperData))
	if err != nil {
		return
	}
	fmt.Println("Client  <====== Server, 长度:", write, "，数据:", upperData)
	accept.Close()
}
