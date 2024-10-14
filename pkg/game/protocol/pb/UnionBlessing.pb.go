// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.26.1
// source: UnionBlessing.proto

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

type UnionBlessingProtocol int32

const (
	UnionBlessingProtocol_UnionBlessing_NULL              UnionBlessingProtocol = 0
	UnionBlessingProtocol_UnionBlessing_sendGift          UnionBlessingProtocol = 210601
	UnionBlessingProtocol_UnionBlessing_receiveGift       UnionBlessingProtocol = 210602
	UnionBlessingProtocol_UnionBlessing_syncGift          UnionBlessingProtocol = 210603
	UnionBlessingProtocol_UnionBlessing_syncBlessingCount UnionBlessingProtocol = 210604
)

// Enum value maps for UnionBlessingProtocol.
var (
	UnionBlessingProtocol_name = map[int32]string{
		0:      "UnionBlessing_NULL",
		210601: "UnionBlessing_sendGift",
		210602: "UnionBlessing_receiveGift",
		210603: "UnionBlessing_syncGift",
		210604: "UnionBlessing_syncBlessingCount",
	}
	UnionBlessingProtocol_value = map[string]int32{
		"UnionBlessing_NULL":              0,
		"UnionBlessing_sendGift":          210601,
		"UnionBlessing_receiveGift":       210602,
		"UnionBlessing_syncGift":          210603,
		"UnionBlessing_syncBlessingCount": 210604,
	}
)

func (x UnionBlessingProtocol) Enum() *UnionBlessingProtocol {
	p := new(UnionBlessingProtocol)
	*p = x
	return p
}

func (x UnionBlessingProtocol) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UnionBlessingProtocol) Descriptor() protoreflect.EnumDescriptor {
	return file_UnionBlessing_proto_enumTypes[0].Descriptor()
}

func (UnionBlessingProtocol) Type() protoreflect.EnumType {
	return &file_UnionBlessing_proto_enumTypes[0]
}

func (x UnionBlessingProtocol) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *UnionBlessingProtocol) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = UnionBlessingProtocol(num)
	return nil
}

// Deprecated: Use UnionBlessingProtocol.Descriptor instead.
func (UnionBlessingProtocol) EnumDescriptor() ([]byte, []int) {
	return file_UnionBlessing_proto_rawDescGZIP(), []int{0}
}

type UnionBlessingSendGiftReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FlagId *int32  `protobuf:"varint,1,req,name=flagId" json:"flagId,omitempty"`
	Words  *string `protobuf:"bytes,2,req,name=words" json:"words,omitempty"`
}

func (x *UnionBlessingSendGiftReq) Reset() {
	*x = UnionBlessingSendGiftReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UnionBlessing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnionBlessingSendGiftReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnionBlessingSendGiftReq) ProtoMessage() {}

func (x *UnionBlessingSendGiftReq) ProtoReflect() protoreflect.Message {
	mi := &file_UnionBlessing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnionBlessingSendGiftReq.ProtoReflect.Descriptor instead.
func (*UnionBlessingSendGiftReq) Descriptor() ([]byte, []int) {
	return file_UnionBlessing_proto_rawDescGZIP(), []int{0}
}

func (x *UnionBlessingSendGiftReq) GetFlagId() int32 {
	if x != nil && x.FlagId != nil {
		return *x.FlagId
	}
	return 0
}

func (x *UnionBlessingSendGiftReq) GetWords() string {
	if x != nil && x.Words != nil {
		return *x.Words
	}
	return ""
}

type UnionBlessingSendGiftResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret *int32 `protobuf:"varint,1,req,name=ret" json:"ret,omitempty"`
}

func (x *UnionBlessingSendGiftResp) Reset() {
	*x = UnionBlessingSendGiftResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UnionBlessing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnionBlessingSendGiftResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnionBlessingSendGiftResp) ProtoMessage() {}

func (x *UnionBlessingSendGiftResp) ProtoReflect() protoreflect.Message {
	mi := &file_UnionBlessing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnionBlessingSendGiftResp.ProtoReflect.Descriptor instead.
func (*UnionBlessingSendGiftResp) Descriptor() ([]byte, []int) {
	return file_UnionBlessing_proto_rawDescGZIP(), []int{1}
}

func (x *UnionBlessingSendGiftResp) GetRet() int32 {
	if x != nil && x.Ret != nil {
		return *x.Ret
	}
	return 0
}

type UnionBlessingGiftSyncList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*UnionBlessingGiftSyncMsg `protobuf:"bytes,1,rep,name=data" json:"data,omitempty"`
}

func (x *UnionBlessingGiftSyncList) Reset() {
	*x = UnionBlessingGiftSyncList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UnionBlessing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnionBlessingGiftSyncList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnionBlessingGiftSyncList) ProtoMessage() {}

func (x *UnionBlessingGiftSyncList) ProtoReflect() protoreflect.Message {
	mi := &file_UnionBlessing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnionBlessingGiftSyncList.ProtoReflect.Descriptor instead.
func (*UnionBlessingGiftSyncList) Descriptor() ([]byte, []int) {
	return file_UnionBlessing_proto_rawDescGZIP(), []int{2}
}

func (x *UnionBlessingGiftSyncList) GetData() []*UnionBlessingGiftSyncMsg {
	if x != nil {
		return x.Data
	}
	return nil
}

type UnionBlessingGiftSyncMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          *string                       `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	ExpiredTime *int64                        `protobuf:"varint,2,req,name=expiredTime" json:"expiredTime,omitempty"`
	Words       *string                       `protobuf:"bytes,3,req,name=words" json:"words,omitempty"`
	ServerId    *int64                        `protobuf:"varint,4,req,name=serverId" json:"serverId,omitempty"`
	UnionId     *int64                        `protobuf:"varint,5,req,name=unionId" json:"unionId,omitempty"`
	UnionName   *string                       `protobuf:"bytes,6,req,name=unionName" json:"unionName,omitempty"`
	FlagId      *int32                        `protobuf:"varint,7,req,name=flagId" json:"flagId,omitempty"`
	Members     []*UnionBlessingPlayerShowMsg `protobuf:"bytes,8,rep,name=members" json:"members,omitempty"`
}

func (x *UnionBlessingGiftSyncMsg) Reset() {
	*x = UnionBlessingGiftSyncMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UnionBlessing_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnionBlessingGiftSyncMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnionBlessingGiftSyncMsg) ProtoMessage() {}

func (x *UnionBlessingGiftSyncMsg) ProtoReflect() protoreflect.Message {
	mi := &file_UnionBlessing_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnionBlessingGiftSyncMsg.ProtoReflect.Descriptor instead.
func (*UnionBlessingGiftSyncMsg) Descriptor() ([]byte, []int) {
	return file_UnionBlessing_proto_rawDescGZIP(), []int{3}
}

func (x *UnionBlessingGiftSyncMsg) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *UnionBlessingGiftSyncMsg) GetExpiredTime() int64 {
	if x != nil && x.ExpiredTime != nil {
		return *x.ExpiredTime
	}
	return 0
}

func (x *UnionBlessingGiftSyncMsg) GetWords() string {
	if x != nil && x.Words != nil {
		return *x.Words
	}
	return ""
}

func (x *UnionBlessingGiftSyncMsg) GetServerId() int64 {
	if x != nil && x.ServerId != nil {
		return *x.ServerId
	}
	return 0
}

func (x *UnionBlessingGiftSyncMsg) GetUnionId() int64 {
	if x != nil && x.UnionId != nil {
		return *x.UnionId
	}
	return 0
}

func (x *UnionBlessingGiftSyncMsg) GetUnionName() string {
	if x != nil && x.UnionName != nil {
		return *x.UnionName
	}
	return ""
}

func (x *UnionBlessingGiftSyncMsg) GetFlagId() int32 {
	if x != nil && x.FlagId != nil {
		return *x.FlagId
	}
	return 0
}

func (x *UnionBlessingGiftSyncMsg) GetMembers() []*UnionBlessingPlayerShowMsg {
	if x != nil {
		return x.Members
	}
	return nil
}

type UnionBlessingPlayerShowMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId     *int64  `protobuf:"varint,1,req,name=playerId" json:"playerId,omitempty"`
	NickName     *string `protobuf:"bytes,2,req,name=nickName" json:"nickName,omitempty"`
	AppearanceId *int32  `protobuf:"varint,3,opt,name=appearanceId" json:"appearanceId,omitempty"`
	EquipCloudId *int32  `protobuf:"varint,4,opt,name=equipCloudId" json:"equipCloudId,omitempty"`
	RealmsId     *int32  `protobuf:"varint,5,opt,name=realmsId" json:"realmsId,omitempty"`
	Position     *int32  `protobuf:"varint,6,opt,name=position" json:"position,omitempty"`
}

func (x *UnionBlessingPlayerShowMsg) Reset() {
	*x = UnionBlessingPlayerShowMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UnionBlessing_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnionBlessingPlayerShowMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnionBlessingPlayerShowMsg) ProtoMessage() {}

func (x *UnionBlessingPlayerShowMsg) ProtoReflect() protoreflect.Message {
	mi := &file_UnionBlessing_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnionBlessingPlayerShowMsg.ProtoReflect.Descriptor instead.
func (*UnionBlessingPlayerShowMsg) Descriptor() ([]byte, []int) {
	return file_UnionBlessing_proto_rawDescGZIP(), []int{4}
}

func (x *UnionBlessingPlayerShowMsg) GetPlayerId() int64 {
	if x != nil && x.PlayerId != nil {
		return *x.PlayerId
	}
	return 0
}

func (x *UnionBlessingPlayerShowMsg) GetNickName() string {
	if x != nil && x.NickName != nil {
		return *x.NickName
	}
	return ""
}

func (x *UnionBlessingPlayerShowMsg) GetAppearanceId() int32 {
	if x != nil && x.AppearanceId != nil {
		return *x.AppearanceId
	}
	return 0
}

func (x *UnionBlessingPlayerShowMsg) GetEquipCloudId() int32 {
	if x != nil && x.EquipCloudId != nil {
		return *x.EquipCloudId
	}
	return 0
}

func (x *UnionBlessingPlayerShowMsg) GetRealmsId() int32 {
	if x != nil && x.RealmsId != nil {
		return *x.RealmsId
	}
	return 0
}

func (x *UnionBlessingPlayerShowMsg) GetPosition() int32 {
	if x != nil && x.Position != nil {
		return *x.Position
	}
	return 0
}

type UnionBlessingRewardReqMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *string `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
}

func (x *UnionBlessingRewardReqMsg) Reset() {
	*x = UnionBlessingRewardReqMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UnionBlessing_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnionBlessingRewardReqMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnionBlessingRewardReqMsg) ProtoMessage() {}

func (x *UnionBlessingRewardReqMsg) ProtoReflect() protoreflect.Message {
	mi := &file_UnionBlessing_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnionBlessingRewardReqMsg.ProtoReflect.Descriptor instead.
func (*UnionBlessingRewardReqMsg) Descriptor() ([]byte, []int) {
	return file_UnionBlessing_proto_rawDescGZIP(), []int{5}
}

func (x *UnionBlessingRewardReqMsg) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

type UnionBlessingRewardRespMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret    *int32  `protobuf:"varint,1,req,name=ret" json:"ret,omitempty"`
	Reward *string `protobuf:"bytes,2,opt,name=reward" json:"reward,omitempty"`
	Id     *string `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
}

func (x *UnionBlessingRewardRespMsg) Reset() {
	*x = UnionBlessingRewardRespMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UnionBlessing_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnionBlessingRewardRespMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnionBlessingRewardRespMsg) ProtoMessage() {}

func (x *UnionBlessingRewardRespMsg) ProtoReflect() protoreflect.Message {
	mi := &file_UnionBlessing_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnionBlessingRewardRespMsg.ProtoReflect.Descriptor instead.
func (*UnionBlessingRewardRespMsg) Descriptor() ([]byte, []int) {
	return file_UnionBlessing_proto_rawDescGZIP(), []int{6}
}

func (x *UnionBlessingRewardRespMsg) GetRet() int32 {
	if x != nil && x.Ret != nil {
		return *x.Ret
	}
	return 0
}

func (x *UnionBlessingRewardRespMsg) GetReward() string {
	if x != nil && x.Reward != nil {
		return *x.Reward
	}
	return ""
}

func (x *UnionBlessingRewardRespMsg) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

type UnionBlessingCountSyncList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*UnionBlessingCountSyncMsg `protobuf:"bytes,1,rep,name=data" json:"data,omitempty"`
}

func (x *UnionBlessingCountSyncList) Reset() {
	*x = UnionBlessingCountSyncList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UnionBlessing_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnionBlessingCountSyncList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnionBlessingCountSyncList) ProtoMessage() {}

func (x *UnionBlessingCountSyncList) ProtoReflect() protoreflect.Message {
	mi := &file_UnionBlessing_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnionBlessingCountSyncList.ProtoReflect.Descriptor instead.
func (*UnionBlessingCountSyncList) Descriptor() ([]byte, []int) {
	return file_UnionBlessing_proto_rawDescGZIP(), []int{7}
}

func (x *UnionBlessingCountSyncList) GetData() []*UnionBlessingCountSyncMsg {
	if x != nil {
		return x.Data
	}
	return nil
}

type UnionBlessingCountSyncMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FlagId           *int32 `protobuf:"varint,1,opt,name=flagId" json:"flagId,omitempty"`
	LastBlessingTime *int64 `protobuf:"varint,2,opt,name=lastBlessingTime" json:"lastBlessingTime,omitempty"`
}

func (x *UnionBlessingCountSyncMsg) Reset() {
	*x = UnionBlessingCountSyncMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UnionBlessing_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnionBlessingCountSyncMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnionBlessingCountSyncMsg) ProtoMessage() {}

func (x *UnionBlessingCountSyncMsg) ProtoReflect() protoreflect.Message {
	mi := &file_UnionBlessing_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnionBlessingCountSyncMsg.ProtoReflect.Descriptor instead.
func (*UnionBlessingCountSyncMsg) Descriptor() ([]byte, []int) {
	return file_UnionBlessing_proto_rawDescGZIP(), []int{8}
}

func (x *UnionBlessingCountSyncMsg) GetFlagId() int32 {
	if x != nil && x.FlagId != nil {
		return *x.FlagId
	}
	return 0
}

func (x *UnionBlessingCountSyncMsg) GetLastBlessingTime() int64 {
	if x != nil && x.LastBlessingTime != nil {
		return *x.LastBlessingTime
	}
	return 0
}

var File_UnionBlessing_proto protoreflect.FileDescriptor

var file_UnionBlessing_proto_rawDesc = []byte{
	0x0a, 0x13, 0x55, 0x6e, 0x69, 0x6f, 0x6e, 0x42, 0x6c, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x79, 0x71, 0x2e, 0x6d, 0x73,
	0x67, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x4d, 0x73, 0x67, 0x22, 0x48, 0x0a, 0x18, 0x55, 0x6e, 0x69,
	0x6f, 0x6e, 0x42, 0x6c, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x6e, 0x64, 0x47, 0x69,
	0x66, 0x74, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6c, 0x61, 0x67, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x06, 0x66, 0x6c, 0x61, 0x67, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x05, 0x77, 0x6f,
	0x72, 0x64, 0x73, 0x22, 0x2d, 0x0a, 0x19, 0x55, 0x6e, 0x69, 0x6f, 0x6e, 0x42, 0x6c, 0x65, 0x73,
	0x73, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x6e, 0x64, 0x47, 0x69, 0x66, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x03, 0x72,
	0x65, 0x74, 0x22, 0x5d, 0x0a, 0x19, 0x55, 0x6e, 0x69, 0x6f, 0x6e, 0x42, 0x6c, 0x65, 0x73, 0x73,
	0x69, 0x6e, 0x67, 0x47, 0x69, 0x66, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x40, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x79, 0x71, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x4d,
	0x73, 0x67, 0x2e, 0x55, 0x6e, 0x69, 0x6f, 0x6e, 0x42, 0x6c, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67,
	0x47, 0x69, 0x66, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x4d, 0x73, 0x67, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x98, 0x02, 0x0a, 0x18, 0x55, 0x6e, 0x69, 0x6f, 0x6e, 0x42, 0x6c, 0x65, 0x73, 0x73,
	0x69, 0x6e, 0x67, 0x47, 0x69, 0x66, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x4d, 0x73, 0x67, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20,
	0x0a, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x02, 0x28, 0x03, 0x52, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x03, 0x20, 0x02, 0x28, 0x09, 0x52,
	0x05, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x04, 0x20, 0x02, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x05, 0x20,
	0x02, 0x28, 0x03, 0x52, 0x07, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x02, 0x28, 0x09, 0x52,
	0x09, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6c,
	0x61, 0x67, 0x49, 0x64, 0x18, 0x07, 0x20, 0x02, 0x28, 0x05, 0x52, 0x06, 0x66, 0x6c, 0x61, 0x67,
	0x49, 0x64, 0x12, 0x48, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x79, 0x71, 0x2e, 0x6d, 0x73, 0x67,
	0x2e, 0x43, 0x69, 0x74, 0x79, 0x4d, 0x73, 0x67, 0x2e, 0x55, 0x6e, 0x69, 0x6f, 0x6e, 0x42, 0x6c,
	0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x77,
	0x4d, 0x73, 0x67, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x22, 0xd4, 0x01, 0x0a,
	0x1a, 0x55, 0x6e, 0x69, 0x6f, 0x6e, 0x42, 0x6c, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x77, 0x4d, 0x73, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x03, 0x52, 0x08, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x70, 0x70, 0x65, 0x61, 0x72, 0x61, 0x6e, 0x63,
	0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x61, 0x70, 0x70, 0x65, 0x61,
	0x72, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x71, 0x75, 0x69, 0x70,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x65,
	0x71, 0x75, 0x69, 0x70, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72,
	0x65, 0x61, 0x6c, 0x6d, 0x73, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72,
	0x65, 0x61, 0x6c, 0x6d, 0x73, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x2b, 0x0a, 0x19, 0x55, 0x6e, 0x69, 0x6f, 0x6e, 0x42, 0x6c, 0x65, 0x73,
	0x73, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x4d, 0x73, 0x67,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x56, 0x0a, 0x1a, 0x55, 0x6e, 0x69, 0x6f, 0x6e, 0x42, 0x6c, 0x65, 0x73, 0x73, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x4d, 0x73, 0x67, 0x12, 0x10,
	0x0a, 0x03, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x03, 0x72, 0x65, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5f, 0x0a, 0x1a, 0x55, 0x6e, 0x69, 0x6f,
	0x6e, 0x42, 0x6c, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x79,
	0x6e, 0x63, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x79, 0x71, 0x2e, 0x6d, 0x73,
	0x67, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x4d, 0x73, 0x67, 0x2e, 0x55, 0x6e, 0x69, 0x6f, 0x6e, 0x42,
	0x6c, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x79, 0x6e, 0x63,
	0x4d, 0x73, 0x67, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x5f, 0x0a, 0x19, 0x55, 0x6e, 0x69,
	0x6f, 0x6e, 0x42, 0x6c, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x53,
	0x79, 0x6e, 0x63, 0x4d, 0x73, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6c, 0x61, 0x67, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x66, 0x6c, 0x61, 0x67, 0x49, 0x64, 0x12, 0x2a,
	0x0a, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x42, 0x6c, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x42, 0x6c,
	0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x2a, 0xb3, 0x01, 0x0a, 0x15, 0x55,
	0x6e, 0x69, 0x6f, 0x6e, 0x42, 0x6c, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x6e, 0x69, 0x6f, 0x6e, 0x42, 0x6c, 0x65,
	0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x4e, 0x55, 0x4c, 0x4c, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x16,
	0x55, 0x6e, 0x69, 0x6f, 0x6e, 0x42, 0x6c, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65,
	0x6e, 0x64, 0x47, 0x69, 0x66, 0x74, 0x10, 0xa9, 0xed, 0x0c, 0x12, 0x1f, 0x0a, 0x19, 0x55, 0x6e,
	0x69, 0x6f, 0x6e, 0x42, 0x6c, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x47, 0x69, 0x66, 0x74, 0x10, 0xaa, 0xed, 0x0c, 0x12, 0x1c, 0x0a, 0x16, 0x55,
	0x6e, 0x69, 0x6f, 0x6e, 0x42, 0x6c, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x79, 0x6e,
	0x63, 0x47, 0x69, 0x66, 0x74, 0x10, 0xab, 0xed, 0x0c, 0x12, 0x25, 0x0a, 0x1f, 0x55, 0x6e, 0x69,
	0x6f, 0x6e, 0x42, 0x6c, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x42,
	0x6c, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x10, 0xac, 0xed, 0x0c,
	0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69,
	0x6d, 0x7a, 0x68, 0x6f, 0x6e, 0x67, 0x71, 0x69, 0x2f, 0x78, 0x64, 0x64, 0x71, 0x2d, 0x67, 0x6f,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62,
}

var (
	file_UnionBlessing_proto_rawDescOnce sync.Once
	file_UnionBlessing_proto_rawDescData = file_UnionBlessing_proto_rawDesc
)

func file_UnionBlessing_proto_rawDescGZIP() []byte {
	file_UnionBlessing_proto_rawDescOnce.Do(func() {
		file_UnionBlessing_proto_rawDescData = protoimpl.X.CompressGZIP(file_UnionBlessing_proto_rawDescData)
	})
	return file_UnionBlessing_proto_rawDescData
}

var file_UnionBlessing_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_UnionBlessing_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_UnionBlessing_proto_goTypes = []any{
	(UnionBlessingProtocol)(0),         // 0: com.yq.msg.CityMsg.UnionBlessingProtocol
	(*UnionBlessingSendGiftReq)(nil),   // 1: com.yq.msg.CityMsg.UnionBlessingSendGiftReq
	(*UnionBlessingSendGiftResp)(nil),  // 2: com.yq.msg.CityMsg.UnionBlessingSendGiftResp
	(*UnionBlessingGiftSyncList)(nil),  // 3: com.yq.msg.CityMsg.UnionBlessingGiftSyncList
	(*UnionBlessingGiftSyncMsg)(nil),   // 4: com.yq.msg.CityMsg.UnionBlessingGiftSyncMsg
	(*UnionBlessingPlayerShowMsg)(nil), // 5: com.yq.msg.CityMsg.UnionBlessingPlayerShowMsg
	(*UnionBlessingRewardReqMsg)(nil),  // 6: com.yq.msg.CityMsg.UnionBlessingRewardReqMsg
	(*UnionBlessingRewardRespMsg)(nil), // 7: com.yq.msg.CityMsg.UnionBlessingRewardRespMsg
	(*UnionBlessingCountSyncList)(nil), // 8: com.yq.msg.CityMsg.UnionBlessingCountSyncList
	(*UnionBlessingCountSyncMsg)(nil),  // 9: com.yq.msg.CityMsg.UnionBlessingCountSyncMsg
}
var file_UnionBlessing_proto_depIdxs = []int32{
	4, // 0: com.yq.msg.CityMsg.UnionBlessingGiftSyncList.data:type_name -> com.yq.msg.CityMsg.UnionBlessingGiftSyncMsg
	5, // 1: com.yq.msg.CityMsg.UnionBlessingGiftSyncMsg.members:type_name -> com.yq.msg.CityMsg.UnionBlessingPlayerShowMsg
	9, // 2: com.yq.msg.CityMsg.UnionBlessingCountSyncList.data:type_name -> com.yq.msg.CityMsg.UnionBlessingCountSyncMsg
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_UnionBlessing_proto_init() }
func file_UnionBlessing_proto_init() {
	if File_UnionBlessing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_UnionBlessing_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*UnionBlessingSendGiftReq); i {
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
		file_UnionBlessing_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*UnionBlessingSendGiftResp); i {
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
		file_UnionBlessing_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*UnionBlessingGiftSyncList); i {
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
		file_UnionBlessing_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UnionBlessingGiftSyncMsg); i {
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
		file_UnionBlessing_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UnionBlessingPlayerShowMsg); i {
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
		file_UnionBlessing_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*UnionBlessingRewardReqMsg); i {
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
		file_UnionBlessing_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*UnionBlessingRewardRespMsg); i {
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
		file_UnionBlessing_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*UnionBlessingCountSyncList); i {
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
		file_UnionBlessing_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*UnionBlessingCountSyncMsg); i {
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
			RawDescriptor: file_UnionBlessing_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_UnionBlessing_proto_goTypes,
		DependencyIndexes: file_UnionBlessing_proto_depIdxs,
		EnumInfos:         file_UnionBlessing_proto_enumTypes,
		MessageInfos:      file_UnionBlessing_proto_msgTypes,
	}.Build()
	File_UnionBlessing_proto = out.File
	file_UnionBlessing_proto_rawDesc = nil
	file_UnionBlessing_proto_goTypes = nil
	file_UnionBlessing_proto_depIdxs = nil
}
