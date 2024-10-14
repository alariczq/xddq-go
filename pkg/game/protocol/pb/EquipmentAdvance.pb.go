// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.26.1
// source: EquipmentAdvance.proto

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

type EquipmentAdvanceDataMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefineLvList     []int32 `protobuf:"varint,1,rep,name=refineLvList" json:"refineLvList,omitempty"`
	CurType          *int32  `protobuf:"varint,2,opt,name=curType" json:"curType,omitempty"`
	UnlockSkillRange *int32  `protobuf:"varint,3,opt,name=unlockSkillRange" json:"unlockSkillRange,omitempty"`
	AdvanceLv        *int32  `protobuf:"varint,4,opt,name=advanceLv" json:"advanceLv,omitempty"`
}

func (x *EquipmentAdvanceDataMsg) Reset() {
	*x = EquipmentAdvanceDataMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EquipmentAdvance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EquipmentAdvanceDataMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EquipmentAdvanceDataMsg) ProtoMessage() {}

func (x *EquipmentAdvanceDataMsg) ProtoReflect() protoreflect.Message {
	mi := &file_EquipmentAdvance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EquipmentAdvanceDataMsg.ProtoReflect.Descriptor instead.
func (*EquipmentAdvanceDataMsg) Descriptor() ([]byte, []int) {
	return file_EquipmentAdvance_proto_rawDescGZIP(), []int{0}
}

func (x *EquipmentAdvanceDataMsg) GetRefineLvList() []int32 {
	if x != nil {
		return x.RefineLvList
	}
	return nil
}

func (x *EquipmentAdvanceDataMsg) GetCurType() int32 {
	if x != nil && x.CurType != nil {
		return *x.CurType
	}
	return 0
}

func (x *EquipmentAdvanceDataMsg) GetUnlockSkillRange() int32 {
	if x != nil && x.UnlockSkillRange != nil {
		return *x.UnlockSkillRange
	}
	return 0
}

func (x *EquipmentAdvanceDataMsg) GetAdvanceLv() int32 {
	if x != nil && x.AdvanceLv != nil {
		return *x.AdvanceLv
	}
	return 0
}

type ReqEquipmentRefine struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsOneKey *bool `protobuf:"varint,1,opt,name=isOneKey" json:"isOneKey,omitempty"`
}

func (x *ReqEquipmentRefine) Reset() {
	*x = ReqEquipmentRefine{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EquipmentAdvance_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqEquipmentRefine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqEquipmentRefine) ProtoMessage() {}

func (x *ReqEquipmentRefine) ProtoReflect() protoreflect.Message {
	mi := &file_EquipmentAdvance_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqEquipmentRefine.ProtoReflect.Descriptor instead.
func (*ReqEquipmentRefine) Descriptor() ([]byte, []int) {
	return file_EquipmentAdvance_proto_rawDescGZIP(), []int{1}
}

func (x *ReqEquipmentRefine) GetIsOneKey() bool {
	if x != nil && x.IsOneKey != nil {
		return *x.IsOneKey
	}
	return false
}

type RespEquipmentRefine struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret          *int32  `protobuf:"varint,1,req,name=ret" json:"ret,omitempty"`
	CurType      *int32  `protobuf:"varint,2,opt,name=curType" json:"curType,omitempty"`
	RefineLvList []int32 `protobuf:"varint,3,rep,name=refineLvList" json:"refineLvList,omitempty"`
}

func (x *RespEquipmentRefine) Reset() {
	*x = RespEquipmentRefine{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EquipmentAdvance_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespEquipmentRefine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespEquipmentRefine) ProtoMessage() {}

func (x *RespEquipmentRefine) ProtoReflect() protoreflect.Message {
	mi := &file_EquipmentAdvance_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespEquipmentRefine.ProtoReflect.Descriptor instead.
func (*RespEquipmentRefine) Descriptor() ([]byte, []int) {
	return file_EquipmentAdvance_proto_rawDescGZIP(), []int{2}
}

func (x *RespEquipmentRefine) GetRet() int32 {
	if x != nil && x.Ret != nil {
		return *x.Ret
	}
	return 0
}

func (x *RespEquipmentRefine) GetCurType() int32 {
	if x != nil && x.CurType != nil {
		return *x.CurType
	}
	return 0
}

func (x *RespEquipmentRefine) GetRefineLvList() []int32 {
	if x != nil {
		return x.RefineLvList
	}
	return nil
}

type ReqEquipmentAdvance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReqEquipmentAdvance) Reset() {
	*x = ReqEquipmentAdvance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EquipmentAdvance_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqEquipmentAdvance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqEquipmentAdvance) ProtoMessage() {}

func (x *ReqEquipmentAdvance) ProtoReflect() protoreflect.Message {
	mi := &file_EquipmentAdvance_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqEquipmentAdvance.ProtoReflect.Descriptor instead.
func (*ReqEquipmentAdvance) Descriptor() ([]byte, []int) {
	return file_EquipmentAdvance_proto_rawDescGZIP(), []int{3}
}

type RespEquipmentAdvance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret       *int32 `protobuf:"varint,1,req,name=ret" json:"ret,omitempty"`
	CurType   *int32 `protobuf:"varint,2,opt,name=curType" json:"curType,omitempty"`
	AdvanceLv *int32 `protobuf:"varint,3,opt,name=advanceLv" json:"advanceLv,omitempty"`
}

func (x *RespEquipmentAdvance) Reset() {
	*x = RespEquipmentAdvance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EquipmentAdvance_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespEquipmentAdvance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespEquipmentAdvance) ProtoMessage() {}

func (x *RespEquipmentAdvance) ProtoReflect() protoreflect.Message {
	mi := &file_EquipmentAdvance_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespEquipmentAdvance.ProtoReflect.Descriptor instead.
func (*RespEquipmentAdvance) Descriptor() ([]byte, []int) {
	return file_EquipmentAdvance_proto_rawDescGZIP(), []int{4}
}

func (x *RespEquipmentAdvance) GetRet() int32 {
	if x != nil && x.Ret != nil {
		return *x.Ret
	}
	return 0
}

func (x *RespEquipmentAdvance) GetCurType() int32 {
	if x != nil && x.CurType != nil {
		return *x.CurType
	}
	return 0
}

func (x *RespEquipmentAdvance) GetAdvanceLv() int32 {
	if x != nil && x.AdvanceLv != nil {
		return *x.AdvanceLv
	}
	return 0
}

type ReqEquipmentActivate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReqEquipmentActivate) Reset() {
	*x = ReqEquipmentActivate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EquipmentAdvance_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqEquipmentActivate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqEquipmentActivate) ProtoMessage() {}

func (x *ReqEquipmentActivate) ProtoReflect() protoreflect.Message {
	mi := &file_EquipmentAdvance_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqEquipmentActivate.ProtoReflect.Descriptor instead.
func (*ReqEquipmentActivate) Descriptor() ([]byte, []int) {
	return file_EquipmentAdvance_proto_rawDescGZIP(), []int{5}
}

type RespEquipmentActivate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret              *int32 `protobuf:"varint,1,req,name=ret" json:"ret,omitempty"`
	UnlockSkillRange *int32 `protobuf:"varint,2,opt,name=unlockSkillRange" json:"unlockSkillRange,omitempty"`
}

func (x *RespEquipmentActivate) Reset() {
	*x = RespEquipmentActivate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EquipmentAdvance_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespEquipmentActivate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespEquipmentActivate) ProtoMessage() {}

func (x *RespEquipmentActivate) ProtoReflect() protoreflect.Message {
	mi := &file_EquipmentAdvance_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespEquipmentActivate.ProtoReflect.Descriptor instead.
func (*RespEquipmentActivate) Descriptor() ([]byte, []int) {
	return file_EquipmentAdvance_proto_rawDescGZIP(), []int{6}
}

func (x *RespEquipmentActivate) GetRet() int32 {
	if x != nil && x.Ret != nil {
		return *x.Ret
	}
	return 0
}

func (x *RespEquipmentActivate) GetUnlockSkillRange() int32 {
	if x != nil && x.UnlockSkillRange != nil {
		return *x.UnlockSkillRange
	}
	return 0
}

var File_EquipmentAdvance_proto protoreflect.FileDescriptor

var file_EquipmentAdvance_proto_rawDesc = []byte{
	0x0a, 0x16, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x76, 0x61, 0x6e,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x79, 0x71,
	0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x4d, 0x73, 0x67, 0x22, 0xa1, 0x01, 0x0a,
	0x17, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x73, 0x67, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x69,
	0x6e, 0x65, 0x4c, 0x76, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0c,
	0x72, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x4c, 0x76, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x75, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63,
	0x75, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b,
	0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x10, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x4c, 0x76, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x4c, 0x76,
	0x22, 0x30, 0x0a, 0x12, 0x52, 0x65, 0x71, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x4f, 0x6e, 0x65, 0x4b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x4f, 0x6e, 0x65, 0x4b,
	0x65, 0x79, 0x22, 0x65, 0x0a, 0x13, 0x52, 0x65, 0x73, 0x70, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x74,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x03, 0x72, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x75, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x75,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x4c,
	0x76, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0c, 0x72, 0x65, 0x66,
	0x69, 0x6e, 0x65, 0x4c, 0x76, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x52, 0x65, 0x71,
	0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65,
	0x22, 0x60, 0x0a, 0x14, 0x52, 0x65, 0x73, 0x70, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e,
	0x74, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x03, 0x72, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x75,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x75, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x4c,
	0x76, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65,
	0x4c, 0x76, 0x22, 0x16, 0x0a, 0x14, 0x52, 0x65, 0x71, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65,
	0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x22, 0x55, 0x0a, 0x15, 0x52, 0x65,
	0x73, 0x70, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05,
	0x52, 0x03, 0x72, 0x65, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x53,
	0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x10, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x69, 0x6d, 0x7a, 0x68, 0x6f, 0x6e, 0x67, 0x71, 0x69, 0x2f, 0x78, 0x64, 0x64, 0x71, 0x2d, 0x67,
	0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62,
}

var (
	file_EquipmentAdvance_proto_rawDescOnce sync.Once
	file_EquipmentAdvance_proto_rawDescData = file_EquipmentAdvance_proto_rawDesc
)

func file_EquipmentAdvance_proto_rawDescGZIP() []byte {
	file_EquipmentAdvance_proto_rawDescOnce.Do(func() {
		file_EquipmentAdvance_proto_rawDescData = protoimpl.X.CompressGZIP(file_EquipmentAdvance_proto_rawDescData)
	})
	return file_EquipmentAdvance_proto_rawDescData
}

var file_EquipmentAdvance_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_EquipmentAdvance_proto_goTypes = []any{
	(*EquipmentAdvanceDataMsg)(nil), // 0: com.yq.msg.CityMsg.EquipmentAdvanceDataMsg
	(*ReqEquipmentRefine)(nil),      // 1: com.yq.msg.CityMsg.ReqEquipmentRefine
	(*RespEquipmentRefine)(nil),     // 2: com.yq.msg.CityMsg.RespEquipmentRefine
	(*ReqEquipmentAdvance)(nil),     // 3: com.yq.msg.CityMsg.ReqEquipmentAdvance
	(*RespEquipmentAdvance)(nil),    // 4: com.yq.msg.CityMsg.RespEquipmentAdvance
	(*ReqEquipmentActivate)(nil),    // 5: com.yq.msg.CityMsg.ReqEquipmentActivate
	(*RespEquipmentActivate)(nil),   // 6: com.yq.msg.CityMsg.RespEquipmentActivate
}
var file_EquipmentAdvance_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_EquipmentAdvance_proto_init() }
func file_EquipmentAdvance_proto_init() {
	if File_EquipmentAdvance_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_EquipmentAdvance_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*EquipmentAdvanceDataMsg); i {
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
		file_EquipmentAdvance_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ReqEquipmentRefine); i {
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
		file_EquipmentAdvance_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*RespEquipmentRefine); i {
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
		file_EquipmentAdvance_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ReqEquipmentAdvance); i {
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
		file_EquipmentAdvance_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*RespEquipmentAdvance); i {
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
		file_EquipmentAdvance_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ReqEquipmentActivate); i {
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
		file_EquipmentAdvance_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*RespEquipmentActivate); i {
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
			RawDescriptor: file_EquipmentAdvance_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EquipmentAdvance_proto_goTypes,
		DependencyIndexes: file_EquipmentAdvance_proto_depIdxs,
		MessageInfos:      file_EquipmentAdvance_proto_msgTypes,
	}.Build()
	File_EquipmentAdvance_proto = out.File
	file_EquipmentAdvance_proto_rawDesc = nil
	file_EquipmentAdvance_proto_goTypes = nil
	file_EquipmentAdvance_proto_depIdxs = nil
}
