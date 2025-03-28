// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.26.1
// source: MiniGames.proto

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

type MiniGamesPb int32

const (
	MiniGamesPb_MiniGamesPb_ReceiveMiniGamesReward       MiniGamesPb = 200021
	MiniGamesPb_MiniGamesPb_StageMapChallenge            MiniGamesPb = 200022
	MiniGamesPb_MiniGamesPb_GetWestJourneyReward         MiniGamesPb = 200023
	MiniGamesPb_MiniGamesPb_Enter_WestJourney            MiniGamesPb = 200024
	MiniGamesPb_MiniGamesPb_Sync_WestJourney_Player_Data MiniGamesPb = 200025
)

// Enum value maps for MiniGamesPb.
var (
	MiniGamesPb_name = map[int32]string{
		200021: "MiniGamesPb_ReceiveMiniGamesReward",
		200022: "MiniGamesPb_StageMapChallenge",
		200023: "MiniGamesPb_GetWestJourneyReward",
		200024: "MiniGamesPb_Enter_WestJourney",
		200025: "MiniGamesPb_Sync_WestJourney_Player_Data",
	}
	MiniGamesPb_value = map[string]int32{
		"MiniGamesPb_ReceiveMiniGamesReward":       200021,
		"MiniGamesPb_StageMapChallenge":            200022,
		"MiniGamesPb_GetWestJourneyReward":         200023,
		"MiniGamesPb_Enter_WestJourney":            200024,
		"MiniGamesPb_Sync_WestJourney_Player_Data": 200025,
	}
)

func (x MiniGamesPb) Enum() *MiniGamesPb {
	p := new(MiniGamesPb)
	*p = x
	return p
}

func (x MiniGamesPb) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MiniGamesPb) Descriptor() protoreflect.EnumDescriptor {
	return file_MiniGames_proto_enumTypes[0].Descriptor()
}

func (MiniGamesPb) Type() protoreflect.EnumType {
	return &file_MiniGames_proto_enumTypes[0]
}

func (x MiniGamesPb) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *MiniGamesPb) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = MiniGamesPb(num)
	return nil
}

// Deprecated: Use MiniGamesPb.Descriptor instead.
func (MiniGamesPb) EnumDescriptor() ([]byte, []int) {
	return file_MiniGames_proto_rawDescGZIP(), []int{0}
}

type MiniGamesType int32

const (
	MiniGamesType_SaveThePig   MiniGamesType = 1
	MiniGamesType_RelicExplore MiniGamesType = 2
	MiniGamesType_Shoot        MiniGamesType = 3
	MiniGamesType_WestJourney  MiniGamesType = 4
)

// Enum value maps for MiniGamesType.
var (
	MiniGamesType_name = map[int32]string{
		1: "SaveThePig",
		2: "RelicExplore",
		3: "Shoot",
		4: "WestJourney",
	}
	MiniGamesType_value = map[string]int32{
		"SaveThePig":   1,
		"RelicExplore": 2,
		"Shoot":        3,
		"WestJourney":  4,
	}
)

func (x MiniGamesType) Enum() *MiniGamesType {
	p := new(MiniGamesType)
	*p = x
	return p
}

func (x MiniGamesType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MiniGamesType) Descriptor() protoreflect.EnumDescriptor {
	return file_MiniGames_proto_enumTypes[1].Descriptor()
}

func (MiniGamesType) Type() protoreflect.EnumType {
	return &file_MiniGames_proto_enumTypes[1]
}

func (x MiniGamesType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *MiniGamesType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = MiniGamesType(num)
	return nil
}

// Deprecated: Use MiniGamesType.Descriptor instead.
func (MiniGamesType) EnumDescriptor() ([]byte, []int) {
	return file_MiniGames_proto_rawDescGZIP(), []int{1}
}

type ReceiveMiniGamesRewardReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  *MiniGamesType `protobuf:"varint,1,req,name=type,enum=com.yq.msg.CityMsg.MiniGamesType" json:"type,omitempty"`
	Stage *int32         `protobuf:"varint,2,req,name=stage" json:"stage,omitempty"`
}

func (x *ReceiveMiniGamesRewardReq) Reset() {
	*x = ReceiveMiniGamesRewardReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MiniGames_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveMiniGamesRewardReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveMiniGamesRewardReq) ProtoMessage() {}

func (x *ReceiveMiniGamesRewardReq) ProtoReflect() protoreflect.Message {
	mi := &file_MiniGames_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveMiniGamesRewardReq.ProtoReflect.Descriptor instead.
func (*ReceiveMiniGamesRewardReq) Descriptor() ([]byte, []int) {
	return file_MiniGames_proto_rawDescGZIP(), []int{0}
}

func (x *ReceiveMiniGamesRewardReq) GetType() MiniGamesType {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return MiniGamesType_SaveThePig
}

func (x *ReceiveMiniGamesRewardReq) GetStage() int32 {
	if x != nil && x.Stage != nil {
		return *x.Stage
	}
	return 0
}

type ReceiveMiniGamesRewardResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret    *int32  `protobuf:"varint,1,req,name=ret" json:"ret,omitempty"`
	Reward *string `protobuf:"bytes,2,opt,name=reward" json:"reward,omitempty"`
}

func (x *ReceiveMiniGamesRewardResp) Reset() {
	*x = ReceiveMiniGamesRewardResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MiniGames_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveMiniGamesRewardResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveMiniGamesRewardResp) ProtoMessage() {}

func (x *ReceiveMiniGamesRewardResp) ProtoReflect() protoreflect.Message {
	mi := &file_MiniGames_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveMiniGamesRewardResp.ProtoReflect.Descriptor instead.
func (*ReceiveMiniGamesRewardResp) Descriptor() ([]byte, []int) {
	return file_MiniGames_proto_rawDescGZIP(), []int{1}
}

func (x *ReceiveMiniGamesRewardResp) GetRet() int32 {
	if x != nil && x.Ret != nil {
		return *x.Ret
	}
	return 0
}

func (x *ReceiveMiniGamesRewardResp) GetReward() string {
	if x != nil && x.Reward != nil {
		return *x.Reward
	}
	return ""
}

type StageMapChallengeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stage *int32 `protobuf:"varint,1,req,name=stage" json:"stage,omitempty"`
	Index *int32 `protobuf:"varint,2,req,name=index" json:"index,omitempty"`
}

func (x *StageMapChallengeReq) Reset() {
	*x = StageMapChallengeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MiniGames_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StageMapChallengeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StageMapChallengeReq) ProtoMessage() {}

func (x *StageMapChallengeReq) ProtoReflect() protoreflect.Message {
	mi := &file_MiniGames_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StageMapChallengeReq.ProtoReflect.Descriptor instead.
func (*StageMapChallengeReq) Descriptor() ([]byte, []int) {
	return file_MiniGames_proto_rawDescGZIP(), []int{2}
}

func (x *StageMapChallengeReq) GetStage() int32 {
	if x != nil && x.Stage != nil {
		return *x.Stage
	}
	return 0
}

func (x *StageMapChallengeReq) GetIndex() int32 {
	if x != nil && x.Index != nil {
		return *x.Index
	}
	return 0
}

type StageMapChallengeRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret              *int32           `protobuf:"varint,1,req,name=ret" json:"ret,omitempty"`
	ChallengeSuccess *bool            `protobuf:"varint,2,opt,name=challengeSuccess" json:"challengeSuccess,omitempty"`
	AllBattleRecord  *BattleRecordMsg `protobuf:"bytes,3,opt,name=allBattleRecord" json:"allBattleRecord,omitempty"`
}

func (x *StageMapChallengeRsp) Reset() {
	*x = StageMapChallengeRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MiniGames_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StageMapChallengeRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StageMapChallengeRsp) ProtoMessage() {}

func (x *StageMapChallengeRsp) ProtoReflect() protoreflect.Message {
	mi := &file_MiniGames_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StageMapChallengeRsp.ProtoReflect.Descriptor instead.
func (*StageMapChallengeRsp) Descriptor() ([]byte, []int) {
	return file_MiniGames_proto_rawDescGZIP(), []int{3}
}

func (x *StageMapChallengeRsp) GetRet() int32 {
	if x != nil && x.Ret != nil {
		return *x.Ret
	}
	return 0
}

func (x *StageMapChallengeRsp) GetChallengeSuccess() bool {
	if x != nil && x.ChallengeSuccess != nil {
		return *x.ChallengeSuccess
	}
	return false
}

func (x *StageMapChallengeRsp) GetAllBattleRecord() *BattleRecordMsg {
	if x != nil {
		return x.AllBattleRecord
	}
	return nil
}

type GetWestJourneyRewardReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stage *int32 `protobuf:"varint,1,req,name=stage" json:"stage,omitempty"`
}

func (x *GetWestJourneyRewardReq) Reset() {
	*x = GetWestJourneyRewardReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MiniGames_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWestJourneyRewardReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWestJourneyRewardReq) ProtoMessage() {}

func (x *GetWestJourneyRewardReq) ProtoReflect() protoreflect.Message {
	mi := &file_MiniGames_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWestJourneyRewardReq.ProtoReflect.Descriptor instead.
func (*GetWestJourneyRewardReq) Descriptor() ([]byte, []int) {
	return file_MiniGames_proto_rawDescGZIP(), []int{4}
}

func (x *GetWestJourneyRewardReq) GetStage() int32 {
	if x != nil && x.Stage != nil {
		return *x.Stage
	}
	return 0
}

type GetWestJourneyRewardResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret    *int32  `protobuf:"varint,1,req,name=ret" json:"ret,omitempty"`
	Reward *string `protobuf:"bytes,2,opt,name=reward" json:"reward,omitempty"`
}

func (x *GetWestJourneyRewardResp) Reset() {
	*x = GetWestJourneyRewardResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MiniGames_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWestJourneyRewardResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWestJourneyRewardResp) ProtoMessage() {}

func (x *GetWestJourneyRewardResp) ProtoReflect() protoreflect.Message {
	mi := &file_MiniGames_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWestJourneyRewardResp.ProtoReflect.Descriptor instead.
func (*GetWestJourneyRewardResp) Descriptor() ([]byte, []int) {
	return file_MiniGames_proto_rawDescGZIP(), []int{5}
}

func (x *GetWestJourneyRewardResp) GetRet() int32 {
	if x != nil && x.Ret != nil {
		return *x.Ret
	}
	return 0
}

func (x *GetWestJourneyRewardResp) GetReward() string {
	if x != nil && x.Reward != nil {
		return *x.Reward
	}
	return ""
}

type EnterWestJourneyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EnterWestJourneyReq) Reset() {
	*x = EnterWestJourneyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MiniGames_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnterWestJourneyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnterWestJourneyReq) ProtoMessage() {}

func (x *EnterWestJourneyReq) ProtoReflect() protoreflect.Message {
	mi := &file_MiniGames_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnterWestJourneyReq.ProtoReflect.Descriptor instead.
func (*EnterWestJourneyReq) Descriptor() ([]byte, []int) {
	return file_MiniGames_proto_rawDescGZIP(), []int{6}
}

type EnterWestJourneyResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret                 *int32 `protobuf:"varint,1,req,name=ret" json:"ret,omitempty"`
	WestJourneyMaxStage *int32 `protobuf:"varint,2,opt,name=westJourneyMaxStage" json:"westJourneyMaxStage,omitempty"`
}

func (x *EnterWestJourneyResp) Reset() {
	*x = EnterWestJourneyResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MiniGames_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnterWestJourneyResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnterWestJourneyResp) ProtoMessage() {}

func (x *EnterWestJourneyResp) ProtoReflect() protoreflect.Message {
	mi := &file_MiniGames_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnterWestJourneyResp.ProtoReflect.Descriptor instead.
func (*EnterWestJourneyResp) Descriptor() ([]byte, []int) {
	return file_MiniGames_proto_rawDescGZIP(), []int{7}
}

func (x *EnterWestJourneyResp) GetRet() int32 {
	if x != nil && x.Ret != nil {
		return *x.Ret
	}
	return 0
}

func (x *EnterWestJourneyResp) GetWestJourneyMaxStage() int32 {
	if x != nil && x.WestJourneyMaxStage != nil {
		return *x.WestJourneyMaxStage
	}
	return 0
}

type MiniGamesWestJourneyData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxStage *int32 `protobuf:"varint,1,opt,name=maxStage" json:"maxStage,omitempty"`
}

func (x *MiniGamesWestJourneyData) Reset() {
	*x = MiniGamesWestJourneyData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MiniGames_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MiniGamesWestJourneyData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MiniGamesWestJourneyData) ProtoMessage() {}

func (x *MiniGamesWestJourneyData) ProtoReflect() protoreflect.Message {
	mi := &file_MiniGames_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MiniGamesWestJourneyData.ProtoReflect.Descriptor instead.
func (*MiniGamesWestJourneyData) Descriptor() ([]byte, []int) {
	return file_MiniGames_proto_rawDescGZIP(), []int{8}
}

func (x *MiniGamesWestJourneyData) GetMaxStage() int32 {
	if x != nil && x.MaxStage != nil {
		return *x.MaxStage
	}
	return 0
}

var File_MiniGames_proto protoreflect.FileDescriptor

var file_MiniGames_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x4d, 0x69, 0x6e, 0x69, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x79, 0x71, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x43, 0x69,
	0x74, 0x79, 0x4d, 0x73, 0x67, 0x1a, 0x0c, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x68, 0x0a, 0x19, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x4d, 0x69,
	0x6e, 0x69, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71,
	0x12, 0x35, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0e, 0x32, 0x21,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x79, 0x71, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x43, 0x69, 0x74, 0x79,
	0x4d, 0x73, 0x67, 0x2e, 0x4d, 0x69, 0x6e, 0x69, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x02, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x22, 0x46, 0x0a,
	0x1a, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x4d, 0x69, 0x6e, 0x69, 0x47, 0x61, 0x6d, 0x65,
	0x73, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x72,
	0x65, 0x74, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x03, 0x72, 0x65, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x22, 0x42, 0x0a, 0x14, 0x53, 0x74, 0x61, 0x67, 0x65, 0x4d, 0x61,
	0x70, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x02,
	0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x22, 0xa3, 0x01, 0x0a, 0x14, 0x53, 0x74,
	0x61, 0x67, 0x65, 0x4d, 0x61, 0x70, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52,
	0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52,
	0x03, 0x72, 0x65, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67,
	0x65, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10,
	0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x12, 0x4d, 0x0a, 0x0f, 0x61, 0x6c, 0x6c, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x79, 0x71, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x4d, 0x73, 0x67, 0x2e, 0x42,
	0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4d, 0x73, 0x67, 0x52, 0x0f,
	0x61, 0x6c, 0x6c, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22,
	0x2f, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x57, 0x65, 0x73, 0x74, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x65,
	0x79, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65,
	0x22, 0x44, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x57, 0x65, 0x73, 0x74, 0x4a, 0x6f, 0x75, 0x72, 0x6e,
	0x65, 0x79, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03,
	0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x03, 0x72, 0x65, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x57,
	0x65, 0x73, 0x74, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x65, 0x79, 0x52, 0x65, 0x71, 0x22, 0x5a, 0x0a,
	0x14, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x57, 0x65, 0x73, 0x74, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x65,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x02,
	0x28, 0x05, 0x52, 0x03, 0x72, 0x65, 0x74, 0x12, 0x30, 0x0a, 0x13, 0x77, 0x65, 0x73, 0x74, 0x4a,
	0x6f, 0x75, 0x72, 0x6e, 0x65, 0x79, 0x4d, 0x61, 0x78, 0x53, 0x74, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x77, 0x65, 0x73, 0x74, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x65,
	0x79, 0x4d, 0x61, 0x78, 0x53, 0x74, 0x61, 0x67, 0x65, 0x22, 0x36, 0x0a, 0x18, 0x4d, 0x69, 0x6e,
	0x69, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x57, 0x65, 0x73, 0x74, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x65,
	0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x61, 0x78, 0x53, 0x74, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x53, 0x74, 0x61, 0x67,
	0x65, 0x2a, 0xd9, 0x01, 0x0a, 0x0b, 0x4d, 0x69, 0x6e, 0x69, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x50,
	0x62, 0x12, 0x28, 0x0a, 0x22, 0x4d, 0x69, 0x6e, 0x69, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x50, 0x62,
	0x5f, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x4d, 0x69, 0x6e, 0x69, 0x47, 0x61, 0x6d, 0x65,
	0x73, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x10, 0xd5, 0x9a, 0x0c, 0x12, 0x23, 0x0a, 0x1d, 0x4d,
	0x69, 0x6e, 0x69, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x50, 0x62, 0x5f, 0x53, 0x74, 0x61, 0x67, 0x65,
	0x4d, 0x61, 0x70, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x10, 0xd6, 0x9a, 0x0c,
	0x12, 0x26, 0x0a, 0x20, 0x4d, 0x69, 0x6e, 0x69, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x50, 0x62, 0x5f,
	0x47, 0x65, 0x74, 0x57, 0x65, 0x73, 0x74, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x65, 0x79, 0x52, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x10, 0xd7, 0x9a, 0x0c, 0x12, 0x23, 0x0a, 0x1d, 0x4d, 0x69, 0x6e, 0x69,
	0x47, 0x61, 0x6d, 0x65, 0x73, 0x50, 0x62, 0x5f, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x57, 0x65,
	0x73, 0x74, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x65, 0x79, 0x10, 0xd8, 0x9a, 0x0c, 0x12, 0x2e, 0x0a,
	0x28, 0x4d, 0x69, 0x6e, 0x69, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x50, 0x62, 0x5f, 0x53, 0x79, 0x6e,
	0x63, 0x5f, 0x57, 0x65, 0x73, 0x74, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x65, 0x79, 0x5f, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x5f, 0x44, 0x61, 0x74, 0x61, 0x10, 0xd9, 0x9a, 0x0c, 0x2a, 0x4d, 0x0a,
	0x0d, 0x4d, 0x69, 0x6e, 0x69, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e,
	0x0a, 0x0a, 0x53, 0x61, 0x76, 0x65, 0x54, 0x68, 0x65, 0x50, 0x69, 0x67, 0x10, 0x01, 0x12, 0x10,
	0x0a, 0x0c, 0x52, 0x65, 0x6c, 0x69, 0x63, 0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x10, 0x02,
	0x12, 0x09, 0x0a, 0x05, 0x53, 0x68, 0x6f, 0x6f, 0x74, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x57,
	0x65, 0x73, 0x74, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x65, 0x79, 0x10, 0x04, 0x42, 0x36, 0x5a, 0x34,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6d, 0x7a, 0x68, 0x6f,
	0x6e, 0x67, 0x71, 0x69, 0x2f, 0x78, 0x64, 0x64, 0x71, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70,
	0x62, 0x3b, 0x70, 0x62,
}

var (
	file_MiniGames_proto_rawDescOnce sync.Once
	file_MiniGames_proto_rawDescData = file_MiniGames_proto_rawDesc
)

func file_MiniGames_proto_rawDescGZIP() []byte {
	file_MiniGames_proto_rawDescOnce.Do(func() {
		file_MiniGames_proto_rawDescData = protoimpl.X.CompressGZIP(file_MiniGames_proto_rawDescData)
	})
	return file_MiniGames_proto_rawDescData
}

var file_MiniGames_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_MiniGames_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_MiniGames_proto_goTypes = []any{
	(MiniGamesPb)(0),                   // 0: com.yq.msg.CityMsg.MiniGamesPb
	(MiniGamesType)(0),                 // 1: com.yq.msg.CityMsg.MiniGamesType
	(*ReceiveMiniGamesRewardReq)(nil),  // 2: com.yq.msg.CityMsg.ReceiveMiniGamesRewardReq
	(*ReceiveMiniGamesRewardResp)(nil), // 3: com.yq.msg.CityMsg.ReceiveMiniGamesRewardResp
	(*StageMapChallengeReq)(nil),       // 4: com.yq.msg.CityMsg.StageMapChallengeReq
	(*StageMapChallengeRsp)(nil),       // 5: com.yq.msg.CityMsg.StageMapChallengeRsp
	(*GetWestJourneyRewardReq)(nil),    // 6: com.yq.msg.CityMsg.GetWestJourneyRewardReq
	(*GetWestJourneyRewardResp)(nil),   // 7: com.yq.msg.CityMsg.GetWestJourneyRewardResp
	(*EnterWestJourneyReq)(nil),        // 8: com.yq.msg.CityMsg.EnterWestJourneyReq
	(*EnterWestJourneyResp)(nil),       // 9: com.yq.msg.CityMsg.EnterWestJourneyResp
	(*MiniGamesWestJourneyData)(nil),   // 10: com.yq.msg.CityMsg.MiniGamesWestJourneyData
	(*BattleRecordMsg)(nil),            // 11: com.yq.msg.CityMsg.BattleRecordMsg
}
var file_MiniGames_proto_depIdxs = []int32{
	1,  // 0: com.yq.msg.CityMsg.ReceiveMiniGamesRewardReq.type:type_name -> com.yq.msg.CityMsg.MiniGamesType
	11, // 1: com.yq.msg.CityMsg.StageMapChallengeRsp.allBattleRecord:type_name -> com.yq.msg.CityMsg.BattleRecordMsg
	2,  // [2:2] is the sub-list for method output_type
	2,  // [2:2] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_MiniGames_proto_init() }
func file_MiniGames_proto_init() {
	if File_MiniGames_proto != nil {
		return
	}
	file_Battle_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_MiniGames_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ReceiveMiniGamesRewardReq); i {
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
		file_MiniGames_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ReceiveMiniGamesRewardResp); i {
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
		file_MiniGames_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*StageMapChallengeReq); i {
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
		file_MiniGames_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*StageMapChallengeRsp); i {
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
		file_MiniGames_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetWestJourneyRewardReq); i {
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
		file_MiniGames_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetWestJourneyRewardResp); i {
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
		file_MiniGames_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*EnterWestJourneyReq); i {
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
		file_MiniGames_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*EnterWestJourneyResp); i {
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
		file_MiniGames_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*MiniGamesWestJourneyData); i {
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
			RawDescriptor: file_MiniGames_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MiniGames_proto_goTypes,
		DependencyIndexes: file_MiniGames_proto_depIdxs,
		EnumInfos:         file_MiniGames_proto_enumTypes,
		MessageInfos:      file_MiniGames_proto_msgTypes,
	}.Build()
	File_MiniGames_proto = out.File
	file_MiniGames_proto_rawDesc = nil
	file_MiniGames_proto_goTypes = nil
	file_MiniGames_proto_depIdxs = nil
}
