// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.26.1
// source: Cloud.proto

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

type PlayerCloudDataMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EquipCloudId     *int32              `protobuf:"varint,1,opt,name=equipCloudId" json:"equipCloudId,omitempty"`
	CloudDataMsg     []*CloudDataMsg     `protobuf:"bytes,2,rep,name=cloudDataMsg" json:"cloudDataMsg,omitempty"`
	Stage            *int32              `protobuf:"varint,3,req,name=stage" json:"stage,omitempty"`
	Lv               *int32              `protobuf:"varint,4,req,name=lv" json:"lv,omitempty"`
	Exp              *int32              `protobuf:"varint,5,req,name=exp" json:"exp,omitempty"`
	CloudSkinDataMsg []*CloudSkinDataMsg `protobuf:"bytes,6,rep,name=cloudSkinDataMsg" json:"cloudSkinDataMsg,omitempty"`
	EquipCloudSkinId *int32              `protobuf:"varint,7,opt,name=equipCloudSkinId" json:"equipCloudSkinId,omitempty"`
}

func (x *PlayerCloudDataMsg) Reset() {
	*x = PlayerCloudDataMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Cloud_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerCloudDataMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerCloudDataMsg) ProtoMessage() {}

func (x *PlayerCloudDataMsg) ProtoReflect() protoreflect.Message {
	mi := &file_Cloud_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerCloudDataMsg.ProtoReflect.Descriptor instead.
func (*PlayerCloudDataMsg) Descriptor() ([]byte, []int) {
	return file_Cloud_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerCloudDataMsg) GetEquipCloudId() int32 {
	if x != nil && x.EquipCloudId != nil {
		return *x.EquipCloudId
	}
	return 0
}

func (x *PlayerCloudDataMsg) GetCloudDataMsg() []*CloudDataMsg {
	if x != nil {
		return x.CloudDataMsg
	}
	return nil
}

func (x *PlayerCloudDataMsg) GetStage() int32 {
	if x != nil && x.Stage != nil {
		return *x.Stage
	}
	return 0
}

func (x *PlayerCloudDataMsg) GetLv() int32 {
	if x != nil && x.Lv != nil {
		return *x.Lv
	}
	return 0
}

func (x *PlayerCloudDataMsg) GetExp() int32 {
	if x != nil && x.Exp != nil {
		return *x.Exp
	}
	return 0
}

func (x *PlayerCloudDataMsg) GetCloudSkinDataMsg() []*CloudSkinDataMsg {
	if x != nil {
		return x.CloudSkinDataMsg
	}
	return nil
}

func (x *PlayerCloudDataMsg) GetEquipCloudSkinId() int32 {
	if x != nil && x.EquipCloudSkinId != nil {
		return *x.EquipCloudSkinId
	}
	return 0
}

type CloudSkinDataMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CloudSkinId    *int32 `protobuf:"varint,1,req,name=cloudSkinId" json:"cloudSkinId,omitempty"`
	ExpirationTime *int64 `protobuf:"varint,2,opt,name=expirationTime" json:"expirationTime,omitempty"`
	Lv             *int32 `protobuf:"varint,3,opt,name=lv" json:"lv,omitempty"`
}

func (x *CloudSkinDataMsg) Reset() {
	*x = CloudSkinDataMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Cloud_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudSkinDataMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudSkinDataMsg) ProtoMessage() {}

func (x *CloudSkinDataMsg) ProtoReflect() protoreflect.Message {
	mi := &file_Cloud_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudSkinDataMsg.ProtoReflect.Descriptor instead.
func (*CloudSkinDataMsg) Descriptor() ([]byte, []int) {
	return file_Cloud_proto_rawDescGZIP(), []int{1}
}

func (x *CloudSkinDataMsg) GetCloudSkinId() int32 {
	if x != nil && x.CloudSkinId != nil {
		return *x.CloudSkinId
	}
	return 0
}

func (x *CloudSkinDataMsg) GetExpirationTime() int64 {
	if x != nil && x.ExpirationTime != nil {
		return *x.ExpirationTime
	}
	return 0
}

func (x *CloudSkinDataMsg) GetLv() int32 {
	if x != nil && x.Lv != nil {
		return *x.Lv
	}
	return 0
}

type CloudDataMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CloudId *int32 `protobuf:"varint,1,req,name=cloudId" json:"cloudId,omitempty"`
}

func (x *CloudDataMsg) Reset() {
	*x = CloudDataMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Cloud_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudDataMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudDataMsg) ProtoMessage() {}

func (x *CloudDataMsg) ProtoReflect() protoreflect.Message {
	mi := &file_Cloud_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudDataMsg.ProtoReflect.Descriptor instead.
func (*CloudDataMsg) Descriptor() ([]byte, []int) {
	return file_Cloud_proto_rawDescGZIP(), []int{2}
}

func (x *CloudDataMsg) GetCloudId() int32 {
	if x != nil && x.CloudId != nil {
		return *x.CloudId
	}
	return 0
}

type UnlockCloudReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CloudId *int32 `protobuf:"varint,1,req,name=cloudId" json:"cloudId,omitempty"`
}

func (x *UnlockCloudReq) Reset() {
	*x = UnlockCloudReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Cloud_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnlockCloudReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnlockCloudReq) ProtoMessage() {}

func (x *UnlockCloudReq) ProtoReflect() protoreflect.Message {
	mi := &file_Cloud_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnlockCloudReq.ProtoReflect.Descriptor instead.
func (*UnlockCloudReq) Descriptor() ([]byte, []int) {
	return file_Cloud_proto_rawDescGZIP(), []int{3}
}

func (x *UnlockCloudReq) GetCloudId() int32 {
	if x != nil && x.CloudId != nil {
		return *x.CloudId
	}
	return 0
}

type UnlockCloudResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret                *int32              `protobuf:"varint,1,req,name=ret" json:"ret,omitempty"`
	PlayerCloudDataMsg *PlayerCloudDataMsg `protobuf:"bytes,2,opt,name=playerCloudDataMsg" json:"playerCloudDataMsg,omitempty"`
}

func (x *UnlockCloudResp) Reset() {
	*x = UnlockCloudResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Cloud_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnlockCloudResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnlockCloudResp) ProtoMessage() {}

func (x *UnlockCloudResp) ProtoReflect() protoreflect.Message {
	mi := &file_Cloud_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnlockCloudResp.ProtoReflect.Descriptor instead.
func (*UnlockCloudResp) Descriptor() ([]byte, []int) {
	return file_Cloud_proto_rawDescGZIP(), []int{4}
}

func (x *UnlockCloudResp) GetRet() int32 {
	if x != nil && x.Ret != nil {
		return *x.Ret
	}
	return 0
}

func (x *UnlockCloudResp) GetPlayerCloudDataMsg() *PlayerCloudDataMsg {
	if x != nil {
		return x.PlayerCloudDataMsg
	}
	return nil
}

type CloudLvUpReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Quick *bool `protobuf:"varint,1,req,name=quick" json:"quick,omitempty"`
}

func (x *CloudLvUpReq) Reset() {
	*x = CloudLvUpReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Cloud_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudLvUpReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudLvUpReq) ProtoMessage() {}

func (x *CloudLvUpReq) ProtoReflect() protoreflect.Message {
	mi := &file_Cloud_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudLvUpReq.ProtoReflect.Descriptor instead.
func (*CloudLvUpReq) Descriptor() ([]byte, []int) {
	return file_Cloud_proto_rawDescGZIP(), []int{5}
}

func (x *CloudLvUpReq) GetQuick() bool {
	if x != nil && x.Quick != nil {
		return *x.Quick
	}
	return false
}

type CloudLvUpResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret                *int32              `protobuf:"varint,1,req,name=ret" json:"ret,omitempty"`
	AddExp             *int32              `protobuf:"varint,2,opt,name=addExp" json:"addExp,omitempty"`
	PlayerCloudDataMsg *PlayerCloudDataMsg `protobuf:"bytes,3,opt,name=playerCloudDataMsg" json:"playerCloudDataMsg,omitempty"`
}

func (x *CloudLvUpResp) Reset() {
	*x = CloudLvUpResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Cloud_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudLvUpResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudLvUpResp) ProtoMessage() {}

func (x *CloudLvUpResp) ProtoReflect() protoreflect.Message {
	mi := &file_Cloud_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudLvUpResp.ProtoReflect.Descriptor instead.
func (*CloudLvUpResp) Descriptor() ([]byte, []int) {
	return file_Cloud_proto_rawDescGZIP(), []int{6}
}

func (x *CloudLvUpResp) GetRet() int32 {
	if x != nil && x.Ret != nil {
		return *x.Ret
	}
	return 0
}

func (x *CloudLvUpResp) GetAddExp() int32 {
	if x != nil && x.AddExp != nil {
		return *x.AddExp
	}
	return 0
}

func (x *CloudLvUpResp) GetPlayerCloudDataMsg() *PlayerCloudDataMsg {
	if x != nil {
		return x.PlayerCloudDataMsg
	}
	return nil
}

type CloudStageUpReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CloudStageUpReq) Reset() {
	*x = CloudStageUpReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Cloud_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudStageUpReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudStageUpReq) ProtoMessage() {}

func (x *CloudStageUpReq) ProtoReflect() protoreflect.Message {
	mi := &file_Cloud_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudStageUpReq.ProtoReflect.Descriptor instead.
func (*CloudStageUpReq) Descriptor() ([]byte, []int) {
	return file_Cloud_proto_rawDescGZIP(), []int{7}
}

type CloudStageUpResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret                *int32              `protobuf:"varint,1,req,name=ret" json:"ret,omitempty"`
	PlayerCloudDataMsg *PlayerCloudDataMsg `protobuf:"bytes,2,opt,name=playerCloudDataMsg" json:"playerCloudDataMsg,omitempty"`
}

func (x *CloudStageUpResp) Reset() {
	*x = CloudStageUpResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Cloud_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudStageUpResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudStageUpResp) ProtoMessage() {}

func (x *CloudStageUpResp) ProtoReflect() protoreflect.Message {
	mi := &file_Cloud_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudStageUpResp.ProtoReflect.Descriptor instead.
func (*CloudStageUpResp) Descriptor() ([]byte, []int) {
	return file_Cloud_proto_rawDescGZIP(), []int{8}
}

func (x *CloudStageUpResp) GetRet() int32 {
	if x != nil && x.Ret != nil {
		return *x.Ret
	}
	return 0
}

func (x *CloudStageUpResp) GetPlayerCloudDataMsg() *PlayerCloudDataMsg {
	if x != nil {
		return x.PlayerCloudDataMsg
	}
	return nil
}

type CloudEquipReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CloudId *int32 `protobuf:"varint,1,req,name=cloudId" json:"cloudId,omitempty"`
}

func (x *CloudEquipReq) Reset() {
	*x = CloudEquipReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Cloud_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudEquipReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudEquipReq) ProtoMessage() {}

func (x *CloudEquipReq) ProtoReflect() protoreflect.Message {
	mi := &file_Cloud_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudEquipReq.ProtoReflect.Descriptor instead.
func (*CloudEquipReq) Descriptor() ([]byte, []int) {
	return file_Cloud_proto_rawDescGZIP(), []int{9}
}

func (x *CloudEquipReq) GetCloudId() int32 {
	if x != nil && x.CloudId != nil {
		return *x.CloudId
	}
	return 0
}

type CloudEquipResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret                *int32              `protobuf:"varint,1,req,name=ret" json:"ret,omitempty"`
	PlayerCloudDataMsg *PlayerCloudDataMsg `protobuf:"bytes,2,opt,name=playerCloudDataMsg" json:"playerCloudDataMsg,omitempty"`
}

func (x *CloudEquipResp) Reset() {
	*x = CloudEquipResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Cloud_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudEquipResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudEquipResp) ProtoMessage() {}

func (x *CloudEquipResp) ProtoReflect() protoreflect.Message {
	mi := &file_Cloud_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudEquipResp.ProtoReflect.Descriptor instead.
func (*CloudEquipResp) Descriptor() ([]byte, []int) {
	return file_Cloud_proto_rawDescGZIP(), []int{10}
}

func (x *CloudEquipResp) GetRet() int32 {
	if x != nil && x.Ret != nil {
		return *x.Ret
	}
	return 0
}

func (x *CloudEquipResp) GetPlayerCloudDataMsg() *PlayerCloudDataMsg {
	if x != nil {
		return x.PlayerCloudDataMsg
	}
	return nil
}

type CloudEquipSkinReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CloudSkinId *int32 `protobuf:"varint,1,req,name=cloudSkinId" json:"cloudSkinId,omitempty"`
}

func (x *CloudEquipSkinReq) Reset() {
	*x = CloudEquipSkinReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Cloud_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudEquipSkinReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudEquipSkinReq) ProtoMessage() {}

func (x *CloudEquipSkinReq) ProtoReflect() protoreflect.Message {
	mi := &file_Cloud_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudEquipSkinReq.ProtoReflect.Descriptor instead.
func (*CloudEquipSkinReq) Descriptor() ([]byte, []int) {
	return file_Cloud_proto_rawDescGZIP(), []int{11}
}

func (x *CloudEquipSkinReq) GetCloudSkinId() int32 {
	if x != nil && x.CloudSkinId != nil {
		return *x.CloudSkinId
	}
	return 0
}

type CloudEquipSkinResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret                *int32              `protobuf:"varint,1,req,name=ret" json:"ret,omitempty"`
	PlayerCloudDataMsg *PlayerCloudDataMsg `protobuf:"bytes,2,opt,name=playerCloudDataMsg" json:"playerCloudDataMsg,omitempty"`
}

func (x *CloudEquipSkinResp) Reset() {
	*x = CloudEquipSkinResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Cloud_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudEquipSkinResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudEquipSkinResp) ProtoMessage() {}

func (x *CloudEquipSkinResp) ProtoReflect() protoreflect.Message {
	mi := &file_Cloud_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudEquipSkinResp.ProtoReflect.Descriptor instead.
func (*CloudEquipSkinResp) Descriptor() ([]byte, []int) {
	return file_Cloud_proto_rawDescGZIP(), []int{12}
}

func (x *CloudEquipSkinResp) GetRet() int32 {
	if x != nil && x.Ret != nil {
		return *x.Ret
	}
	return 0
}

func (x *CloudEquipSkinResp) GetPlayerCloudDataMsg() *PlayerCloudDataMsg {
	if x != nil {
		return x.PlayerCloudDataMsg
	}
	return nil
}

type CloudSkinLvUpReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CloudSkinId *int32 `protobuf:"varint,1,req,name=cloudSkinId" json:"cloudSkinId,omitempty"`
}

func (x *CloudSkinLvUpReq) Reset() {
	*x = CloudSkinLvUpReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Cloud_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudSkinLvUpReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudSkinLvUpReq) ProtoMessage() {}

func (x *CloudSkinLvUpReq) ProtoReflect() protoreflect.Message {
	mi := &file_Cloud_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudSkinLvUpReq.ProtoReflect.Descriptor instead.
func (*CloudSkinLvUpReq) Descriptor() ([]byte, []int) {
	return file_Cloud_proto_rawDescGZIP(), []int{13}
}

func (x *CloudSkinLvUpReq) GetCloudSkinId() int32 {
	if x != nil && x.CloudSkinId != nil {
		return *x.CloudSkinId
	}
	return 0
}

type CloudSkinLvUpResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret                *int32              `protobuf:"varint,1,req,name=ret" json:"ret,omitempty"`
	PlayerCloudDataMsg *PlayerCloudDataMsg `protobuf:"bytes,2,opt,name=playerCloudDataMsg" json:"playerCloudDataMsg,omitempty"`
}

func (x *CloudSkinLvUpResp) Reset() {
	*x = CloudSkinLvUpResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Cloud_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudSkinLvUpResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudSkinLvUpResp) ProtoMessage() {}

func (x *CloudSkinLvUpResp) ProtoReflect() protoreflect.Message {
	mi := &file_Cloud_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudSkinLvUpResp.ProtoReflect.Descriptor instead.
func (*CloudSkinLvUpResp) Descriptor() ([]byte, []int) {
	return file_Cloud_proto_rawDescGZIP(), []int{14}
}

func (x *CloudSkinLvUpResp) GetRet() int32 {
	if x != nil && x.Ret != nil {
		return *x.Ret
	}
	return 0
}

func (x *CloudSkinLvUpResp) GetPlayerCloudDataMsg() *PlayerCloudDataMsg {
	if x != nil {
		return x.PlayerCloudDataMsg
	}
	return nil
}

var File_Cloud_proto protoreflect.FileDescriptor

var file_Cloud_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x63,
	0x6f, 0x6d, 0x2e, 0x79, 0x71, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x4d, 0x73,
	0x67, 0x22, 0xb4, 0x02, 0x0a, 0x12, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x73, 0x67, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x71, 0x75, 0x69,
	0x70, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x65, 0x71, 0x75, 0x69, 0x70, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x64, 0x12, 0x44, 0x0a, 0x0c,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x79, 0x71, 0x2e, 0x6d, 0x73, 0x67, 0x2e,
	0x43, 0x69, 0x74, 0x79, 0x4d, 0x73, 0x67, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x44, 0x61, 0x74,
	0x61, 0x4d, 0x73, 0x67, 0x52, 0x0c, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x44, 0x61, 0x74, 0x61, 0x4d,
	0x73, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x02, 0x28,
	0x05, 0x52, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6c, 0x76, 0x18, 0x04,
	0x20, 0x02, 0x28, 0x05, 0x52, 0x02, 0x6c, 0x76, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78, 0x70, 0x18,
	0x05, 0x20, 0x02, 0x28, 0x05, 0x52, 0x03, 0x65, 0x78, 0x70, 0x12, 0x50, 0x0a, 0x10, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x53, 0x6b, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x73, 0x67, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x79, 0x71, 0x2e, 0x6d, 0x73,
	0x67, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x4d, 0x73, 0x67, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x53,
	0x6b, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x73, 0x67, 0x52, 0x10, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x53, 0x6b, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x73, 0x67, 0x12, 0x2a, 0x0a, 0x10,
	0x65, 0x71, 0x75, 0x69, 0x70, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x53, 0x6b, 0x69, 0x6e, 0x49, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x65, 0x71, 0x75, 0x69, 0x70, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x53, 0x6b, 0x69, 0x6e, 0x49, 0x64, 0x22, 0x6c, 0x0a, 0x10, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x53, 0x6b, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x73, 0x67, 0x12, 0x20, 0x0a, 0x0b,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x53, 0x6b, 0x69, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28,
	0x05, 0x52, 0x0b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x53, 0x6b, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x26,
	0x0a, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6c, 0x76, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x6c, 0x76, 0x22, 0x28, 0x0a, 0x0c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x44,
	0x61, 0x74, 0x61, 0x4d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x07, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x64,
	0x22, 0x2a, 0x0a, 0x0e, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x52,
	0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x02, 0x28, 0x05, 0x52, 0x07, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x64, 0x22, 0x7b, 0x0a, 0x0f,
	0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x10, 0x0a, 0x03, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x03, 0x72, 0x65,
	0x74, 0x12, 0x56, 0x0a, 0x12, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x44, 0x61, 0x74, 0x61, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x79, 0x71, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x4d,
	0x73, 0x67, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x44, 0x61,
	0x74, 0x61, 0x4d, 0x73, 0x67, 0x52, 0x12, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x73, 0x67, 0x22, 0x24, 0x0a, 0x0c, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x4c, 0x76, 0x55, 0x70, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x69,
	0x63, 0x6b, 0x18, 0x01, 0x20, 0x02, 0x28, 0x08, 0x52, 0x05, 0x71, 0x75, 0x69, 0x63, 0x6b, 0x22,
	0x91, 0x01, 0x0a, 0x0d, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x4c, 0x76, 0x55, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x03,
	0x72, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x64, 0x64, 0x45, 0x78, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x64, 0x64, 0x45, 0x78, 0x70, 0x12, 0x56, 0x0a, 0x12, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x73,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x79, 0x71,
	0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x4d, 0x73, 0x67, 0x2e, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x73, 0x67, 0x52,
	0x12, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x44, 0x61, 0x74, 0x61,
	0x4d, 0x73, 0x67, 0x22, 0x11, 0x0a, 0x0f, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x53, 0x74, 0x61, 0x67,
	0x65, 0x55, 0x70, 0x52, 0x65, 0x71, 0x22, 0x7c, 0x0a, 0x10, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x53,
	0x74, 0x61, 0x67, 0x65, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65,
	0x74, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x03, 0x72, 0x65, 0x74, 0x12, 0x56, 0x0a, 0x12,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x44, 0x61, 0x74, 0x61, 0x4d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x79,
	0x71, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x4d, 0x73, 0x67, 0x2e, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x73, 0x67,
	0x52, 0x12, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x44, 0x61, 0x74,
	0x61, 0x4d, 0x73, 0x67, 0x22, 0x29, 0x0a, 0x0d, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x45, 0x71, 0x75,
	0x69, 0x70, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x07, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x64, 0x22,
	0x7a, 0x0a, 0x0e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x45, 0x71, 0x75, 0x69, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x03,
	0x72, 0x65, 0x74, 0x12, 0x56, 0x0a, 0x12, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x79, 0x71, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x43, 0x69, 0x74,
	0x79, 0x4d, 0x73, 0x67, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x44, 0x61, 0x74, 0x61, 0x4d, 0x73, 0x67, 0x52, 0x12, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x73, 0x67, 0x22, 0x35, 0x0a, 0x11, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x45, 0x71, 0x75, 0x69, 0x70, 0x53, 0x6b, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x53, 0x6b, 0x69, 0x6e, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x53, 0x6b, 0x69, 0x6e,
	0x49, 0x64, 0x22, 0x7e, 0x0a, 0x12, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x45, 0x71, 0x75, 0x69, 0x70,
	0x53, 0x6b, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x03, 0x72, 0x65, 0x74, 0x12, 0x56, 0x0a, 0x12, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x79, 0x71, 0x2e,
	0x6d, 0x73, 0x67, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x4d, 0x73, 0x67, 0x2e, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x73, 0x67, 0x52, 0x12,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x44, 0x61, 0x74, 0x61, 0x4d,
	0x73, 0x67, 0x22, 0x34, 0x0a, 0x10, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x53, 0x6b, 0x69, 0x6e, 0x4c,
	0x76, 0x55, 0x70, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x53,
	0x6b, 0x69, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x53, 0x6b, 0x69, 0x6e, 0x49, 0x64, 0x22, 0x7d, 0x0a, 0x11, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x53, 0x6b, 0x69, 0x6e, 0x4c, 0x76, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a,
	0x03, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x03, 0x72, 0x65, 0x74, 0x12,
	0x56, 0x0a, 0x12, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x44, 0x61,
	0x74, 0x61, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x79, 0x71, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x4d, 0x73, 0x67,
	0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x44, 0x61, 0x74, 0x61,
	0x4d, 0x73, 0x67, 0x52, 0x12, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x44, 0x61, 0x74, 0x61, 0x4d, 0x73, 0x67, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6d, 0x7a, 0x68, 0x6f, 0x6e, 0x67, 0x71, 0x69, 0x2f,
	0x78, 0x64, 0x64, 0x71, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x61, 0x6d, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62,
}

var (
	file_Cloud_proto_rawDescOnce sync.Once
	file_Cloud_proto_rawDescData = file_Cloud_proto_rawDesc
)

func file_Cloud_proto_rawDescGZIP() []byte {
	file_Cloud_proto_rawDescOnce.Do(func() {
		file_Cloud_proto_rawDescData = protoimpl.X.CompressGZIP(file_Cloud_proto_rawDescData)
	})
	return file_Cloud_proto_rawDescData
}

var file_Cloud_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_Cloud_proto_goTypes = []any{
	(*PlayerCloudDataMsg)(nil), // 0: com.yq.msg.CityMsg.PlayerCloudDataMsg
	(*CloudSkinDataMsg)(nil),   // 1: com.yq.msg.CityMsg.CloudSkinDataMsg
	(*CloudDataMsg)(nil),       // 2: com.yq.msg.CityMsg.CloudDataMsg
	(*UnlockCloudReq)(nil),     // 3: com.yq.msg.CityMsg.UnlockCloudReq
	(*UnlockCloudResp)(nil),    // 4: com.yq.msg.CityMsg.UnlockCloudResp
	(*CloudLvUpReq)(nil),       // 5: com.yq.msg.CityMsg.CloudLvUpReq
	(*CloudLvUpResp)(nil),      // 6: com.yq.msg.CityMsg.CloudLvUpResp
	(*CloudStageUpReq)(nil),    // 7: com.yq.msg.CityMsg.CloudStageUpReq
	(*CloudStageUpResp)(nil),   // 8: com.yq.msg.CityMsg.CloudStageUpResp
	(*CloudEquipReq)(nil),      // 9: com.yq.msg.CityMsg.CloudEquipReq
	(*CloudEquipResp)(nil),     // 10: com.yq.msg.CityMsg.CloudEquipResp
	(*CloudEquipSkinReq)(nil),  // 11: com.yq.msg.CityMsg.CloudEquipSkinReq
	(*CloudEquipSkinResp)(nil), // 12: com.yq.msg.CityMsg.CloudEquipSkinResp
	(*CloudSkinLvUpReq)(nil),   // 13: com.yq.msg.CityMsg.CloudSkinLvUpReq
	(*CloudSkinLvUpResp)(nil),  // 14: com.yq.msg.CityMsg.CloudSkinLvUpResp
}
var file_Cloud_proto_depIdxs = []int32{
	2, // 0: com.yq.msg.CityMsg.PlayerCloudDataMsg.cloudDataMsg:type_name -> com.yq.msg.CityMsg.CloudDataMsg
	1, // 1: com.yq.msg.CityMsg.PlayerCloudDataMsg.cloudSkinDataMsg:type_name -> com.yq.msg.CityMsg.CloudSkinDataMsg
	0, // 2: com.yq.msg.CityMsg.UnlockCloudResp.playerCloudDataMsg:type_name -> com.yq.msg.CityMsg.PlayerCloudDataMsg
	0, // 3: com.yq.msg.CityMsg.CloudLvUpResp.playerCloudDataMsg:type_name -> com.yq.msg.CityMsg.PlayerCloudDataMsg
	0, // 4: com.yq.msg.CityMsg.CloudStageUpResp.playerCloudDataMsg:type_name -> com.yq.msg.CityMsg.PlayerCloudDataMsg
	0, // 5: com.yq.msg.CityMsg.CloudEquipResp.playerCloudDataMsg:type_name -> com.yq.msg.CityMsg.PlayerCloudDataMsg
	0, // 6: com.yq.msg.CityMsg.CloudEquipSkinResp.playerCloudDataMsg:type_name -> com.yq.msg.CityMsg.PlayerCloudDataMsg
	0, // 7: com.yq.msg.CityMsg.CloudSkinLvUpResp.playerCloudDataMsg:type_name -> com.yq.msg.CityMsg.PlayerCloudDataMsg
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_Cloud_proto_init() }
func file_Cloud_proto_init() {
	if File_Cloud_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Cloud_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*PlayerCloudDataMsg); i {
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
		file_Cloud_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CloudSkinDataMsg); i {
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
		file_Cloud_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CloudDataMsg); i {
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
		file_Cloud_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UnlockCloudReq); i {
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
		file_Cloud_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UnlockCloudResp); i {
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
		file_Cloud_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*CloudLvUpReq); i {
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
		file_Cloud_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*CloudLvUpResp); i {
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
		file_Cloud_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*CloudStageUpReq); i {
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
		file_Cloud_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*CloudStageUpResp); i {
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
		file_Cloud_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*CloudEquipReq); i {
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
		file_Cloud_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*CloudEquipResp); i {
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
		file_Cloud_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*CloudEquipSkinReq); i {
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
		file_Cloud_proto_msgTypes[12].Exporter = func(v any, i int) any {
			switch v := v.(*CloudEquipSkinResp); i {
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
		file_Cloud_proto_msgTypes[13].Exporter = func(v any, i int) any {
			switch v := v.(*CloudSkinLvUpReq); i {
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
		file_Cloud_proto_msgTypes[14].Exporter = func(v any, i int) any {
			switch v := v.(*CloudSkinLvUpResp); i {
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
			RawDescriptor: file_Cloud_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Cloud_proto_goTypes,
		DependencyIndexes: file_Cloud_proto_depIdxs,
		MessageInfos:      file_Cloud_proto_msgTypes,
	}.Build()
	File_Cloud_proto = out.File
	file_Cloud_proto_rawDesc = nil
	file_Cloud_proto_goTypes = nil
	file_Cloud_proto_depIdxs = nil
}
