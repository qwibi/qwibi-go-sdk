// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: layer.proto

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

type QPBxLayer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LayerId    string            `protobuf:"bytes,1,opt,name=layer_id,json=layerId,proto3" json:"layer_id,omitempty"`
	Name       string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Public     bool              `protobuf:"varint,3,opt,name=public,proto3" json:"public,omitempty"`
	Properties []byte            `protobuf:"bytes,4,opt,name=properties,proto3" json:"properties,omitempty"`
	Commands   map[string]string `protobuf:"bytes,5,rep,name=commands,proto3" json:"commands,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *QPBxLayer) Reset() {
	*x = QPBxLayer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_layer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QPBxLayer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QPBxLayer) ProtoMessage() {}

func (x *QPBxLayer) ProtoReflect() protoreflect.Message {
	mi := &file_layer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QPBxLayer.ProtoReflect.Descriptor instead.
func (*QPBxLayer) Descriptor() ([]byte, []int) {
	return file_layer_proto_rawDescGZIP(), []int{0}
}

func (x *QPBxLayer) GetLayerId() string {
	if x != nil {
		return x.LayerId
	}
	return ""
}

func (x *QPBxLayer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *QPBxLayer) GetPublic() bool {
	if x != nil {
		return x.Public
	}
	return false
}

func (x *QPBxLayer) GetProperties() []byte {
	if x != nil {
		return x.Properties
	}
	return nil
}

func (x *QPBxLayer) GetCommands() map[string]string {
	if x != nil {
		return x.Commands
	}
	return nil
}

var File_layer_proto protoreflect.FileDescriptor

var file_layer_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe5, 0x01,
	0x0a, 0x09, 0x51, 0x50, 0x42, 0x78, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69,
	0x65, 0x73, 0x12, 0x34, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x51, 0x50, 0x42, 0x78, 0x4c, 0x61, 0x79, 0x65, 0x72,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x1a, 0x3b, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_layer_proto_rawDescOnce sync.Once
	file_layer_proto_rawDescData = file_layer_proto_rawDesc
)

func file_layer_proto_rawDescGZIP() []byte {
	file_layer_proto_rawDescOnce.Do(func() {
		file_layer_proto_rawDescData = protoimpl.X.CompressGZIP(file_layer_proto_rawDescData)
	})
	return file_layer_proto_rawDescData
}

var file_layer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_layer_proto_goTypes = []interface{}{
	(*QPBxLayer)(nil), // 0: QPBxLayer
	nil,               // 1: QPBxLayer.CommandsEntry
}
var file_layer_proto_depIdxs = []int32{
	1, // 0: QPBxLayer.commands:type_name -> QPBxLayer.CommandsEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_layer_proto_init() }
func file_layer_proto_init() {
	if File_layer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_layer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QPBxLayer); i {
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
			RawDescriptor: file_layer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_layer_proto_goTypes,
		DependencyIndexes: file_layer_proto_depIdxs,
		MessageInfos:      file_layer_proto_msgTypes,
	}.Build()
	File_layer_proto = out.File
	file_layer_proto_rawDesc = nil
	file_layer_proto_goTypes = nil
	file_layer_proto_depIdxs = nil
}
