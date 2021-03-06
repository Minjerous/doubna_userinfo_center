// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: user_center.proto

package proto

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

//密码注册
type LoginByPWResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OK bool `protobuf:"varint,1,opt,name=OK,proto3" json:"OK,omitempty"`
}

func (x *LoginByPWResp) Reset() {
	*x = LoginByPWResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_center_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginByPWResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginByPWResp) ProtoMessage() {}

func (x *LoginByPWResp) ProtoReflect() protoreflect.Message {
	mi := &file_user_center_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginByPWResp.ProtoReflect.Descriptor instead.
func (*LoginByPWResp) Descriptor() ([]byte, []int) {
	return file_user_center_proto_rawDescGZIP(), []int{0}
}

func (x *LoginByPWResp) GetOK() bool {
	if x != nil {
		return x.OK
	}
	return false
}

type LoginByPWReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LoginAccount string `protobuf:"bytes,1,opt,name=loginAccount,proto3" json:"loginAccount,omitempty"`
	PassWord     string `protobuf:"bytes,2,opt,name=PassWord,proto3" json:"PassWord,omitempty"`
}

func (x *LoginByPWReq) Reset() {
	*x = LoginByPWReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_center_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginByPWReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginByPWReq) ProtoMessage() {}

func (x *LoginByPWReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_center_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginByPWReq.ProtoReflect.Descriptor instead.
func (*LoginByPWReq) Descriptor() ([]byte, []int) {
	return file_user_center_proto_rawDescGZIP(), []int{1}
}

func (x *LoginByPWReq) GetLoginAccount() string {
	if x != nil {
		return x.LoginAccount
	}
	return ""
}

func (x *LoginByPWReq) GetPassWord() string {
	if x != nil {
		return x.PassWord
	}
	return ""
}

//短信登录
type LoginBySmsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone string `protobuf:"bytes,1,opt,name=Phone,proto3" json:"Phone,omitempty"`
	Code  string `protobuf:"bytes,2,opt,name=Code,proto3" json:"Code,omitempty"`
}

func (x *LoginBySmsReq) Reset() {
	*x = LoginBySmsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_center_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginBySmsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginBySmsReq) ProtoMessage() {}

func (x *LoginBySmsReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_center_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginBySmsReq.ProtoReflect.Descriptor instead.
func (*LoginBySmsReq) Descriptor() ([]byte, []int) {
	return file_user_center_proto_rawDescGZIP(), []int{2}
}

func (x *LoginBySmsReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *LoginBySmsReq) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

//注册
type RegisterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	Pwd      string `protobuf:"bytes,2,opt,name=Pwd,proto3" json:"Pwd,omitempty"`
	Phone    string `protobuf:"bytes,3,opt,name=Phone,proto3" json:"Phone,omitempty"`
}

func (x *RegisterReq) Reset() {
	*x = RegisterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_center_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReq) ProtoMessage() {}

func (x *RegisterReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_center_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReq.ProtoReflect.Descriptor instead.
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return file_user_center_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RegisterReq) GetPwd() string {
	if x != nil {
		return x.Pwd
	}
	return ""
}

func (x *RegisterReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type RegisterResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OK bool `protobuf:"varint,1,opt,name=OK,proto3" json:"OK,omitempty"`
}

func (x *RegisterResp) Reset() {
	*x = RegisterResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_center_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResp) ProtoMessage() {}

func (x *RegisterResp) ProtoReflect() protoreflect.Message {
	mi := &file_user_center_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResp.ProtoReflect.Descriptor instead.
func (*RegisterResp) Descriptor() ([]byte, []int) {
	return file_user_center_proto_rawDescGZIP(), []int{4}
}

func (x *RegisterResp) GetOK() bool {
	if x != nil {
		return x.OK
	}
	return false
}

var File_user_center_proto protoreflect.FileDescriptor

var file_user_center_proto_rawDesc = []byte{
	0x0a, 0x11, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x67, 0x72, 0x70, 0x63, 0x69, 0x6e, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x22, 0x1f, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79, 0x50, 0x57, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x0e, 0x0a, 0x02, 0x4f, 0x4b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x4f,
	0x4b, 0x22, 0x4e, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79, 0x50, 0x57, 0x52, 0x65,
	0x71, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x57, 0x6f, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x57, 0x6f, 0x72,
	0x64, 0x22, 0x39, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79, 0x53, 0x6d, 0x73, 0x52,
	0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x51, 0x0a, 0x0b,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x50, 0x77, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x50, 0x77, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x22,
	0x1e, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x0e, 0x0a, 0x02, 0x4f, 0x4b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x4f, 0x4b, 0x32,
	0xdc, 0x01, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x44, 0x0a,
	0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79, 0x50, 0x57, 0x12, 0x19, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x69, 0x6e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79,
	0x50, 0x57, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x69, 0x6e, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79, 0x50, 0x57, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x59, 0x53, 0x4d,
	0x73, 0x12, 0x1a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x69, 0x6e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2e,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79, 0x53, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x69, 0x6e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x42, 0x79, 0x53, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x08, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x69, 0x6e,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x1a, 0x19, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x69, 0x6e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x08,
	0x5a, 0x06, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_center_proto_rawDescOnce sync.Once
	file_user_center_proto_rawDescData = file_user_center_proto_rawDesc
)

func file_user_center_proto_rawDescGZIP() []byte {
	file_user_center_proto_rawDescOnce.Do(func() {
		file_user_center_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_center_proto_rawDescData)
	})
	return file_user_center_proto_rawDescData
}

var file_user_center_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_user_center_proto_goTypes = []interface{}{
	(*LoginByPWResp)(nil), // 0: grpcinclass.LoginByPWResp
	(*LoginByPWReq)(nil),  // 1: grpcinclass.LoginByPWReq
	(*LoginBySmsReq)(nil), // 2: grpcinclass.LoginBySmsReq
	(*RegisterReq)(nil),   // 3: grpcinclass.RegisterReq
	(*RegisterResp)(nil),  // 4: grpcinclass.RegisterResp
}
var file_user_center_proto_depIdxs = []int32{
	1, // 0: grpcinclass.UseCenter.LoginByPW:input_type -> grpcinclass.LoginByPWReq
	2, // 1: grpcinclass.UseCenter.LoginBYSMs:input_type -> grpcinclass.LoginBySmsReq
	3, // 2: grpcinclass.UseCenter.Register:input_type -> grpcinclass.RegisterReq
	0, // 3: grpcinclass.UseCenter.LoginByPW:output_type -> grpcinclass.LoginByPWResp
	2, // 4: grpcinclass.UseCenter.LoginBYSMs:output_type -> grpcinclass.LoginBySmsReq
	4, // 5: grpcinclass.UseCenter.Register:output_type -> grpcinclass.RegisterResp
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_center_proto_init() }
func file_user_center_proto_init() {
	if File_user_center_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_center_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginByPWResp); i {
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
		file_user_center_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginByPWReq); i {
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
		file_user_center_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginBySmsReq); i {
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
		file_user_center_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterReq); i {
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
		file_user_center_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResp); i {
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
			RawDescriptor: file_user_center_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_center_proto_goTypes,
		DependencyIndexes: file_user_center_proto_depIdxs,
		MessageInfos:      file_user_center_proto_msgTypes,
	}.Build()
	File_user_center_proto = out.File
	file_user_center_proto_rawDesc = nil
	file_user_center_proto_goTypes = nil
	file_user_center_proto_depIdxs = nil
}
