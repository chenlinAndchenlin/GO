// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/bj38.proto

package bj38

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Bj38 service

func NewBj38Endpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Bj38 service

type Bj38Service interface {
	Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error)
	ClientStream(ctx context.Context, opts ...client.CallOption) (Bj38_ClientStreamService, error)
	ServerStream(ctx context.Context, in *ServerStreamRequest, opts ...client.CallOption) (Bj38_ServerStreamService, error)
	BidiStream(ctx context.Context, opts ...client.CallOption) (Bj38_BidiStreamService, error)
}

type bj38Service struct {
	c    client.Client
	name string
}

func NewBj38Service(name string, c client.Client) Bj38Service {
	return &bj38Service{
		c:    c,
		name: name,
	}
}

func (c *bj38Service) Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error) {
	req := c.c.NewRequest(c.name, "Bj38.Call", in)
	out := new(CallResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bj38Service) ClientStream(ctx context.Context, opts ...client.CallOption) (Bj38_ClientStreamService, error) {
	req := c.c.NewRequest(c.name, "Bj38.ClientStream", &ClientStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &bj38ServiceClientStream{stream}, nil
}

type Bj38_ClientStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*ClientStreamRequest) error
}

type bj38ServiceClientStream struct {
	stream client.Stream
}

func (x *bj38ServiceClientStream) Close() error {
	return x.stream.Close()
}

func (x *bj38ServiceClientStream) Context() context.Context {
	return x.stream.Context()
}

func (x *bj38ServiceClientStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *bj38ServiceClientStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *bj38ServiceClientStream) Send(m *ClientStreamRequest) error {
	return x.stream.Send(m)
}

func (c *bj38Service) ServerStream(ctx context.Context, in *ServerStreamRequest, opts ...client.CallOption) (Bj38_ServerStreamService, error) {
	req := c.c.NewRequest(c.name, "Bj38.ServerStream", &ServerStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &bj38ServiceServerStream{stream}, nil
}

type Bj38_ServerStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*ServerStreamResponse, error)
}

type bj38ServiceServerStream struct {
	stream client.Stream
}

func (x *bj38ServiceServerStream) Close() error {
	return x.stream.Close()
}

func (x *bj38ServiceServerStream) Context() context.Context {
	return x.stream.Context()
}

func (x *bj38ServiceServerStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *bj38ServiceServerStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *bj38ServiceServerStream) Recv() (*ServerStreamResponse, error) {
	m := new(ServerStreamResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bj38Service) BidiStream(ctx context.Context, opts ...client.CallOption) (Bj38_BidiStreamService, error) {
	req := c.c.NewRequest(c.name, "Bj38.BidiStream", &BidiStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &bj38ServiceBidiStream{stream}, nil
}

type Bj38_BidiStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*BidiStreamRequest) error
	Recv() (*BidiStreamResponse, error)
}

type bj38ServiceBidiStream struct {
	stream client.Stream
}

func (x *bj38ServiceBidiStream) Close() error {
	return x.stream.Close()
}

func (x *bj38ServiceBidiStream) Context() context.Context {
	return x.stream.Context()
}

func (x *bj38ServiceBidiStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *bj38ServiceBidiStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *bj38ServiceBidiStream) Send(m *BidiStreamRequest) error {
	return x.stream.Send(m)
}

func (x *bj38ServiceBidiStream) Recv() (*BidiStreamResponse, error) {
	m := new(BidiStreamResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Bj38 service

type Bj38Handler interface {
	Call(context.Context, *CallRequest, *CallResponse) error
	ClientStream(context.Context, Bj38_ClientStreamStream) error
	ServerStream(context.Context, *ServerStreamRequest, Bj38_ServerStreamStream) error
	BidiStream(context.Context, Bj38_BidiStreamStream) error
}

func RegisterBj38Handler(s server.Server, hdlr Bj38Handler, opts ...server.HandlerOption) error {
	type bj38 interface {
		Call(ctx context.Context, in *CallRequest, out *CallResponse) error
		ClientStream(ctx context.Context, stream server.Stream) error
		ServerStream(ctx context.Context, stream server.Stream) error
		BidiStream(ctx context.Context, stream server.Stream) error
	}
	type Bj38 struct {
		bj38
	}
	h := &bj38Handler{hdlr}
	return s.Handle(s.NewHandler(&Bj38{h}, opts...))
}

type bj38Handler struct {
	Bj38Handler
}

func (h *bj38Handler) Call(ctx context.Context, in *CallRequest, out *CallResponse) error {
	return h.Bj38Handler.Call(ctx, in, out)
}

func (h *bj38Handler) ClientStream(ctx context.Context, stream server.Stream) error {
	return h.Bj38Handler.ClientStream(ctx, &bj38ClientStreamStream{stream})
}

type Bj38_ClientStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*ClientStreamRequest, error)
}

type bj38ClientStreamStream struct {
	stream server.Stream
}

func (x *bj38ClientStreamStream) Close() error {
	return x.stream.Close()
}

func (x *bj38ClientStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *bj38ClientStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *bj38ClientStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *bj38ClientStreamStream) Recv() (*ClientStreamRequest, error) {
	m := new(ClientStreamRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (h *bj38Handler) ServerStream(ctx context.Context, stream server.Stream) error {
	m := new(ServerStreamRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.Bj38Handler.ServerStream(ctx, m, &bj38ServerStreamStream{stream})
}

type Bj38_ServerStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*ServerStreamResponse) error
}

type bj38ServerStreamStream struct {
	stream server.Stream
}

func (x *bj38ServerStreamStream) Close() error {
	return x.stream.Close()
}

func (x *bj38ServerStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *bj38ServerStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *bj38ServerStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *bj38ServerStreamStream) Send(m *ServerStreamResponse) error {
	return x.stream.Send(m)
}

func (h *bj38Handler) BidiStream(ctx context.Context, stream server.Stream) error {
	return h.Bj38Handler.BidiStream(ctx, &bj38BidiStreamStream{stream})
}

type Bj38_BidiStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*BidiStreamResponse) error
	Recv() (*BidiStreamRequest, error)
}

type bj38BidiStreamStream struct {
	stream server.Stream
}

func (x *bj38BidiStreamStream) Close() error {
	return x.stream.Close()
}

func (x *bj38BidiStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *bj38BidiStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *bj38BidiStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *bj38BidiStreamStream) Send(m *BidiStreamResponse) error {
	return x.stream.Send(m)
}

func (x *bj38BidiStreamStream) Recv() (*BidiStreamRequest, error) {
	m := new(BidiStreamRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
