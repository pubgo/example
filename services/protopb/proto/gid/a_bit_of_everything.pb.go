// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: proto/gid/a_bit_of_everything.proto

package gid

import (
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

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The entered username
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// The entered password
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gid_a_bit_of_everything_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gid_a_bit_of_everything_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_proto_gid_a_bit_of_everything_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	// Whether you have access or not
	Access bool `protobuf:"varint,2,opt,name=access,proto3" json:"access,omitempty"`
}

func (x *LoginReply) Reset() {
	*x = LoginReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gid_a_bit_of_everything_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReply) ProtoMessage() {}

func (x *LoginReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gid_a_bit_of_everything_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReply.ProtoReflect.Descriptor instead.
func (*LoginReply) Descriptor() ([]byte, []int) {
	return file_proto_gid_a_bit_of_everything_proto_rawDescGZIP(), []int{1}
}

func (x *LoginReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *LoginReply) GetAccess() bool {
	if x != nil {
		return x.Access
	}
	return false
}

type LogoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The time the logout was registered
	Timeoflogout string `protobuf:"bytes,1,opt,name=timeoflogout,proto3" json:"timeoflogout,omitempty"`
	// This is the title
	//
	// This is the "Description" of field test
	// you can use as many newlines as you want
	//
	//
	// it will still format the same in the table
	Test int32 `protobuf:"varint,2,opt,name=test,proto3" json:"test,omitempty"`
	// This is an array
	//
	// It displays that using [] infront of the type
	Stringarray []string `protobuf:"bytes,3,rep,name=stringarray,proto3" json:"stringarray,omitempty"`
}

func (x *LogoutRequest) Reset() {
	*x = LogoutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gid_a_bit_of_everything_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutRequest) ProtoMessage() {}

func (x *LogoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gid_a_bit_of_everything_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutRequest.ProtoReflect.Descriptor instead.
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return file_proto_gid_a_bit_of_everything_proto_rawDescGZIP(), []int{2}
}

func (x *LogoutRequest) GetTimeoflogout() string {
	if x != nil {
		return x.Timeoflogout
	}
	return ""
}

func (x *LogoutRequest) GetTest() int32 {
	if x != nil {
		return x.Test
	}
	return 0
}

func (x *LogoutRequest) GetStringarray() []string {
	if x != nil {
		return x.Stringarray
	}
	return nil
}

type LogoutReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Message that tells you whether your
	// logout was succesful or not
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *LogoutReply) Reset() {
	*x = LogoutReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gid_a_bit_of_everything_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutReply) ProtoMessage() {}

func (x *LogoutReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gid_a_bit_of_everything_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutReply.ProtoReflect.Descriptor instead.
func (*LogoutReply) Descriptor() ([]byte, []int) {
	return file_proto_gid_a_bit_of_everything_proto_rawDescGZIP(), []int{3}
}

func (x *LogoutReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_gid_a_bit_of_everything_proto protoreflect.FileDescriptor

var file_proto_gid_a_bit_of_everything_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x69, 0x64, 0x2f, 0x61, 0x5f, 0x62, 0x69,
	0x74, 0x5f, 0x6f, 0x66, 0x5f, 0x65, 0x76, 0x65, 0x72, 0x79, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x67, 0x69, 0x64, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x22, 0x3e, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x22, 0x69, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x66, 0x6c, 0x6f, 0x67, 0x6f, 0x75,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x66, 0x6c,
	0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x61, 0x72, 0x72, 0x61, 0x79, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x61, 0x72, 0x72, 0x61, 0x79, 0x22, 0x27, 0x0a, 0x0b, 0x4c,
	0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x32, 0xa8, 0x01, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x11,
	0x2e, 0x67, 0x69, 0x64, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0f, 0x2e, 0x67, 0x69, 0x64, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x11, 0x2f, 0x76, 0x31, 0x2f,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x3a, 0x01, 0x2a,
	0x12, 0x4d, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x12, 0x2e, 0x67, 0x69, 0x64,
	0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10,
	0x2e, 0x67, 0x69, 0x64, 0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x6c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x3a, 0x01, 0x2a, 0x42,
	0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x3b, 0x67, 0x69, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_gid_a_bit_of_everything_proto_rawDescOnce sync.Once
	file_proto_gid_a_bit_of_everything_proto_rawDescData = file_proto_gid_a_bit_of_everything_proto_rawDesc
)

func file_proto_gid_a_bit_of_everything_proto_rawDescGZIP() []byte {
	file_proto_gid_a_bit_of_everything_proto_rawDescOnce.Do(func() {
		file_proto_gid_a_bit_of_everything_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_gid_a_bit_of_everything_proto_rawDescData)
	})
	return file_proto_gid_a_bit_of_everything_proto_rawDescData
}

var file_proto_gid_a_bit_of_everything_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_gid_a_bit_of_everything_proto_goTypes = []interface{}{
	(*LoginRequest)(nil),  // 0: gid.LoginRequest
	(*LoginReply)(nil),    // 1: gid.LoginReply
	(*LogoutRequest)(nil), // 2: gid.LogoutRequest
	(*LogoutReply)(nil),   // 3: gid.LogoutReply
}
var file_proto_gid_a_bit_of_everything_proto_depIdxs = []int32{
	0, // 0: gid.LoginService.Login:input_type -> gid.LoginRequest
	2, // 1: gid.LoginService.Logout:input_type -> gid.LogoutRequest
	1, // 2: gid.LoginService.Login:output_type -> gid.LoginReply
	3, // 3: gid.LoginService.Logout:output_type -> gid.LogoutReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_gid_a_bit_of_everything_proto_init() }
func file_proto_gid_a_bit_of_everything_proto_init() {
	if File_proto_gid_a_bit_of_everything_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_gid_a_bit_of_everything_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_proto_gid_a_bit_of_everything_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReply); i {
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
		file_proto_gid_a_bit_of_everything_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoutRequest); i {
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
		file_proto_gid_a_bit_of_everything_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoutReply); i {
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
			RawDescriptor: file_proto_gid_a_bit_of_everything_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_gid_a_bit_of_everything_proto_goTypes,
		DependencyIndexes: file_proto_gid_a_bit_of_everything_proto_depIdxs,
		MessageInfos:      file_proto_gid_a_bit_of_everything_proto_msgTypes,
	}.Build()
	File_proto_gid_a_bit_of_everything_proto = out.File
	file_proto_gid_a_bit_of_everything_proto_rawDesc = nil
	file_proto_gid_a_bit_of_everything_proto_goTypes = nil
	file_proto_gid_a_bit_of_everything_proto_depIdxs = nil
}
