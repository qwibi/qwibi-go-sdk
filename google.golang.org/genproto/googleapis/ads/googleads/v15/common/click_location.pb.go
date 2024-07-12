// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: click_location.proto

package common

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

// Location criteria associated with a click.
type ClickLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The city location criterion associated with the impression.
	City *string `protobuf:"bytes,6,opt,name=city,proto3,oneof" json:"city,omitempty"`
	// The country location criterion associated with the impression.
	Country *string `protobuf:"bytes,7,opt,name=country,proto3,oneof" json:"country,omitempty"`
	// The metro location criterion associated with the impression.
	Metro *string `protobuf:"bytes,8,opt,name=metro,proto3,oneof" json:"metro,omitempty"`
	// The most specific location criterion associated with the impression.
	MostSpecific *string `protobuf:"bytes,9,opt,name=most_specific,json=mostSpecific,proto3,oneof" json:"most_specific,omitempty"`
	// The region location criterion associated with the impression.
	Region *string `protobuf:"bytes,10,opt,name=region,proto3,oneof" json:"region,omitempty"`
}

func (x *ClickLocation) Reset() {
	*x = ClickLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_click_location_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClickLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClickLocation) ProtoMessage() {}

func (x *ClickLocation) ProtoReflect() protoreflect.Message {
	mi := &file_click_location_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClickLocation.ProtoReflect.Descriptor instead.
func (*ClickLocation) Descriptor() ([]byte, []int) {
	return file_click_location_proto_rawDescGZIP(), []int{0}
}

func (x *ClickLocation) GetCity() string {
	if x != nil && x.City != nil {
		return *x.City
	}
	return ""
}

func (x *ClickLocation) GetCountry() string {
	if x != nil && x.Country != nil {
		return *x.Country
	}
	return ""
}

func (x *ClickLocation) GetMetro() string {
	if x != nil && x.Metro != nil {
		return *x.Metro
	}
	return ""
}

func (x *ClickLocation) GetMostSpecific() string {
	if x != nil && x.MostSpecific != nil {
		return *x.MostSpecific
	}
	return ""
}

func (x *ClickLocation) GetRegion() string {
	if x != nil && x.Region != nil {
		return *x.Region
	}
	return ""
}

var File_click_location_proto protoreflect.FileDescriptor

var file_click_location_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x35,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0xe5, 0x01, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x63,
	0x6b, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x04, 0x63, 0x69, 0x74,
	0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x88,
	0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x88, 0x01,
	0x01, 0x12, 0x19, 0x0a, 0x05, 0x6d, 0x65, 0x74, 0x72, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x02, 0x52, 0x05, 0x6d, 0x65, 0x74, 0x72, 0x6f, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x0d,
	0x6d, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x0c, 0x6d, 0x6f, 0x73, 0x74, 0x53, 0x70, 0x65, 0x63, 0x69,
	0x66, 0x69, 0x63, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x63, 0x69, 0x74, 0x79, 0x42, 0x0a, 0x0a, 0x08,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6d, 0x65, 0x74,
	0x72, 0x6f, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x6d, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x70, 0x65, 0x63,
	0x69, 0x66, 0x69, 0x63, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x42,
	0xf2, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x35,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x42, 0x12, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3b, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64,
	0x73, 0x2e, 0x56, 0x31, 0x35, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xca, 0x02, 0x1f, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x35, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xea, 0x02,
	0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x35, 0x3a, 0x3a, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_click_location_proto_rawDescOnce sync.Once
	file_click_location_proto_rawDescData = file_click_location_proto_rawDesc
)

func file_click_location_proto_rawDescGZIP() []byte {
	file_click_location_proto_rawDescOnce.Do(func() {
		file_click_location_proto_rawDescData = protoimpl.X.CompressGZIP(file_click_location_proto_rawDescData)
	})
	return file_click_location_proto_rawDescData
}

var file_click_location_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_click_location_proto_goTypes = []interface{}{
	(*ClickLocation)(nil), // 0: google.ads.googleads.v15.common.ClickLocation
}
var file_click_location_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_click_location_proto_init() }
func file_click_location_proto_init() {
	if File_click_location_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_click_location_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClickLocation); i {
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
	file_click_location_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_click_location_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_click_location_proto_goTypes,
		DependencyIndexes: file_click_location_proto_depIdxs,
		MessageInfos:      file_click_location_proto_msgTypes,
	}.Build()
	File_click_location_proto = out.File
	file_click_location_proto_rawDesc = nil
	file_click_location_proto_goTypes = nil
	file_click_location_proto_depIdxs = nil
}