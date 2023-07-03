// The processing service definition.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: grpc/processing_service.proto

package endpoint

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProcessingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ProcessingRequest) Reset() {
	*x = ProcessingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_processing_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessingRequest) ProtoMessage() {}

func (x *ProcessingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_processing_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessingRequest.ProtoReflect.Descriptor instead.
func (*ProcessingRequest) Descriptor() ([]byte, []int) {
	return file_grpc_processing_service_proto_rawDescGZIP(), []int{0}
}

func (x *ProcessingRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type ProcessingReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ProcessingReply) Reset() {
	*x = ProcessingReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_processing_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessingReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessingReply) ProtoMessage() {}

func (x *ProcessingReply) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_processing_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessingReply.ProtoReflect.Descriptor instead.
func (*ProcessingReply) Descriptor() ([]byte, []int) {
	return file_grpc_processing_service_proto_rawDescGZIP(), []int{1}
}

func (x *ProcessingReply) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type RetrivingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *RetrivingRequest) Reset() {
	*x = RetrivingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_processing_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrivingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrivingRequest) ProtoMessage() {}

func (x *RetrivingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_processing_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrivingRequest.ProtoReflect.Descriptor instead.
func (*RetrivingRequest) Descriptor() ([]byte, []int) {
	return file_grpc_processing_service_proto_rawDescGZIP(), []int{2}
}

func (x *RetrivingRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type RetrivingReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *RetrivingReply) Reset() {
	*x = RetrivingReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_processing_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrivingReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrivingReply) ProtoMessage() {}

func (x *RetrivingReply) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_processing_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrivingReply.ProtoReflect.Descriptor instead.
func (*RetrivingReply) Descriptor() ([]byte, []int) {
	return file_grpc_processing_service_proto_rawDescGZIP(), []int{3}
}

func (x *RetrivingReply) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_grpc_processing_service_proto protoreflect.FileDescriptor

var file_grpc_processing_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e,
	0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x27, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x27, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x28, 0x0a, 0x10, 0x52, 0x65, 0x74, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x24, 0x0a, 0x0e, 0x52,
	0x65, 0x74, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x32, 0x77, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x12, 0x12, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x07, 0x52, 0x65, 0x74,
	0x72, 0x69, 0x76, 0x65, 0x12, 0x11, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x76,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x74, 0x6f,
	0x79, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_processing_service_proto_rawDescOnce sync.Once
	file_grpc_processing_service_proto_rawDescData = file_grpc_processing_service_proto_rawDesc
)

func file_grpc_processing_service_proto_rawDescGZIP() []byte {
	file_grpc_processing_service_proto_rawDescOnce.Do(func() {
		file_grpc_processing_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_processing_service_proto_rawDescData)
	})
	return file_grpc_processing_service_proto_rawDescData
}

var file_grpc_processing_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_grpc_processing_service_proto_goTypes = []interface{}{
	(*ProcessingRequest)(nil), // 0: ProcessingRequest
	(*ProcessingReply)(nil),   // 1: ProcessingReply
	(*RetrivingRequest)(nil),  // 2: RetrivingRequest
	(*RetrivingReply)(nil),    // 3: RetrivingReply
}
var file_grpc_processing_service_proto_depIdxs = []int32{
	0, // 0: ProcessingService.Process:input_type -> ProcessingRequest
	2, // 1: ProcessingService.Retrive:input_type -> RetrivingRequest
	1, // 2: ProcessingService.Process:output_type -> ProcessingReply
	3, // 3: ProcessingService.Retrive:output_type -> RetrivingReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_processing_service_proto_init() }
func file_grpc_processing_service_proto_init() {
	if File_grpc_processing_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_processing_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessingRequest); i {
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
		file_grpc_processing_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessingReply); i {
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
		file_grpc_processing_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrivingRequest); i {
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
		file_grpc_processing_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrivingReply); i {
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
			RawDescriptor: file_grpc_processing_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_processing_service_proto_goTypes,
		DependencyIndexes: file_grpc_processing_service_proto_depIdxs,
		MessageInfos:      file_grpc_processing_service_proto_msgTypes,
	}.Build()
	File_grpc_processing_service_proto = out.File
	file_grpc_processing_service_proto_rawDesc = nil
	file_grpc_processing_service_proto_goTypes = nil
	file_grpc_processing_service_proto_depIdxs = nil
}
