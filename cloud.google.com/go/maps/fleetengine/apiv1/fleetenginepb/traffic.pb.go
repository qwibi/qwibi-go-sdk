// Copyright 2024 Google LLC
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
// source: traffic.proto

package fleetenginepb

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

// The classification of polyline speed based on traffic data.
type SpeedReadingInterval_Speed int32

const (
	// Default value. This value is unused.
	SpeedReadingInterval_SPEED_UNSPECIFIED SpeedReadingInterval_Speed = 0
	// Normal speed, no slowdown is detected.
	SpeedReadingInterval_NORMAL SpeedReadingInterval_Speed = 1
	// Slowdown detected, but no traffic jam formed.
	SpeedReadingInterval_SLOW SpeedReadingInterval_Speed = 2
	// Traffic jam detected.
	SpeedReadingInterval_TRAFFIC_JAM SpeedReadingInterval_Speed = 3
)

// Enum value maps for SpeedReadingInterval_Speed.
var (
	SpeedReadingInterval_Speed_name = map[int32]string{
		0: "SPEED_UNSPECIFIED",
		1: "NORMAL",
		2: "SLOW",
		3: "TRAFFIC_JAM",
	}
	SpeedReadingInterval_Speed_value = map[string]int32{
		"SPEED_UNSPECIFIED": 0,
		"NORMAL":            1,
		"SLOW":              2,
		"TRAFFIC_JAM":       3,
	}
)

func (x SpeedReadingInterval_Speed) Enum() *SpeedReadingInterval_Speed {
	p := new(SpeedReadingInterval_Speed)
	*p = x
	return p
}

func (x SpeedReadingInterval_Speed) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SpeedReadingInterval_Speed) Descriptor() protoreflect.EnumDescriptor {
	return file_traffic_proto_enumTypes[0].Descriptor()
}

func (SpeedReadingInterval_Speed) Type() protoreflect.EnumType {
	return &file_traffic_proto_enumTypes[0]
}

func (x SpeedReadingInterval_Speed) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SpeedReadingInterval_Speed.Descriptor instead.
func (SpeedReadingInterval_Speed) EnumDescriptor() ([]byte, []int) {
	return file_traffic_proto_rawDescGZIP(), []int{0, 0}
}

// Traffic density indicator on a contiguous segment of a path. Given a path
// with points P_0, P_1, ... , P_N (zero-based index), the SpeedReadingInterval
// defines an interval and describes its traffic using the following categories.
type SpeedReadingInterval struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The starting index of this interval in the path.
	// In JSON, when the index is 0, the field will appear to be unpopulated.
	StartPolylinePointIndex int32 `protobuf:"varint,1,opt,name=start_polyline_point_index,json=startPolylinePointIndex,proto3" json:"start_polyline_point_index,omitempty"`
	// The ending index of this interval in the path.
	// In JSON, when the index is 0, the field will appear to be unpopulated.
	EndPolylinePointIndex int32 `protobuf:"varint,2,opt,name=end_polyline_point_index,json=endPolylinePointIndex,proto3" json:"end_polyline_point_index,omitempty"`
	// Traffic speed in this interval.
	Speed SpeedReadingInterval_Speed `protobuf:"varint,3,opt,name=speed,proto3,enum=maps.fleetengine.v1.SpeedReadingInterval_Speed" json:"speed,omitempty"`
}

func (x *SpeedReadingInterval) Reset() {
	*x = SpeedReadingInterval{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traffic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpeedReadingInterval) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpeedReadingInterval) ProtoMessage() {}

func (x *SpeedReadingInterval) ProtoReflect() protoreflect.Message {
	mi := &file_traffic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpeedReadingInterval.ProtoReflect.Descriptor instead.
func (*SpeedReadingInterval) Descriptor() ([]byte, []int) {
	return file_traffic_proto_rawDescGZIP(), []int{0}
}

func (x *SpeedReadingInterval) GetStartPolylinePointIndex() int32 {
	if x != nil {
		return x.StartPolylinePointIndex
	}
	return 0
}

func (x *SpeedReadingInterval) GetEndPolylinePointIndex() int32 {
	if x != nil {
		return x.EndPolylinePointIndex
	}
	return 0
}

func (x *SpeedReadingInterval) GetSpeed() SpeedReadingInterval_Speed {
	if x != nil {
		return x.Speed
	}
	return SpeedReadingInterval_SPEED_UNSPECIFIED
}

// Traffic density along a Vehicle's path.
type ConsumableTrafficPolyline struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Traffic speed along the path from the previous waypoint to the current
	// waypoint.
	SpeedReadingInterval []*SpeedReadingInterval `protobuf:"bytes,1,rep,name=speed_reading_interval,json=speedReadingInterval,proto3" json:"speed_reading_interval,omitempty"`
	// The path the driver is taking from the previous waypoint to the current
	// waypoint. This path has landmarks in it so clients can show traffic markers
	// along the path (see `speed_reading_interval`).  Decoding is not yet
	// supported.
	EncodedPathToWaypoint string `protobuf:"bytes,2,opt,name=encoded_path_to_waypoint,json=encodedPathToWaypoint,proto3" json:"encoded_path_to_waypoint,omitempty"`
}

func (x *ConsumableTrafficPolyline) Reset() {
	*x = ConsumableTrafficPolyline{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traffic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumableTrafficPolyline) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumableTrafficPolyline) ProtoMessage() {}

func (x *ConsumableTrafficPolyline) ProtoReflect() protoreflect.Message {
	mi := &file_traffic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumableTrafficPolyline.ProtoReflect.Descriptor instead.
func (*ConsumableTrafficPolyline) Descriptor() ([]byte, []int) {
	return file_traffic_proto_rawDescGZIP(), []int{1}
}

func (x *ConsumableTrafficPolyline) GetSpeedReadingInterval() []*SpeedReadingInterval {
	if x != nil {
		return x.SpeedReadingInterval
	}
	return nil
}

func (x *ConsumableTrafficPolyline) GetEncodedPathToWaypoint() string {
	if x != nil {
		return x.EncodedPathToWaypoint
	}
	return ""
}

var File_traffic_proto protoreflect.FileDescriptor

var file_traffic_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x13, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2e, 0x76, 0x31, 0x22, 0x9a, 0x02, 0x0a, 0x14, 0x53, 0x70, 0x65, 0x65, 0x64, 0x52, 0x65,
	0x61, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x3b, 0x0a,
	0x1a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x70, 0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65, 0x5f,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x17, 0x73, 0x74, 0x61, 0x72, 0x74, 0x50, 0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x37, 0x0a, 0x18, 0x65, 0x6e,
	0x64, 0x5f, 0x70, 0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x15, 0x65, 0x6e,
	0x64, 0x50, 0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x45, 0x0a, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x52, 0x65,
	0x61, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x2e, 0x53, 0x70,
	0x65, 0x65, 0x64, 0x52, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x22, 0x45, 0x0a, 0x05, 0x53, 0x70,
	0x65, 0x65, 0x64, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x50, 0x45, 0x45, 0x44, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x4f,
	0x52, 0x4d, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x4c, 0x4f, 0x57, 0x10, 0x02,
	0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x52, 0x41, 0x46, 0x46, 0x49, 0x43, 0x5f, 0x4a, 0x41, 0x4d, 0x10,
	0x03, 0x22, 0xb5, 0x01, 0x0a, 0x19, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x61, 0x62, 0x6c, 0x65,
	0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65, 0x12,
	0x5f, 0x0a, 0x16, 0x73, 0x70, 0x65, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67,
	0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x52, 0x65, 0x61, 0x64, 0x69,
	0x6e, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x52, 0x14, 0x73, 0x70, 0x65, 0x65,
	0x64, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x12, 0x37, 0x0a, 0x18, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x5f, 0x70, 0x61, 0x74, 0x68,
	0x5f, 0x74, 0x6f, 0x5f, 0x77, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x15, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x50, 0x61, 0x74, 0x68, 0x54,
	0x6f, 0x57, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x42, 0xd4, 0x01, 0x0a, 0x1a, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69,
	0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x46, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x61,
	0x70, 0x73, 0x2f, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x76, 0x31, 0x2f, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x70, 0x62, 0x3b, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x70, 0x62,
	0xa2, 0x02, 0x03, 0x43, 0x46, 0x45, 0xaa, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x4d, 0x61, 0x70, 0x73, 0x2e, 0x46, 0x6c, 0x65, 0x65, 0x74, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x4d, 0x61, 0x70,
	0x73, 0x5c, 0x46, 0x6c, 0x65, 0x65, 0x74, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5c, 0x56, 0x31,
	0xea, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x4d, 0x61, 0x70, 0x73, 0x3a,
	0x3a, 0x46, 0x6c, 0x65, 0x65, 0x74, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_traffic_proto_rawDescOnce sync.Once
	file_traffic_proto_rawDescData = file_traffic_proto_rawDesc
)

func file_traffic_proto_rawDescGZIP() []byte {
	file_traffic_proto_rawDescOnce.Do(func() {
		file_traffic_proto_rawDescData = protoimpl.X.CompressGZIP(file_traffic_proto_rawDescData)
	})
	return file_traffic_proto_rawDescData
}

var file_traffic_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_traffic_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_traffic_proto_goTypes = []interface{}{
	(SpeedReadingInterval_Speed)(0),   // 0: maps.fleetengine.v1.SpeedReadingInterval.Speed
	(*SpeedReadingInterval)(nil),      // 1: maps.fleetengine.v1.SpeedReadingInterval
	(*ConsumableTrafficPolyline)(nil), // 2: maps.fleetengine.v1.ConsumableTrafficPolyline
}
var file_traffic_proto_depIdxs = []int32{
	0, // 0: maps.fleetengine.v1.SpeedReadingInterval.speed:type_name -> maps.fleetengine.v1.SpeedReadingInterval.Speed
	1, // 1: maps.fleetengine.v1.ConsumableTrafficPolyline.speed_reading_interval:type_name -> maps.fleetengine.v1.SpeedReadingInterval
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_traffic_proto_init() }
func file_traffic_proto_init() {
	if File_traffic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_traffic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpeedReadingInterval); i {
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
		file_traffic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsumableTrafficPolyline); i {
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
			RawDescriptor: file_traffic_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_traffic_proto_goTypes,
		DependencyIndexes: file_traffic_proto_depIdxs,
		EnumInfos:         file_traffic_proto_enumTypes,
		MessageInfos:      file_traffic_proto_msgTypes,
	}.Build()
	File_traffic_proto = out.File
	file_traffic_proto_rawDesc = nil
	file_traffic_proto_goTypes = nil
	file_traffic_proto_depIdxs = nil
}
