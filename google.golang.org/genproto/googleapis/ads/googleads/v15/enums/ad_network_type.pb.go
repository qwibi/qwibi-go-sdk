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
// source: ad_network_type.proto

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

// Enumerates Google Ads network types.
type AdNetworkTypeEnum_AdNetworkType int32

const (
	// Not specified.
	AdNetworkTypeEnum_UNSPECIFIED AdNetworkTypeEnum_AdNetworkType = 0
	// The value is unknown in this version.
	AdNetworkTypeEnum_UNKNOWN AdNetworkTypeEnum_AdNetworkType = 1
	// Google search.
	AdNetworkTypeEnum_SEARCH AdNetworkTypeEnum_AdNetworkType = 2
	// Search partners.
	AdNetworkTypeEnum_SEARCH_PARTNERS AdNetworkTypeEnum_AdNetworkType = 3
	// Display Network.
	AdNetworkTypeEnum_CONTENT AdNetworkTypeEnum_AdNetworkType = 4
	// Cross-network.
	AdNetworkTypeEnum_MIXED AdNetworkTypeEnum_AdNetworkType = 7
	// YouTube
	AdNetworkTypeEnum_YOUTUBE AdNetworkTypeEnum_AdNetworkType = 8
	// Google TV
	AdNetworkTypeEnum_GOOGLE_TV AdNetworkTypeEnum_AdNetworkType = 9
)

// Enum value maps for AdNetworkTypeEnum_AdNetworkType.
var (
	AdNetworkTypeEnum_AdNetworkType_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "SEARCH",
		3: "SEARCH_PARTNERS",
		4: "CONTENT",
		7: "MIXED",
		8: "YOUTUBE",
		9: "GOOGLE_TV",
	}
	AdNetworkTypeEnum_AdNetworkType_value = map[string]int32{
		"UNSPECIFIED":     0,
		"UNKNOWN":         1,
		"SEARCH":          2,
		"SEARCH_PARTNERS": 3,
		"CONTENT":         4,
		"MIXED":           7,
		"YOUTUBE":         8,
		"GOOGLE_TV":       9,
	}
)

func (x AdNetworkTypeEnum_AdNetworkType) Enum() *AdNetworkTypeEnum_AdNetworkType {
	p := new(AdNetworkTypeEnum_AdNetworkType)
	*p = x
	return p
}

func (x AdNetworkTypeEnum_AdNetworkType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AdNetworkTypeEnum_AdNetworkType) Descriptor() protoreflect.EnumDescriptor {
	return file_ad_network_type_proto_enumTypes[0].Descriptor()
}

func (AdNetworkTypeEnum_AdNetworkType) Type() protoreflect.EnumType {
	return &file_ad_network_type_proto_enumTypes[0]
}

func (x AdNetworkTypeEnum_AdNetworkType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AdNetworkTypeEnum_AdNetworkType.Descriptor instead.
func (AdNetworkTypeEnum_AdNetworkType) EnumDescriptor() ([]byte, []int) {
	return file_ad_network_type_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enumeration of Google Ads network types.
type AdNetworkTypeEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AdNetworkTypeEnum) Reset() {
	*x = AdNetworkTypeEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ad_network_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdNetworkTypeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdNetworkTypeEnum) ProtoMessage() {}

func (x *AdNetworkTypeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_ad_network_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdNetworkTypeEnum.ProtoReflect.Descriptor instead.
func (*AdNetworkTypeEnum) Descriptor() ([]byte, []int) {
	return file_ad_network_type_proto_rawDescGZIP(), []int{0}
}

var File_ad_network_type_proto protoreflect.FileDescriptor

var file_ad_network_type_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x64, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x35, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0x98, 0x01, 0x0a, 0x11, 0x41, 0x64, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x82, 0x01,
	0x0a, 0x0d, 0x41, 0x64, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x0a, 0x0a,
	0x06, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x45, 0x41,
	0x52, 0x43, 0x48, 0x5f, 0x50, 0x41, 0x52, 0x54, 0x4e, 0x45, 0x52, 0x53, 0x10, 0x03, 0x12, 0x0b,
	0x0a, 0x07, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x4d,
	0x49, 0x58, 0x45, 0x44, 0x10, 0x07, 0x12, 0x0b, 0x0a, 0x07, 0x59, 0x4f, 0x55, 0x54, 0x55, 0x42,
	0x45, 0x10, 0x08, 0x12, 0x0d, 0x0a, 0x09, 0x47, 0x4f, 0x4f, 0x47, 0x4c, 0x45, 0x5f, 0x54, 0x56,
	0x10, 0x09, 0x42, 0xec, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x35, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x12, 0x41, 0x64, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64,
	0x73, 0x2e, 0x56, 0x31, 0x35, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x5c, 0x56, 0x31, 0x35, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02, 0x22, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x35, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ad_network_type_proto_rawDescOnce sync.Once
	file_ad_network_type_proto_rawDescData = file_ad_network_type_proto_rawDesc
)

func file_ad_network_type_proto_rawDescGZIP() []byte {
	file_ad_network_type_proto_rawDescOnce.Do(func() {
		file_ad_network_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_ad_network_type_proto_rawDescData)
	})
	return file_ad_network_type_proto_rawDescData
}

var file_ad_network_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ad_network_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ad_network_type_proto_goTypes = []interface{}{
	(AdNetworkTypeEnum_AdNetworkType)(0), // 0: google.ads.googleads.v15.enums.AdNetworkTypeEnum.AdNetworkType
	(*AdNetworkTypeEnum)(nil),            // 1: google.ads.googleads.v15.enums.AdNetworkTypeEnum
}
var file_ad_network_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ad_network_type_proto_init() }
func file_ad_network_type_proto_init() {
	if File_ad_network_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ad_network_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdNetworkTypeEnum); i {
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
			RawDescriptor: file_ad_network_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ad_network_type_proto_goTypes,
		DependencyIndexes: file_ad_network_type_proto_depIdxs,
		EnumInfos:         file_ad_network_type_proto_enumTypes,
		MessageInfos:      file_ad_network_type_proto_msgTypes,
	}.Build()
	File_ad_network_type_proto = out.File
	file_ad_network_type_proto_rawDesc = nil
	file_ad_network_type_proto_goTypes = nil
	file_ad_network_type_proto_depIdxs = nil
}
