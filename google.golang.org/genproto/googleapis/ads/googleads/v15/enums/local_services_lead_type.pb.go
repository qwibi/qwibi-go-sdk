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
// source: local_services_lead_type.proto

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

// Possible types of lead.
type LocalServicesLeadTypeEnum_LeadType int32

const (
	// Not specified.
	LocalServicesLeadTypeEnum_UNSPECIFIED LocalServicesLeadTypeEnum_LeadType = 0
	// Used for return value only. Represents value unknown in this version.
	LocalServicesLeadTypeEnum_UNKNOWN LocalServicesLeadTypeEnum_LeadType = 1
	// Message lead.
	LocalServicesLeadTypeEnum_MESSAGE LocalServicesLeadTypeEnum_LeadType = 2
	// Lead created by phone call.
	LocalServicesLeadTypeEnum_PHONE_CALL LocalServicesLeadTypeEnum_LeadType = 3
	// Booking lead.
	LocalServicesLeadTypeEnum_BOOKING LocalServicesLeadTypeEnum_LeadType = 4
)

// Enum value maps for LocalServicesLeadTypeEnum_LeadType.
var (
	LocalServicesLeadTypeEnum_LeadType_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "MESSAGE",
		3: "PHONE_CALL",
		4: "BOOKING",
	}
	LocalServicesLeadTypeEnum_LeadType_value = map[string]int32{
		"UNSPECIFIED": 0,
		"UNKNOWN":     1,
		"MESSAGE":     2,
		"PHONE_CALL":  3,
		"BOOKING":     4,
	}
)

func (x LocalServicesLeadTypeEnum_LeadType) Enum() *LocalServicesLeadTypeEnum_LeadType {
	p := new(LocalServicesLeadTypeEnum_LeadType)
	*p = x
	return p
}

func (x LocalServicesLeadTypeEnum_LeadType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LocalServicesLeadTypeEnum_LeadType) Descriptor() protoreflect.EnumDescriptor {
	return file_local_services_lead_type_proto_enumTypes[0].Descriptor()
}

func (LocalServicesLeadTypeEnum_LeadType) Type() protoreflect.EnumType {
	return &file_local_services_lead_type_proto_enumTypes[0]
}

func (x LocalServicesLeadTypeEnum_LeadType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LocalServicesLeadTypeEnum_LeadType.Descriptor instead.
func (LocalServicesLeadTypeEnum_LeadType) EnumDescriptor() ([]byte, []int) {
	return file_local_services_lead_type_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible types of lead.
type LocalServicesLeadTypeEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LocalServicesLeadTypeEnum) Reset() {
	*x = LocalServicesLeadTypeEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_local_services_lead_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalServicesLeadTypeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalServicesLeadTypeEnum) ProtoMessage() {}

func (x *LocalServicesLeadTypeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_local_services_lead_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalServicesLeadTypeEnum.ProtoReflect.Descriptor instead.
func (*LocalServicesLeadTypeEnum) Descriptor() ([]byte, []int) {
	return file_local_services_lead_type_proto_rawDescGZIP(), []int{0}
}

var File_local_services_lead_type_proto protoreflect.FileDescriptor

var file_local_services_lead_type_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x5f, 0x6c, 0x65, 0x61, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x22, 0x6f, 0x0a, 0x19, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x4c, 0x65, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x52, 0x0a,
	0x08, 0x4c, 0x65, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x45, 0x53, 0x53, 0x41,
	0x47, 0x45, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x48, 0x4f, 0x4e, 0x45, 0x5f, 0x43, 0x41,
	0x4c, 0x4c, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x4f, 0x4f, 0x4b, 0x49, 0x4e, 0x47, 0x10,
	0x04, 0x42, 0xf4, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x35, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x1a, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x4c, 0x65, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64,
	0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x35, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41,
	0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x35, 0x2e, 0x45, 0x6e, 0x75,
	0x6d, 0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x35, 0x5c, 0x45, 0x6e,
	0x75, 0x6d, 0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64,
	0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31,
	0x35, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_local_services_lead_type_proto_rawDescOnce sync.Once
	file_local_services_lead_type_proto_rawDescData = file_local_services_lead_type_proto_rawDesc
)

func file_local_services_lead_type_proto_rawDescGZIP() []byte {
	file_local_services_lead_type_proto_rawDescOnce.Do(func() {
		file_local_services_lead_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_local_services_lead_type_proto_rawDescData)
	})
	return file_local_services_lead_type_proto_rawDescData
}

var file_local_services_lead_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_local_services_lead_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_local_services_lead_type_proto_goTypes = []interface{}{
	(LocalServicesLeadTypeEnum_LeadType)(0), // 0: google.ads.googleads.v15.enums.LocalServicesLeadTypeEnum.LeadType
	(*LocalServicesLeadTypeEnum)(nil),       // 1: google.ads.googleads.v15.enums.LocalServicesLeadTypeEnum
}
var file_local_services_lead_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_local_services_lead_type_proto_init() }
func file_local_services_lead_type_proto_init() {
	if File_local_services_lead_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_local_services_lead_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalServicesLeadTypeEnum); i {
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
			RawDescriptor: file_local_services_lead_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_local_services_lead_type_proto_goTypes,
		DependencyIndexes: file_local_services_lead_type_proto_depIdxs,
		EnumInfos:         file_local_services_lead_type_proto_enumTypes,
		MessageInfos:      file_local_services_lead_type_proto_msgTypes,
	}.Build()
	File_local_services_lead_type_proto = out.File
	file_local_services_lead_type_proto_rawDesc = nil
	file_local_services_lead_type_proto_goTypes = nil
	file_local_services_lead_type_proto_depIdxs = nil
}
