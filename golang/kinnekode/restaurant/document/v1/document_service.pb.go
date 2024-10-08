// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: kinnekode/restaurant/document/v1/document_service.proto

package v1

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

type GeneratePreviewRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestedDocument *RequestedDocument `protobuf:"bytes,1,opt,name=requested_document,json=requestedDocument,proto3" json:"requested_document,omitempty"`
}

func (x *GeneratePreviewRequest) Reset() {
	*x = GeneratePreviewRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kinnekode_restaurant_document_v1_document_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneratePreviewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneratePreviewRequest) ProtoMessage() {}

func (x *GeneratePreviewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kinnekode_restaurant_document_v1_document_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneratePreviewRequest.ProtoReflect.Descriptor instead.
func (*GeneratePreviewRequest) Descriptor() ([]byte, []int) {
	return file_kinnekode_restaurant_document_v1_document_service_proto_rawDescGZIP(), []int{0}
}

func (x *GeneratePreviewRequest) GetRequestedDocument() *RequestedDocument {
	if x != nil {
		return x.RequestedDocument
	}
	return nil
}

type GeneratePreviewResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The first message must be metadata
	// The next messages are only chunks of the file in order
	// On end of stream the file is ended
	//
	// Types that are assignable to File:
	//
	//	*GeneratePreviewResponse_Metadata
	//	*GeneratePreviewResponse_Chunk
	File isGeneratePreviewResponse_File `protobuf_oneof:"file"`
}

func (x *GeneratePreviewResponse) Reset() {
	*x = GeneratePreviewResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kinnekode_restaurant_document_v1_document_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneratePreviewResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneratePreviewResponse) ProtoMessage() {}

func (x *GeneratePreviewResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kinnekode_restaurant_document_v1_document_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneratePreviewResponse.ProtoReflect.Descriptor instead.
func (*GeneratePreviewResponse) Descriptor() ([]byte, []int) {
	return file_kinnekode_restaurant_document_v1_document_service_proto_rawDescGZIP(), []int{1}
}

func (m *GeneratePreviewResponse) GetFile() isGeneratePreviewResponse_File {
	if m != nil {
		return m.File
	}
	return nil
}

func (x *GeneratePreviewResponse) GetMetadata() *GeneratedFileMetadata {
	if x, ok := x.GetFile().(*GeneratePreviewResponse_Metadata); ok {
		return x.Metadata
	}
	return nil
}

func (x *GeneratePreviewResponse) GetChunk() []byte {
	if x, ok := x.GetFile().(*GeneratePreviewResponse_Chunk); ok {
		return x.Chunk
	}
	return nil
}

type isGeneratePreviewResponse_File interface {
	isGeneratePreviewResponse_File()
}

type GeneratePreviewResponse_Metadata struct {
	Metadata *GeneratedFileMetadata `protobuf:"bytes,1,opt,name=metadata,proto3,oneof"`
}

type GeneratePreviewResponse_Chunk struct {
	Chunk []byte `protobuf:"bytes,2,opt,name=chunk,proto3,oneof"`
}

func (*GeneratePreviewResponse_Metadata) isGeneratePreviewResponse_File() {}

func (*GeneratePreviewResponse_Chunk) isGeneratePreviewResponse_File() {}

var File_kinnekode_restaurant_document_v1_document_service_proto protoreflect.FileDescriptor

var file_kinnekode_restaurant_document_v1_document_service_proto_rawDesc = []byte{
	0x0a, 0x37, 0x6b, 0x69, 0x6e, 0x6e, 0x65, 0x6b, 0x6f, 0x64, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x74,
	0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2f,
	0x76, 0x31, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x6b, 0x69, 0x6e, 0x6e, 0x65,
	0x6b, 0x6f, 0x64, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e,
	0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x2f, 0x6b, 0x69, 0x6e,
	0x6e, 0x65, 0x6b, 0x6f, 0x64, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e,
	0x74, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x39, 0x6b, 0x69,
	0x6e, 0x6e, 0x65, 0x6b, 0x6f, 0x64, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61,
	0x6e, 0x74, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7c, 0x0a, 0x16, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x62, 0x0a, 0x12, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x64,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e,
	0x6b, 0x69, 0x6e, 0x6e, 0x65, 0x6b, 0x6f, 0x64, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75,
	0x72, 0x61, 0x6e, 0x74, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x11, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x44, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x90, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x55, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x6b, 0x69, 0x6e, 0x6e, 0x65, 0x6b, 0x6f, 0x64, 0x65, 0x2e,
	0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e,
	0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x42, 0x06, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x32, 0x9c, 0x01, 0x0a, 0x0f, 0x44, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x88, 0x01, 0x0a,
	0x0f, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x12, 0x38, 0x2e, 0x6b, 0x69, 0x6e, 0x6e, 0x65, 0x6b, 0x6f, 0x64, 0x65, 0x2e, 0x72, 0x65, 0x73,
	0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x50, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e, 0x6b, 0x69, 0x6e,
	0x6e, 0x65, 0x6b, 0x6f, 0x64, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e,
	0x74, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x4c, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x69, 0x6e, 0x6e, 0x65, 0x6b, 0x6f, 0x2d, 0x64, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x6b, 0x69, 0x6e, 0x6e, 0x65, 0x6b, 0x6f, 0x64, 0x65, 0x2f, 0x72,
	0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kinnekode_restaurant_document_v1_document_service_proto_rawDescOnce sync.Once
	file_kinnekode_restaurant_document_v1_document_service_proto_rawDescData = file_kinnekode_restaurant_document_v1_document_service_proto_rawDesc
)

func file_kinnekode_restaurant_document_v1_document_service_proto_rawDescGZIP() []byte {
	file_kinnekode_restaurant_document_v1_document_service_proto_rawDescOnce.Do(func() {
		file_kinnekode_restaurant_document_v1_document_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_kinnekode_restaurant_document_v1_document_service_proto_rawDescData)
	})
	return file_kinnekode_restaurant_document_v1_document_service_proto_rawDescData
}

var file_kinnekode_restaurant_document_v1_document_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_kinnekode_restaurant_document_v1_document_service_proto_goTypes = []any{
	(*GeneratePreviewRequest)(nil),  // 0: kinnekode.restaurant.document.v1.GeneratePreviewRequest
	(*GeneratePreviewResponse)(nil), // 1: kinnekode.restaurant.document.v1.GeneratePreviewResponse
	(*RequestedDocument)(nil),       // 2: kinnekode.restaurant.document.v1.RequestedDocument
	(*GeneratedFileMetadata)(nil),   // 3: kinnekode.restaurant.document.v1.GeneratedFileMetadata
}
var file_kinnekode_restaurant_document_v1_document_service_proto_depIdxs = []int32{
	2, // 0: kinnekode.restaurant.document.v1.GeneratePreviewRequest.requested_document:type_name -> kinnekode.restaurant.document.v1.RequestedDocument
	3, // 1: kinnekode.restaurant.document.v1.GeneratePreviewResponse.metadata:type_name -> kinnekode.restaurant.document.v1.GeneratedFileMetadata
	0, // 2: kinnekode.restaurant.document.v1.DocumentService.GeneratePreview:input_type -> kinnekode.restaurant.document.v1.GeneratePreviewRequest
	1, // 3: kinnekode.restaurant.document.v1.DocumentService.GeneratePreview:output_type -> kinnekode.restaurant.document.v1.GeneratePreviewResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_kinnekode_restaurant_document_v1_document_service_proto_init() }
func file_kinnekode_restaurant_document_v1_document_service_proto_init() {
	if File_kinnekode_restaurant_document_v1_document_service_proto != nil {
		return
	}
	file_kinnekode_restaurant_document_v1_document_proto_init()
	file_kinnekode_restaurant_document_v1_requested_document_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_kinnekode_restaurant_document_v1_document_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GeneratePreviewRequest); i {
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
		file_kinnekode_restaurant_document_v1_document_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GeneratePreviewResponse); i {
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
	file_kinnekode_restaurant_document_v1_document_service_proto_msgTypes[1].OneofWrappers = []any{
		(*GeneratePreviewResponse_Metadata)(nil),
		(*GeneratePreviewResponse_Chunk)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_kinnekode_restaurant_document_v1_document_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kinnekode_restaurant_document_v1_document_service_proto_goTypes,
		DependencyIndexes: file_kinnekode_restaurant_document_v1_document_service_proto_depIdxs,
		MessageInfos:      file_kinnekode_restaurant_document_v1_document_service_proto_msgTypes,
	}.Build()
	File_kinnekode_restaurant_document_v1_document_service_proto = out.File
	file_kinnekode_restaurant_document_v1_document_service_proto_rawDesc = nil
	file_kinnekode_restaurant_document_v1_document_service_proto_goTypes = nil
	file_kinnekode_restaurant_document_v1_document_service_proto_depIdxs = nil
}
