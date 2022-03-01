// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: echo/v1/echo.proto

package v1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	common "github.com/yangchnet/skeleton/api/common"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateEchoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// echo request message
	Message *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateEchoRequest) Reset() {
	*x = CreateEchoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_v1_echo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEchoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEchoRequest) ProtoMessage() {}

func (x *CreateEchoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_echo_v1_echo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEchoRequest.ProtoReflect.Descriptor instead.
func (*CreateEchoRequest) Descriptor() ([]byte, []int) {
	return file_echo_v1_echo_proto_rawDescGZIP(), []int{0}
}

func (x *CreateEchoRequest) GetMessage() *wrapperspb.StringValue {
	if x != nil {
		return x.Message
	}
	return nil
}

type CreateEchoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// echo response message
	Echo *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=echo,proto3" json:"echo,omitempty"`
}

func (x *CreateEchoResponse) Reset() {
	*x = CreateEchoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_v1_echo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEchoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEchoResponse) ProtoMessage() {}

func (x *CreateEchoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_echo_v1_echo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEchoResponse.ProtoReflect.Descriptor instead.
func (*CreateEchoResponse) Descriptor() ([]byte, []int) {
	return file_echo_v1_echo_proto_rawDescGZIP(), []int{1}
}

func (x *CreateEchoResponse) GetEcho() *wrapperspb.StringValue {
	if x != nil {
		return x.Echo
	}
	return nil
}

type ListEchoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int64 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *ListEchoRequest) Reset() {
	*x = ListEchoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_v1_echo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEchoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEchoRequest) ProtoMessage() {}

func (x *ListEchoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_echo_v1_echo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEchoRequest.ProtoReflect.Descriptor instead.
func (*ListEchoRequest) Descriptor() ([]byte, []int) {
	return file_echo_v1_echo_proto_rawDescGZIP(), []int{2}
}

func (x *ListEchoRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListEchoRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ListEchoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// echo record list
	Records []*EchoRecord `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
	// list length
	ListLength int32 `protobuf:"varint,2,opt,name=listLength,proto3" json:"listLength,omitempty"`
}

func (x *ListEchoResponse) Reset() {
	*x = ListEchoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_v1_echo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEchoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEchoResponse) ProtoMessage() {}

func (x *ListEchoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_echo_v1_echo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEchoResponse.ProtoReflect.Descriptor instead.
func (*ListEchoResponse) Descriptor() ([]byte, []int) {
	return file_echo_v1_echo_proto_rawDescGZIP(), []int{3}
}

func (x *ListEchoResponse) GetRecords() []*EchoRecord {
	if x != nil {
		return x.Records
	}
	return nil
}

func (x *ListEchoResponse) GetListLength() int32 {
	if x != nil {
		return x.ListLength
	}
	return 0
}

type UpdateEchoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// echo record id
	Id *wrapperspb.Int64Value `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// new echo message
	NewEcho *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=new_echo,json=newEcho,proto3" json:"new_echo,omitempty"`
}

func (x *UpdateEchoRequest) Reset() {
	*x = UpdateEchoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_v1_echo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEchoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEchoRequest) ProtoMessage() {}

func (x *UpdateEchoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_echo_v1_echo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEchoRequest.ProtoReflect.Descriptor instead.
func (*UpdateEchoRequest) Descriptor() ([]byte, []int) {
	return file_echo_v1_echo_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateEchoRequest) GetId() *wrapperspb.Int64Value {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *UpdateEchoRequest) GetNewEcho() *wrapperspb.StringValue {
	if x != nil {
		return x.NewEcho
	}
	return nil
}

type UpdateEchoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// update result
	Result common.OperationResult `protobuf:"varint,1,opt,name=result,proto3,enum=common.OperationResult" json:"result,omitempty"`
}

func (x *UpdateEchoResponse) Reset() {
	*x = UpdateEchoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_v1_echo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEchoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEchoResponse) ProtoMessage() {}

func (x *UpdateEchoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_echo_v1_echo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEchoResponse.ProtoReflect.Descriptor instead.
func (*UpdateEchoResponse) Descriptor() ([]byte, []int) {
	return file_echo_v1_echo_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateEchoResponse) GetResult() common.OperationResult {
	if x != nil {
		return x.Result
	}
	return common.OperationResult(0)
}

type DeleteEchoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// delete record id
	Id *wrapperspb.Int64Value `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteEchoRequest) Reset() {
	*x = DeleteEchoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_v1_echo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEchoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEchoRequest) ProtoMessage() {}

func (x *DeleteEchoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_echo_v1_echo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEchoRequest.ProtoReflect.Descriptor instead.
func (*DeleteEchoRequest) Descriptor() ([]byte, []int) {
	return file_echo_v1_echo_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteEchoRequest) GetId() *wrapperspb.Int64Value {
	if x != nil {
		return x.Id
	}
	return nil
}

type DeleteEchoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// delete result
	Result common.OperationResult `protobuf:"varint,1,opt,name=result,proto3,enum=common.OperationResult" json:"result,omitempty"`
}

func (x *DeleteEchoResponse) Reset() {
	*x = DeleteEchoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_v1_echo_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEchoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEchoResponse) ProtoMessage() {}

func (x *DeleteEchoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_echo_v1_echo_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEchoResponse.ProtoReflect.Descriptor instead.
func (*DeleteEchoResponse) Descriptor() ([]byte, []int) {
	return file_echo_v1_echo_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteEchoResponse) GetResult() common.OperationResult {
	if x != nil {
		return x.Result
	}
	return common.OperationResult(0)
}

type EchoRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// echo record id
	Id *wrapperspb.Int64Value `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// echo request message
	Message *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// echo response message
	Echo *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=echo,proto3" json:"echo,omitempty"`
	// echo request time
	EchoTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=echo_time,json=echoTime,proto3" json:"echo_time,omitempty"`
}

func (x *EchoRecord) Reset() {
	*x = EchoRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_v1_echo_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoRecord) ProtoMessage() {}

func (x *EchoRecord) ProtoReflect() protoreflect.Message {
	mi := &file_echo_v1_echo_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoRecord.ProtoReflect.Descriptor instead.
func (*EchoRecord) Descriptor() ([]byte, []int) {
	return file_echo_v1_echo_proto_rawDescGZIP(), []int{8}
}

func (x *EchoRecord) GetId() *wrapperspb.Int64Value {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *EchoRecord) GetMessage() *wrapperspb.StringValue {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *EchoRecord) GetEcho() *wrapperspb.StringValue {
	if x != nil {
		return x.Echo
	}
	return nil
}

func (x *EchoRecord) GetEchoTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EchoTime
	}
	return nil
}

type GetEchoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetEchoRequest) Reset() {
	*x = GetEchoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_v1_echo_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEchoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEchoRequest) ProtoMessage() {}

func (x *GetEchoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_echo_v1_echo_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEchoRequest.ProtoReflect.Descriptor instead.
func (*GetEchoRequest) Descriptor() ([]byte, []int) {
	return file_echo_v1_echo_proto_rawDescGZIP(), []int{9}
}

func (x *GetEchoRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_echo_v1_echo_proto protoreflect.FileDescriptor

var file_echo_v1_echo_proto_rawDesc = []byte{
	0x0a, 0x12, 0x65, 0x63, 0x68, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76,
	0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x46, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x63, 0x68,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x65, 0x63, 0x68,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x65, 0x63, 0x68, 0x6f, 0x22, 0x3f, 0x0a, 0x0f, 0x4c,
	0x69, 0x73, 0x74, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x61, 0x0a, 0x10,
	0x4c, 0x69, 0x73, 0x74, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2d, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x63, 0x68, 0x6f,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12,
	0x1e, 0x0a, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x22,
	0x79, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x37, 0x0a, 0x08, 0x6e, 0x65, 0x77, 0x5f, 0x65, 0x63, 0x68, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x07, 0x6e, 0x65, 0x77, 0x45, 0x63, 0x68, 0x6f, 0x22, 0x45, 0x0a, 0x12, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2f, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x40, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x63, 0x68, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x45, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x63, 0x68,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0xdc, 0x01, 0x0a, 0x0a, 0x45,
	0x63, 0x68, 0x6f, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x2b, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x02, 0x69, 0x64, 0x12, 0x36, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x30,
	0x0a, 0x04, 0x65, 0x63, 0x68, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x65, 0x63, 0x68, 0x6f,
	0x12, 0x37, 0x0a, 0x09, 0x65, 0x63, 0x68, 0x6f, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x08, 0x65, 0x63, 0x68, 0x6f, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x32, 0xc4, 0x05, 0x0a, 0x0b,
	0x45, 0x63, 0x68, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x99, 0x01, 0x0a, 0x0a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x1a, 0x2e, 0x65, 0x63, 0x68,
	0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x63, 0x68, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x52, 0x92, 0x41, 0x38, 0x12, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x20, 0x65, 0x63, 0x68, 0x6f, 0x1a, 0x27, 0x70, 0x6f, 0x73, 0x74, 0x20, 0x61, 0x6e, 0x79, 0x20,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x20, 0x74, 0x6f, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x67, 0x65, 0x74, 0x20, 0x65, 0x63, 0x68, 0x6f, 0x62, 0x00,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x0c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x65, 0x63, 0x68, 0x6f, 0x3a, 0x01, 0x2a, 0x12, 0x81, 0x01, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74,
	0x45, 0x63, 0x68, 0x6f, 0x12, 0x18, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x63, 0x68,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x40, 0x92, 0x41, 0x28, 0x12, 0x09,
	0x6c, 0x69, 0x73, 0x74, 0x20, 0x65, 0x63, 0x68, 0x6f, 0x1a, 0x19, 0x67, 0x65, 0x74, 0x20, 0x65,
	0x63, 0x68, 0x6f, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x62, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x73, 0x12, 0x91, 0x01, 0x0a, 0x0a,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x1a, 0x2e, 0x65, 0x63, 0x68,
	0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x63, 0x68, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x4a, 0x92, 0x41, 0x29, 0x12, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x20, 0x65, 0x63, 0x68, 0x6f, 0x20, 0x1a, 0x19, 0x67, 0x65, 0x74, 0x20, 0x65, 0x63, 0x68, 0x6f,
	0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x1a, 0x13, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x65, 0x63, 0x68, 0x6f, 0x2f, 0x7b, 0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x3a, 0x01, 0x2a, 0x12,
	0x86, 0x01, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x1a,
	0x2e, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45,
	0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x65, 0x63, 0x68,
	0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x63, 0x68, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3f, 0x92, 0x41, 0x21, 0x12, 0x0b, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x20, 0x65, 0x63, 0x68, 0x6f, 0x1a, 0x12, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x20, 0x65, 0x63, 0x68, 0x6f, 0x20, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x15, 0x2a, 0x13, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x63, 0x68,
	0x6f, 0x2f, 0x7b, 0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x12, 0x78, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x45,
	0x63, 0x68, 0x6f, 0x12, 0x17, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x65,
	0x63, 0x68, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x22, 0x3f, 0x92, 0x41, 0x21, 0x12, 0x08, 0x67, 0x65, 0x74, 0x20, 0x65, 0x63, 0x68, 0x6f,
	0x1a, 0x15, 0x67, 0x65, 0x74, 0x20, 0x65, 0x63, 0x68, 0x6f, 0x20, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x20, 0x62, 0x79, 0x20, 0x69, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x2f, 0x7b, 0x69, 0x64, 0x3d,
	0x2a, 0x7d, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x79, 0x61, 0x6e, 0x67, 0x63, 0x68, 0x6e, 0x65, 0x74, 0x2f, 0x73, 0x6b, 0x65, 0x6c, 0x65,
	0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x2f, 0x76, 0x31, 0x3b,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_echo_v1_echo_proto_rawDescOnce sync.Once
	file_echo_v1_echo_proto_rawDescData = file_echo_v1_echo_proto_rawDesc
)

func file_echo_v1_echo_proto_rawDescGZIP() []byte {
	file_echo_v1_echo_proto_rawDescOnce.Do(func() {
		file_echo_v1_echo_proto_rawDescData = protoimpl.X.CompressGZIP(file_echo_v1_echo_proto_rawDescData)
	})
	return file_echo_v1_echo_proto_rawDescData
}

var file_echo_v1_echo_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_echo_v1_echo_proto_goTypes = []interface{}{
	(*CreateEchoRequest)(nil),      // 0: echo.v1.CreateEchoRequest
	(*CreateEchoResponse)(nil),     // 1: echo.v1.CreateEchoResponse
	(*ListEchoRequest)(nil),        // 2: echo.v1.ListEchoRequest
	(*ListEchoResponse)(nil),       // 3: echo.v1.ListEchoResponse
	(*UpdateEchoRequest)(nil),      // 4: echo.v1.UpdateEchoRequest
	(*UpdateEchoResponse)(nil),     // 5: echo.v1.UpdateEchoResponse
	(*DeleteEchoRequest)(nil),      // 6: echo.v1.DeleteEchoRequest
	(*DeleteEchoResponse)(nil),     // 7: echo.v1.DeleteEchoResponse
	(*EchoRecord)(nil),             // 8: echo.v1.EchoRecord
	(*GetEchoRequest)(nil),         // 9: echo.v1.GetEchoRequest
	(*wrapperspb.StringValue)(nil), // 10: google.protobuf.StringValue
	(*wrapperspb.Int64Value)(nil),  // 11: google.protobuf.Int64Value
	(common.OperationResult)(0),    // 12: common.OperationResult
	(*timestamppb.Timestamp)(nil),  // 13: google.protobuf.Timestamp
}
var file_echo_v1_echo_proto_depIdxs = []int32{
	10, // 0: echo.v1.CreateEchoRequest.message:type_name -> google.protobuf.StringValue
	10, // 1: echo.v1.CreateEchoResponse.echo:type_name -> google.protobuf.StringValue
	8,  // 2: echo.v1.ListEchoResponse.records:type_name -> echo.v1.EchoRecord
	11, // 3: echo.v1.UpdateEchoRequest.id:type_name -> google.protobuf.Int64Value
	10, // 4: echo.v1.UpdateEchoRequest.new_echo:type_name -> google.protobuf.StringValue
	12, // 5: echo.v1.UpdateEchoResponse.result:type_name -> common.OperationResult
	11, // 6: echo.v1.DeleteEchoRequest.id:type_name -> google.protobuf.Int64Value
	12, // 7: echo.v1.DeleteEchoResponse.result:type_name -> common.OperationResult
	11, // 8: echo.v1.EchoRecord.id:type_name -> google.protobuf.Int64Value
	10, // 9: echo.v1.EchoRecord.message:type_name -> google.protobuf.StringValue
	10, // 10: echo.v1.EchoRecord.echo:type_name -> google.protobuf.StringValue
	13, // 11: echo.v1.EchoRecord.echo_time:type_name -> google.protobuf.Timestamp
	0,  // 12: echo.v1.EchoService.CreateEcho:input_type -> echo.v1.CreateEchoRequest
	2,  // 13: echo.v1.EchoService.ListEcho:input_type -> echo.v1.ListEchoRequest
	4,  // 14: echo.v1.EchoService.UpdateEcho:input_type -> echo.v1.UpdateEchoRequest
	6,  // 15: echo.v1.EchoService.DeleteEcho:input_type -> echo.v1.DeleteEchoRequest
	9,  // 16: echo.v1.EchoService.GetEcho:input_type -> echo.v1.GetEchoRequest
	1,  // 17: echo.v1.EchoService.CreateEcho:output_type -> echo.v1.CreateEchoResponse
	3,  // 18: echo.v1.EchoService.ListEcho:output_type -> echo.v1.ListEchoResponse
	5,  // 19: echo.v1.EchoService.UpdateEcho:output_type -> echo.v1.UpdateEchoResponse
	7,  // 20: echo.v1.EchoService.DeleteEcho:output_type -> echo.v1.DeleteEchoResponse
	8,  // 21: echo.v1.EchoService.GetEcho:output_type -> echo.v1.EchoRecord
	17, // [17:22] is the sub-list for method output_type
	12, // [12:17] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_echo_v1_echo_proto_init() }
func file_echo_v1_echo_proto_init() {
	if File_echo_v1_echo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_echo_v1_echo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEchoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_echo_v1_echo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEchoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_echo_v1_echo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEchoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_echo_v1_echo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEchoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_echo_v1_echo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEchoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_echo_v1_echo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEchoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_echo_v1_echo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEchoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_echo_v1_echo_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEchoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_echo_v1_echo_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoRecord); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_echo_v1_echo_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEchoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_echo_v1_echo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_echo_v1_echo_proto_goTypes,
		DependencyIndexes: file_echo_v1_echo_proto_depIdxs,
		MessageInfos:      file_echo_v1_echo_proto_msgTypes,
	}.Build()
	File_echo_v1_echo_proto = out.File
	file_echo_v1_echo_proto_rawDesc = nil
	file_echo_v1_echo_proto_goTypes = nil
	file_echo_v1_echo_proto_depIdxs = nil
}
