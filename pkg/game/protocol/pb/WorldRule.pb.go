// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.26.1
// source: WorldRule.proto

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

type WorldRulePerceptionLogMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProgrammeIndex  *int32                     `protobuf:"varint,1,opt,name=programmeIndex" json:"programmeIndex,omitempty"`
	HoleListMsg     []*WorldRuleHoleDetailsMsg `protobuf:"bytes,2,rep,name=holeListMsg" json:"holeListMsg,omitempty"`
	BaseAtt         *SkillMsg                  `protobuf:"bytes,3,opt,name=baseAtt" json:"baseAtt,omitempty"`
	PerceptionTimes *int32                     `protobuf:"varint,4,opt,name=perceptionTimes" json:"perceptionTimes,omitempty"`
	IsUpgrade       *bool                      `protobuf:"varint,5,opt,name=isUpgrade" json:"isUpgrade,omitempty"`
}

func (x *WorldRulePerceptionLogMsg) Reset() {
	*x = WorldRulePerceptionLogMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WorldRule_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorldRulePerceptionLogMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorldRulePerceptionLogMsg) ProtoMessage() {}

func (x *WorldRulePerceptionLogMsg) ProtoReflect() protoreflect.Message {
	mi := &file_WorldRule_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorldRulePerceptionLogMsg.ProtoReflect.Descriptor instead.
func (*WorldRulePerceptionLogMsg) Descriptor() ([]byte, []int) {
	return file_WorldRule_proto_rawDescGZIP(), []int{0}
}

func (x *WorldRulePerceptionLogMsg) GetProgrammeIndex() int32 {
	if x != nil && x.ProgrammeIndex != nil {
		return *x.ProgrammeIndex
	}
	return 0
}

func (x *WorldRulePerceptionLogMsg) GetHoleListMsg() []*WorldRuleHoleDetailsMsg {
	if x != nil {
		return x.HoleListMsg
	}
	return nil
}

func (x *WorldRulePerceptionLogMsg) GetBaseAtt() *SkillMsg {
	if x != nil {
		return x.BaseAtt
	}
	return nil
}

func (x *WorldRulePerceptionLogMsg) GetPerceptionTimes() int32 {
	if x != nil && x.PerceptionTimes != nil {
		return *x.PerceptionTimes
	}
	return 0
}

func (x *WorldRulePerceptionLogMsg) GetIsUpgrade() bool {
	if x != nil && x.IsUpgrade != nil {
		return *x.IsUpgrade
	}
	return false
}

type WorldRulePerceptionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RuleId            *int32  `protobuf:"varint,1,req,name=ruleId" json:"ruleId,omitempty"`
	LockHoleIndex     []int32 `protobuf:"varint,2,rep,name=lockHoleIndex" json:"lockHoleIndex,omitempty"`
	IsAdRealize       *bool   `protobuf:"varint,3,opt,name=isAdRealize" json:"isAdRealize,omitempty"`
	IsUseKillTimeItem *bool   `protobuf:"varint,4,opt,name=isUseKillTimeItem" json:"isUseKillTimeItem,omitempty"`
}

func (x *WorldRulePerceptionReq) Reset() {
	*x = WorldRulePerceptionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WorldRule_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorldRulePerceptionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorldRulePerceptionReq) ProtoMessage() {}

func (x *WorldRulePerceptionReq) ProtoReflect() protoreflect.Message {
	mi := &file_WorldRule_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorldRulePerceptionReq.ProtoReflect.Descriptor instead.
func (*WorldRulePerceptionReq) Descriptor() ([]byte, []int) {
	return file_WorldRule_proto_rawDescGZIP(), []int{1}
}

func (x *WorldRulePerceptionReq) GetRuleId() int32 {
	if x != nil && x.RuleId != nil {
		return *x.RuleId
	}
	return 0
}

func (x *WorldRulePerceptionReq) GetLockHoleIndex() []int32 {
	if x != nil {
		return x.LockHoleIndex
	}
	return nil
}

func (x *WorldRulePerceptionReq) GetIsAdRealize() bool {
	if x != nil && x.IsAdRealize != nil {
		return *x.IsAdRealize
	}
	return false
}

func (x *WorldRulePerceptionReq) GetIsUseKillTimeItem() bool {
	if x != nil && x.IsUseKillTimeItem != nil {
		return *x.IsUseKillTimeItem
	}
	return false
}

type WorldRulePerceptionResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret *int32 `protobuf:"varint,1,req,name=ret" json:"ret,omitempty"`
}

func (x *WorldRulePerceptionResp) Reset() {
	*x = WorldRulePerceptionResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WorldRule_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorldRulePerceptionResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorldRulePerceptionResp) ProtoMessage() {}

func (x *WorldRulePerceptionResp) ProtoReflect() protoreflect.Message {
	mi := &file_WorldRule_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorldRulePerceptionResp.ProtoReflect.Descriptor instead.
func (*WorldRulePerceptionResp) Descriptor() ([]byte, []int) {
	return file_WorldRule_proto_rawDescGZIP(), []int{2}
}

func (x *WorldRulePerceptionResp) GetRet() int32 {
	if x != nil && x.Ret != nil {
		return *x.Ret
	}
	return 0
}

type WorldRulePerceptionReplaceReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RuleId *int32 `protobuf:"varint,1,req,name=ruleId" json:"ruleId,omitempty"`
}

func (x *WorldRulePerceptionReplaceReq) Reset() {
	*x = WorldRulePerceptionReplaceReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WorldRule_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorldRulePerceptionReplaceReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorldRulePerceptionReplaceReq) ProtoMessage() {}

func (x *WorldRulePerceptionReplaceReq) ProtoReflect() protoreflect.Message {
	mi := &file_WorldRule_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorldRulePerceptionReplaceReq.ProtoReflect.Descriptor instead.
func (*WorldRulePerceptionReplaceReq) Descriptor() ([]byte, []int) {
	return file_WorldRule_proto_rawDescGZIP(), []int{3}
}

func (x *WorldRulePerceptionReplaceReq) GetRuleId() int32 {
	if x != nil && x.RuleId != nil {
		return *x.RuleId
	}
	return 0
}

type WorldRulePerceptionReplaceResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret *int32 `protobuf:"varint,1,req,name=ret" json:"ret,omitempty"`
}

func (x *WorldRulePerceptionReplaceResp) Reset() {
	*x = WorldRulePerceptionReplaceResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WorldRule_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorldRulePerceptionReplaceResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorldRulePerceptionReplaceResp) ProtoMessage() {}

func (x *WorldRulePerceptionReplaceResp) ProtoReflect() protoreflect.Message {
	mi := &file_WorldRule_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorldRulePerceptionReplaceResp.ProtoReflect.Descriptor instead.
func (*WorldRulePerceptionReplaceResp) Descriptor() ([]byte, []int) {
	return file_WorldRule_proto_rawDescGZIP(), []int{4}
}

func (x *WorldRulePerceptionReplaceResp) GetRet() int32 {
	if x != nil && x.Ret != nil {
		return *x.Ret
	}
	return 0
}

type WorldRuleGetPerceptionLogReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RuleId *int32 `protobuf:"varint,1,req,name=ruleId" json:"ruleId,omitempty"`
}

func (x *WorldRuleGetPerceptionLogReq) Reset() {
	*x = WorldRuleGetPerceptionLogReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WorldRule_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorldRuleGetPerceptionLogReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorldRuleGetPerceptionLogReq) ProtoMessage() {}

func (x *WorldRuleGetPerceptionLogReq) ProtoReflect() protoreflect.Message {
	mi := &file_WorldRule_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorldRuleGetPerceptionLogReq.ProtoReflect.Descriptor instead.
func (*WorldRuleGetPerceptionLogReq) Descriptor() ([]byte, []int) {
	return file_WorldRule_proto_rawDescGZIP(), []int{5}
}

func (x *WorldRuleGetPerceptionLogReq) GetRuleId() int32 {
	if x != nil && x.RuleId != nil {
		return *x.RuleId
	}
	return 0
}

type WorldRuleGetPerceptionLogResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret              *int32                       `protobuf:"varint,1,req,name=ret" json:"ret,omitempty"`
	RealizeRecordMsg []*WorldRulePerceptionLogMsg `protobuf:"bytes,2,rep,name=realizeRecordMsg" json:"realizeRecordMsg,omitempty"`
}

func (x *WorldRuleGetPerceptionLogResp) Reset() {
	*x = WorldRuleGetPerceptionLogResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WorldRule_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorldRuleGetPerceptionLogResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorldRuleGetPerceptionLogResp) ProtoMessage() {}

func (x *WorldRuleGetPerceptionLogResp) ProtoReflect() protoreflect.Message {
	mi := &file_WorldRule_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorldRuleGetPerceptionLogResp.ProtoReflect.Descriptor instead.
func (*WorldRuleGetPerceptionLogResp) Descriptor() ([]byte, []int) {
	return file_WorldRule_proto_rawDescGZIP(), []int{6}
}

func (x *WorldRuleGetPerceptionLogResp) GetRet() int32 {
	if x != nil && x.Ret != nil {
		return *x.Ret
	}
	return 0
}

func (x *WorldRuleGetPerceptionLogResp) GetRealizeRecordMsg() []*WorldRulePerceptionLogMsg {
	if x != nil {
		return x.RealizeRecordMsg
	}
	return nil
}

type WorldRuleSwitchProgrammeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProgrammeIndex *int32 `protobuf:"varint,1,req,name=programmeIndex" json:"programmeIndex,omitempty"`
}

func (x *WorldRuleSwitchProgrammeReq) Reset() {
	*x = WorldRuleSwitchProgrammeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WorldRule_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorldRuleSwitchProgrammeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorldRuleSwitchProgrammeReq) ProtoMessage() {}

func (x *WorldRuleSwitchProgrammeReq) ProtoReflect() protoreflect.Message {
	mi := &file_WorldRule_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorldRuleSwitchProgrammeReq.ProtoReflect.Descriptor instead.
func (*WorldRuleSwitchProgrammeReq) Descriptor() ([]byte, []int) {
	return file_WorldRule_proto_rawDescGZIP(), []int{7}
}

func (x *WorldRuleSwitchProgrammeReq) GetProgrammeIndex() int32 {
	if x != nil && x.ProgrammeIndex != nil {
		return *x.ProgrammeIndex
	}
	return 0
}

type WorldRuleSwitchProgrammeResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret *int32 `protobuf:"varint,1,req,name=ret" json:"ret,omitempty"`
}

func (x *WorldRuleSwitchProgrammeResp) Reset() {
	*x = WorldRuleSwitchProgrammeResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WorldRule_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorldRuleSwitchProgrammeResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorldRuleSwitchProgrammeResp) ProtoMessage() {}

func (x *WorldRuleSwitchProgrammeResp) ProtoReflect() protoreflect.Message {
	mi := &file_WorldRule_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorldRuleSwitchProgrammeResp.ProtoReflect.Descriptor instead.
func (*WorldRuleSwitchProgrammeResp) Descriptor() ([]byte, []int) {
	return file_WorldRule_proto_rawDescGZIP(), []int{8}
}

func (x *WorldRuleSwitchProgrammeResp) GetRet() int32 {
	if x != nil && x.Ret != nil {
		return *x.Ret
	}
	return 0
}

var File_WorldRule_proto protoreflect.FileDescriptor

var file_WorldRule_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x79, 0x71, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x43, 0x69,
	0x74, 0x79, 0x4d, 0x73, 0x67, 0x1a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x92, 0x02, 0x0a, 0x19, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x75, 0x6c,
	0x65, 0x50, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x4d, 0x73,
	0x67, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x6d, 0x65, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x67, 0x72,
	0x61, 0x6d, 0x6d, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x4d, 0x0a, 0x0b, 0x68, 0x6f, 0x6c,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x79, 0x71, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x43, 0x69, 0x74, 0x79,
	0x4d, 0x73, 0x67, 0x2e, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x75, 0x6c, 0x65, 0x48, 0x6f, 0x6c,
	0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x4d, 0x73, 0x67, 0x52, 0x0b, 0x68, 0x6f, 0x6c,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x36, 0x0a, 0x07, 0x62, 0x61, 0x73, 0x65,
	0x41, 0x74, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x79, 0x71, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x4d, 0x73, 0x67, 0x2e, 0x53,
	0x6b, 0x69, 0x6c, 0x6c, 0x4d, 0x73, 0x67, 0x52, 0x07, 0x62, 0x61, 0x73, 0x65, 0x41, 0x74, 0x74,
	0x12, 0x28, 0x0a, 0x0f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x70, 0x65, 0x72, 0x63, 0x65,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73,
	0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69,
	0x73, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x22, 0xa6, 0x01, 0x0a, 0x16, 0x57, 0x6f, 0x72,
	0x6c, 0x64, 0x52, 0x75, 0x6c, 0x65, 0x50, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x02, 0x28, 0x05, 0x52, 0x06, 0x72, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x6c,
	0x6f, 0x63, 0x6b, 0x48, 0x6f, 0x6c, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x05, 0x52, 0x0d, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x6f, 0x6c, 0x65, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x73, 0x41, 0x64, 0x52, 0x65, 0x61, 0x6c, 0x69, 0x7a, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x41, 0x64, 0x52, 0x65, 0x61, 0x6c,
	0x69, 0x7a, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x69, 0x73, 0x55, 0x73, 0x65, 0x4b, 0x69, 0x6c, 0x6c,
	0x54, 0x69, 0x6d, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11,
	0x69, 0x73, 0x55, 0x73, 0x65, 0x4b, 0x69, 0x6c, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x49, 0x74, 0x65,
	0x6d, 0x22, 0x2b, 0x0a, 0x17, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x75, 0x6c, 0x65, 0x50, 0x65,
	0x72, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03,
	0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x03, 0x72, 0x65, 0x74, 0x22, 0x37,
	0x0a, 0x1d, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x75, 0x6c, 0x65, 0x50, 0x65, 0x72, 0x63, 0x65,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52,
	0x06, 0x72, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x32, 0x0a, 0x1e, 0x57, 0x6f, 0x72, 0x6c, 0x64,
	0x52, 0x75, 0x6c, 0x65, 0x50, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x74,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x03, 0x72, 0x65, 0x74, 0x22, 0x36, 0x0a, 0x1c, 0x57,
	0x6f, 0x72, 0x6c, 0x64, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x63, 0x65,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x75, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x06, 0x72, 0x75, 0x6c,
	0x65, 0x49, 0x64, 0x22, 0x8c, 0x01, 0x0a, 0x1d, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x75, 0x6c,
	0x65, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x02,
	0x28, 0x05, 0x52, 0x03, 0x72, 0x65, 0x74, 0x12, 0x59, 0x0a, 0x10, 0x72, 0x65, 0x61, 0x6c, 0x69,
	0x7a, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x79, 0x71, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x43,
	0x69, 0x74, 0x79, 0x4d, 0x73, 0x67, 0x2e, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x75, 0x6c, 0x65,
	0x50, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x4d, 0x73, 0x67,
	0x52, 0x10, 0x72, 0x65, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4d,
	0x73, 0x67, 0x22, 0x45, 0x0a, 0x1b, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x75, 0x6c, 0x65, 0x53,
	0x77, 0x69, 0x74, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x6d, 0x65, 0x52, 0x65,
	0x71, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x6d, 0x65, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x67, 0x72,
	0x61, 0x6d, 0x6d, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x30, 0x0a, 0x1c, 0x57, 0x6f, 0x72,
	0x6c, 0x64, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x67,
	0x72, 0x61, 0x6d, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x74,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x03, 0x72, 0x65, 0x74, 0x42, 0x36, 0x5a, 0x34, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6d, 0x7a, 0x68, 0x6f, 0x6e,
	0x67, 0x71, 0x69, 0x2f, 0x78, 0x64, 0x64, 0x71, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x67, 0x61, 0x6d, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x62,
	0x3b, 0x70, 0x62,
}

var (
	file_WorldRule_proto_rawDescOnce sync.Once
	file_WorldRule_proto_rawDescData = file_WorldRule_proto_rawDesc
)

func file_WorldRule_proto_rawDescGZIP() []byte {
	file_WorldRule_proto_rawDescOnce.Do(func() {
		file_WorldRule_proto_rawDescData = protoimpl.X.CompressGZIP(file_WorldRule_proto_rawDescData)
	})
	return file_WorldRule_proto_rawDescData
}

var file_WorldRule_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_WorldRule_proto_goTypes = []any{
	(*WorldRulePerceptionLogMsg)(nil),      // 0: com.yq.msg.CityMsg.WorldRulePerceptionLogMsg
	(*WorldRulePerceptionReq)(nil),         // 1: com.yq.msg.CityMsg.WorldRulePerceptionReq
	(*WorldRulePerceptionResp)(nil),        // 2: com.yq.msg.CityMsg.WorldRulePerceptionResp
	(*WorldRulePerceptionReplaceReq)(nil),  // 3: com.yq.msg.CityMsg.WorldRulePerceptionReplaceReq
	(*WorldRulePerceptionReplaceResp)(nil), // 4: com.yq.msg.CityMsg.WorldRulePerceptionReplaceResp
	(*WorldRuleGetPerceptionLogReq)(nil),   // 5: com.yq.msg.CityMsg.WorldRuleGetPerceptionLogReq
	(*WorldRuleGetPerceptionLogResp)(nil),  // 6: com.yq.msg.CityMsg.WorldRuleGetPerceptionLogResp
	(*WorldRuleSwitchProgrammeReq)(nil),    // 7: com.yq.msg.CityMsg.WorldRuleSwitchProgrammeReq
	(*WorldRuleSwitchProgrammeResp)(nil),   // 8: com.yq.msg.CityMsg.WorldRuleSwitchProgrammeResp
	(*WorldRuleHoleDetailsMsg)(nil),        // 9: com.yq.msg.CityMsg.WorldRuleHoleDetailsMsg
	(*SkillMsg)(nil),                       // 10: com.yq.msg.CityMsg.SkillMsg
}
var file_WorldRule_proto_depIdxs = []int32{
	9,  // 0: com.yq.msg.CityMsg.WorldRulePerceptionLogMsg.holeListMsg:type_name -> com.yq.msg.CityMsg.WorldRuleHoleDetailsMsg
	10, // 1: com.yq.msg.CityMsg.WorldRulePerceptionLogMsg.baseAtt:type_name -> com.yq.msg.CityMsg.SkillMsg
	0,  // 2: com.yq.msg.CityMsg.WorldRuleGetPerceptionLogResp.realizeRecordMsg:type_name -> com.yq.msg.CityMsg.WorldRulePerceptionLogMsg
	3,  // [3:3] is the sub-list for method output_type
	3,  // [3:3] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_WorldRule_proto_init() }
func file_WorldRule_proto_init() {
	if File_WorldRule_proto != nil {
		return
	}
	file_Common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_WorldRule_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*WorldRulePerceptionLogMsg); i {
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
		file_WorldRule_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*WorldRulePerceptionReq); i {
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
		file_WorldRule_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*WorldRulePerceptionResp); i {
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
		file_WorldRule_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*WorldRulePerceptionReplaceReq); i {
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
		file_WorldRule_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*WorldRulePerceptionReplaceResp); i {
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
		file_WorldRule_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*WorldRuleGetPerceptionLogReq); i {
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
		file_WorldRule_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*WorldRuleGetPerceptionLogResp); i {
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
		file_WorldRule_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*WorldRuleSwitchProgrammeReq); i {
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
		file_WorldRule_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*WorldRuleSwitchProgrammeResp); i {
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
			RawDescriptor: file_WorldRule_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_WorldRule_proto_goTypes,
		DependencyIndexes: file_WorldRule_proto_depIdxs,
		MessageInfos:      file_WorldRule_proto_msgTypes,
	}.Build()
	File_WorldRule_proto = out.File
	file_WorldRule_proto_rawDesc = nil
	file_WorldRule_proto_goTypes = nil
	file_WorldRule_proto_depIdxs = nil
}
