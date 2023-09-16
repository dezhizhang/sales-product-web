// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: banner.proto

package proto

import (
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

type CreateBannerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image string `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	Url   string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Index int32  `protobuf:"varint,4,opt,name=index,proto3" json:"index,omitempty"`
}

func (x *CreateBannerRequest) Reset() {
	*x = CreateBannerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBannerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBannerRequest) ProtoMessage() {}

func (x *CreateBannerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_banner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBannerRequest.ProtoReflect.Descriptor instead.
func (*CreateBannerRequest) Descriptor() ([]byte, []int) {
	return file_banner_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBannerRequest) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *CreateBannerRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *CreateBannerRequest) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

type UpdateBannerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Image string `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	Url   string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Index int32  `protobuf:"varint,4,opt,name=index,proto3" json:"index,omitempty"`
}

func (x *UpdateBannerRequest) Reset() {
	*x = UpdateBannerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banner_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBannerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBannerRequest) ProtoMessage() {}

func (x *UpdateBannerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_banner_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBannerRequest.ProtoReflect.Descriptor instead.
func (*UpdateBannerRequest) Descriptor() ([]byte, []int) {
	return file_banner_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateBannerRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateBannerRequest) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *UpdateBannerRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *UpdateBannerRequest) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

type DeleteBannerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteBannerRequest) Reset() {
	*x = DeleteBannerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banner_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBannerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBannerRequest) ProtoMessage() {}

func (x *DeleteBannerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_banner_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBannerRequest.ProtoReflect.Descriptor instead.
func (*DeleteBannerRequest) Descriptor() ([]byte, []int) {
	return file_banner_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteBannerRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type BannerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageIndex int32 `protobuf:"varint,1,opt,name=pageIndex,proto3" json:"pageIndex,omitempty"`
	PageSize  int32 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *BannerRequest) Reset() {
	*x = BannerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banner_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BannerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BannerRequest) ProtoMessage() {}

func (x *BannerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_banner_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BannerRequest.ProtoReflect.Descriptor instead.
func (*BannerRequest) Descriptor() ([]byte, []int) {
	return file_banner_proto_rawDescGZIP(), []int{3}
}

func (x *BannerRequest) GetPageIndex() int32 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *BannerRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type BannerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Image string `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	Url   string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Index int32  `protobuf:"varint,4,opt,name=index,proto3" json:"index,omitempty"`
}

func (x *BannerResponse) Reset() {
	*x = BannerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banner_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BannerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BannerResponse) ProtoMessage() {}

func (x *BannerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_banner_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BannerResponse.ProtoReflect.Descriptor instead.
func (*BannerResponse) Descriptor() ([]byte, []int) {
	return file_banner_proto_rawDescGZIP(), []int{4}
}

func (x *BannerResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BannerResponse) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *BannerResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *BannerResponse) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

type BannerResponseList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int32             `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Data  []*BannerResponse `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *BannerResponseList) Reset() {
	*x = BannerResponseList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banner_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BannerResponseList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BannerResponseList) ProtoMessage() {}

func (x *BannerResponseList) ProtoReflect() protoreflect.Message {
	mi := &file_banner_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BannerResponseList.ProtoReflect.Descriptor instead.
func (*BannerResponseList) Descriptor() ([]byte, []int) {
	return file_banner_proto_rawDescGZIP(), []int{5}
}

func (x *BannerResponseList) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *BannerResponseList) GetData() []*BannerResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_banner_proto protoreflect.FileDescriptor

var file_banner_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x53, 0x0a, 0x13, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x22, 0x63, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x25, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42,
	0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x49, 0x0a, 0x0d,
	0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x5e, 0x0a, 0x0e, 0x42, 0x61, 0x6e, 0x6e, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x4f, 0x0a, 0x12, 0x42, 0x61, 0x6e, 0x6e, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x12, 0x23, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xf8, 0x01, 0x0a, 0x06, 0x42, 0x61, 0x6e,
	0x6e, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x0e, 0x2e, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x0c, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3c, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3c, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42,
	0x61, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61,
	0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_banner_proto_rawDescOnce sync.Once
	file_banner_proto_rawDescData = file_banner_proto_rawDesc
)

func file_banner_proto_rawDescGZIP() []byte {
	file_banner_proto_rawDescOnce.Do(func() {
		file_banner_proto_rawDescData = protoimpl.X.CompressGZIP(file_banner_proto_rawDescData)
	})
	return file_banner_proto_rawDescData
}

var file_banner_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_banner_proto_goTypes = []interface{}{
	(*CreateBannerRequest)(nil), // 0: CreateBannerRequest
	(*UpdateBannerRequest)(nil), // 1: UpdateBannerRequest
	(*DeleteBannerRequest)(nil), // 2: DeleteBannerRequest
	(*BannerRequest)(nil),       // 3: BannerRequest
	(*BannerResponse)(nil),      // 4: BannerResponse
	(*BannerResponseList)(nil),  // 5: BannerResponseList
	(*emptypb.Empty)(nil),       // 6: google.protobuf.Empty
}
var file_banner_proto_depIdxs = []int32{
	4, // 0: BannerResponseList.data:type_name -> BannerResponse
	3, // 1: Banner.GetBannerList:input_type -> BannerRequest
	0, // 2: Banner.CreateBanner:input_type -> CreateBannerRequest
	2, // 3: Banner.DeleteBanner:input_type -> DeleteBannerRequest
	1, // 4: Banner.UpdateBanner:input_type -> UpdateBannerRequest
	5, // 5: Banner.GetBannerList:output_type -> BannerResponseList
	6, // 6: Banner.CreateBanner:output_type -> google.protobuf.Empty
	6, // 7: Banner.DeleteBanner:output_type -> google.protobuf.Empty
	6, // 8: Banner.UpdateBanner:output_type -> google.protobuf.Empty
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_banner_proto_init() }
func file_banner_proto_init() {
	if File_banner_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_banner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBannerRequest); i {
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
		file_banner_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBannerRequest); i {
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
		file_banner_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBannerRequest); i {
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
		file_banner_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BannerRequest); i {
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
		file_banner_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BannerResponse); i {
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
		file_banner_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BannerResponseList); i {
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
			RawDescriptor: file_banner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_banner_proto_goTypes,
		DependencyIndexes: file_banner_proto_depIdxs,
		MessageInfos:      file_banner_proto_msgTypes,
	}.Build()
	File_banner_proto = out.File
	file_banner_proto_rawDesc = nil
	file_banner_proto_goTypes = nil
	file_banner_proto_depIdxs = nil
}