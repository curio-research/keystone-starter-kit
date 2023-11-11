// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.9.2
// source: base.proto

package pb_base

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

type CMD int32

const (
	CMD_S2C_Error CMD = 0
)

// Enum value maps for CMD.
var (
	CMD_name = map[int32]string{
		0: "S2C_Error",
	}
	CMD_value = map[string]int32{
		"S2C_Error": 0,
	}
)

func (x CMD) Enum() *CMD {
	p := new(CMD)
	*p = x
	return p
}

func (x CMD) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CMD) Descriptor() protoreflect.EnumDescriptor {
	return file_base_proto_enumTypes[0].Descriptor()
}

func (CMD) Type() protoreflect.EnumType {
	return &file_base_proto_enumTypes[0]
}

func (x CMD) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CMD.Descriptor instead.
func (CMD) EnumDescriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{0}
}

type S2C_ErrorMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=Content,proto3" json:"Content,omitempty"`
}

func (x *S2C_ErrorMessage) Reset() {
	*x = S2C_ErrorMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_ErrorMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_ErrorMessage) ProtoMessage() {}

func (x *S2C_ErrorMessage) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_ErrorMessage.ProtoReflect.Descriptor instead.
func (*S2C_ErrorMessage) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{0}
}

func (x *S2C_ErrorMessage) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_base_proto protoreflect.FileDescriptor

var file_base_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x70, 0x62, 0x22, 0x2c, 0x0a, 0x10, 0x53, 0x32, 0x43, 0x5f, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x2a, 0x14, 0x0a, 0x03, 0x43, 0x4d, 0x44, 0x12, 0x0d, 0x0a, 0x09, 0x53,
	0x32, 0x43, 0x5f, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x70, 0x62,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_base_proto_rawDescOnce sync.Once
	file_base_proto_rawDescData = file_base_proto_rawDesc
)

func file_base_proto_rawDescGZIP() []byte {
	file_base_proto_rawDescOnce.Do(func() {
		file_base_proto_rawDescData = protoimpl.X.CompressGZIP(file_base_proto_rawDescData)
	})
	return file_base_proto_rawDescData
}

var file_base_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_base_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_base_proto_goTypes = []interface{}{
	(CMD)(0),                 // 0: serverpb.CMD
	(*S2C_ErrorMessage)(nil), // 1: serverpb.S2C_ErrorMessage
}
var file_base_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_base_proto_init() }
func file_base_proto_init() {
	if File_base_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_base_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_ErrorMessage); i {
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
			RawDescriptor: file_base_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_base_proto_goTypes,
		DependencyIndexes: file_base_proto_depIdxs,
		EnumInfos:         file_base_proto_enumTypes,
		MessageInfos:      file_base_proto_msgTypes,
	}.Build()
	File_base_proto = out.File
	file_base_proto_rawDesc = nil
	file_base_proto_goTypes = nil
	file_base_proto_depIdxs = nil
}
