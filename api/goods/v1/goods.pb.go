// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: goods/v1/goods.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price float32 `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_v1_goods_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_goods_v1_goods_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_goods_v1_goods_proto_rawDescGZIP(), []int{0}
}

func (x *Item) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Item) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type CreateGoodsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Price float32 `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *CreateGoodsRequest) Reset() {
	*x = CreateGoodsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_v1_goods_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGoodsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGoodsRequest) ProtoMessage() {}

func (x *CreateGoodsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goods_v1_goods_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGoodsRequest.ProtoReflect.Descriptor instead.
func (*CreateGoodsRequest) Descriptor() ([]byte, []int) {
	return file_goods_v1_goods_proto_rawDescGZIP(), []int{1}
}

func (x *CreateGoodsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateGoodsRequest) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type CreateGoodsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price float32 `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *CreateGoodsReply) Reset() {
	*x = CreateGoodsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_v1_goods_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGoodsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGoodsReply) ProtoMessage() {}

func (x *CreateGoodsReply) ProtoReflect() protoreflect.Message {
	mi := &file_goods_v1_goods_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGoodsReply.ProtoReflect.Descriptor instead.
func (*CreateGoodsReply) Descriptor() ([]byte, []int) {
	return file_goods_v1_goods_proto_rawDescGZIP(), []int{2}
}

func (x *CreateGoodsReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateGoodsReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateGoodsReply) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type UpdateGoodsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price float32 `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *UpdateGoodsRequest) Reset() {
	*x = UpdateGoodsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_v1_goods_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGoodsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGoodsRequest) ProtoMessage() {}

func (x *UpdateGoodsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goods_v1_goods_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGoodsRequest.ProtoReflect.Descriptor instead.
func (*UpdateGoodsRequest) Descriptor() ([]byte, []int) {
	return file_goods_v1_goods_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateGoodsRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateGoodsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateGoodsRequest) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type UpdateGoodsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price float32 `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *UpdateGoodsReply) Reset() {
	*x = UpdateGoodsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_v1_goods_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGoodsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGoodsReply) ProtoMessage() {}

func (x *UpdateGoodsReply) ProtoReflect() protoreflect.Message {
	mi := &file_goods_v1_goods_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGoodsReply.ProtoReflect.Descriptor instead.
func (*UpdateGoodsReply) Descriptor() ([]byte, []int) {
	return file_goods_v1_goods_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateGoodsReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateGoodsReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateGoodsReply) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type DeleteGoodsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteGoodsRequest) Reset() {
	*x = DeleteGoodsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_v1_goods_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGoodsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGoodsRequest) ProtoMessage() {}

func (x *DeleteGoodsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goods_v1_goods_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGoodsRequest.ProtoReflect.Descriptor instead.
func (*DeleteGoodsRequest) Descriptor() ([]byte, []int) {
	return file_goods_v1_goods_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteGoodsRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteGoodsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteGoodsReply) Reset() {
	*x = DeleteGoodsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_v1_goods_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGoodsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGoodsReply) ProtoMessage() {}

func (x *DeleteGoodsReply) ProtoReflect() protoreflect.Message {
	mi := &file_goods_v1_goods_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGoodsReply.ProtoReflect.Descriptor instead.
func (*DeleteGoodsReply) Descriptor() ([]byte, []int) {
	return file_goods_v1_goods_proto_rawDescGZIP(), []int{6}
}

type GetGoodsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetGoodsRequest) Reset() {
	*x = GetGoodsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_v1_goods_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGoodsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGoodsRequest) ProtoMessage() {}

func (x *GetGoodsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goods_v1_goods_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGoodsRequest.ProtoReflect.Descriptor instead.
func (*GetGoodsRequest) Descriptor() ([]byte, []int) {
	return file_goods_v1_goods_proto_rawDescGZIP(), []int{7}
}

func (x *GetGoodsRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetGoodsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price float32 `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *GetGoodsReply) Reset() {
	*x = GetGoodsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_v1_goods_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGoodsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGoodsReply) ProtoMessage() {}

func (x *GetGoodsReply) ProtoReflect() protoreflect.Message {
	mi := &file_goods_v1_goods_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGoodsReply.ProtoReflect.Descriptor instead.
func (*GetGoodsReply) Descriptor() ([]byte, []int) {
	return file_goods_v1_goods_proto_rawDescGZIP(), []int{8}
}

func (x *GetGoodsReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetGoodsReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetGoodsReply) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type ListGoodsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListGoodsRequest) Reset() {
	*x = ListGoodsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_v1_goods_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGoodsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGoodsRequest) ProtoMessage() {}

func (x *ListGoodsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goods_v1_goods_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGoodsRequest.ProtoReflect.Descriptor instead.
func (*ListGoodsRequest) Descriptor() ([]byte, []int) {
	return file_goods_v1_goods_proto_rawDescGZIP(), []int{9}
}

type ListGoodsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*Item `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *ListGoodsReply) Reset() {
	*x = ListGoodsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_v1_goods_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGoodsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGoodsReply) ProtoMessage() {}

func (x *ListGoodsReply) ProtoReflect() protoreflect.Message {
	mi := &file_goods_v1_goods_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGoodsReply.ProtoReflect.Descriptor instead.
func (*ListGoodsReply) Descriptor() ([]byte, []int) {
	return file_goods_v1_goods_proto_rawDescGZIP(), []int{10}
}

func (x *ListGoodsReply) GetList() []*Item {
	if x != nil {
		return x.List
	}
	return nil
}

var File_goods_v1_goods_proto protoreflect.FileDescriptor

var file_goods_v1_goods_proto_rawDesc = []byte{
	0x0a, 0x14, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6f, 0x6f, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x40, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x22, 0x3e, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x6f,
	0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x22, 0x4c, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x6f,
	0x6f, 0x64, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x22, 0x4e, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x22, 0x4c, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64,
	0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x22, 0x24, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x12, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x21, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x49, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74,
	0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x38, 0x0a, 0x0e,
	0x4c, 0x69, 0x73, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x26,
	0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x32, 0xc5, 0x01, 0x0a, 0x05, 0x47, 0x6f, 0x6f, 0x64, 0x73,
	0x12, 0x5e, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x12, 0x1d, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x47,
	0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x6f,
	0x6f, 0x64, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10,
	0x12, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x12, 0x5c, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x12, 0x1e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x11, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x42, 0x17,
	0x5a, 0x15, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x6f, 0x64,
	0x73, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_goods_v1_goods_proto_rawDescOnce sync.Once
	file_goods_v1_goods_proto_rawDescData = file_goods_v1_goods_proto_rawDesc
)

func file_goods_v1_goods_proto_rawDescGZIP() []byte {
	file_goods_v1_goods_proto_rawDescOnce.Do(func() {
		file_goods_v1_goods_proto_rawDescData = protoimpl.X.CompressGZIP(file_goods_v1_goods_proto_rawDescData)
	})
	return file_goods_v1_goods_proto_rawDescData
}

var file_goods_v1_goods_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_goods_v1_goods_proto_goTypes = []interface{}{
	(*Item)(nil),               // 0: api.goods.v1.Item
	(*CreateGoodsRequest)(nil), // 1: api.goods.v1.CreateGoodsRequest
	(*CreateGoodsReply)(nil),   // 2: api.goods.v1.CreateGoodsReply
	(*UpdateGoodsRequest)(nil), // 3: api.goods.v1.UpdateGoodsRequest
	(*UpdateGoodsReply)(nil),   // 4: api.goods.v1.UpdateGoodsReply
	(*DeleteGoodsRequest)(nil), // 5: api.goods.v1.DeleteGoodsRequest
	(*DeleteGoodsReply)(nil),   // 6: api.goods.v1.DeleteGoodsReply
	(*GetGoodsRequest)(nil),    // 7: api.goods.v1.GetGoodsRequest
	(*GetGoodsReply)(nil),      // 8: api.goods.v1.GetGoodsReply
	(*ListGoodsRequest)(nil),   // 9: api.goods.v1.ListGoodsRequest
	(*ListGoodsReply)(nil),     // 10: api.goods.v1.ListGoodsReply
}
var file_goods_v1_goods_proto_depIdxs = []int32{
	0,  // 0: api.goods.v1.ListGoodsReply.list:type_name -> api.goods.v1.Item
	7,  // 1: api.goods.v1.Goods.GetGoods:input_type -> api.goods.v1.GetGoodsRequest
	9,  // 2: api.goods.v1.Goods.ListGoods:input_type -> api.goods.v1.ListGoodsRequest
	8,  // 3: api.goods.v1.Goods.GetGoods:output_type -> api.goods.v1.GetGoodsReply
	10, // 4: api.goods.v1.Goods.ListGoods:output_type -> api.goods.v1.ListGoodsReply
	3,  // [3:5] is the sub-list for method output_type
	1,  // [1:3] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_goods_v1_goods_proto_init() }
func file_goods_v1_goods_proto_init() {
	if File_goods_v1_goods_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_goods_v1_goods_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
		file_goods_v1_goods_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGoodsRequest); i {
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
		file_goods_v1_goods_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGoodsReply); i {
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
		file_goods_v1_goods_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGoodsRequest); i {
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
		file_goods_v1_goods_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGoodsReply); i {
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
		file_goods_v1_goods_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGoodsRequest); i {
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
		file_goods_v1_goods_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGoodsReply); i {
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
		file_goods_v1_goods_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGoodsRequest); i {
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
		file_goods_v1_goods_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGoodsReply); i {
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
		file_goods_v1_goods_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGoodsRequest); i {
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
		file_goods_v1_goods_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGoodsReply); i {
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
			RawDescriptor: file_goods_v1_goods_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_goods_v1_goods_proto_goTypes,
		DependencyIndexes: file_goods_v1_goods_proto_depIdxs,
		MessageInfos:      file_goods_v1_goods_proto_msgTypes,
	}.Build()
	File_goods_v1_goods_proto = out.File
	file_goods_v1_goods_proto_rawDesc = nil
	file_goods_v1_goods_proto_goTypes = nil
	file_goods_v1_goods_proto_depIdxs = nil
}
