package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", ":8848")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	fmt.Println("cliet与server连接建立成功!")
	sendData := []byte("helloworld")

	for {
		cnt, err := conn.Write(sendData)
		if err != nil {
			fmt.Println("conn.Write err:", err)

			return
		}
		fmt.Println("Client ===> Server cnt:", cnt, ", data:", string(sendData))

		buf := make([]byte, 1024)
		cnt, err = conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}

		fmt.Println("Client <==== Server , cnt:", cnt, ", data:", string(buf[0:cnt]))
		time.Sleep(1 * time.Second)
	}

	conn.Close()
}
