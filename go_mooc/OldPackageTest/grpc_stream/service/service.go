package main

import (
	"OldPackageTest/grpc_stream/proto"
	"fmt"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
)

const PROT = ":50052"

type server struct {
}

//	type GreeterServer interface {
//		GetStream(*StreamReqData, Greeter_GetStreamServer) error
//		PutStream(Greeter_PutStreamServer) error
//		AllStream(Greeter_AllStreamServer) error
//	}//func (s *server) GetStream(ctx context.Context, req *proto.StreamReqData) {
//
//		return nil, nil
//	}
func (s *server) GetStream(req *proto.StreamReqData, res proto.Greeter_GetStreamServer) error {
	i := 0
	for {
		i++
		err := res.Send(&proto.StreamResData{
			Data: fmt.Sprintf("%v", time.Now().Unix()),
		})
		if err != nil {
			return err
		}
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}
	return nil
}
func (s *server) PutStream(cliStr proto.Greeter_PutStreamServer) error {
	for {
		a, err := cliStr.Recv()
		if err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Println(a.Data)
		}
	}
	return nil

}
func (s *server) AllStream(allStr proto.Greeter_AllStreamServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {

			recv, err := allStr.Recv()
			if err != nil {
				return
			}
			fmt.Println("收到client的消息" + recv.Data)
		}

	}()
	go func() {
		defer wg.Done()
		i := 0
		for {
			i++
			err := allStr.Send(&proto.StreamResData{Data: fmt.Sprintf("从service to client 发送数据：%v", time.Now().Unix())})
			if err != nil {
				return
			}
			time.Sleep(time.Second)
			if i > 10 {
				break
			}
		}

	}()
	wg.Wait()
	return nil
}
func main() {
	//监听断口
	listener, err := net.Listen("tcp", PROT)
	if err != nil {
		panic(err)
	}
	//创建grpc服务器
	s := grpc.NewServer()
	//注册事件
	proto.RegisterGreeterServer(s, &server{})
	//处理链接
	err = s.Serve(listener)
	if err != nil {
		panic(err)
	}
}
