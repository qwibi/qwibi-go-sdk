// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: geo.proto

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

type QPBxGeoObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gid string `protobuf:"bytes,1,opt,name=gid,proto3" json:"gid,omitempty"`
	// Types that are assignable to Geo:
	//
	//	*QPBxGeoObject_Layer
	//	*QPBxGeoObject_Point
	Geo isQPBxGeoObject_Geo `protobuf_oneof:"geo"`
}

func (x *QPBxGeoObject) Reset() {
	*x = QPBxGeoObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QPBxGeoObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QPBxGeoObject) ProtoMessage() {}

func (x *QPBxGeoObject) ProtoReflect() protoreflect.Message {
	mi := &file_geo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QPBxGeoObject.ProtoReflect.Descriptor instead.
func (*QPBxGeoObject) Descriptor() ([]byte, []int) {
	return file_geo_proto_rawDescGZIP(), []int{0}
}

func (x *QPBxGeoObject) GetGid() string {
	if x != nil {
		return x.Gid
	}
	return ""
}

func (m *QPBxGeoObject) GetGeo() isQPBxGeoObject_Geo {
	if m != nil {
		return m.Geo
	}
	return nil
}

func (x *QPBxGeoObject) GetLayer() *QPBxGeoLayer {
	if x, ok := x.GetGeo().(*QPBxGeoObject_Layer); ok {
		return x.Layer
	}
	return nil
}

func (x *QPBxGeoObject) GetPoint() *QPBxGeoPoint {
	if x, ok := x.GetGeo().(*QPBxGeoObject_Point); ok {
		return x.Point
	}
	return nil
}

type isQPBxGeoObject_Geo interface {
	isQPBxGeoObject_Geo()
}

type QPBxGeoObject_Layer struct {
	Layer *QPBxGeoLayer `protobuf:"bytes,11,opt,name=layer,proto3,oneof"`
}

type QPBxGeoObject_Point struct {
	Point *QPBxGeoPoint `protobuf:"bytes,12,opt,name=point,proto3,oneof"`
}

func (*QPBxGeoObject_Layer) isQPBxGeoObject_Geo() {}

func (*QPBxGeoObject_Point) isQPBxGeoObject_Geo() {}

type QPBxGeoLayer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Feature    *QPBxLayerFeature `protobuf:"bytes,2,opt,name=feature,proto3" json:"feature,omitempty"`
	Properties []byte            `protobuf:"bytes,3,opt,name=properties,proto3" json:"properties,omitempty"`
}

func (x *QPBxGeoLayer) Reset() {
	*x = QPBxGeoLayer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QPBxGeoLayer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QPBxGeoLayer) ProtoMessage() {}

func (x *QPBxGeoLayer) ProtoReflect() protoreflect.Message {
	mi := &file_geo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QPBxGeoLayer.ProtoReflect.Descriptor instead.
func (*QPBxGeoLayer) Descriptor() ([]byte, []int) {
	return file_geo_proto_rawDescGZIP(), []int{1}
}

func (x *QPBxGeoLayer) GetFeature() *QPBxLayerFeature {
	if x != nil {
		return x.Feature
	}
	return nil
}

func (x *QPBxGeoLayer) GetProperties() []byte {
	if x != nil {
		return x.Properties
	}
	return nil
}

type QPBxGeoPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Feature    *QPBxPointFeature `protobuf:"bytes,2,opt,name=feature,proto3" json:"feature,omitempty"`
	Properties []byte            `protobuf:"bytes,3,opt,name=properties,proto3" json:"properties,omitempty"`
}

func (x *QPBxGeoPoint) Reset() {
	*x = QPBxGeoPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QPBxGeoPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QPBxGeoPoint) ProtoMessage() {}

func (x *QPBxGeoPoint) ProtoReflect() protoreflect.Message {
	mi := &file_geo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QPBxGeoPoint.ProtoReflect.Descriptor instead.
func (*QPBxGeoPoint) Descriptor() ([]byte, []int) {
	return file_geo_proto_rawDescGZIP(), []int{2}
}

func (x *QPBxGeoPoint) GetFeature() *QPBxPointFeature {
	if x != nil {
		return x.Feature
	}
	return nil
}

func (x *QPBxGeoPoint) GetProperties() []byte {
	if x != nil {
		return x.Properties
	}
	return nil
}

var File_geo_proto protoreflect.FileDescriptor

var file_geo_proto_rawDesc = []byte{
	0x0a, 0x09, 0x67, 0x65, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x66, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x76, 0x0a, 0x0d, 0x51, 0x50,
	0x42, 0x78, 0x47, 0x65, 0x6f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x67,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x67, 0x69, 0x64, 0x12, 0x25, 0x0a,
	0x05, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x51,
	0x50, 0x42, 0x78, 0x47, 0x65, 0x6f, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x48, 0x00, 0x52, 0x05, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x51, 0x50, 0x42, 0x78, 0x47, 0x65, 0x6f, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x48, 0x00, 0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x42, 0x05, 0x0a, 0x03, 0x67,
	0x65, 0x6f, 0x22, 0x5b, 0x0a, 0x0c, 0x51, 0x50, 0x42, 0x78, 0x47, 0x65, 0x6f, 0x4c, 0x61, 0x79,
	0x65, 0x72, 0x12, 0x2b, 0x0a, 0x07, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x51, 0x50, 0x42, 0x78, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x46,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x07, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x22,
	0x5b, 0x0a, 0x0c, 0x51, 0x50, 0x42, 0x78, 0x47, 0x65, 0x6f, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12,
	0x2b, 0x0a, 0x07, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x51, 0x50, 0x42, 0x78, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x46, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x52, 0x07, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x42, 0x0f, 0x5a, 0x0d,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_geo_proto_rawDescOnce sync.Once
	file_geo_proto_rawDescData = file_geo_proto_rawDesc
)

func file_geo_proto_rawDescGZIP() []byte {
	file_geo_proto_rawDescOnce.Do(func() {
		file_geo_proto_rawDescData = protoimpl.X.CompressGZIP(file_geo_proto_rawDescData)
	})
	return file_geo_proto_rawDescData
}

var file_geo_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_geo_proto_goTypes = []interface{}{
	(*QPBxGeoObject)(nil),    // 0: QPBxGeoObject
	(*QPBxGeoLayer)(nil),     // 1: QPBxGeoLayer
	(*QPBxGeoPoint)(nil),     // 2: QPBxGeoPoint
	(*QPBxLayerFeature)(nil), // 3: QPBxLayerFeature
	(*QPBxPointFeature)(nil), // 4: QPBxPointFeature
}
var file_geo_proto_depIdxs = []int32{
	1, // 0: QPBxGeoObject.layer:type_name -> QPBxGeoLayer
	2, // 1: QPBxGeoObject.point:type_name -> QPBxGeoPoint
	3, // 2: QPBxGeoLayer.feature:type_name -> QPBxLayerFeature
	4, // 3: QPBxGeoPoint.feature:type_name -> QPBxPointFeature
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_geo_proto_init() }
func file_geo_proto_init() {
	if File_geo_proto != nil {
		return
	}
	file_feature_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_geo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QPBxGeoObject); i {
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
		file_geo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QPBxGeoLayer); i {
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
		file_geo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QPBxGeoPoint); i {
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
	file_geo_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*QPBxGeoObject_Layer)(nil),
		(*QPBxGeoObject_Point)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_geo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_geo_proto_goTypes,
		DependencyIndexes: file_geo_proto_depIdxs,
		MessageInfos:      file_geo_proto_msgTypes,
	}.Build()
	File_geo_proto = out.File
	file_geo_proto_rawDesc = nil
	file_geo_proto_goTypes = nil
	file_geo_proto_depIdxs = nil
}
