// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: qwibi/join.proto

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

type QPBxJoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LayerID string `protobuf:"bytes,1,opt,name=layerID,proto3" json:"layerID,omitempty"`
}

func (x *QPBxJoinRequest) Reset() {
	*x = QPBxJoinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qwibi_join_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QPBxJoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QPBxJoinRequest) ProtoMessage() {}

func (x *QPBxJoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qwibi_join_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QPBxJoinRequest.ProtoReflect.Descriptor instead.
func (*QPBxJoinRequest) Descriptor() ([]byte, []int) {
	return file_qwibi_join_proto_rawDescGZIP(), []int{0}
}

func (x *QPBxJoinRequest) GetLayerID() string {
	if x != nil {
		return x.LayerID
	}
	return ""
}

type QPBxJoinResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Layer *QPBxLayer `protobuf:"bytes,1,opt,name=layer,proto3" json:"layer,omitempty"`
}

func (x *QPBxJoinResponse) Reset() {
	*x = QPBxJoinResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qwibi_join_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QPBxJoinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QPBxJoinResponse) ProtoMessage() {}

func (x *QPBxJoinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_qwibi_join_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QPBxJoinResponse.ProtoReflect.Descriptor instead.
func (*QPBxJoinResponse) Descriptor() ([]byte, []int) {
	return file_qwibi_join_proto_rawDescGZIP(), []int{1}
}

func (x *QPBxJoinResponse) GetLayer() *QPBxLayer {
	if x != nil {
		return x.Layer
	}
	return nil
}

var File_qwibi_join_proto protoreflect.FileDescriptor

var file_qwibi_join_proto_rawDesc = []byte{
	0x0a, 0x10, 0x71, 0x77, 0x69, 0x62, 0x69, 0x2f, 0x6a, 0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0f, 0x71, 0x77, 0x69, 0x62, 0x69, 0x2f, 0x67, 0x65, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x0f, 0x51, 0x50, 0x42, 0x78, 0x4a, 0x6f, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44,
	0x22, 0x34, 0x0a, 0x10, 0x51, 0x50, 0x42, 0x78, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x05, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x51, 0x50, 0x42, 0x78, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x52,
	0x05, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_qwibi_join_proto_rawDescOnce sync.Once
	file_qwibi_join_proto_rawDescData = file_qwibi_join_proto_rawDesc
)

func file_qwibi_join_proto_rawDescGZIP() []byte {
	file_qwibi_join_proto_rawDescOnce.Do(func() {
		file_qwibi_join_proto_rawDescData = protoimpl.X.CompressGZIP(file_qwibi_join_proto_rawDescData)
	})
	return file_qwibi_join_proto_rawDescData
}

var file_qwibi_join_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_qwibi_join_proto_goTypes = []interface{}{
	(*QPBxJoinRequest)(nil),  // 0: QPBxJoinRequest
	(*QPBxJoinResponse)(nil), // 1: QPBxJoinResponse
	(*QPBxLayer)(nil),        // 2: QPBxLayer
}
var file_qwibi_join_proto_depIdxs = []int32{
	2, // 0: QPBxJoinResponse.layer:type_name -> QPBxLayer
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_qwibi_join_proto_init() }
func file_qwibi_join_proto_init() {
	if File_qwibi_join_proto != nil {
		return
	}
	file_qwibi_geo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_qwibi_join_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QPBxJoinRequest); i {
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
		file_qwibi_join_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QPBxJoinResponse); i {
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
			RawDescriptor: file_qwibi_join_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_qwibi_join_proto_goTypes,
		DependencyIndexes: file_qwibi_join_proto_depIdxs,
		MessageInfos:      file_qwibi_join_proto_msgTypes,
	}.Build()
	File_qwibi_join_proto = out.File
	file_qwibi_join_proto_rawDesc = nil
	file_qwibi_join_proto_goTypes = nil
	file_qwibi_join_proto_depIdxs = nil
}
