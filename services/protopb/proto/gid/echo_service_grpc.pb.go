// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package gid

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// EchoServiceClient is the client API for EchoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EchoServiceClient interface {
	// Echo method receives a simple message and returns it.
	//
	// The message posted as the id parameter will also be
	// returned.
	Echo(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error)
	// EchoBody method receives a simple message and returns it.
	EchoBody(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error)
	// EchoDelete method receives a simple message and returns it.
	EchoDelete(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error)
	// EchoPatch method receives a NonStandardUpdateRequest and returns it.
	EchoPatch(ctx context.Context, in *DynamicMessageUpdate, opts ...grpc.CallOption) (*DynamicMessageUpdate, error)
	// EchoUnauthorized method receives a simple message and returns it. It must
	// always return a google.rpc.Code of `UNAUTHENTICATED` and a HTTP Status code
	// of 401.
	EchoUnauthorized(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error)
}

type echoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEchoServiceClient(cc grpc.ClientConnInterface) EchoServiceClient {
	return &echoServiceClient{cc}
}

func (c *echoServiceClient) Echo(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error) {
	out := new(SimpleMessage)
	err := c.cc.Invoke(ctx, "/gid.EchoService/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoServiceClient) EchoBody(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error) {
	out := new(SimpleMessage)
	err := c.cc.Invoke(ctx, "/gid.EchoService/EchoBody", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoServiceClient) EchoDelete(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error) {
	out := new(SimpleMessage)
	err := c.cc.Invoke(ctx, "/gid.EchoService/EchoDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoServiceClient) EchoPatch(ctx context.Context, in *DynamicMessageUpdate, opts ...grpc.CallOption) (*DynamicMessageUpdate, error) {
	out := new(DynamicMessageUpdate)
	err := c.cc.Invoke(ctx, "/gid.EchoService/EchoPatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoServiceClient) EchoUnauthorized(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error) {
	out := new(SimpleMessage)
	err := c.cc.Invoke(ctx, "/gid.EchoService/EchoUnauthorized", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EchoServiceServer is the server API for EchoService service.
// All implementations should embed UnimplementedEchoServiceServer
// for forward compatibility
type EchoServiceServer interface {
	// Echo method receives a simple message and returns it.
	//
	// The message posted as the id parameter will also be
	// returned.
	Echo(context.Context, *SimpleMessage) (*SimpleMessage, error)
	// EchoBody method receives a simple message and returns it.
	EchoBody(context.Context, *SimpleMessage) (*SimpleMessage, error)
	// EchoDelete method receives a simple message and returns it.
	EchoDelete(context.Context, *SimpleMessage) (*SimpleMessage, error)
	// EchoPatch method receives a NonStandardUpdateRequest and returns it.
	EchoPatch(context.Context, *DynamicMessageUpdate) (*DynamicMessageUpdate, error)
	// EchoUnauthorized method receives a simple message and returns it. It must
	// always return a google.rpc.Code of `UNAUTHENTICATED` and a HTTP Status code
	// of 401.
	EchoUnauthorized(context.Context, *SimpleMessage) (*SimpleMessage, error)
}

// UnimplementedEchoServiceServer should be embedded to have forward compatible implementations.
type UnimplementedEchoServiceServer struct {
}

func (UnimplementedEchoServiceServer) Echo(context.Context, *SimpleMessage) (*SimpleMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedEchoServiceServer) EchoBody(context.Context, *SimpleMessage) (*SimpleMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EchoBody not implemented")
}
func (UnimplementedEchoServiceServer) EchoDelete(context.Context, *SimpleMessage) (*SimpleMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EchoDelete not implemented")
}
func (UnimplementedEchoServiceServer) EchoPatch(context.Context, *DynamicMessageUpdate) (*DynamicMessageUpdate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EchoPatch not implemented")
}
func (UnimplementedEchoServiceServer) EchoUnauthorized(context.Context, *SimpleMessage) (*SimpleMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EchoUnauthorized not implemented")
}

// UnsafeEchoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EchoServiceServer will
// result in compilation errors.
type UnsafeEchoServiceServer interface {
	mustEmbedUnimplementedEchoServiceServer()
}

func RegisterEchoServiceServer(s grpc.ServiceRegistrar, srv EchoServiceServer) {
	s.RegisterService(&_EchoService_serviceDesc, srv)
}

func _EchoService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gid.EchoService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).Echo(ctx, req.(*SimpleMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _EchoService_EchoBody_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).EchoBody(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gid.EchoService/EchoBody",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).EchoBody(ctx, req.(*SimpleMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _EchoService_EchoDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).EchoDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gid.EchoService/EchoDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).EchoDelete(ctx, req.(*SimpleMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _EchoService_EchoPatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DynamicMessageUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).EchoPatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gid.EchoService/EchoPatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).EchoPatch(ctx, req.(*DynamicMessageUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _EchoService_EchoUnauthorized_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).EchoUnauthorized(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gid.EchoService/EchoUnauthorized",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).EchoUnauthorized(ctx, req.(*SimpleMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _EchoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gid.EchoService",
	HandlerType: (*EchoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _EchoService_Echo_Handler,
		},
		{
			MethodName: "EchoBody",
			Handler:    _EchoService_EchoBody_Handler,
		},
		{
			MethodName: "EchoDelete",
			Handler:    _EchoService_EchoDelete_Handler,
		},
		{
			MethodName: "EchoPatch",
			Handler:    _EchoService_EchoPatch_Handler,
		},
		{
			MethodName: "EchoUnauthorized",
			Handler:    _EchoService_EchoUnauthorized_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/gid/echo_service.proto",
}
