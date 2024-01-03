package web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	var h Server = &HTTPServer{}

	h.addRoute(http.MethodGet, "/user", func(ctx Context) {
		fmt.Println("处理第一件事")
		fmt.Println("处理第二件事")
	})

	/*handle1 := func(ctx Context) {
		fmt.Println("处理第一件事")
	}
	handle2 := func(ctx Context) {
		fmt.Println("处理第二件事")
	}
	h.addRoute(http.MethodGet, "/user", func(ctx Context) {
		handle1(ctx)
		handle2(ctx)
	})
	//这两种方法等效
	h.AddRoute1(http.MethodGet, "/user", handle1, handle2)*/

	//用法1 委托给http
	http.ListenAndServe("localhost:8080", h)
	//http.ListenAndServe(":9090",这里搞一个handler)
	//用法2 自己手动管理
	h.Start(":8080")

}
