// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.26.1
// source: ADTimeActivity.proto

package pb

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

type ADTimeTriggerReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ADTimeTriggerReq) Reset() {
	*x = ADTimeTriggerReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ADTimeActivity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ADTimeTriggerReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ADTimeTriggerReq) ProtoMessage() {}

func (x *ADTimeTriggerReq) ProtoReflect() protoreflect.Message {
	mi := &file_ADTimeActivity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ADTimeTriggerReq.ProtoReflect.Descriptor instead.
func (*ADTimeTriggerReq) Descriptor() ([]byte, []int) {
	return file_ADTimeActivity_proto_rawDescGZIP(), []int{0}
}

type ADTimeUseReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ADTimeUseReq) Reset() {
	*x = ADTimeUseReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ADTimeActivity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ADTimeUseReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ADTimeUseReq) ProtoMessage() {}

func (x *ADTimeUseReq) ProtoReflect() protoreflect.Message {
	mi := &file_ADTimeActivity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ADTimeUseReq.ProtoReflect.Descriptor instead.
func (*ADTimeUseReq) Descriptor() ([]byte, []int) {
	return file_ADTimeActivity_proto_rawDescGZIP(), []int{1}
}

var File_ADTimeActivity_proto protoreflect.FileDescriptor

var file_ADTimeActivity_proto_rawDesc = []byte{
	0x0a, 0x14, 0x41, 0x44, 0x54, 0x69, 0x6d, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x79, 0x71, 0x2e, 0x6d,
	0x73, 0x67, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x4d, 0x73, 0x67, 0x22, 0x12, 0x0a, 0x10, 0x41, 0x44,
	0x54, 0x69, 0x6d, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x22, 0x0e,
	0x0a, 0x0c, 0x41, 0x44, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x73, 0x65, 0x52, 0x65, 0x71, 0x42, 0x36,
	0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6d, 0x7a,
	0x68, 0x6f, 0x6e, 0x67, 0x71, 0x69, 0x2f, 0x78, 0x64, 0x64, 0x71, 0x2d, 0x67, 0x6f, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62,
}

var (
	file_ADTimeActivity_proto_rawDescOnce sync.Once
	file_ADTimeActivity_proto_rawDescData = file_ADTimeActivity_proto_rawDesc
)

func file_ADTimeActivity_proto_rawDescGZIP() []byte {
	file_ADTimeActivity_proto_rawDescOnce.Do(func() {
		file_ADTimeActivity_proto_rawDescData = protoimpl.X.CompressGZIP(file_ADTimeActivity_proto_rawDescData)
	})
	return file_ADTimeActivity_proto_rawDescData
}

var file_ADTimeActivity_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ADTimeActivity_proto_goTypes = []any{
	(*ADTimeTriggerReq)(nil), // 0: com.yq.msg.CityMsg.ADTimeTriggerReq
	(*ADTimeUseReq)(nil),     // 1: com.yq.msg.CityMsg.ADTimeUseReq
}
var file_ADTimeActivity_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ADTimeActivity_proto_init() }
func file_ADTimeActivity_proto_init() {
	if File_ADTimeActivity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ADTimeActivity_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ADTimeTriggerReq); i {
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
		file_ADTimeActivity_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ADTimeUseReq); i {
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
			RawDescriptor: file_ADTimeActivity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ADTimeActivity_proto_goTypes,
		DependencyIndexes: file_ADTimeActivity_proto_depIdxs,
		MessageInfos:      file_ADTimeActivity_proto_msgTypes,
	}.Build()
	File_ADTimeActivity_proto = out.File
	file_ADTimeActivity_proto_rawDesc = nil
	file_ADTimeActivity_proto_goTypes = nil
	file_ADTimeActivity_proto_depIdxs = nil
}
