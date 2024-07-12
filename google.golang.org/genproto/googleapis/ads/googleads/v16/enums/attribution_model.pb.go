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
// source: attribution_model.proto

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

// The attribution model that describes how to distribute credit for a
// particular conversion across potentially many prior interactions.
type AttributionModelEnum_AttributionModel int32

const (
	// Not specified.
	AttributionModelEnum_UNSPECIFIED AttributionModelEnum_AttributionModel = 0
	// Used for return value only. Represents value unknown in this version.
	AttributionModelEnum_UNKNOWN AttributionModelEnum_AttributionModel = 1
	// Uses external attribution.
	AttributionModelEnum_EXTERNAL AttributionModelEnum_AttributionModel = 100
	// Attributes all credit for a conversion to its last click.
	AttributionModelEnum_GOOGLE_ADS_LAST_CLICK AttributionModelEnum_AttributionModel = 101
	// Attributes all credit for a conversion to its first click using Google
	// Search attribution.
	AttributionModelEnum_GOOGLE_SEARCH_ATTRIBUTION_FIRST_CLICK AttributionModelEnum_AttributionModel = 102
	// Attributes credit for a conversion equally across all of its clicks using
	// Google Search attribution.
	AttributionModelEnum_GOOGLE_SEARCH_ATTRIBUTION_LINEAR AttributionModelEnum_AttributionModel = 103
	// Attributes exponentially more credit for a conversion to its more recent
	// clicks using Google Search attribution (half-life is 1 week).
	AttributionModelEnum_GOOGLE_SEARCH_ATTRIBUTION_TIME_DECAY AttributionModelEnum_AttributionModel = 104
	// Attributes 40% of the credit for a conversion to its first and last
	// clicks. Remaining 20% is evenly distributed across all other clicks. This
	// uses Google Search attribution.
	AttributionModelEnum_GOOGLE_SEARCH_ATTRIBUTION_POSITION_BASED AttributionModelEnum_AttributionModel = 105
	// Flexible model that uses machine learning to determine the appropriate
	// distribution of credit among clicks using Google Search attribution.
	AttributionModelEnum_GOOGLE_SEARCH_ATTRIBUTION_DATA_DRIVEN AttributionModelEnum_AttributionModel = 106
)

// Enum value maps for AttributionModelEnum_AttributionModel.
var (
	AttributionModelEnum_AttributionModel_name = map[int32]string{
		0:   "UNSPECIFIED",
		1:   "UNKNOWN",
		100: "EXTERNAL",
		101: "GOOGLE_ADS_LAST_CLICK",
		102: "GOOGLE_SEARCH_ATTRIBUTION_FIRST_CLICK",
		103: "GOOGLE_SEARCH_ATTRIBUTION_LINEAR",
		104: "GOOGLE_SEARCH_ATTRIBUTION_TIME_DECAY",
		105: "GOOGLE_SEARCH_ATTRIBUTION_POSITION_BASED",
		106: "GOOGLE_SEARCH_ATTRIBUTION_DATA_DRIVEN",
	}
	AttributionModelEnum_AttributionModel_value = map[string]int32{
		"UNSPECIFIED":                              0,
		"UNKNOWN":                                  1,
		"EXTERNAL":                                 100,
		"GOOGLE_ADS_LAST_CLICK":                    101,
		"GOOGLE_SEARCH_ATTRIBUTION_FIRST_CLICK":    102,
		"GOOGLE_SEARCH_ATTRIBUTION_LINEAR":         103,
		"GOOGLE_SEARCH_ATTRIBUTION_TIME_DECAY":     104,
		"GOOGLE_SEARCH_ATTRIBUTION_POSITION_BASED": 105,
		"GOOGLE_SEARCH_ATTRIBUTION_DATA_DRIVEN":    106,
	}
)

func (x AttributionModelEnum_AttributionModel) Enum() *AttributionModelEnum_AttributionModel {
	p := new(AttributionModelEnum_AttributionModel)
	*p = x
	return p
}

func (x AttributionModelEnum_AttributionModel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AttributionModelEnum_AttributionModel) Descriptor() protoreflect.EnumDescriptor {
	return file_attribution_model_proto_enumTypes[0].Descriptor()
}

func (AttributionModelEnum_AttributionModel) Type() protoreflect.EnumType {
	return &file_attribution_model_proto_enumTypes[0]
}

func (x AttributionModelEnum_AttributionModel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AttributionModelEnum_AttributionModel.Descriptor instead.
func (AttributionModelEnum_AttributionModel) EnumDescriptor() ([]byte, []int) {
	return file_attribution_model_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum representing the attribution model that describes how to
// distribute credit for a particular conversion across potentially many prior
// interactions.
type AttributionModelEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AttributionModelEnum) Reset() {
	*x = AttributionModelEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attribution_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttributionModelEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttributionModelEnum) ProtoMessage() {}

func (x *AttributionModelEnum) ProtoReflect() protoreflect.Message {
	mi := &file_attribution_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttributionModelEnum.ProtoReflect.Descriptor instead.
func (*AttributionModelEnum) Descriptor() ([]byte, []int) {
	return file_attribution_model_proto_rawDescGZIP(), []int{0}
}

var File_attribution_model_proto protoreflect.FileDescriptor

var file_attribution_model_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x36, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0xc6, 0x02, 0x0a, 0x14, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x45, 0x6e,
	0x75, 0x6d, 0x22, 0xad, 0x02, 0x0a, 0x10, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x58, 0x54, 0x45, 0x52, 0x4e, 0x41,
	0x4c, 0x10, 0x64, 0x12, 0x19, 0x0a, 0x15, 0x47, 0x4f, 0x4f, 0x47, 0x4c, 0x45, 0x5f, 0x41, 0x44,
	0x53, 0x5f, 0x4c, 0x41, 0x53, 0x54, 0x5f, 0x43, 0x4c, 0x49, 0x43, 0x4b, 0x10, 0x65, 0x12, 0x29,
	0x0a, 0x25, 0x47, 0x4f, 0x4f, 0x47, 0x4c, 0x45, 0x5f, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x5f,
	0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x49, 0x52, 0x53,
	0x54, 0x5f, 0x43, 0x4c, 0x49, 0x43, 0x4b, 0x10, 0x66, 0x12, 0x24, 0x0a, 0x20, 0x47, 0x4f, 0x4f,
	0x47, 0x4c, 0x45, 0x5f, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x49,
	0x42, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4c, 0x49, 0x4e, 0x45, 0x41, 0x52, 0x10, 0x67, 0x12,
	0x28, 0x0a, 0x24, 0x47, 0x4f, 0x4f, 0x47, 0x4c, 0x45, 0x5f, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48,
	0x5f, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x49, 0x4d,
	0x45, 0x5f, 0x44, 0x45, 0x43, 0x41, 0x59, 0x10, 0x68, 0x12, 0x2c, 0x0a, 0x28, 0x47, 0x4f, 0x4f,
	0x47, 0x4c, 0x45, 0x5f, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x49,
	0x42, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x42, 0x41, 0x53, 0x45, 0x44, 0x10, 0x69, 0x12, 0x29, 0x0a, 0x25, 0x47, 0x4f, 0x4f, 0x47, 0x4c,
	0x45, 0x5f, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x44, 0x52, 0x49, 0x56, 0x45, 0x4e,
	0x10, 0x6a, 0x42, 0xef, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x36, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x15, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x36, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02,
	0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x36, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea,
	0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x36, 0x3a, 0x3a, 0x45,
	0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_attribution_model_proto_rawDescOnce sync.Once
	file_attribution_model_proto_rawDescData = file_attribution_model_proto_rawDesc
)

func file_attribution_model_proto_rawDescGZIP() []byte {
	file_attribution_model_proto_rawDescOnce.Do(func() {
		file_attribution_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_attribution_model_proto_rawDescData)
	})
	return file_attribution_model_proto_rawDescData
}

var file_attribution_model_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_attribution_model_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_attribution_model_proto_goTypes = []interface{}{
	(AttributionModelEnum_AttributionModel)(0), // 0: google.ads.googleads.v16.enums.AttributionModelEnum.AttributionModel
	(*AttributionModelEnum)(nil),               // 1: google.ads.googleads.v16.enums.AttributionModelEnum
}
var file_attribution_model_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_attribution_model_proto_init() }
func file_attribution_model_proto_init() {
	if File_attribution_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_attribution_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttributionModelEnum); i {
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
			RawDescriptor: file_attribution_model_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_attribution_model_proto_goTypes,
		DependencyIndexes: file_attribution_model_proto_depIdxs,
		EnumInfos:         file_attribution_model_proto_enumTypes,
		MessageInfos:      file_attribution_model_proto_msgTypes,
	}.Build()
	File_attribution_model_proto = out.File
	file_attribution_model_proto_rawDesc = nil
	file_attribution_model_proto_goTypes = nil
	file_attribution_model_proto_depIdxs = nil
}
