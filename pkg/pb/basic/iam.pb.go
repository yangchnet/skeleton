// Code generated by protoc-gen-go. DO NOT EDIT.
// source: iam.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type CreateUserRequest struct {
	// 用户名
	Username *wrappers.StringValue `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// 密码
	Password *wrappers.StringValue `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	// 邮箱
	Email                *wrappers.StringValue `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a2c201915207782, []int{0}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetUsername() *wrappers.StringValue {
	if m != nil {
		return m.Username
	}
	return nil
}

func (m *CreateUserRequest) GetPassword() *wrappers.StringValue {
	if m != nil {
		return m.Password
	}
	return nil
}

func (m *CreateUserRequest) GetEmail() *wrappers.StringValue {
	if m != nil {
		return m.Email
	}
	return nil
}

type CreateUserResponse struct {
	// 用户名
	Username *wrappers.StringValue `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// 用户编号
	UserId               *wrappers.StringValue `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CreateUserResponse) Reset()         { *m = CreateUserResponse{} }
func (m *CreateUserResponse) String() string { return proto.CompactTextString(m) }
func (*CreateUserResponse) ProtoMessage()    {}
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a2c201915207782, []int{1}
}

func (m *CreateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResponse.Unmarshal(m, b)
}
func (m *CreateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResponse.Marshal(b, m, deterministic)
}
func (m *CreateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResponse.Merge(m, src)
}
func (m *CreateUserResponse) XXX_Size() int {
	return xxx_messageInfo_CreateUserResponse.Size(m)
}
func (m *CreateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResponse proto.InternalMessageInfo

func (m *CreateUserResponse) GetUsername() *wrappers.StringValue {
	if m != nil {
		return m.Username
	}
	return nil
}

func (m *CreateUserResponse) GetUserId() *wrappers.StringValue {
	if m != nil {
		return m.UserId
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateUserRequest)(nil), "skeleton.basic.iam.v1.CreateUserRequest")
	proto.RegisterType((*CreateUserResponse)(nil), "skeleton.basic.iam.v1.CreateUserResponse")
}

func init() { proto.RegisterFile("iam.proto", fileDescriptor_0a2c201915207782) }

var fileDescriptor_0a2c201915207782 = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x40, 0x4d, 0xd5, 0xaa, 0xe3, 0xc9, 0x05, 0x21, 0x14, 0x11, 0xc9, 0xa9, 0x5e, 0xa6, 0x18,
	0x11, 0x3c, 0xeb, 0xa9, 0xd7, 0x16, 0x3d, 0x78, 0x91, 0x49, 0x3a, 0x86, 0xc5, 0x24, 0xbb, 0x9d,
	0xd9, 0xb4, 0x5f, 0xe0, 0x57, 0xf9, 0x73, 0x92, 0x86, 0xaa, 0xa0, 0x87, 0x40, 0x6f, 0x0b, 0xfb,
	0xde, 0xf0, 0x76, 0x07, 0x4e, 0x2c, 0x55, 0xe8, 0xc5, 0x05, 0x67, 0xce, 0xf5, 0x9d, 0x4b, 0x0e,
	0xae, 0xc6, 0x8c, 0xd4, 0xe6, 0xd8, 0xde, 0xac, 0x6e, 0x46, 0x97, 0x85, 0x73, 0x45, 0xc9, 0x93,
	0x0d, 0x94, 0x35, 0x6f, 0x93, 0xb5, 0x90, 0xf7, 0x2c, 0xda, 0x69, 0xc9, 0x67, 0x04, 0x67, 0x8f,
	0xc2, 0x14, 0xf8, 0x49, 0x59, 0x66, 0xbc, 0x6c, 0x58, 0x83, 0xb9, 0x87, 0xe3, 0x46, 0x59, 0x6a,
	0xaa, 0x38, 0x8e, 0xae, 0xa2, 0xf1, 0x69, 0x7a, 0x81, 0xdd, 0x20, 0xdc, 0x0e, 0xc2, 0x79, 0x10,
	0x5b, 0x17, 0xcf, 0x54, 0x36, 0x3c, 0xfb, 0xa6, 0x5b, 0xd3, 0x93, 0xea, 0xda, 0xc9, 0x22, 0x1e,
	0xf4, 0x31, 0xb7, 0xb4, 0x49, 0xe1, 0x90, 0x2b, 0xb2, 0x65, 0xbc, 0xdf, 0x43, 0xeb, 0xd0, 0xe4,
	0x23, 0x02, 0xf3, 0xbb, 0x5e, 0xbd, 0xab, 0x95, 0x77, 0xc8, 0xbf, 0x83, 0xa3, 0xf6, 0xfc, 0x6a,
	0xfb, 0xd5, 0x0f, 0x5b, 0x78, 0xba, 0x48, 0x97, 0x00, 0x53, 0xaa, 0xe6, 0x2c, 0x2b, 0x9b, 0xb3,
	0xc9, 0x01, 0x7e, 0xa2, 0xcc, 0x18, 0xff, 0xdd, 0x0c, 0xfe, 0xf9, 0xf5, 0xd1, 0x75, 0x0f, 0xb2,
	0x7b, 0x61, 0xb2, 0xf7, 0x70, 0xf0, 0x32, 0xf0, 0x59, 0x36, 0xdc, 0x64, 0xdd, 0x7e, 0x05, 0x00,
	0x00, 0xff, 0xff, 0x89, 0xfe, 0xd9, 0x4e, 0x09, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IamServiceClient is the client API for IamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IamServiceClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
}

type iamServiceClient struct {
	cc *grpc.ClientConn
}

func NewIamServiceClient(cc *grpc.ClientConn) IamServiceClient {
	return &iamServiceClient{cc}
}

func (c *iamServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/skeleton.basic.iam.v1.IamService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IamServiceServer is the server API for IamService service.
type IamServiceServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
}

// UnimplementedIamServiceServer can be embedded to have forward compatible implementations.
type UnimplementedIamServiceServer struct {
}

func (*UnimplementedIamServiceServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}

func RegisterIamServiceServer(s *grpc.Server, srv IamServiceServer) {
	s.RegisterService(&_IamService_serviceDesc, srv)
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
		FullMethod: "/skeleton.basic.iam.v1.IamService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IamServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IamService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "skeleton.basic.iam.v1.IamService",
	HandlerType: (*IamServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _IamService_CreateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "iam.proto",
}
