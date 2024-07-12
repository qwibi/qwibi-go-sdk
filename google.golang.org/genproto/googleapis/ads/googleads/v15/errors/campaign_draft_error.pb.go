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
// source: campaign_draft_error.proto

package errors

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

// Enum describing possible campaign draft errors.
type CampaignDraftErrorEnum_CampaignDraftError int32

const (
	// Enum unspecified.
	CampaignDraftErrorEnum_UNSPECIFIED CampaignDraftErrorEnum_CampaignDraftError = 0
	// The received error code is not known in this version.
	CampaignDraftErrorEnum_UNKNOWN CampaignDraftErrorEnum_CampaignDraftError = 1
	// A draft with this name already exists for this campaign.
	CampaignDraftErrorEnum_DUPLICATE_DRAFT_NAME CampaignDraftErrorEnum_CampaignDraftError = 2
	// The draft is removed and cannot be transitioned to another status.
	CampaignDraftErrorEnum_INVALID_STATUS_TRANSITION_FROM_REMOVED CampaignDraftErrorEnum_CampaignDraftError = 3
	// The draft has been promoted and cannot be transitioned to the specified
	// status.
	CampaignDraftErrorEnum_INVALID_STATUS_TRANSITION_FROM_PROMOTED CampaignDraftErrorEnum_CampaignDraftError = 4
	// The draft has failed to be promoted and cannot be transitioned to the
	// specified status.
	CampaignDraftErrorEnum_INVALID_STATUS_TRANSITION_FROM_PROMOTE_FAILED CampaignDraftErrorEnum_CampaignDraftError = 5
	// This customer is not allowed to create drafts.
	CampaignDraftErrorEnum_CUSTOMER_CANNOT_CREATE_DRAFT CampaignDraftErrorEnum_CampaignDraftError = 6
	// This campaign is not allowed to create drafts.
	CampaignDraftErrorEnum_CAMPAIGN_CANNOT_CREATE_DRAFT CampaignDraftErrorEnum_CampaignDraftError = 7
	// This modification cannot be made on a draft.
	CampaignDraftErrorEnum_INVALID_DRAFT_CHANGE CampaignDraftErrorEnum_CampaignDraftError = 8
	// The draft cannot be transitioned to the specified status from its
	// current status.
	CampaignDraftErrorEnum_INVALID_STATUS_TRANSITION CampaignDraftErrorEnum_CampaignDraftError = 9
	// The campaign has reached the maximum number of drafts that can be created
	// for a campaign throughout its lifetime. No additional drafts can be
	// created for this campaign. Removed drafts also count towards this limit.
	CampaignDraftErrorEnum_MAX_NUMBER_OF_DRAFTS_PER_CAMPAIGN_REACHED CampaignDraftErrorEnum_CampaignDraftError = 10
	// ListAsyncErrors was called without first promoting the draft.
	CampaignDraftErrorEnum_LIST_ERRORS_FOR_PROMOTED_DRAFT_ONLY CampaignDraftErrorEnum_CampaignDraftError = 11
)

// Enum value maps for CampaignDraftErrorEnum_CampaignDraftError.
var (
	CampaignDraftErrorEnum_CampaignDraftError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "DUPLICATE_DRAFT_NAME",
		3:  "INVALID_STATUS_TRANSITION_FROM_REMOVED",
		4:  "INVALID_STATUS_TRANSITION_FROM_PROMOTED",
		5:  "INVALID_STATUS_TRANSITION_FROM_PROMOTE_FAILED",
		6:  "CUSTOMER_CANNOT_CREATE_DRAFT",
		7:  "CAMPAIGN_CANNOT_CREATE_DRAFT",
		8:  "INVALID_DRAFT_CHANGE",
		9:  "INVALID_STATUS_TRANSITION",
		10: "MAX_NUMBER_OF_DRAFTS_PER_CAMPAIGN_REACHED",
		11: "LIST_ERRORS_FOR_PROMOTED_DRAFT_ONLY",
	}
	CampaignDraftErrorEnum_CampaignDraftError_value = map[string]int32{
		"UNSPECIFIED":                                   0,
		"UNKNOWN":                                       1,
		"DUPLICATE_DRAFT_NAME":                          2,
		"INVALID_STATUS_TRANSITION_FROM_REMOVED":        3,
		"INVALID_STATUS_TRANSITION_FROM_PROMOTED":       4,
		"INVALID_STATUS_TRANSITION_FROM_PROMOTE_FAILED": 5,
		"CUSTOMER_CANNOT_CREATE_DRAFT":                  6,
		"CAMPAIGN_CANNOT_CREATE_DRAFT":                  7,
		"INVALID_DRAFT_CHANGE":                          8,
		"INVALID_STATUS_TRANSITION":                     9,
		"MAX_NUMBER_OF_DRAFTS_PER_CAMPAIGN_REACHED":     10,
		"LIST_ERRORS_FOR_PROMOTED_DRAFT_ONLY":           11,
	}
)

func (x CampaignDraftErrorEnum_CampaignDraftError) Enum() *CampaignDraftErrorEnum_CampaignDraftError {
	p := new(CampaignDraftErrorEnum_CampaignDraftError)
	*p = x
	return p
}

func (x CampaignDraftErrorEnum_CampaignDraftError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CampaignDraftErrorEnum_CampaignDraftError) Descriptor() protoreflect.EnumDescriptor {
	return file_campaign_draft_error_proto_enumTypes[0].Descriptor()
}

func (CampaignDraftErrorEnum_CampaignDraftError) Type() protoreflect.EnumType {
	return &file_campaign_draft_error_proto_enumTypes[0]
}

func (x CampaignDraftErrorEnum_CampaignDraftError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CampaignDraftErrorEnum_CampaignDraftError.Descriptor instead.
func (CampaignDraftErrorEnum_CampaignDraftError) EnumDescriptor() ([]byte, []int) {
	return file_campaign_draft_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible campaign draft errors.
type CampaignDraftErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CampaignDraftErrorEnum) Reset() {
	*x = CampaignDraftErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_campaign_draft_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CampaignDraftErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CampaignDraftErrorEnum) ProtoMessage() {}

func (x *CampaignDraftErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_campaign_draft_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CampaignDraftErrorEnum.ProtoReflect.Descriptor instead.
func (*CampaignDraftErrorEnum) Descriptor() ([]byte, []int) {
	return file_campaign_draft_error_proto_rawDescGZIP(), []int{0}
}

var File_campaign_draft_error_proto protoreflect.FileDescriptor

var file_campaign_draft_error_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x64, 0x72, 0x61, 0x66, 0x74,
	0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0xc8, 0x03,
	0x0a, 0x16, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x44, 0x72, 0x61, 0x66, 0x74, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xad, 0x03, 0x0a, 0x12, 0x43, 0x61, 0x6d,
	0x70, 0x61, 0x69, 0x67, 0x6e, 0x44, 0x72, 0x61, 0x66, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x18, 0x0a,
	0x14, 0x44, 0x55, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x45, 0x5f, 0x44, 0x52, 0x41, 0x46, 0x54,
	0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x02, 0x12, 0x2a, 0x0a, 0x26, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x49,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x52, 0x4f, 0x4d, 0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45,
	0x44, 0x10, 0x03, 0x12, 0x2b, 0x0a, 0x27, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x46, 0x52, 0x4f, 0x4d, 0x5f, 0x50, 0x52, 0x4f, 0x4d, 0x4f, 0x54, 0x45, 0x44, 0x10, 0x04,
	0x12, 0x31, 0x0a, 0x2d, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x52,
	0x4f, 0x4d, 0x5f, 0x50, 0x52, 0x4f, 0x4d, 0x4f, 0x54, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45,
	0x44, 0x10, 0x05, 0x12, 0x20, 0x0a, 0x1c, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x5f,
	0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x44, 0x52,
	0x41, 0x46, 0x54, 0x10, 0x06, 0x12, 0x20, 0x0a, 0x1c, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47,
	0x4e, 0x5f, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f,
	0x44, 0x52, 0x41, 0x46, 0x54, 0x10, 0x07, 0x12, 0x18, 0x0a, 0x14, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x5f, 0x44, 0x52, 0x41, 0x46, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x10,
	0x08, 0x12, 0x1d, 0x0a, 0x19, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x09,
	0x12, 0x2d, 0x0a, 0x29, 0x4d, 0x41, 0x58, 0x5f, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x4f,
	0x46, 0x5f, 0x44, 0x52, 0x41, 0x46, 0x54, 0x53, 0x5f, 0x50, 0x45, 0x52, 0x5f, 0x43, 0x41, 0x4d,
	0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x52, 0x45, 0x41, 0x43, 0x48, 0x45, 0x44, 0x10, 0x0a, 0x12,
	0x27, 0x0a, 0x23, 0x4c, 0x49, 0x53, 0x54, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x53, 0x5f, 0x46,
	0x4f, 0x52, 0x5f, 0x50, 0x52, 0x4f, 0x4d, 0x4f, 0x54, 0x45, 0x44, 0x5f, 0x44, 0x52, 0x41, 0x46,
	0x54, 0x5f, 0x4f, 0x4e, 0x4c, 0x59, 0x10, 0x0b, 0x42, 0xf7, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x42, 0x17, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x44, 0x72, 0x61, 0x66, 0x74, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2f, 0x76, 0x31, 0x35, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e,
	0x56, 0x31, 0x35, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64,
	0x73, 0x5c, 0x56, 0x31, 0x35, 0x5c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xea, 0x02, 0x23, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x35, 0x3a, 0x3a, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_campaign_draft_error_proto_rawDescOnce sync.Once
	file_campaign_draft_error_proto_rawDescData = file_campaign_draft_error_proto_rawDesc
)

func file_campaign_draft_error_proto_rawDescGZIP() []byte {
	file_campaign_draft_error_proto_rawDescOnce.Do(func() {
		file_campaign_draft_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_campaign_draft_error_proto_rawDescData)
	})
	return file_campaign_draft_error_proto_rawDescData
}

var file_campaign_draft_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_campaign_draft_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_campaign_draft_error_proto_goTypes = []interface{}{
	(CampaignDraftErrorEnum_CampaignDraftError)(0), // 0: google.ads.googleads.v15.errors.CampaignDraftErrorEnum.CampaignDraftError
	(*CampaignDraftErrorEnum)(nil),                 // 1: google.ads.googleads.v15.errors.CampaignDraftErrorEnum
}
var file_campaign_draft_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_campaign_draft_error_proto_init() }
func file_campaign_draft_error_proto_init() {
	if File_campaign_draft_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_campaign_draft_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CampaignDraftErrorEnum); i {
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
			RawDescriptor: file_campaign_draft_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_campaign_draft_error_proto_goTypes,
		DependencyIndexes: file_campaign_draft_error_proto_depIdxs,
		EnumInfos:         file_campaign_draft_error_proto_enumTypes,
		MessageInfos:      file_campaign_draft_error_proto_msgTypes,
	}.Build()
	File_campaign_draft_error_proto = out.File
	file_campaign_draft_error_proto_rawDesc = nil
	file_campaign_draft_error_proto_goTypes = nil
	file_campaign_draft_error_proto_depIdxs = nil
}
