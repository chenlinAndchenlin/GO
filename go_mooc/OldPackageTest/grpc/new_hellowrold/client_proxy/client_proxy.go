package client_proxy

import (
	"OldPackageTest/grpc/new_hellowrold/handle"
	"net/rpc"
)

type HelloServiceStub struct {
	*rpc.Client
}

func NewHelloServiceClient(protol, address string) HelloServiceStub {
	conn, err := rpc.Dial(protol, address)
	if err != nil {
		panic(err)
	}
	return HelloServiceStub{conn}
}
func (c *HelloServiceStub) Hello(request string, reply *string) error {
	err := c.Call(handle.HelloServiceName+".Hello", request, reply)
	if err != nil {
		return err
	}
	return err
}
