package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//注册路由
	http.HandleFunc("/user", func(w http.ResponseWriter, request *http.Request) {
		//request : ===> 包含客户端发来的数据
		fmt.Println("用户请求的详情：")
		fmt.Println("request:", request)
		fmt.Println("method:", request.Method)

		//writer :  ===> 通过writer将数据返回给客户端
		_, err := io.WriteString(w, "这是给/user请求返回的数据")
		if err != nil {
			return
		}
	})
	http.HandleFunc("/name", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = io.WriteString(writer, "这是/name请求返回的数据!")
	})
	//https://127.0.0.1:8080/id
	http.HandleFunc("/id", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = io.WriteString(writer, "这是/id请求返回的数据!")
	})
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		fmt.Println("http start failed, err:", err)
		return
	}
}
