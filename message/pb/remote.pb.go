// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: remote.proto

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

type TestMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *TestMsg) Reset() {
	*x = TestMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestMsg) ProtoMessage() {}

func (x *TestMsg) ProtoReflect() protoreflect.Message {
	mi := &file_remote_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestMsg.ProtoReflect.Descriptor instead.
func (*TestMsg) Descriptor() ([]byte, []int) {
	return file_remote_proto_rawDescGZIP(), []int{0}
}

func (x *TestMsg) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type SyncMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    uint32         `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	SyncTrans *SyncTransForm `protobuf:"bytes,2,opt,name=sync_trans,json=syncTrans,proto3" json:"sync_trans,omitempty"`
	SyncKey   *SyncKey       `protobuf:"bytes,3,opt,name=sync_key,json=syncKey,proto3" json:"sync_key,omitempty"`
	SyncAxis  []*SyncAxis    `protobuf:"bytes,4,rep,name=sync_axis,json=syncAxis,proto3" json:"sync_axis,omitempty"`
}

func (x *SyncMsg) Reset() {
	*x = SyncMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncMsg) ProtoMessage() {}

func (x *SyncMsg) ProtoReflect() protoreflect.Message {
	mi := &file_remote_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncMsg.ProtoReflect.Descriptor instead.
func (*SyncMsg) Descriptor() ([]byte, []int) {
	return file_remote_proto_rawDescGZIP(), []int{1}
}

func (x *SyncMsg) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *SyncMsg) GetSyncTrans() *SyncTransForm {
	if x != nil {
		return x.SyncTrans
	}
	return nil
}

func (x *SyncMsg) GetSyncKey() *SyncKey {
	if x != nil {
		return x.SyncKey
	}
	return nil
}

func (x *SyncMsg) GetSyncAxis() []*SyncAxis {
	if x != nil {
		return x.SyncAxis
	}
	return nil
}

type SyncTransForm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PositionX float32 `protobuf:"fixed32,1,opt,name=position_x,json=positionX,proto3" json:"position_x,omitempty"`
	PositionY float32 `protobuf:"fixed32,2,opt,name=position_y,json=positionY,proto3" json:"position_y,omitempty"`
	PositionZ float32 `protobuf:"fixed32,3,opt,name=position_z,json=positionZ,proto3" json:"position_z,omitempty"`
	RotationX float32 `protobuf:"fixed32,4,opt,name=rotation_x,json=rotationX,proto3" json:"rotation_x,omitempty"`
	RotationY float32 `protobuf:"fixed32,5,opt,name=rotation_y,json=rotationY,proto3" json:"rotation_y,omitempty"`
	RotationZ float32 `protobuf:"fixed32,6,opt,name=rotation_z,json=rotationZ,proto3" json:"rotation_z,omitempty"`
}

func (x *SyncTransForm) Reset() {
	*x = SyncTransForm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncTransForm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncTransForm) ProtoMessage() {}

func (x *SyncTransForm) ProtoReflect() protoreflect.Message {
	mi := &file_remote_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncTransForm.ProtoReflect.Descriptor instead.
func (*SyncTransForm) Descriptor() ([]byte, []int) {
	return file_remote_proto_rawDescGZIP(), []int{2}
}

func (x *SyncTransForm) GetPositionX() float32 {
	if x != nil {
		return x.PositionX
	}
	return 0
}

func (x *SyncTransForm) GetPositionY() float32 {
	if x != nil {
		return x.PositionY
	}
	return 0
}

func (x *SyncTransForm) GetPositionZ() float32 {
	if x != nil {
		return x.PositionZ
	}
	return 0
}

func (x *SyncTransForm) GetRotationX() float32 {
	if x != nil {
		return x.RotationX
	}
	return 0
}

func (x *SyncTransForm) GetRotationY() float32 {
	if x != nil {
		return x.RotationY
	}
	return 0
}

func (x *SyncTransForm) GetRotationZ() float32 {
	if x != nil {
		return x.RotationZ
	}
	return 0
}

type SyncKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyCode uint32 `protobuf:"varint,1,opt,name=key_code,json=keyCode,proto3" json:"key_code,omitempty"`
	KeyDown bool   `protobuf:"varint,2,opt,name=key_down,json=keyDown,proto3" json:"key_down,omitempty"`
}

func (x *SyncKey) Reset() {
	*x = SyncKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncKey) ProtoMessage() {}

func (x *SyncKey) ProtoReflect() protoreflect.Message {
	mi := &file_remote_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncKey.ProtoReflect.Descriptor instead.
func (*SyncKey) Descriptor() ([]byte, []int) {
	return file_remote_proto_rawDescGZIP(), []int{3}
}

func (x *SyncKey) GetKeyCode() uint32 {
	if x != nil {
		return x.KeyCode
	}
	return 0
}

func (x *SyncKey) GetKeyDown() bool {
	if x != nil {
		return x.KeyDown
	}
	return false
}

type SyncAxis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value float32 `protobuf:"fixed32,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SyncAxis) Reset() {
	*x = SyncAxis{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncAxis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncAxis) ProtoMessage() {}

func (x *SyncAxis) ProtoReflect() protoreflect.Message {
	mi := &file_remote_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncAxis.ProtoReflect.Descriptor instead.
func (*SyncAxis) Descriptor() ([]byte, []int) {
	return file_remote_proto_rawDescGZIP(), []int{4}
}

func (x *SyncAxis) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SyncAxis) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_remote_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_remote_proto_rawDescGZIP(), []int{5}
}

func (x *LoginReq) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type LoginRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Result bool   `protobuf:"varint,2,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *LoginRes) Reset() {
	*x = LoginRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRes) ProtoMessage() {}

func (x *LoginRes) ProtoReflect() protoreflect.Message {
	mi := &file_remote_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRes.ProtoReflect.Descriptor instead.
func (*LoginRes) Descriptor() ([]byte, []int) {
	return file_remote_proto_rawDescGZIP(), []int{6}
}

func (x *LoginRes) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *LoginRes) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

var File_remote_proto protoreflect.FileDescriptor

var file_remote_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15,
	0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33, 0x22, 0x23, 0x0a, 0x07, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x73, 0x67,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0xe0, 0x01, 0x0a, 0x07, 0x53,
	0x79, 0x6e, 0x63, 0x4d, 0x73, 0x67, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x43, 0x0a, 0x0a, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x2e, 0x53, 0x79, 0x6e, 0x63,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x46, 0x6f, 0x72, 0x6d, 0x52, 0x09, 0x73, 0x79, 0x6e, 0x63, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x12, 0x39, 0x0a, 0x08, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x2e, 0x53,
	0x79, 0x6e, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x07, 0x73, 0x79, 0x6e, 0x63, 0x4b, 0x65, 0x79, 0x12,
	0x3c, 0x0a, 0x09, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x61, 0x78, 0x69, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x41,
	0x78, 0x69, 0x73, 0x52, 0x08, 0x73, 0x79, 0x6e, 0x63, 0x41, 0x78, 0x69, 0x73, 0x22, 0xc9, 0x01,
	0x0a, 0x0d, 0x53, 0x79, 0x6e, 0x63, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x46, 0x6f, 0x72, 0x6d, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x78, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x58, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x59, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x7a, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5a, 0x12, 0x1d, 0x0a, 0x0a,
	0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x09, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x58, 0x12, 0x1d, 0x0a, 0x0a, 0x72,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x09, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x59, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x7a, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09,
	0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5a, 0x22, 0x3f, 0x0a, 0x07, 0x53, 0x79, 0x6e,
	0x63, 0x4b, 0x65, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x64, 0x6f, 0x77, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x44, 0x6f, 0x77, 0x6e, 0x22, 0x34, 0x0a, 0x08, 0x53, 0x79,
	0x6e, 0x63, 0x41, 0x78, 0x69, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x23, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3b, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x73, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_remote_proto_rawDescOnce sync.Once
	file_remote_proto_rawDescData = file_remote_proto_rawDesc
)

func file_remote_proto_rawDescGZIP() []byte {
	file_remote_proto_rawDescOnce.Do(func() {
		file_remote_proto_rawDescData = protoimpl.X.CompressGZIP(file_remote_proto_rawDescData)
	})
	return file_remote_proto_rawDescData
}

var file_remote_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_remote_proto_goTypes = []interface{}{
	(*TestMsg)(nil),       // 0: goproto.protoc.proto3.TestMsg
	(*SyncMsg)(nil),       // 1: goproto.protoc.proto3.SyncMsg
	(*SyncTransForm)(nil), // 2: goproto.protoc.proto3.SyncTransForm
	(*SyncKey)(nil),       // 3: goproto.protoc.proto3.SyncKey
	(*SyncAxis)(nil),      // 4: goproto.protoc.proto3.SyncAxis
	(*LoginReq)(nil),      // 5: goproto.protoc.proto3.LoginReq
	(*LoginRes)(nil),      // 6: goproto.protoc.proto3.LoginRes
}
var file_remote_proto_depIdxs = []int32{
	2, // 0: goproto.protoc.proto3.SyncMsg.sync_trans:type_name -> goproto.protoc.proto3.SyncTransForm
	3, // 1: goproto.protoc.proto3.SyncMsg.sync_key:type_name -> goproto.protoc.proto3.SyncKey
	4, // 2: goproto.protoc.proto3.SyncMsg.sync_axis:type_name -> goproto.protoc.proto3.SyncAxis
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_remote_proto_init() }
func file_remote_proto_init() {
	if File_remote_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_remote_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestMsg); i {
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
		file_remote_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncMsg); i {
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
		file_remote_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncTransForm); i {
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
		file_remote_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncKey); i {
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
		file_remote_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncAxis); i {
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
		file_remote_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReq); i {
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
		file_remote_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRes); i {
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
			RawDescriptor: file_remote_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_remote_proto_goTypes,
		DependencyIndexes: file_remote_proto_depIdxs,
		MessageInfos:      file_remote_proto_msgTypes,
	}.Build()
	File_remote_proto = out.File
	file_remote_proto_rawDesc = nil
	file_remote_proto_goTypes = nil
	file_remote_proto_depIdxs = nil
}
