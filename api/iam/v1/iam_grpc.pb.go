// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: iam/v1/iam.proto

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

// IamServiceClient is the client API for IamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IamServiceClient interface {
	// Token
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	// rpc RefreshToken(RefreshTokenRequest) returns (RefreshTokenResponse) {};
	// ====================================================================
	// User Manage
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	// rpc UpdateUser(UpdateUserRequest) returns (UpdateUserResponse) {};
	// rpc DeleteUser(DeleteUserRequest) returns (DeleteUserResponse) {};
	// rpc GetUser(GetUserRequest) returns (GetUserResponse) {};
	// ====================================================================
	// Role Manage
	CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*CreateRoleResponse, error)
	// rpc UpdateRole(UpdateRoleRequest) returns (UpdateRoleResponse) {};
	// rpc DeleteRole(DeleteRoleRequest) returns (DeleteRoleResponse) {};
	// rpc GetRole(GetRoleRequest) returns (GetRoleResponse) {};
	// ====================================================================
	// Tenant Manage
	CreateTenant(ctx context.Context, in *CreateTenantRequest, opts ...grpc.CallOption) (*CreateTenantResponse, error)
	// rpc UpdateTenant(UpdateTenantRequest) returns (UpdateTenantResponse) {};
	// rpc DeleteTenant(DeleteTenantRequest) returns (DeleteTenantResponse) {};
	// rpc GetTenant(GetTenantRequest) returns (GetTenantResponse) {};
	// ====================================================================
	// Bind Manage
	CreateBinding(ctx context.Context, in *CreateBindingRequest, opts ...grpc.CallOption) (*CreateBindingResponse, error)
	// rpc UpdateBinding(UpdateBindingRequest) returns (UpdateBindingResponse) {};
	// rpc DeleteBinding(DeleteBindingRequest) returns (DeleteBindingResponse) {};
	// rpc GetBinding(GetBindingRequest) returns (GetBindingResponse) {};
	// ====================================================================
	// Policy Manage
	CreatePolicy(ctx context.Context, in *CreatePolicyRequest, opts ...grpc.CallOption) (*CreatePolicyResponse, error)
}

type iamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIamServiceClient(cc grpc.ClientConnInterface) IamServiceClient {
	return &iamServiceClient{cc}
}

func (c *iamServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/iam.v1.IamService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iamServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/iam.v1.IamService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iamServiceClient) CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*CreateRoleResponse, error) {
	out := new(CreateRoleResponse)
	err := c.cc.Invoke(ctx, "/iam.v1.IamService/CreateRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iamServiceClient) CreateTenant(ctx context.Context, in *CreateTenantRequest, opts ...grpc.CallOption) (*CreateTenantResponse, error) {
	out := new(CreateTenantResponse)
	err := c.cc.Invoke(ctx, "/iam.v1.IamService/CreateTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iamServiceClient) CreateBinding(ctx context.Context, in *CreateBindingRequest, opts ...grpc.CallOption) (*CreateBindingResponse, error) {
	out := new(CreateBindingResponse)
	err := c.cc.Invoke(ctx, "/iam.v1.IamService/CreateBinding", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iamServiceClient) CreatePolicy(ctx context.Context, in *CreatePolicyRequest, opts ...grpc.CallOption) (*CreatePolicyResponse, error) {
	out := new(CreatePolicyResponse)
	err := c.cc.Invoke(ctx, "/iam.v1.IamService/CreatePolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IamServiceServer is the server API for IamService service.
// All implementations must embed UnimplementedIamServiceServer
// for forward compatibility
type IamServiceServer interface {
	// Token
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	// rpc RefreshToken(RefreshTokenRequest) returns (RefreshTokenResponse) {};
	// ====================================================================
	// User Manage
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	// rpc UpdateUser(UpdateUserRequest) returns (UpdateUserResponse) {};
	// rpc DeleteUser(DeleteUserRequest) returns (DeleteUserResponse) {};
	// rpc GetUser(GetUserRequest) returns (GetUserResponse) {};
	// ====================================================================
	// Role Manage
	CreateRole(context.Context, *CreateRoleRequest) (*CreateRoleResponse, error)
	// rpc UpdateRole(UpdateRoleRequest) returns (UpdateRoleResponse) {};
	// rpc DeleteRole(DeleteRoleRequest) returns (DeleteRoleResponse) {};
	// rpc GetRole(GetRoleRequest) returns (GetRoleResponse) {};
	// ====================================================================
	// Tenant Manage
	CreateTenant(context.Context, *CreateTenantRequest) (*CreateTenantResponse, error)
	// rpc UpdateTenant(UpdateTenantRequest) returns (UpdateTenantResponse) {};
	// rpc DeleteTenant(DeleteTenantRequest) returns (DeleteTenantResponse) {};
	// rpc GetTenant(GetTenantRequest) returns (GetTenantResponse) {};
	// ====================================================================
	// Bind Manage
	CreateBinding(context.Context, *CreateBindingRequest) (*CreateBindingResponse, error)
	// rpc UpdateBinding(UpdateBindingRequest) returns (UpdateBindingResponse) {};
	// rpc DeleteBinding(DeleteBindingRequest) returns (DeleteBindingResponse) {};
	// rpc GetBinding(GetBindingRequest) returns (GetBindingResponse) {};
	// ====================================================================
	// Policy Manage
	CreatePolicy(context.Context, *CreatePolicyRequest) (*CreatePolicyResponse, error)
	mustEmbedUnimplementedIamServiceServer()
}

// UnimplementedIamServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIamServiceServer struct {
}

func (UnimplementedIamServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedIamServiceServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedIamServiceServer) CreateRole(context.Context, *CreateRoleRequest) (*CreateRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRole not implemented")
}
func (UnimplementedIamServiceServer) CreateTenant(context.Context, *CreateTenantRequest) (*CreateTenantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTenant not implemented")
}
func (UnimplementedIamServiceServer) CreateBinding(context.Context, *CreateBindingRequest) (*CreateBindingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBinding not implemented")
}
func (UnimplementedIamServiceServer) CreatePolicy(context.Context, *CreatePolicyRequest) (*CreatePolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePolicy not implemented")
}
func (UnimplementedIamServiceServer) mustEmbedUnimplementedIamServiceServer() {}

// UnsafeIamServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IamServiceServer will
// result in compilation errors.
type UnsafeIamServiceServer interface {
	mustEmbedUnimplementedIamServiceServer()
}

func RegisterIamServiceServer(s grpc.ServiceRegistrar, srv IamServiceServer) {
	s.RegisterService(&IamService_ServiceDesc, srv)
}

func _IamService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IamServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iam.v1.IamService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IamServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IamService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IamServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iam.v1.IamService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IamServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IamService_CreateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IamServiceServer).CreateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iam.v1.IamService/CreateRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IamServiceServer).CreateRole(ctx, req.(*CreateRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IamService_CreateTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IamServiceServer).CreateTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iam.v1.IamService/CreateTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IamServiceServer).CreateTenant(ctx, req.(*CreateTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IamService_CreateBinding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBindingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IamServiceServer).CreateBinding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iam.v1.IamService/CreateBinding",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IamServiceServer).CreateBinding(ctx, req.(*CreateBindingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IamService_CreatePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IamServiceServer).CreatePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iam.v1.IamService/CreatePolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IamServiceServer).CreatePolicy(ctx, req.(*CreatePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IamService_ServiceDesc is the grpc.ServiceDesc for IamService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IamService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "iam.v1.IamService",
	HandlerType: (*IamServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _IamService_Login_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _IamService_CreateUser_Handler,
		},
		{
			MethodName: "CreateRole",
			Handler:    _IamService_CreateRole_Handler,
		},
		{
			MethodName: "CreateTenant",
			Handler:    _IamService_CreateTenant_Handler,
		},
		{
			MethodName: "CreateBinding",
			Handler:    _IamService_CreateBinding_Handler,
		},
		{
			MethodName: "CreatePolicy",
			Handler:    _IamService_CreatePolicy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "iam/v1/iam.proto",
}
