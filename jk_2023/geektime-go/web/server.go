package web

import (
	"net"
	"net/http"
)

// 确保结构体实现所以的接口方法
var _ Server = &HTTPServer{}

type HandleFunc func(ctx Context)
type Server interface {
	http.Handler
	Start(add string) error
	// AddRoute 路由注册功能
	// method 是 HTTP 方法
	// path 是路由
	// handleFunc 是你的业务逻辑
	addRoute(method string, path string, handleFunc HandleFunc)
	//addRoute(method string, path string, handleFunc HandleFunc, ms ...Middleware)
	// 这种允许注册多个，没有必要提供
	// 让用户自己去管 切片 提供多个handles
	//AddRoute1(method string, path string, handles ...HandleFunc)
}

type HTTPServer struct {
}

/*func (h *HTTPServer) AddRoute1(method string, path string, handles ...HandleFunc) {
	//TODO implement me
	panic("implement me")
}*/

/*
	type HTTPSServer struct {
		HTTPServer
	}
*/
//ServeHTTP 核心请求的入口
func (h *HTTPServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	// 你的框架代码就是在这里
	// 你的框架代码就在这里
	ctx := &Context{
		Req:  request,
		Resp: writer,
		//tplEngine: h.tplEngine,
	}
	//接下来就是查询路由，并且执行命中的业务逻辑
	h.serve(ctx)

}
func (h *HTTPServer) serve(ctx *Context) {
	// 接下来就是查找路由，并且执行命中的业务逻辑
}
func (h *HTTPServer) addRoute(method string, path string, handleFunc HandleFunc) {

}

func (h *HTTPServer) Start(addr string) error {
	//TODO implement me
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	//区别：在这里用户可以注册所谓的 after start 回调
	//比如说往你的admin注册自己的实例
	//执行一些业务所需要的前置条件
	return http.Serve(l, h)
}

func (h *HTTPServer) Start1(addr string) error {

	/*l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}*/
	return http.ListenAndServe(addr, h)
}
