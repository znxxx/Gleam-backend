// Version

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.2
// source: modules/user/userPb/userPb.proto

package gleam_backend

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

// Structure
type SearchUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *SearchUserReq) Reset() {
	*x = SearchUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_user_userPb_userPb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchUserReq) ProtoMessage() {}

func (x *SearchUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_modules_user_userPb_userPb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchUserReq.ProtoReflect.Descriptor instead.
func (*SearchUserReq) Descriptor() ([]byte, []int) {
	return file_modules_user_userPb_userPb_proto_rawDescGZIP(), []int{0}
}

func (x *SearchUserReq) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type SearchUserRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Valid  bool  `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
}

func (x *SearchUserRes) Reset() {
	*x = SearchUserRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_user_userPb_userPb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchUserRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchUserRes) ProtoMessage() {}

func (x *SearchUserRes) ProtoReflect() protoreflect.Message {
	mi := &file_modules_user_userPb_userPb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchUserRes.ProtoReflect.Descriptor instead.
func (*SearchUserRes) Descriptor() ([]byte, []int) {
	return file_modules_user_userPb_userPb_proto_rawDescGZIP(), []int{1}
}

func (x *SearchUserRes) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *SearchUserRes) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

var File_modules_user_userPb_userPb_proto protoreflect.FileDescriptor

var file_modules_user_userPb_userPb_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x50, 0x62, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x50, 0x62, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x27, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3d, 0x0a, 0x0d, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x32, 0x3f, 0x0a, 0x0f, 0x55, 0x73,
	0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a,
	0x0a, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x42, 0x21, 0x5a, 0x1f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x57, 0x69, 0x6e, 0x2d, 0x54, 0x53,
	0x2f, 0x67, 0x6c, 0x65, 0x61, 0x6d, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_modules_user_userPb_userPb_proto_rawDescOnce sync.Once
	file_modules_user_userPb_userPb_proto_rawDescData = file_modules_user_userPb_userPb_proto_rawDesc
)

func file_modules_user_userPb_userPb_proto_rawDescGZIP() []byte {
	file_modules_user_userPb_userPb_proto_rawDescOnce.Do(func() {
		file_modules_user_userPb_userPb_proto_rawDescData = protoimpl.X.CompressGZIP(file_modules_user_userPb_userPb_proto_rawDescData)
	})
	return file_modules_user_userPb_userPb_proto_rawDescData
}

var file_modules_user_userPb_userPb_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_modules_user_userPb_userPb_proto_goTypes = []interface{}{
	(*SearchUserReq)(nil), // 0: SearchUserReq
	(*SearchUserRes)(nil), // 1: SearchUserRes
}
var file_modules_user_userPb_userPb_proto_depIdxs = []int32{
	0, // 0: UserGrpcService.SearchUser:input_type -> SearchUserReq
	1, // 1: UserGrpcService.SearchUser:output_type -> SearchUserRes
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_modules_user_userPb_userPb_proto_init() }
func file_modules_user_userPb_userPb_proto_init() {
	if File_modules_user_userPb_userPb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_modules_user_userPb_userPb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchUserReq); i {
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
		file_modules_user_userPb_userPb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchUserRes); i {
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
			RawDescriptor: file_modules_user_userPb_userPb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_modules_user_userPb_userPb_proto_goTypes,
		DependencyIndexes: file_modules_user_userPb_userPb_proto_depIdxs,
		MessageInfos:      file_modules_user_userPb_userPb_proto_msgTypes,
	}.Build()
	File_modules_user_userPb_userPb_proto = out.File
	file_modules_user_userPb_userPb_proto_rawDesc = nil
	file_modules_user_userPb_userPb_proto_goTypes = nil
	file_modules_user_userPb_userPb_proto_depIdxs = nil
}
