// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: v1/account.proto

package account

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CreateUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NickName string `protobuf:"bytes,1,opt,name=nickName,proto3" json:"nickName,omitempty"`
	Gender   string `protobuf:"bytes,2,opt,name=gender,proto3" json:"gender,omitempty"`
	Age      int32  `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	Phone    string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *CreateUserReq) Reset() {
	*x = CreateUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserReq) ProtoMessage() {}

func (x *CreateUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserReq.ProtoReflect.Descriptor instead.
func (*CreateUserReq) Descriptor() ([]byte, []int) {
	return file_v1_account_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserReq) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *CreateUserReq) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *CreateUserReq) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *CreateUserReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type CreateUserResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *CreateUserResp) Reset() {
	*x = CreateUserResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResp) ProtoMessage() {}

func (x *CreateUserResp) ProtoReflect() protoreflect.Message {
	mi := &file_v1_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResp.ProtoReflect.Descriptor instead.
func (*CreateUserResp) Descriptor() ([]byte, []int) {
	return file_v1_account_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CreateUserResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type DelUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID   int64  `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	NickName string `protobuf:"bytes,2,opt,name=nickName,proto3" json:"nickName,omitempty"`
	Bender   string `protobuf:"bytes,3,opt,name=bender,proto3" json:"bender,omitempty"`
	Age      int32  `protobuf:"varint,4,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *DelUserReq) Reset() {
	*x = DelUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelUserReq) ProtoMessage() {}

func (x *DelUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelUserReq.ProtoReflect.Descriptor instead.
func (*DelUserReq) Descriptor() ([]byte, []int) {
	return file_v1_account_proto_rawDescGZIP(), []int{2}
}

func (x *DelUserReq) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *DelUserReq) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *DelUserReq) GetBender() string {
	if x != nil {
		return x.Bender
	}
	return ""
}

func (x *DelUserReq) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

type DelUserResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *DelUserResp) Reset() {
	*x = DelUserResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_account_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelUserResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelUserResp) ProtoMessage() {}

func (x *DelUserResp) ProtoReflect() protoreflect.Message {
	mi := &file_v1_account_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelUserResp.ProtoReflect.Descriptor instead.
func (*DelUserResp) Descriptor() ([]byte, []int) {
	return file_v1_account_proto_rawDescGZIP(), []int{3}
}

func (x *DelUserResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *DelUserResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_v1_account_proto protoreflect.FileDescriptor

var file_v1_account_proto_rawDesc = []byte{
	0x0a, 0x10, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x11, 0x77, 0x65, 0x65, 0x6b, 0x30, 0x34, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x6b, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x22, 0x36, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x6a, 0x0a, 0x0a, 0x44, 0x65,
	0x6c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x62, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x22, 0x33, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0xaa, 0x01, 0x0a, 0x07,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x53, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x77, 0x65, 0x65, 0x6b, 0x30, 0x34, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x21, 0x2e, 0x77, 0x65, 0x65, 0x6b, 0x30, 0x34,
	0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x07,
	0x44, 0x65, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x77, 0x65, 0x65, 0x6b, 0x30, 0x34,
	0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x77, 0x65, 0x65, 0x6b, 0x30, 0x34, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x2e,
	0x73, 0x79, 0x6b, 0x2e, 0x74, 0x76, 0x2f, 0x77, 0x65, 0x65, 0x6b, 0x30, 0x34, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_account_proto_rawDescOnce sync.Once
	file_v1_account_proto_rawDescData = file_v1_account_proto_rawDesc
)

func file_v1_account_proto_rawDescGZIP() []byte {
	file_v1_account_proto_rawDescOnce.Do(func() {
		file_v1_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_account_proto_rawDescData)
	})
	return file_v1_account_proto_rawDescData
}

var file_v1_account_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_v1_account_proto_goTypes = []interface{}{
	(*CreateUserReq)(nil),  // 0: week04.account.v1.CreateUserReq
	(*CreateUserResp)(nil), // 1: week04.account.v1.CreateUserResp
	(*DelUserReq)(nil),     // 2: week04.account.v1.DelUserReq
	(*DelUserResp)(nil),    // 3: week04.account.v1.DelUserResp
}
var file_v1_account_proto_depIdxs = []int32{
	0, // 0: week04.account.v1.Account.CreateUser:input_type -> week04.account.v1.CreateUserReq
	2, // 1: week04.account.v1.Account.DelUser:input_type -> week04.account.v1.DelUserReq
	1, // 2: week04.account.v1.Account.CreateUser:output_type -> week04.account.v1.CreateUserResp
	3, // 3: week04.account.v1.Account.DelUser:output_type -> week04.account.v1.DelUserResp
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_account_proto_init() }
func file_v1_account_proto_init() {
	if File_v1_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserReq); i {
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
		file_v1_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserResp); i {
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
		file_v1_account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelUserReq); i {
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
		file_v1_account_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelUserResp); i {
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
			RawDescriptor: file_v1_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_account_proto_goTypes,
		DependencyIndexes: file_v1_account_proto_depIdxs,
		MessageInfos:      file_v1_account_proto_msgTypes,
	}.Build()
	File_v1_account_proto = out.File
	file_v1_account_proto_rawDesc = nil
	file_v1_account_proto_goTypes = nil
	file_v1_account_proto_depIdxs = nil
}
