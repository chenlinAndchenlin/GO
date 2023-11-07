package main

import (
	"OldPackageTest/grpc_stream/proto"
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"google.golang.org/grpc"
)

//	type GreeterClient interface {
//		GetStream(ctx context.Co'ntext, in *StreamReqData, opts ...grpc.CallOption) (Greeter_GetStreamClient, error)
//		PutStream(ctx context.Context, opts ...grpc.CallOption) (Greeter_PutStreamClient, error)
//		AllStream(ctx context.Context, opts ...grpc.CallOption) (Greeter_AllStreamClient, error)
//	}
func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	//reqstreamData := &proto.StreamReqData{Data:"aaa"}
	////调用服务端推送流
	//res, err2 := c.GetStream(context.Background(), &proto.StreamReqData{Data: "慕课网"})
	//if err2 != nil {
	//	fmt.Println(err2)
	//}
	//for {
	//	recv, err := res.Recv()
	//	if err != nil {
	//		fmt.Println(err)
	//		break
	//	}
	//	log.Println(recv)
	//}
	////客户端 推送 流
	//stream, err := c.PutStream(context.Background())
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//i := 0
	//for {
	//	i++
	//	stream.Send(&proto.StreamReqData{Data: fmt.Sprintf("慕课网%v", i)})
	//	time.Sleep(time.Second)
	//	if i > 10 {
	//		break
	//	}
	//}
	//服务端 客户端 双向
	allStream, err := c.AllStream(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			recv, err := allStream.Recv()
			if err != nil {
				fmt.Println(err)
				break
			}
			log.Println(recv.Data)
		}

	}()
	go func() {
		defer wg.Done()
		i := 0
		for {
			i++
			allStream.Send(&proto.StreamReqData{Data: fmt.Sprintf("慕课网 从client to server:%v", i)})
			time.Sleep(time.Second)
			if i > 10 {
				break
			}
		}
	}()
	wg.Wait()

}
