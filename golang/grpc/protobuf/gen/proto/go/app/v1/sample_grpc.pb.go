// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// HogeClient is the client API for Hoge service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HogeClient interface {
	AtoB(ctx context.Context, in *AA, opts ...grpc.CallOption) (*BB, error)
	AtoBstream(ctx context.Context, opts ...grpc.CallOption) (Hoge_AtoBstreamClient, error)
}

type hogeClient struct {
	cc grpc.ClientConnInterface
}

func NewHogeClient(cc grpc.ClientConnInterface) HogeClient {
	return &hogeClient{cc}
}

func (c *hogeClient) AtoB(ctx context.Context, in *AA, opts ...grpc.CallOption) (*BB, error) {
	out := new(BB)
	err := c.cc.Invoke(ctx, "/app.v1.Hoge/AtoB", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hogeClient) AtoBstream(ctx context.Context, opts ...grpc.CallOption) (Hoge_AtoBstreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Hoge_ServiceDesc.Streams[0], "/app.v1.Hoge/AtoBstream", opts...)
	if err != nil {
		return nil, err
	}
	x := &hogeAtoBstreamClient{stream}
	return x, nil
}

type Hoge_AtoBstreamClient interface {
	Send(*AA) error
	Recv() (*BB, error)
	grpc.ClientStream
}

type hogeAtoBstreamClient struct {
	grpc.ClientStream
}

func (x *hogeAtoBstreamClient) Send(m *AA) error {
	return x.ClientStream.SendMsg(m)
}

func (x *hogeAtoBstreamClient) Recv() (*BB, error) {
	m := new(BB)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HogeServer is the server API for Hoge service.
// All implementations must embed UnimplementedHogeServer
// for forward compatibility
type HogeServer interface {
	AtoB(context.Context, *AA) (*BB, error)
	AtoBstream(Hoge_AtoBstreamServer) error
	mustEmbedUnimplementedHogeServer()
}

// UnimplementedHogeServer must be embedded to have forward compatible implementations.
type UnimplementedHogeServer struct {
}

func (UnimplementedHogeServer) AtoB(context.Context, *AA) (*BB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AtoB not implemented")
}
func (UnimplementedHogeServer) AtoBstream(Hoge_AtoBstreamServer) error {
	return status.Errorf(codes.Unimplemented, "method AtoBstream not implemented")
}
func (UnimplementedHogeServer) mustEmbedUnimplementedHogeServer() {}

// UnsafeHogeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HogeServer will
// result in compilation errors.
type UnsafeHogeServer interface {
	mustEmbedUnimplementedHogeServer()
}

func RegisterHogeServer(s grpc.ServiceRegistrar, srv HogeServer) {
	s.RegisterService(&Hoge_ServiceDesc, srv)
}

func _Hoge_AtoB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AA)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HogeServer).AtoB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.v1.Hoge/AtoB",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HogeServer).AtoB(ctx, req.(*AA))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hoge_AtoBstream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HogeServer).AtoBstream(&hogeAtoBstreamServer{stream})
}

type Hoge_AtoBstreamServer interface {
	Send(*BB) error
	Recv() (*AA, error)
	grpc.ServerStream
}

type hogeAtoBstreamServer struct {
	grpc.ServerStream
}

func (x *hogeAtoBstreamServer) Send(m *BB) error {
	return x.ServerStream.SendMsg(m)
}

func (x *hogeAtoBstreamServer) Recv() (*AA, error) {
	m := new(AA)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Hoge_ServiceDesc is the grpc.ServiceDesc for Hoge service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Hoge_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "app.v1.Hoge",
	HandlerType: (*HogeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AtoB",
			Handler:    _Hoge_AtoB_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AtoBstream",
			Handler:       _Hoge_AtoBstream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "app/v1/sample.proto",
}
