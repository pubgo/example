// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: proto/login/merge.proto

// 账户合并相关

package login

import (
	_ "github.com/gogo/protobuf/gogoproto"
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

type TelephoneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 登陆用户
	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	// 新手机号
	TargetTelephone string `protobuf:"bytes,2,opt,name=targetTelephone,proto3" json:"targetTelephone,omitempty"`
	// 是否走新流程
	IsNewProcess bool `protobuf:"varint,3,opt,name=isNewProcess,proto3" json:"isNewProcess,omitempty"`
}

func (x *TelephoneRequest) Reset() {
	*x = TelephoneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_login_merge_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TelephoneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TelephoneRequest) ProtoMessage() {}

func (x *TelephoneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_login_merge_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TelephoneRequest.ProtoReflect.Descriptor instead.
func (*TelephoneRequest) Descriptor() ([]byte, []int) {
	return file_proto_login_merge_proto_rawDescGZIP(), []int{0}
}

func (x *TelephoneRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *TelephoneRequest) GetTargetTelephone() string {
	if x != nil {
		return x.TargetTelephone
	}
	return ""
}

func (x *TelephoneRequest) GetIsNewProcess() bool {
	if x != nil {
		return x.IsNewProcess
	}
	return false
}

type WeChatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 登陆用户
	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	// 要合并的用户
	TargetUid int64 `protobuf:"varint,2,opt,name=targetUid,proto3" json:"targetUid,omitempty"`
}

func (x *WeChatRequest) Reset() {
	*x = WeChatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_login_merge_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeChatRequest) ProtoMessage() {}

func (x *WeChatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_login_merge_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeChatRequest.ProtoReflect.Descriptor instead.
func (*WeChatRequest) Descriptor() ([]byte, []int) {
	return file_proto_login_merge_proto_rawDescGZIP(), []int{1}
}

func (x *WeChatRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *WeChatRequest) GetTargetUid() int64 {
	if x != nil {
		return x.TargetUid
	}
	return 0
}

type WeChatUnMergeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 登陆用户
	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *WeChatUnMergeRequest) Reset() {
	*x = WeChatUnMergeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_login_merge_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeChatUnMergeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeChatUnMergeRequest) ProtoMessage() {}

func (x *WeChatUnMergeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_login_merge_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeChatUnMergeRequest.ProtoReflect.Descriptor instead.
func (*WeChatUnMergeRequest) Descriptor() ([]byte, []int) {
	return file_proto_login_merge_proto_rawDescGZIP(), []int{2}
}

func (x *WeChatUnMergeRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// code
	Code int64 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// msg
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	// 时间戳
	NowTime int64 `protobuf:"varint,3,opt,name=nowTime,proto3" json:"nowTime,omitempty"`
	// 数据
	Data map[string]string `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Reply) Reset() {
	*x = Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_login_merge_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_login_merge_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply.ProtoReflect.Descriptor instead.
func (*Reply) Descriptor() ([]byte, []int) {
	return file_proto_login_merge_proto_rawDescGZIP(), []int{3}
}

func (x *Reply) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Reply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *Reply) GetNowTime() int64 {
	if x != nil {
		return x.NowTime
	}
	return 0
}

func (x *Reply) GetData() map[string]string {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_proto_login_merge_proto protoreflect.FileDescriptor

var file_proto_login_merge_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2f, 0x6d, 0x65,
	0x72, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x01, 0x0a, 0x10, 0x54, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x03, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xea, 0xde, 0x1f, 0x03, 0x75, 0x69, 0x64, 0x52,
	0x03, 0x75, 0x69, 0x64, 0x12, 0x3d, 0x0a, 0x0f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x65,
	0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x13, 0xea,
	0xde, 0x1f, 0x0f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x52, 0x0f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x65, 0x6c, 0x65, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x12, 0x34, 0x0a, 0x0c, 0x69, 0x73, 0x4e, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x42, 0x10, 0xea, 0xde, 0x1f, 0x0c, 0x69,
	0x73, 0x4e, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x0c, 0x69, 0x73, 0x4e,
	0x65, 0x77, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x22, 0x57, 0x0a, 0x0d, 0x57, 0x65, 0x43,
	0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x03, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xea, 0xde, 0x1f, 0x03, 0x75, 0x69, 0x64,
	0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x2b, 0x0a, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0d, 0xea, 0xde, 0x1f, 0x09, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x55, 0x69, 0x64, 0x52, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55,
	0x69, 0x64, 0x22, 0x31, 0x0a, 0x14, 0x57, 0x65, 0x43, 0x68, 0x61, 0x74, 0x55, 0x6e, 0x4d, 0x65,
	0x72, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x03, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xea, 0xde, 0x1f, 0x03, 0x75, 0x69, 0x64,
	0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0xd6, 0x01, 0x0a, 0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x1c, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x08, 0xea,
	0xde, 0x1f, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xea, 0xde, 0x1f, 0x03,
	0x6d, 0x73, 0x67, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x25, 0x0a, 0x07, 0x6e, 0x6f, 0x77, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0b, 0xea, 0xde, 0x1f, 0x07, 0x6e,
	0x6f, 0x77, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x07, 0x6e, 0x6f, 0x77, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x34, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x08, 0xea, 0xde, 0x1f, 0x04, 0x64, 0x61, 0x74, 0x61, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x37, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xca,
	0x03, 0x0a, 0x05, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x12, 0x54, 0x0a, 0x09, 0x54, 0x65, 0x6c, 0x65,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x17, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x54, 0x65,
	0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c,
	0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x20, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6d, 0x65, 0x72, 0x67,
	0x65, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x5f,
	0x0a, 0x0e, 0x54, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x12, 0x17, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x54, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x22,
	0x1b, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x2f, 0x74, 0x65, 0x6c,
	0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x2d, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x3a, 0x01, 0x2a, 0x12,
	0x4c, 0x0a, 0x06, 0x57, 0x65, 0x43, 0x68, 0x61, 0x74, 0x12, 0x14, 0x2e, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x2e, 0x57, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0c, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1e, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22, 0x13, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6d, 0x65, 0x72,
	0x67, 0x65, 0x2f, 0x77, 0x65, 0x2d, 0x63, 0x68, 0x61, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x57, 0x0a,
	0x0b, 0x57, 0x65, 0x43, 0x68, 0x61, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x14, 0x2e, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x57, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x22, 0x19, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x6d, 0x65, 0x72, 0x67, 0x65, 0x2f, 0x77, 0x65, 0x2d, 0x63, 0x68, 0x61, 0x74, 0x2d, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x3a, 0x01, 0x2a, 0x12, 0x63, 0x0a, 0x0d, 0x57, 0x65, 0x43, 0x68, 0x61, 0x74,
	0x55, 0x6e, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x12, 0x1b, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e,
	0x57, 0x65, 0x43, 0x68, 0x61, 0x74, 0x55, 0x6e, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x22, 0x1c, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2f, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x2f, 0x77, 0x65, 0x2d, 0x63, 0x68, 0x61, 0x74, 0x2d,
	0x75, 0x6e, 0x2d, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x3a, 0x01, 0x2a, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_login_merge_proto_rawDescOnce sync.Once
	file_proto_login_merge_proto_rawDescData = file_proto_login_merge_proto_rawDesc
)

func file_proto_login_merge_proto_rawDescGZIP() []byte {
	file_proto_login_merge_proto_rawDescOnce.Do(func() {
		file_proto_login_merge_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_login_merge_proto_rawDescData)
	})
	return file_proto_login_merge_proto_rawDescData
}

var file_proto_login_merge_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_login_merge_proto_goTypes = []interface{}{
	(*TelephoneRequest)(nil),     // 0: login.TelephoneRequest
	(*WeChatRequest)(nil),        // 1: login.WeChatRequest
	(*WeChatUnMergeRequest)(nil), // 2: login.WeChatUnMergeRequest
	(*Reply)(nil),                // 3: login.Reply
	nil,                          // 4: login.Reply.DataEntry
}
var file_proto_login_merge_proto_depIdxs = []int32{
	4, // 0: login.Reply.data:type_name -> login.Reply.DataEntry
	0, // 1: login.Merge.Telephone:input_type -> login.TelephoneRequest
	0, // 2: login.Merge.TelephoneCheck:input_type -> login.TelephoneRequest
	1, // 3: login.Merge.WeChat:input_type -> login.WeChatRequest
	1, // 4: login.Merge.WeChatCheck:input_type -> login.WeChatRequest
	2, // 5: login.Merge.WeChatUnMerge:input_type -> login.WeChatUnMergeRequest
	3, // 6: login.Merge.Telephone:output_type -> login.Reply
	3, // 7: login.Merge.TelephoneCheck:output_type -> login.Reply
	3, // 8: login.Merge.WeChat:output_type -> login.Reply
	3, // 9: login.Merge.WeChatCheck:output_type -> login.Reply
	3, // 10: login.Merge.WeChatUnMerge:output_type -> login.Reply
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_login_merge_proto_init() }
func file_proto_login_merge_proto_init() {
	if File_proto_login_merge_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_login_merge_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TelephoneRequest); i {
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
		file_proto_login_merge_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeChatRequest); i {
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
		file_proto_login_merge_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeChatUnMergeRequest); i {
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
		file_proto_login_merge_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply); i {
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
			RawDescriptor: file_proto_login_merge_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_login_merge_proto_goTypes,
		DependencyIndexes: file_proto_login_merge_proto_depIdxs,
		MessageInfos:      file_proto_login_merge_proto_msgTypes,
	}.Build()
	File_proto_login_merge_proto = out.File
	file_proto_login_merge_proto_rawDesc = nil
	file_proto_login_merge_proto_goTypes = nil
	file_proto_login_merge_proto_depIdxs = nil
}
