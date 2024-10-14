// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.26.1
// source: Invade.proto

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

type InvadeDataMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurInvadeId *int32 `protobuf:"varint,1,req,name=curInvadeId" json:"curInvadeId,omitempty"`
	Count       *int32 `protobuf:"varint,3,req,name=count" json:"count,omitempty"`
}

func (x *InvadeDataMsg) Reset() {
	*x = InvadeDataMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Invade_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvadeDataMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvadeDataMsg) ProtoMessage() {}

func (x *InvadeDataMsg) ProtoReflect() protoreflect.Message {
	mi := &file_Invade_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvadeDataMsg.ProtoReflect.Descriptor instead.
func (*InvadeDataMsg) Descriptor() ([]byte, []int) {
	return file_Invade_proto_rawDescGZIP(), []int{0}
}

func (x *InvadeDataMsg) GetCurInvadeId() int32 {
	if x != nil && x.CurInvadeId != nil {
		return *x.CurInvadeId
	}
	return 0
}

func (x *InvadeDataMsg) GetCount() int32 {
	if x != nil && x.Count != nil {
		return *x.Count
	}
	return 0
}

type InvadeChallengeResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret             *int32           `protobuf:"varint,1,req,name=ret" json:"ret,omitempty"`
	AllBattleRecord *BattleRecordMsg `protobuf:"bytes,2,opt,name=allBattleRecord" json:"allBattleRecord,omitempty"`
	Damage          *int64           `protobuf:"varint,3,opt,name=damage" json:"damage,omitempty"`
	Rewards         *string          `protobuf:"bytes,4,opt,name=rewards" json:"rewards,omitempty"`
	InvadeDataSync  *InvadeDataMsg   `protobuf:"bytes,5,opt,name=invadeDataSync" json:"invadeDataSync,omitempty"`
}

func (x *InvadeChallengeResp) Reset() {
	*x = InvadeChallengeResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Invade_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvadeChallengeResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvadeChallengeResp) ProtoMessage() {}

func (x *InvadeChallengeResp) ProtoReflect() protoreflect.Message {
	mi := &file_Invade_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvadeChallengeResp.ProtoReflect.Descriptor instead.
func (*InvadeChallengeResp) Descriptor() ([]byte, []int) {
	return file_Invade_proto_rawDescGZIP(), []int{1}
}

func (x *InvadeChallengeResp) GetRet() int32 {
	if x != nil && x.Ret != nil {
		return *x.Ret
	}
	return 0
}

func (x *InvadeChallengeResp) GetAllBattleRecord() *BattleRecordMsg {
	if x != nil {
		return x.AllBattleRecord
	}
	return nil
}

func (x *InvadeChallengeResp) GetDamage() int64 {
	if x != nil && x.Damage != nil {
		return *x.Damage
	}
	return 0
}

func (x *InvadeChallengeResp) GetRewards() string {
	if x != nil && x.Rewards != nil {
		return *x.Rewards
	}
	return ""
}

func (x *InvadeChallengeResp) GetInvadeDataSync() *InvadeDataMsg {
	if x != nil {
		return x.InvadeDataSync
	}
	return nil
}

var File_Invade_proto protoreflect.FileDescriptor

var file_Invade_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x49, 0x6e, 0x76, 0x61, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12,
	0x63, 0x6f, 0x6d, 0x2e, 0x79, 0x71, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x4d,
	0x73, 0x67, 0x1a, 0x0c, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x47, 0x0a, 0x0d, 0x49, 0x6e, 0x76, 0x61, 0x64, 0x65, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x73,
	0x67, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x75, 0x72, 0x49, 0x6e, 0x76, 0x61, 0x64, 0x65, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x49, 0x6e, 0x76, 0x61, 0x64,
	0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x02,
	0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xf3, 0x01, 0x0a, 0x13, 0x49, 0x6e,
	0x76, 0x61, 0x64, 0x65, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x03,
	0x72, 0x65, 0x74, 0x12, 0x4d, 0x0a, 0x0f, 0x61, 0x6c, 0x6c, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x79, 0x71, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x4d, 0x73,
	0x67, 0x2e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4d, 0x73,
	0x67, 0x52, 0x0f, 0x61, 0x6c, 0x6c, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x73, 0x12, 0x49, 0x0a, 0x0e, 0x69, 0x6e, 0x76, 0x61, 0x64, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x79, 0x6e, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x79, 0x71, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x4d, 0x73,
	0x67, 0x2e, 0x49, 0x6e, 0x76, 0x61, 0x64, 0x65, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x73, 0x67, 0x52,
	0x0e, 0x69, 0x6e, 0x76, 0x61, 0x64, 0x65, 0x44, 0x61, 0x74, 0x61, 0x53, 0x79, 0x6e, 0x63, 0x42,
	0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6d,
	0x7a, 0x68, 0x6f, 0x6e, 0x67, 0x71, 0x69, 0x2f, 0x78, 0x64, 0x64, 0x71, 0x2d, 0x67, 0x6f, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62,
}

var (
	file_Invade_proto_rawDescOnce sync.Once
	file_Invade_proto_rawDescData = file_Invade_proto_rawDesc
)

func file_Invade_proto_rawDescGZIP() []byte {
	file_Invade_proto_rawDescOnce.Do(func() {
		file_Invade_proto_rawDescData = protoimpl.X.CompressGZIP(file_Invade_proto_rawDescData)
	})
	return file_Invade_proto_rawDescData
}

var file_Invade_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Invade_proto_goTypes = []any{
	(*InvadeDataMsg)(nil),       // 0: com.yq.msg.CityMsg.InvadeDataMsg
	(*InvadeChallengeResp)(nil), // 1: com.yq.msg.CityMsg.InvadeChallengeResp
	(*BattleRecordMsg)(nil),     // 2: com.yq.msg.CityMsg.BattleRecordMsg
}
var file_Invade_proto_depIdxs = []int32{
	2, // 0: com.yq.msg.CityMsg.InvadeChallengeResp.allBattleRecord:type_name -> com.yq.msg.CityMsg.BattleRecordMsg
	0, // 1: com.yq.msg.CityMsg.InvadeChallengeResp.invadeDataSync:type_name -> com.yq.msg.CityMsg.InvadeDataMsg
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_Invade_proto_init() }
func file_Invade_proto_init() {
	if File_Invade_proto != nil {
		return
	}
	file_Battle_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Invade_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*InvadeDataMsg); i {
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
		file_Invade_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*InvadeChallengeResp); i {
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
			RawDescriptor: file_Invade_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Invade_proto_goTypes,
		DependencyIndexes: file_Invade_proto_depIdxs,
		MessageInfos:      file_Invade_proto_msgTypes,
	}.Build()
	File_Invade_proto = out.File
	file_Invade_proto_rawDesc = nil
	file_Invade_proto_goTypes = nil
	file_Invade_proto_depIdxs = nil
}
