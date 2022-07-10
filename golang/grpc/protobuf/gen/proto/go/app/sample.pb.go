// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: app/sample.proto

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

type AA struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AA) Reset() {
	*x = AA{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_sample_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AA) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AA) ProtoMessage() {}

func (x *AA) ProtoReflect() protoreflect.Message {
	mi := &file_app_sample_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AA.ProtoReflect.Descriptor instead.
func (*AA) Descriptor() ([]byte, []int) {
	return file_app_sample_proto_rawDescGZIP(), []int{0}
}

func (x *AA) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type BB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *BB) Reset() {
	*x = BB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_sample_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BB) ProtoMessage() {}

func (x *BB) ProtoReflect() protoreflect.Message {
	mi := &file_app_sample_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BB.ProtoReflect.Descriptor instead.
func (*BB) Descriptor() ([]byte, []int) {
	return file_app_sample_proto_rawDescGZIP(), []int{1}
}

func (x *BB) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_app_sample_proto protoreflect.FileDescriptor

var file_app_sample_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x70, 0x70, 0x2f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x18, 0x0a, 0x02, 0x41, 0x41, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x18, 0x0a, 0x02, 0x42, 0x42, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x6f, 0x0a, 0x04,
	0x48, 0x6f, 0x67, 0x65, 0x12, 0x3b, 0x0a, 0x04, 0x41, 0x74, 0x6f, 0x42, 0x12, 0x0a, 0x2e, 0x73,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x41, 0x41, 0x1a, 0x0a, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x42, 0x42, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22,
	0x10, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x65, 0x63, 0x68,
	0x6f, 0x12, 0x2a, 0x0a, 0x0a, 0x41, 0x74, 0x6f, 0x42, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12,
	0x0a, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x41, 0x41, 0x1a, 0x0a, 0x2e, 0x73, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x42, 0x42, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x8e, 0x01,
	0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x42, 0x0b, 0x53, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x38, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6e, 0x74, 0x6b, 0x73, 0x2f, 0x73, 0x61,
	0x6e, 0x64, 0x62, 0x6f, 0x78, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x61,
	0x70, 0x70, 0x2f, 0x76, 0x31, 0xf8, 0x01, 0x00, 0xa2, 0x02, 0x03, 0x53, 0x58, 0x58, 0xaa, 0x02,
	0x06, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0xca, 0x02, 0x06, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0xe2, 0x02, 0x12, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x06, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_sample_proto_rawDescOnce sync.Once
	file_app_sample_proto_rawDescData = file_app_sample_proto_rawDesc
)

func file_app_sample_proto_rawDescGZIP() []byte {
	file_app_sample_proto_rawDescOnce.Do(func() {
		file_app_sample_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_sample_proto_rawDescData)
	})
	return file_app_sample_proto_rawDescData
}

var file_app_sample_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_app_sample_proto_goTypes = []interface{}{
	(*AA)(nil), // 0: sample.AA
	(*BB)(nil), // 1: sample.BB
}
var file_app_sample_proto_depIdxs = []int32{
	0, // 0: sample.Hoge.AtoB:input_type -> sample.AA
	0, // 1: sample.Hoge.AtoBstream:input_type -> sample.AA
	1, // 2: sample.Hoge.AtoB:output_type -> sample.BB
	1, // 3: sample.Hoge.AtoBstream:output_type -> sample.BB
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_app_sample_proto_init() }
func file_app_sample_proto_init() {
	if File_app_sample_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_sample_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AA); i {
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
		file_app_sample_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BB); i {
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
			RawDescriptor: file_app_sample_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_sample_proto_goTypes,
		DependencyIndexes: file_app_sample_proto_depIdxs,
		MessageInfos:      file_app_sample_proto_msgTypes,
	}.Build()
	File_app_sample_proto = out.File
	file_app_sample_proto_rawDesc = nil
	file_app_sample_proto_goTypes = nil
	file_app_sample_proto_depIdxs = nil
}
