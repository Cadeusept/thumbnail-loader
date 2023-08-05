// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: downloader.proto

package downloader_proto

import (
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

type DownloadTRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Link string `protobuf:"bytes,1,opt,name=Link,proto3" json:"Link,omitempty"`
}

func (x *DownloadTRequest) Reset() {
	*x = DownloadTRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_downloader_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadTRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadTRequest) ProtoMessage() {}

func (x *DownloadTRequest) ProtoReflect() protoreflect.Message {
	mi := &file_downloader_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadTRequest.ProtoReflect.Descriptor instead.
func (*DownloadTRequest) Descriptor() ([]byte, []int) {
	return file_downloader_proto_rawDescGZIP(), []int{0}
}

func (x *DownloadTRequest) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

type DownloadTResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Picture string `protobuf:"bytes,1,opt,name=Picture,proto3" json:"Picture,omitempty"`
}

func (x *DownloadTResponse) Reset() {
	*x = DownloadTResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_downloader_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadTResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadTResponse) ProtoMessage() {}

func (x *DownloadTResponse) ProtoReflect() protoreflect.Message {
	mi := &file_downloader_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadTResponse.ProtoReflect.Descriptor instead.
func (*DownloadTResponse) Descriptor() ([]byte, []int) {
	return file_downloader_proto_rawDescGZIP(), []int{1}
}

func (x *DownloadTResponse) GetPicture() string {
	if x != nil {
		return x.Picture
	}
	return ""
}

var File_downloader_proto protoreflect.FileDescriptor

var file_downloader_proto_rawDesc = []byte{
	0x0a, 0x10, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x10, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x10, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x54, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4c, 0x69, 0x6e, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x22, 0x2d, 0x0a, 0x11,
	0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x32, 0x73, 0x0a, 0x11, 0x44,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x5e, 0x0a, 0x11, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x68, 0x75, 0x6d,
	0x62, 0x6e, 0x61, 0x69, 0x6c, 0x12, 0x22, 0x2e, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61,
	0x64, 0x54, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x64, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x14, 0x5a, 0x12, 0x2e, 0x2f, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_downloader_proto_rawDescOnce sync.Once
	file_downloader_proto_rawDescData = file_downloader_proto_rawDesc
)

func file_downloader_proto_rawDescGZIP() []byte {
	file_downloader_proto_rawDescOnce.Do(func() {
		file_downloader_proto_rawDescData = protoimpl.X.CompressGZIP(file_downloader_proto_rawDescData)
	})
	return file_downloader_proto_rawDescData
}

var file_downloader_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_downloader_proto_goTypes = []interface{}{
	(*DownloadTRequest)(nil),  // 0: downloader_proto.DownloadTRequest
	(*DownloadTResponse)(nil), // 1: downloader_proto.DownloadTResponse
}
var file_downloader_proto_depIdxs = []int32{
	0, // 0: downloader_proto.DownloaderService.DownloadThumbnail:input_type -> downloader_proto.DownloadTRequest
	1, // 1: downloader_proto.DownloaderService.DownloadThumbnail:output_type -> downloader_proto.DownloadTResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_downloader_proto_init() }
func file_downloader_proto_init() {
	if File_downloader_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_downloader_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadTRequest); i {
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
		file_downloader_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadTResponse); i {
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
			RawDescriptor: file_downloader_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_downloader_proto_goTypes,
		DependencyIndexes: file_downloader_proto_depIdxs,
		MessageInfos:      file_downloader_proto_msgTypes,
	}.Build()
	File_downloader_proto = out.File
	file_downloader_proto_rawDesc = nil
	file_downloader_proto_goTypes = nil
	file_downloader_proto_depIdxs = nil
}
