// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: echo/v1/echo.proto

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

// EchoServiceClient is the client API for EchoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EchoServiceClient interface {
	CreateEcho(ctx context.Context, in *CreateEchoRequest, opts ...grpc.CallOption) (*CreateEchoResponse, error)
	ListEcho(ctx context.Context, in *ListEchoRequest, opts ...grpc.CallOption) (*ListEchoResponse, error)
	UpdateEcho(ctx context.Context, in *UpdateEchoRequest, opts ...grpc.CallOption) (*UpdateEchoResponse, error)
	DeleteEcho(ctx context.Context, in *DeleteEchoRequest, opts ...grpc.CallOption) (*DeleteEchoResponse, error)
	GetEcho(ctx context.Context, in *GetEchoRequest, opts ...grpc.CallOption) (*EchoRecord, error)
}

type echoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEchoServiceClient(cc grpc.ClientConnInterface) EchoServiceClient {
	return &echoServiceClient{cc}
}

func (c *echoServiceClient) CreateEcho(ctx context.Context, in *CreateEchoRequest, opts ...grpc.CallOption) (*CreateEchoResponse, error) {
	out := new(CreateEchoResponse)
	err := c.cc.Invoke(ctx, "/echo.v1.EchoService/CreateEcho", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoServiceClient) ListEcho(ctx context.Context, in *ListEchoRequest, opts ...grpc.CallOption) (*ListEchoResponse, error) {
	out := new(ListEchoResponse)
	err := c.cc.Invoke(ctx, "/echo.v1.EchoService/ListEcho", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoServiceClient) UpdateEcho(ctx context.Context, in *UpdateEchoRequest, opts ...grpc.CallOption) (*UpdateEchoResponse, error) {
	out := new(UpdateEchoResponse)
	err := c.cc.Invoke(ctx, "/echo.v1.EchoService/UpdateEcho", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoServiceClient) DeleteEcho(ctx context.Context, in *DeleteEchoRequest, opts ...grpc.CallOption) (*DeleteEchoResponse, error) {
	out := new(DeleteEchoResponse)
	err := c.cc.Invoke(ctx, "/echo.v1.EchoService/DeleteEcho", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoServiceClient) GetEcho(ctx context.Context, in *GetEchoRequest, opts ...grpc.CallOption) (*EchoRecord, error) {
	out := new(EchoRecord)
	err := c.cc.Invoke(ctx, "/echo.v1.EchoService/GetEcho", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EchoServiceServer is the server API for EchoService service.
// All implementations must embed UnimplementedEchoServiceServer
// for forward compatibility
type EchoServiceServer interface {
	CreateEcho(context.Context, *CreateEchoRequest) (*CreateEchoResponse, error)
	ListEcho(context.Context, *ListEchoRequest) (*ListEchoResponse, error)
	UpdateEcho(context.Context, *UpdateEchoRequest) (*UpdateEchoResponse, error)
	DeleteEcho(context.Context, *DeleteEchoRequest) (*DeleteEchoResponse, error)
	GetEcho(context.Context, *GetEchoRequest) (*EchoRecord, error)
	mustEmbedUnimplementedEchoServiceServer()
}

// UnimplementedEchoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEchoServiceServer struct {
}

func (UnimplementedEchoServiceServer) CreateEcho(context.Context, *CreateEchoRequest) (*CreateEchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEcho not implemented")
}
func (UnimplementedEchoServiceServer) ListEcho(context.Context, *ListEchoRequest) (*ListEchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEcho not implemented")
}
func (UnimplementedEchoServiceServer) UpdateEcho(context.Context, *UpdateEchoRequest) (*UpdateEchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEcho not implemented")
}
func (UnimplementedEchoServiceServer) DeleteEcho(context.Context, *DeleteEchoRequest) (*DeleteEchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEcho not implemented")
}
func (UnimplementedEchoServiceServer) GetEcho(context.Context, *GetEchoRequest) (*EchoRecord, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEcho not implemented")
}
func (UnimplementedEchoServiceServer) mustEmbedUnimplementedEchoServiceServer() {}

// UnsafeEchoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EchoServiceServer will
// result in compilation errors.
type UnsafeEchoServiceServer interface {
	mustEmbedUnimplementedEchoServiceServer()
}

func RegisterEchoServiceServer(s grpc.ServiceRegistrar, srv EchoServiceServer) {
	s.RegisterService(&EchoService_ServiceDesc, srv)
}

func _EchoService_CreateEcho_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).CreateEcho(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echo.v1.EchoService/CreateEcho",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).CreateEcho(ctx, req.(*CreateEchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EchoService_ListEcho_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).ListEcho(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echo.v1.EchoService/ListEcho",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).ListEcho(ctx, req.(*ListEchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EchoService_UpdateEcho_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).UpdateEcho(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echo.v1.EchoService/UpdateEcho",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).UpdateEcho(ctx, req.(*UpdateEchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EchoService_DeleteEcho_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).DeleteEcho(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echo.v1.EchoService/DeleteEcho",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).DeleteEcho(ctx, req.(*DeleteEchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EchoService_GetEcho_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).GetEcho(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echo.v1.EchoService/GetEcho",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).GetEcho(ctx, req.(*GetEchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EchoService_ServiceDesc is the grpc.ServiceDesc for EchoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EchoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "echo.v1.EchoService",
	HandlerType: (*EchoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEcho",
			Handler:    _EchoService_CreateEcho_Handler,
		},
		{
			MethodName: "ListEcho",
			Handler:    _EchoService_ListEcho_Handler,
		},
		{
			MethodName: "UpdateEcho",
			Handler:    _EchoService_UpdateEcho_Handler,
		},
		{
			MethodName: "DeleteEcho",
			Handler:    _EchoService_DeleteEcho_Handler,
		},
		{
			MethodName: "GetEcho",
			Handler:    _EchoService_GetEcho_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "echo/v1/echo.proto",
}
