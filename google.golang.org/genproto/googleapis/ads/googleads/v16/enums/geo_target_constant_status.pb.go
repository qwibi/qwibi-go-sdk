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
// source: geo_target_constant_status.proto

package enums

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

// The possible statuses of a geo target constant.
type GeoTargetConstantStatusEnum_GeoTargetConstantStatus int32

const (
	// No value has been specified.
	GeoTargetConstantStatusEnum_UNSPECIFIED GeoTargetConstantStatusEnum_GeoTargetConstantStatus = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	GeoTargetConstantStatusEnum_UNKNOWN GeoTargetConstantStatusEnum_GeoTargetConstantStatus = 1
	// The geo target constant is valid.
	GeoTargetConstantStatusEnum_ENABLED GeoTargetConstantStatusEnum_GeoTargetConstantStatus = 2
	// The geo target constant is obsolete and will be removed.
	GeoTargetConstantStatusEnum_REMOVAL_PLANNED GeoTargetConstantStatusEnum_GeoTargetConstantStatus = 3
)

// Enum value maps for GeoTargetConstantStatusEnum_GeoTargetConstantStatus.
var (
	GeoTargetConstantStatusEnum_GeoTargetConstantStatus_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "ENABLED",
		3: "REMOVAL_PLANNED",
	}
	GeoTargetConstantStatusEnum_GeoTargetConstantStatus_value = map[string]int32{
		"UNSPECIFIED":     0,
		"UNKNOWN":         1,
		"ENABLED":         2,
		"REMOVAL_PLANNED": 3,
	}
)

func (x GeoTargetConstantStatusEnum_GeoTargetConstantStatus) Enum() *GeoTargetConstantStatusEnum_GeoTargetConstantStatus {
	p := new(GeoTargetConstantStatusEnum_GeoTargetConstantStatus)
	*p = x
	return p
}

func (x GeoTargetConstantStatusEnum_GeoTargetConstantStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GeoTargetConstantStatusEnum_GeoTargetConstantStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_geo_target_constant_status_proto_enumTypes[0].Descriptor()
}

func (GeoTargetConstantStatusEnum_GeoTargetConstantStatus) Type() protoreflect.EnumType {
	return &file_geo_target_constant_status_proto_enumTypes[0]
}

func (x GeoTargetConstantStatusEnum_GeoTargetConstantStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GeoTargetConstantStatusEnum_GeoTargetConstantStatus.Descriptor instead.
func (GeoTargetConstantStatusEnum_GeoTargetConstantStatus) EnumDescriptor() ([]byte, []int) {
	return file_geo_target_constant_status_proto_rawDescGZIP(), []int{0, 0}
}

// Container for describing the status of a geo target constant.
type GeoTargetConstantStatusEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GeoTargetConstantStatusEnum) Reset() {
	*x = GeoTargetConstantStatusEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geo_target_constant_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeoTargetConstantStatusEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoTargetConstantStatusEnum) ProtoMessage() {}

func (x *GeoTargetConstantStatusEnum) ProtoReflect() protoreflect.Message {
	mi := &file_geo_target_constant_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeoTargetConstantStatusEnum.ProtoReflect.Descriptor instead.
func (*GeoTargetConstantStatusEnum) Descriptor() ([]byte, []int) {
	return file_geo_target_constant_status_proto_rawDescGZIP(), []int{0}
}

var File_geo_target_constant_status_proto protoreflect.FileDescriptor

var file_geo_target_constant_status_proto_rawDesc = []byte{
	0x0a, 0x20, 0x67, 0x65, 0x6f, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x22, 0x78, 0x0a, 0x1b, 0x47, 0x65, 0x6f, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75,
	0x6d, 0x22, 0x59, 0x0a, 0x17, 0x47, 0x65, 0x6f, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0f, 0x0a, 0x0b,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x4e,
	0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x45, 0x4d, 0x4f, 0x56,
	0x41, 0x4c, 0x5f, 0x50, 0x4c, 0x41, 0x4e, 0x4e, 0x45, 0x44, 0x10, 0x03, 0x42, 0xf6, 0x01, 0x0a,
	0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x42, 0x1c, 0x47, 0x65, 0x6f, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02,
	0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x36, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca,
	0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x36, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73,
	0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x36, 0x3a, 0x3a,
	0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_geo_target_constant_status_proto_rawDescOnce sync.Once
	file_geo_target_constant_status_proto_rawDescData = file_geo_target_constant_status_proto_rawDesc
)

func file_geo_target_constant_status_proto_rawDescGZIP() []byte {
	file_geo_target_constant_status_proto_rawDescOnce.Do(func() {
		file_geo_target_constant_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_geo_target_constant_status_proto_rawDescData)
	})
	return file_geo_target_constant_status_proto_rawDescData
}

var file_geo_target_constant_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_geo_target_constant_status_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_geo_target_constant_status_proto_goTypes = []interface{}{
	(GeoTargetConstantStatusEnum_GeoTargetConstantStatus)(0), // 0: google.ads.googleads.v16.enums.GeoTargetConstantStatusEnum.GeoTargetConstantStatus
	(*GeoTargetConstantStatusEnum)(nil),                      // 1: google.ads.googleads.v16.enums.GeoTargetConstantStatusEnum
}
var file_geo_target_constant_status_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_geo_target_constant_status_proto_init() }
func file_geo_target_constant_status_proto_init() {
	if File_geo_target_constant_status_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_geo_target_constant_status_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeoTargetConstantStatusEnum); i {
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
			RawDescriptor: file_geo_target_constant_status_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_geo_target_constant_status_proto_goTypes,
		DependencyIndexes: file_geo_target_constant_status_proto_depIdxs,
		EnumInfos:         file_geo_target_constant_status_proto_enumTypes,
		MessageInfos:      file_geo_target_constant_status_proto_msgTypes,
	}.Build()
	File_geo_target_constant_status_proto = out.File
	file_geo_target_constant_status_proto_rawDesc = nil
	file_geo_target_constant_status_proto_goTypes = nil
	file_geo_target_constant_status_proto_depIdxs = nil
}
