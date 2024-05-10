// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: grpcUserServer.proto

package grpcUserServer

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

type GetProfileDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserUUID        string `protobuf:"bytes,1,opt,name=UserUUID,proto3" json:"UserUUID,omitempty"`
	BlockedUserUUID string `protobuf:"bytes,2,opt,name=BlockedUserUUID,proto3" json:"BlockedUserUUID,omitempty"`
	Type            string `protobuf:"bytes,3,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (x *GetProfileDataRequest) Reset() {
	*x = GetProfileDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcUserServer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProfileDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfileDataRequest) ProtoMessage() {}

func (x *GetProfileDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpcUserServer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfileDataRequest.ProtoReflect.Descriptor instead.
func (*GetProfileDataRequest) Descriptor() ([]byte, []int) {
	return file_grpcUserServer_proto_rawDescGZIP(), []int{0}
}

func (x *GetProfileDataRequest) GetUserUUID() string {
	if x != nil {
		return x.UserUUID
	}
	return ""
}

func (x *GetProfileDataRequest) GetBlockedUserUUID() string {
	if x != nil {
		return x.BlockedUserUUID
	}
	return ""
}

func (x *GetProfileDataRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type GetProfileDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserUUID        string `protobuf:"bytes,1,opt,name=UserUUID,proto3" json:"UserUUID,omitempty"`
	BlockedUserUUID string `protobuf:"bytes,2,opt,name=BlockedUserUUID,proto3" json:"BlockedUserUUID,omitempty"`
	Type            string `protobuf:"bytes,3,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (x *GetProfileDataResponse) Reset() {
	*x = GetProfileDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcUserServer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProfileDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfileDataResponse) ProtoMessage() {}

func (x *GetProfileDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpcUserServer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfileDataResponse.ProtoReflect.Descriptor instead.
func (*GetProfileDataResponse) Descriptor() ([]byte, []int) {
	return file_grpcUserServer_proto_rawDescGZIP(), []int{1}
}

func (x *GetProfileDataResponse) GetUserUUID() string {
	if x != nil {
		return x.UserUUID
	}
	return ""
}

func (x *GetProfileDataResponse) GetBlockedUserUUID() string {
	if x != nil {
		return x.BlockedUserUUID
	}
	return ""
}

func (x *GetProfileDataResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

var File_grpcUserServer_proto protoreflect.FileDescriptor

var file_grpcUserServer_proto_rawDesc = []byte{
	0x0a, 0x14, 0x67, 0x72, 0x70, 0x63, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x67, 0x72, 0x70, 0x63, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x71, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x12, 0x28, 0x0a, 0x0f, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x55, 0x73, 0x65,
	0x72, 0x55, 0x55, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x22, 0x72, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x12,
	0x28, 0x0a, 0x0f, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x55, 0x55,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65,
	0x64, 0x55, 0x73, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x32, 0x74, 0x0a,
	0x11, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x5f, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x25, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x47, 0x75, 0x72, 0x62, 0x61, 0x6e, 0x6f, 0x77, 0x47, 0x61, 0x64, 0x61, 0x6d, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpcUserServer_proto_rawDescOnce sync.Once
	file_grpcUserServer_proto_rawDescData = file_grpcUserServer_proto_rawDesc
)

func file_grpcUserServer_proto_rawDescGZIP() []byte {
	file_grpcUserServer_proto_rawDescOnce.Do(func() {
		file_grpcUserServer_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpcUserServer_proto_rawDescData)
	})
	return file_grpcUserServer_proto_rawDescData
}

var file_grpcUserServer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_grpcUserServer_proto_goTypes = []interface{}{
	(*GetProfileDataRequest)(nil),  // 0: grpcUserServer.GetProfileDataRequest
	(*GetProfileDataResponse)(nil), // 1: grpcUserServer.GetProfileDataResponse
}
var file_grpcUserServer_proto_depIdxs = []int32{
	0, // 0: grpcUserServer.userServerService.GetProfileData:input_type -> grpcUserServer.GetProfileDataRequest
	1, // 1: grpcUserServer.userServerService.GetProfileData:output_type -> grpcUserServer.GetProfileDataResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpcUserServer_proto_init() }
func file_grpcUserServer_proto_init() {
	if File_grpcUserServer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpcUserServer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProfileDataRequest); i {
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
		file_grpcUserServer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProfileDataResponse); i {
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
			RawDescriptor: file_grpcUserServer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpcUserServer_proto_goTypes,
		DependencyIndexes: file_grpcUserServer_proto_depIdxs,
		MessageInfos:      file_grpcUserServer_proto_msgTypes,
	}.Build()
	File_grpcUserServer_proto = out.File
	file_grpcUserServer_proto_rawDesc = nil
	file_grpcUserServer_proto_goTypes = nil
	file_grpcUserServer_proto_depIdxs = nil
}