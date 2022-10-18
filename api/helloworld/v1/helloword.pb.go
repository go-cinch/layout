// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: helloworld/v1/helloword.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Page struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num     uint64 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	Size    uint64 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Total   int64  `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	Disable bool   `protobuf:"varint,4,opt,name=disable,proto3" json:"disable,omitempty"`
}

func (x *Page) Reset() {
	*x = Page{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_v1_helloword_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Page) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Page) ProtoMessage() {}

func (x *Page) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_v1_helloword_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Page.ProtoReflect.Descriptor instead.
func (*Page) Descriptor() ([]byte, []int) {
	return file_helloworld_v1_helloword_proto_rawDescGZIP(), []int{0}
}

func (x *Page) GetNum() uint64 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *Page) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Page) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Page) GetDisable() bool {
	if x != nil {
		return x.Disable
	}
	return false
}

type IdsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids string `protobuf:"bytes,1,opt,name=ids,proto3" json:"ids,omitempty"`
}

func (x *IdsRequest) Reset() {
	*x = IdsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_v1_helloword_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdsRequest) ProtoMessage() {}

func (x *IdsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_v1_helloword_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdsRequest.ProtoReflect.Descriptor instead.
func (*IdsRequest) Descriptor() ([]byte, []int) {
	return file_helloworld_v1_helloword_proto_rawDescGZIP(), []int{1}
}

func (x *IdsRequest) GetIds() string {
	if x != nil {
		return x.Ids
	}
	return ""
}

type Greeter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age  int32  `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *Greeter) Reset() {
	*x = Greeter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_v1_helloword_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Greeter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Greeter) ProtoMessage() {}

func (x *Greeter) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_v1_helloword_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Greeter.ProtoReflect.Descriptor instead.
func (*Greeter) Descriptor() ([]byte, []int) {
	return file_helloworld_v1_helloword_proto_rawDescGZIP(), []int{2}
}

func (x *Greeter) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Greeter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Greeter) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

type CreateGreeterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` // the name of string must be between 2 and 50 character
	Age  int32  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`  // the age of int32 must be >= 0
}

func (x *CreateGreeterRequest) Reset() {
	*x = CreateGreeterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_v1_helloword_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGreeterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGreeterRequest) ProtoMessage() {}

func (x *CreateGreeterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_v1_helloword_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGreeterRequest.ProtoReflect.Descriptor instead.
func (*CreateGreeterRequest) Descriptor() ([]byte, []int) {
	return file_helloworld_v1_helloword_proto_rawDescGZIP(), []int{3}
}

func (x *CreateGreeterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateGreeterRequest) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

type GetGreeterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetGreeterRequest) Reset() {
	*x = GetGreeterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_v1_helloword_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGreeterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGreeterRequest) ProtoMessage() {}

func (x *GetGreeterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_v1_helloword_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGreeterRequest.ProtoReflect.Descriptor instead.
func (*GetGreeterRequest) Descriptor() ([]byte, []int) {
	return file_helloworld_v1_helloword_proto_rawDescGZIP(), []int{4}
}

func (x *GetGreeterRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetGreeterReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age  int32  `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *GetGreeterReply) Reset() {
	*x = GetGreeterReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_v1_helloword_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGreeterReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGreeterReply) ProtoMessage() {}

func (x *GetGreeterReply) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_v1_helloword_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGreeterReply.ProtoReflect.Descriptor instead.
func (*GetGreeterReply) Descriptor() ([]byte, []int) {
	return file_helloworld_v1_helloword_proto_rawDescGZIP(), []int{5}
}

func (x *GetGreeterReply) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetGreeterReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetGreeterReply) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

type FindGreeterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page *Page   `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	Name *string `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Age  *int32  `protobuf:"varint,3,opt,name=age,proto3,oneof" json:"age,omitempty"`
}

func (x *FindGreeterRequest) Reset() {
	*x = FindGreeterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_v1_helloword_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindGreeterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindGreeterRequest) ProtoMessage() {}

func (x *FindGreeterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_v1_helloword_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindGreeterRequest.ProtoReflect.Descriptor instead.
func (*FindGreeterRequest) Descriptor() ([]byte, []int) {
	return file_helloworld_v1_helloword_proto_rawDescGZIP(), []int{6}
}

func (x *FindGreeterRequest) GetPage() *Page {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *FindGreeterRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *FindGreeterRequest) GetAge() int32 {
	if x != nil && x.Age != nil {
		return *x.Age
	}
	return 0
}

type FindGreeterReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page *Page      `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	List []*Greeter `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *FindGreeterReply) Reset() {
	*x = FindGreeterReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_v1_helloword_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindGreeterReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindGreeterReply) ProtoMessage() {}

func (x *FindGreeterReply) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_v1_helloword_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindGreeterReply.ProtoReflect.Descriptor instead.
func (*FindGreeterReply) Descriptor() ([]byte, []int) {
	return file_helloworld_v1_helloword_proto_rawDescGZIP(), []int{7}
}

func (x *FindGreeterReply) GetPage() *Page {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *FindGreeterReply) GetList() []*Greeter {
	if x != nil {
		return x.List
	}
	return nil
}

type UpdateGreeterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name *string `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"` // the name of string must be between 2 and 50 character
	Age  *int32  `protobuf:"varint,3,opt,name=age,proto3,oneof" json:"age,omitempty"`  // the age of int32 must be >= 0
}

func (x *UpdateGreeterRequest) Reset() {
	*x = UpdateGreeterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_v1_helloword_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGreeterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGreeterRequest) ProtoMessage() {}

func (x *UpdateGreeterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_v1_helloword_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGreeterRequest.ProtoReflect.Descriptor instead.
func (*UpdateGreeterRequest) Descriptor() ([]byte, []int) {
	return file_helloworld_v1_helloword_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateGreeterRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateGreeterRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *UpdateGreeterRequest) GetAge() int32 {
	if x != nil && x.Age != nil {
		return *x.Age
	}
	return 0
}

var File_helloworld_v1_helloword_proto protoreflect.FileDescriptor

var file_helloworld_v1_helloword_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x76, 0x31, 0x2f,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0d, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x5c, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65,
	0x22, 0x1e, 0x0a, 0x0a, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73,
	0x22, 0x3f, 0x0a, 0x07, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67,
	0x65, 0x22, 0x50, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x65, 0x65, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x02,
	0x18, 0x32, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x28, 0x00, 0x52, 0x03,
	0x61, 0x67, 0x65, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x47, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x47,
	0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67,
	0x65, 0x22, 0x7e, 0x0a, 0x12, 0x46, 0x69, 0x6e, 0x64, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x03, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01,
	0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x61, 0x67,
	0x65, 0x22, 0x67, 0x0a, 0x10, 0x46, 0x69, 0x6e, 0x64, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x27, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x2a,
	0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x65,
	0x65, 0x74, 0x65, 0x72, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x7b, 0x0a, 0x14, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x22, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x02, 0x18, 0x32, 0x48, 0x00, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1e, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x28, 0x00, 0x48, 0x01, 0x52, 0x03,
	0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42,
	0x06, 0x0a, 0x04, 0x5f, 0x61, 0x67, 0x65, 0x32, 0x93, 0x04, 0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x12, 0x61, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x12, 0x23, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72,
	0x65, 0x65, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x22, 0x08, 0x2f, 0x67,
	0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x12, 0x65, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x65, 0x65, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x65,
	0x65, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0f, 0x12, 0x0d, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x12, 0x63, 0x0a, 0x0b, 0x46, 0x69, 0x6e, 0x64, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x12,
	0x21, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x46, 0x69, 0x6e, 0x64, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f, 0x67, 0x72,
	0x65, 0x65, 0x74, 0x65, 0x72, 0x12, 0x7a, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47,
	0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x12, 0x23, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x65,
	0x65, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x2c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x1a, 0x0d, 0x2f, 0x67, 0x72,
	0x65, 0x65, 0x74, 0x65, 0x72, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x5a, 0x12, 0x32,
	0x0d, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x01,
	0x2a, 0x12, 0x5a, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x65, 0x65, 0x74,
	0x65, 0x72, 0x12, 0x19, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x2a, 0x0e, 0x2f,
	0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2f, 0x7b, 0x69, 0x64, 0x73, 0x7d, 0x42, 0x55, 0x0a,
	0x0d, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x42, 0x11,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56,
	0x31, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x6f, 0x2d, 0x63, 0x69, 0x6e, 0x63, 0x68, 0x2f, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x76,
	0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_helloworld_v1_helloword_proto_rawDescOnce sync.Once
	file_helloworld_v1_helloword_proto_rawDescData = file_helloworld_v1_helloword_proto_rawDesc
)

func file_helloworld_v1_helloword_proto_rawDescGZIP() []byte {
	file_helloworld_v1_helloword_proto_rawDescOnce.Do(func() {
		file_helloworld_v1_helloword_proto_rawDescData = protoimpl.X.CompressGZIP(file_helloworld_v1_helloword_proto_rawDescData)
	})
	return file_helloworld_v1_helloword_proto_rawDescData
}

var file_helloworld_v1_helloword_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_helloworld_v1_helloword_proto_goTypes = []interface{}{
	(*Page)(nil),                 // 0: helloworld.v1.Page
	(*IdsRequest)(nil),           // 1: helloworld.v1.IdsRequest
	(*Greeter)(nil),              // 2: helloworld.v1.Greeter
	(*CreateGreeterRequest)(nil), // 3: helloworld.v1.CreateGreeterRequest
	(*GetGreeterRequest)(nil),    // 4: helloworld.v1.GetGreeterRequest
	(*GetGreeterReply)(nil),      // 5: helloworld.v1.GetGreeterReply
	(*FindGreeterRequest)(nil),   // 6: helloworld.v1.FindGreeterRequest
	(*FindGreeterReply)(nil),     // 7: helloworld.v1.FindGreeterReply
	(*UpdateGreeterRequest)(nil), // 8: helloworld.v1.UpdateGreeterRequest
	(*emptypb.Empty)(nil),        // 9: google.protobuf.Empty
}
var file_helloworld_v1_helloword_proto_depIdxs = []int32{
	0, // 0: helloworld.v1.FindGreeterRequest.page:type_name -> helloworld.v1.Page
	0, // 1: helloworld.v1.FindGreeterReply.page:type_name -> helloworld.v1.Page
	2, // 2: helloworld.v1.FindGreeterReply.list:type_name -> helloworld.v1.Greeter
	3, // 3: helloworld.v1.Helloworld.CreateGreeter:input_type -> helloworld.v1.CreateGreeterRequest
	4, // 4: helloworld.v1.Helloworld.GetGreeter:input_type -> helloworld.v1.GetGreeterRequest
	6, // 5: helloworld.v1.Helloworld.FindGreeter:input_type -> helloworld.v1.FindGreeterRequest
	8, // 6: helloworld.v1.Helloworld.UpdateGreeter:input_type -> helloworld.v1.UpdateGreeterRequest
	1, // 7: helloworld.v1.Helloworld.DeleteGreeter:input_type -> helloworld.v1.IdsRequest
	9, // 8: helloworld.v1.Helloworld.CreateGreeter:output_type -> google.protobuf.Empty
	5, // 9: helloworld.v1.Helloworld.GetGreeter:output_type -> helloworld.v1.GetGreeterReply
	7, // 10: helloworld.v1.Helloworld.FindGreeter:output_type -> helloworld.v1.FindGreeterReply
	9, // 11: helloworld.v1.Helloworld.UpdateGreeter:output_type -> google.protobuf.Empty
	9, // 12: helloworld.v1.Helloworld.DeleteGreeter:output_type -> google.protobuf.Empty
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_helloworld_v1_helloword_proto_init() }
func file_helloworld_v1_helloword_proto_init() {
	if File_helloworld_v1_helloword_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_helloworld_v1_helloword_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Page); i {
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
		file_helloworld_v1_helloword_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdsRequest); i {
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
		file_helloworld_v1_helloword_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Greeter); i {
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
		file_helloworld_v1_helloword_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGreeterRequest); i {
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
		file_helloworld_v1_helloword_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGreeterRequest); i {
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
		file_helloworld_v1_helloword_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGreeterReply); i {
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
		file_helloworld_v1_helloword_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindGreeterRequest); i {
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
		file_helloworld_v1_helloword_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindGreeterReply); i {
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
		file_helloworld_v1_helloword_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGreeterRequest); i {
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
	file_helloworld_v1_helloword_proto_msgTypes[6].OneofWrappers = []interface{}{}
	file_helloworld_v1_helloword_proto_msgTypes[8].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_helloworld_v1_helloword_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_helloworld_v1_helloword_proto_goTypes,
		DependencyIndexes: file_helloworld_v1_helloword_proto_depIdxs,
		MessageInfos:      file_helloworld_v1_helloword_proto_msgTypes,
	}.Build()
	File_helloworld_v1_helloword_proto = out.File
	file_helloworld_v1_helloword_proto_rawDesc = nil
	file_helloworld_v1_helloword_proto_goTypes = nil
	file_helloworld_v1_helloword_proto_depIdxs = nil
}
