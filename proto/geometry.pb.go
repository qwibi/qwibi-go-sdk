// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.28.2
// source: geometry.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
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

	Gid        string           `protobuf:"bytes,1,opt,name=gid,proto3" json:"gid,omitempty"`
	Features   []*QPBxFeature   `protobuf:"bytes,2,rep,name=features,proto3" json:"features,omitempty"`
	Properties *structpb.Struct `protobuf:"bytes,3,opt,name=Properties,proto3" json:"Properties,omitempty"`
}

func (x *QPBxGeoObject) Reset() {
	*x = QPBxGeoObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geometry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QPBxGeoObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QPBxGeoObject) ProtoMessage() {}

func (x *QPBxGeoObject) ProtoReflect() protoreflect.Message {
	mi := &file_geometry_proto_msgTypes[0]
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
	return file_geometry_proto_rawDescGZIP(), []int{0}
}

func (x *QPBxGeoObject) GetGid() string {
	if x != nil {
		return x.Gid
	}
	return ""
}

func (x *QPBxGeoObject) GetFeatures() []*QPBxFeature {
	if x != nil {
		return x.Features
	}
	return nil
}

func (x *QPBxGeoObject) GetProperties() *structpb.Struct {
	if x != nil {
		return x.Properties
	}
	return nil
}

type QPBxFeature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//
	//	*QPBxFeature_Point
	//	*QPBxFeature_Polyline
	Type isQPBxFeature_Type `protobuf_oneof:"type"`
}

func (x *QPBxFeature) Reset() {
	*x = QPBxFeature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geometry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QPBxFeature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QPBxFeature) ProtoMessage() {}

func (x *QPBxFeature) ProtoReflect() protoreflect.Message {
	mi := &file_geometry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QPBxFeature.ProtoReflect.Descriptor instead.
func (*QPBxFeature) Descriptor() ([]byte, []int) {
	return file_geometry_proto_rawDescGZIP(), []int{1}
}

func (m *QPBxFeature) GetType() isQPBxFeature_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *QPBxFeature) GetPoint() *QPBxPointFeature {
	if x, ok := x.GetType().(*QPBxFeature_Point); ok {
		return x.Point
	}
	return nil
}

func (x *QPBxFeature) GetPolyline() *QPBxPolylineFeature {
	if x, ok := x.GetType().(*QPBxFeature_Polyline); ok {
		return x.Polyline
	}
	return nil
}

type isQPBxFeature_Type interface {
	isQPBxFeature_Type()
}

type QPBxFeature_Point struct {
	Point *QPBxPointFeature `protobuf:"bytes,1,opt,name=point,proto3,oneof"`
}

type QPBxFeature_Polyline struct {
	Polyline *QPBxPolylineFeature `protobuf:"bytes,2,opt,name=polyline,proto3,oneof"`
}

func (*QPBxFeature_Point) isQPBxFeature_Type() {}

func (*QPBxFeature_Polyline) isQPBxFeature_Type() {}

type QPBxPointFeature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fid        string                       `protobuf:"bytes,1,opt,name=fid,proto3" json:"fid,omitempty"`
	Geometry   *QPBxPoint                   `protobuf:"bytes,2,opt,name=geometry,proto3" json:"geometry,omitempty"`
	Properties *QPBxPointFeature_Properties `protobuf:"bytes,3,opt,name=properties,proto3" json:"properties,omitempty"`
}

func (x *QPBxPointFeature) Reset() {
	*x = QPBxPointFeature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geometry_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QPBxPointFeature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QPBxPointFeature) ProtoMessage() {}

func (x *QPBxPointFeature) ProtoReflect() protoreflect.Message {
	mi := &file_geometry_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QPBxPointFeature.ProtoReflect.Descriptor instead.
func (*QPBxPointFeature) Descriptor() ([]byte, []int) {
	return file_geometry_proto_rawDescGZIP(), []int{2}
}

func (x *QPBxPointFeature) GetFid() string {
	if x != nil {
		return x.Fid
	}
	return ""
}

func (x *QPBxPointFeature) GetGeometry() *QPBxPoint {
	if x != nil {
		return x.Geometry
	}
	return nil
}

func (x *QPBxPointFeature) GetProperties() *QPBxPointFeature_Properties {
	if x != nil {
		return x.Properties
	}
	return nil
}

type QPBxPolylineFeature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fid        string                          `protobuf:"bytes,1,opt,name=fid,proto3" json:"fid,omitempty"`
	Polyline   []*QPBxPolyLine                 `protobuf:"bytes,2,rep,name=polyline,proto3" json:"polyline,omitempty"`
	Properties *QPBxPolylineFeature_Properties `protobuf:"bytes,3,opt,name=properties,proto3" json:"properties,omitempty"`
}

func (x *QPBxPolylineFeature) Reset() {
	*x = QPBxPolylineFeature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geometry_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QPBxPolylineFeature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QPBxPolylineFeature) ProtoMessage() {}

func (x *QPBxPolylineFeature) ProtoReflect() protoreflect.Message {
	mi := &file_geometry_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QPBxPolylineFeature.ProtoReflect.Descriptor instead.
func (*QPBxPolylineFeature) Descriptor() ([]byte, []int) {
	return file_geometry_proto_rawDescGZIP(), []int{3}
}

func (x *QPBxPolylineFeature) GetFid() string {
	if x != nil {
		return x.Fid
	}
	return ""
}

func (x *QPBxPolylineFeature) GetPolyline() []*QPBxPolyLine {
	if x != nil {
		return x.Polyline
	}
	return nil
}

func (x *QPBxPolylineFeature) GetProperties() *QPBxPolylineFeature_Properties {
	if x != nil {
		return x.Properties
	}
	return nil
}

type QPBxGeometry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//
	//	*QPBxGeometry_Point
	//	*QPBxGeometry_Polyline
	Type isQPBxGeometry_Type `protobuf_oneof:"type"`
}

func (x *QPBxGeometry) Reset() {
	*x = QPBxGeometry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geometry_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QPBxGeometry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QPBxGeometry) ProtoMessage() {}

func (x *QPBxGeometry) ProtoReflect() protoreflect.Message {
	mi := &file_geometry_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QPBxGeometry.ProtoReflect.Descriptor instead.
func (*QPBxGeometry) Descriptor() ([]byte, []int) {
	return file_geometry_proto_rawDescGZIP(), []int{4}
}

func (m *QPBxGeometry) GetType() isQPBxGeometry_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *QPBxGeometry) GetPoint() *QPBxPoint {
	if x, ok := x.GetType().(*QPBxGeometry_Point); ok {
		return x.Point
	}
	return nil
}

func (x *QPBxGeometry) GetPolyline() *QPBxPolyLine {
	if x, ok := x.GetType().(*QPBxGeometry_Polyline); ok {
		return x.Polyline
	}
	return nil
}

type isQPBxGeometry_Type interface {
	isQPBxGeometry_Type()
}

type QPBxGeometry_Point struct {
	Point *QPBxPoint `protobuf:"bytes,1,opt,name=point,proto3,oneof"`
}

type QPBxGeometry_Polyline struct {
	Polyline *QPBxPolyLine `protobuf:"bytes,2,opt,name=polyline,proto3,oneof"`
}

func (*QPBxGeometry_Point) isQPBxGeometry_Type() {}

func (*QPBxGeometry_Polyline) isQPBxGeometry_Type() {}

type QPBxPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Coordinates []float64 `protobuf:"fixed64,1,rep,packed,name=coordinates,proto3" json:"coordinates,omitempty"`
}

func (x *QPBxPoint) Reset() {
	*x = QPBxPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geometry_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QPBxPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QPBxPoint) ProtoMessage() {}

func (x *QPBxPoint) ProtoReflect() protoreflect.Message {
	mi := &file_geometry_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QPBxPoint.ProtoReflect.Descriptor instead.
func (*QPBxPoint) Descriptor() ([]byte, []int) {
	return file_geometry_proto_rawDescGZIP(), []int{5}
}

func (x *QPBxPoint) GetCoordinates() []float64 {
	if x != nil {
		return x.Coordinates
	}
	return nil
}

type QPBxPolyLine struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Polyline []*QPBxPoint `protobuf:"bytes,1,rep,name=polyline,proto3" json:"polyline,omitempty"`
}

func (x *QPBxPolyLine) Reset() {
	*x = QPBxPolyLine{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geometry_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QPBxPolyLine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QPBxPolyLine) ProtoMessage() {}

func (x *QPBxPolyLine) ProtoReflect() protoreflect.Message {
	mi := &file_geometry_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QPBxPolyLine.ProtoReflect.Descriptor instead.
func (*QPBxPolyLine) Descriptor() ([]byte, []int) {
	return file_geometry_proto_rawDescGZIP(), []int{6}
}

func (x *QPBxPolyLine) GetPolyline() []*QPBxPoint {
	if x != nil {
		return x.Polyline
	}
	return nil
}

type QPBxPolygon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Polygon []*QPBxPolyLine `protobuf:"bytes,1,rep,name=polygon,proto3" json:"polygon,omitempty"`
}

func (x *QPBxPolygon) Reset() {
	*x = QPBxPolygon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geometry_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QPBxPolygon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QPBxPolygon) ProtoMessage() {}

func (x *QPBxPolygon) ProtoReflect() protoreflect.Message {
	mi := &file_geometry_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QPBxPolygon.ProtoReflect.Descriptor instead.
func (*QPBxPolygon) Descriptor() ([]byte, []int) {
	return file_geometry_proto_rawDescGZIP(), []int{7}
}

func (x *QPBxPolygon) GetPolygon() []*QPBxPolyLine {
	if x != nil {
		return x.Polygon
	}
	return nil
}

type QPBxPointFeature_Properties struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageUri     *string  `protobuf:"bytes,1,opt,name=image_uri,json=imageUri,proto3,oneof" json:"image_uri,omitempty"`
	ImageData    []byte   `protobuf:"bytes,2,opt,name=image_data,json=imageData,proto3,oneof" json:"image_data,omitempty"`
	ImageScale   *float64 `protobuf:"fixed64,3,opt,name=image_scale,json=imageScale,proto3,oneof" json:"image_scale,omitempty"`
	ImageAnchor  *string  `protobuf:"bytes,4,opt,name=image_anchor,json=imageAnchor,proto3,oneof" json:"image_anchor,omitempty"`
	ImageInverse *bool    `protobuf:"varint,5,opt,name=image_inverse,json=imageInverse,proto3,oneof" json:"image_inverse,omitempty"`
}

func (x *QPBxPointFeature_Properties) Reset() {
	*x = QPBxPointFeature_Properties{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geometry_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QPBxPointFeature_Properties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QPBxPointFeature_Properties) ProtoMessage() {}

func (x *QPBxPointFeature_Properties) ProtoReflect() protoreflect.Message {
	mi := &file_geometry_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QPBxPointFeature_Properties.ProtoReflect.Descriptor instead.
func (*QPBxPointFeature_Properties) Descriptor() ([]byte, []int) {
	return file_geometry_proto_rawDescGZIP(), []int{2, 0}
}

func (x *QPBxPointFeature_Properties) GetImageUri() string {
	if x != nil && x.ImageUri != nil {
		return *x.ImageUri
	}
	return ""
}

func (x *QPBxPointFeature_Properties) GetImageData() []byte {
	if x != nil {
		return x.ImageData
	}
	return nil
}

func (x *QPBxPointFeature_Properties) GetImageScale() float64 {
	if x != nil && x.ImageScale != nil {
		return *x.ImageScale
	}
	return 0
}

func (x *QPBxPointFeature_Properties) GetImageAnchor() string {
	if x != nil && x.ImageAnchor != nil {
		return *x.ImageAnchor
	}
	return ""
}

func (x *QPBxPointFeature_Properties) GetImageInverse() bool {
	if x != nil && x.ImageInverse != nil {
		return *x.ImageInverse
	}
	return false
}

type QPBxPolylineFeature_Properties struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LineWidth float64 `protobuf:"fixed64,1,opt,name=line_width,json=lineWidth,proto3" json:"line_width,omitempty"`
}

func (x *QPBxPolylineFeature_Properties) Reset() {
	*x = QPBxPolylineFeature_Properties{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geometry_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QPBxPolylineFeature_Properties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QPBxPolylineFeature_Properties) ProtoMessage() {}

func (x *QPBxPolylineFeature_Properties) ProtoReflect() protoreflect.Message {
	mi := &file_geometry_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QPBxPolylineFeature_Properties.ProtoReflect.Descriptor instead.
func (*QPBxPolylineFeature_Properties) Descriptor() ([]byte, []int) {
	return file_geometry_proto_rawDescGZIP(), []int{3, 0}
}

func (x *QPBxPolylineFeature_Properties) GetLineWidth() float64 {
	if x != nil {
		return x.LineWidth
	}
	return 0
}

var File_geometry_proto protoreflect.FileDescriptor

var file_geometry_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x67, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84,
	0x01, 0x0a, 0x0d, 0x51, 0x50, 0x42, 0x78, 0x47, 0x65, 0x6f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x67, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x67,
	0x69, 0x64, 0x12, 0x28, 0x0a, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x51, 0x50, 0x42, 0x78, 0x46, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x52, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x37, 0x0a, 0x0a,
	0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0a, 0x50, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x69, 0x65, 0x73, 0x22, 0x74, 0x0a, 0x0b, 0x51, 0x50, 0x42, 0x78, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x51, 0x50, 0x42, 0x78, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x46,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x48, 0x00, 0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12,
	0x32, 0x0a, 0x08, 0x70, 0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x51, 0x50, 0x42, 0x78, 0x50, 0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65,
	0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x48, 0x00, 0x52, 0x08, 0x70, 0x6f, 0x6c, 0x79, 0x6c,
	0x69, 0x6e, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0xa7, 0x03, 0x0a, 0x10,
	0x51, 0x50, 0x42, 0x78, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x66, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x66,
	0x69, 0x64, 0x12, 0x26, 0x0a, 0x08, 0x67, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x51, 0x50, 0x42, 0x78, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x52, 0x08, 0x67, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x12, 0x3c, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x51, 0x50, 0x42, 0x78, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x52, 0x0a, 0x70, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x1a, 0x9a, 0x02, 0x0a, 0x0a, 0x50, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x5f, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x55, 0x72, 0x69, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x01, 0x52,
	0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a,
	0x0b, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x01, 0x48, 0x02, 0x52, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x63, 0x61, 0x6c, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x6e, 0x63,
	0x68, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x0b, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x0d, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x48, 0x04, 0x52, 0x0c, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x76, 0x65, 0x72,
	0x73, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f,
	0x75, 0x72, 0x69, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x63, 0x61,
	0x6c, 0x65, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x6e, 0x63,
	0x68, 0x6f, 0x72, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x6e,
	0x76, 0x65, 0x72, 0x73, 0x65, 0x22, 0xc0, 0x01, 0x0a, 0x13, 0x51, 0x50, 0x42, 0x78, 0x50, 0x6f,
	0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x66, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x66, 0x69, 0x64, 0x12,
	0x29, 0x0a, 0x08, 0x70, 0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x51, 0x50, 0x42, 0x78, 0x50, 0x6f, 0x6c, 0x79, 0x4c, 0x69, 0x6e, 0x65,
	0x52, 0x08, 0x70, 0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x51, 0x50, 0x42, 0x78, 0x50, 0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x52,
	0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x1a, 0x2b, 0x0a, 0x0a, 0x50,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x69, 0x6e,
	0x65, 0x5f, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c,
	0x69, 0x6e, 0x65, 0x57, 0x69, 0x64, 0x74, 0x68, 0x22, 0x67, 0x0a, 0x0c, 0x51, 0x50, 0x42, 0x78,
	0x47, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x12, 0x22, 0x0a, 0x05, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x51, 0x50, 0x42, 0x78, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x08,
	0x70, 0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x51, 0x50, 0x42, 0x78, 0x50, 0x6f, 0x6c, 0x79, 0x4c, 0x69, 0x6e, 0x65, 0x48, 0x00, 0x52,
	0x08, 0x70, 0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x22, 0x2d, 0x0a, 0x09, 0x51, 0x50, 0x42, 0x78, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x20,
	0x0a, 0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x01, 0x52, 0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73,
	0x22, 0x36, 0x0a, 0x0c, 0x51, 0x50, 0x42, 0x78, 0x50, 0x6f, 0x6c, 0x79, 0x4c, 0x69, 0x6e, 0x65,
	0x12, 0x26, 0x0a, 0x08, 0x70, 0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x51, 0x50, 0x42, 0x78, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x08,
	0x70, 0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0x36, 0x0a, 0x0b, 0x51, 0x50, 0x42, 0x78,
	0x50, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x07, 0x70, 0x6f, 0x6c, 0x79, 0x67,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x51, 0x50, 0x42, 0x78, 0x50,
	0x6f, 0x6c, 0x79, 0x4c, 0x69, 0x6e, 0x65, 0x52, 0x07, 0x70, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e,
	0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_geometry_proto_rawDescOnce sync.Once
	file_geometry_proto_rawDescData = file_geometry_proto_rawDesc
)

func file_geometry_proto_rawDescGZIP() []byte {
	file_geometry_proto_rawDescOnce.Do(func() {
		file_geometry_proto_rawDescData = protoimpl.X.CompressGZIP(file_geometry_proto_rawDescData)
	})
	return file_geometry_proto_rawDescData
}

var file_geometry_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_geometry_proto_goTypes = []interface{}{
	(*QPBxGeoObject)(nil),                  // 0: QPBxGeoObject
	(*QPBxFeature)(nil),                    // 1: QPBxFeature
	(*QPBxPointFeature)(nil),               // 2: QPBxPointFeature
	(*QPBxPolylineFeature)(nil),            // 3: QPBxPolylineFeature
	(*QPBxGeometry)(nil),                   // 4: QPBxGeometry
	(*QPBxPoint)(nil),                      // 5: QPBxPoint
	(*QPBxPolyLine)(nil),                   // 6: QPBxPolyLine
	(*QPBxPolygon)(nil),                    // 7: QPBxPolygon
	(*QPBxPointFeature_Properties)(nil),    // 8: QPBxPointFeature.Properties
	(*QPBxPolylineFeature_Properties)(nil), // 9: QPBxPolylineFeature.Properties
	(*structpb.Struct)(nil),                // 10: google.protobuf.Struct
}
var file_geometry_proto_depIdxs = []int32{
	1,  // 0: QPBxGeoObject.features:type_name -> QPBxFeature
	10, // 1: QPBxGeoObject.Properties:type_name -> google.protobuf.Struct
	2,  // 2: QPBxFeature.point:type_name -> QPBxPointFeature
	3,  // 3: QPBxFeature.polyline:type_name -> QPBxPolylineFeature
	5,  // 4: QPBxPointFeature.geometry:type_name -> QPBxPoint
	8,  // 5: QPBxPointFeature.properties:type_name -> QPBxPointFeature.Properties
	6,  // 6: QPBxPolylineFeature.polyline:type_name -> QPBxPolyLine
	9,  // 7: QPBxPolylineFeature.properties:type_name -> QPBxPolylineFeature.Properties
	5,  // 8: QPBxGeometry.point:type_name -> QPBxPoint
	6,  // 9: QPBxGeometry.polyline:type_name -> QPBxPolyLine
	5,  // 10: QPBxPolyLine.polyline:type_name -> QPBxPoint
	6,  // 11: QPBxPolygon.polygon:type_name -> QPBxPolyLine
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_geometry_proto_init() }
func file_geometry_proto_init() {
	if File_geometry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_geometry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_geometry_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QPBxFeature); i {
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
		file_geometry_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QPBxPointFeature); i {
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
		file_geometry_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QPBxPolylineFeature); i {
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
		file_geometry_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QPBxGeometry); i {
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
		file_geometry_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QPBxPoint); i {
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
		file_geometry_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QPBxPolyLine); i {
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
		file_geometry_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QPBxPolygon); i {
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
		file_geometry_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QPBxPointFeature_Properties); i {
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
		file_geometry_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QPBxPolylineFeature_Properties); i {
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
	file_geometry_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*QPBxFeature_Point)(nil),
		(*QPBxFeature_Polyline)(nil),
	}
	file_geometry_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*QPBxGeometry_Point)(nil),
		(*QPBxGeometry_Polyline)(nil),
	}
	file_geometry_proto_msgTypes[8].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_geometry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_geometry_proto_goTypes,
		DependencyIndexes: file_geometry_proto_depIdxs,
		MessageInfos:      file_geometry_proto_msgTypes,
	}.Build()
	File_geometry_proto = out.File
	file_geometry_proto_rawDesc = nil
	file_geometry_proto_goTypes = nil
	file_geometry_proto_depIdxs = nil
}
