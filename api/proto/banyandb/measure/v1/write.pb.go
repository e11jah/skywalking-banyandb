// Licensed to Apache Software Foundation (ASF) under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Apache Software Foundation (ASF) licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: banyandb/measure/v1/write.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	v11 "github.com/apache/skywalking-banyandb/api/proto/banyandb/common/v1"
	v1 "github.com/apache/skywalking-banyandb/api/proto/banyandb/model/v1"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//DataPointValue is the data point for writing. It only contains values.
type DataPointValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// timestamp is in the timeunit of nanoseconds.
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// the order of tag_families' items match the measure schema
	TagFamilies []*v1.TagFamilyForWrite `protobuf:"bytes,2,rep,name=tag_families,json=tagFamilies,proto3" json:"tag_families,omitempty"`
	// the order of fields match the measure schema
	Fields []*v1.FieldValue `protobuf:"bytes,3,rep,name=fields,proto3" json:"fields,omitempty"`
}

func (x *DataPointValue) Reset() {
	*x = DataPointValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banyandb_measure_v1_write_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataPointValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataPointValue) ProtoMessage() {}

func (x *DataPointValue) ProtoReflect() protoreflect.Message {
	mi := &file_banyandb_measure_v1_write_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataPointValue.ProtoReflect.Descriptor instead.
func (*DataPointValue) Descriptor() ([]byte, []int) {
	return file_banyandb_measure_v1_write_proto_rawDescGZIP(), []int{0}
}

func (x *DataPointValue) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *DataPointValue) GetTagFamilies() []*v1.TagFamilyForWrite {
	if x != nil {
		return x.TagFamilies
	}
	return nil
}

func (x *DataPointValue) GetFields() []*v1.FieldValue {
	if x != nil {
		return x.Fields
	}
	return nil
}

// WriteRequest is the request contract for write
type WriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the metadata is required.
	Metadata *v11.Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// the data_point is required.
	DataPoint *DataPointValue `protobuf:"bytes,2,opt,name=data_point,json=dataPoint,proto3" json:"data_point,omitempty"`
}

func (x *WriteRequest) Reset() {
	*x = WriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banyandb_measure_v1_write_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteRequest) ProtoMessage() {}

func (x *WriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_banyandb_measure_v1_write_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteRequest.ProtoReflect.Descriptor instead.
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return file_banyandb_measure_v1_write_proto_rawDescGZIP(), []int{1}
}

func (x *WriteRequest) GetMetadata() *v11.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *WriteRequest) GetDataPoint() *DataPointValue {
	if x != nil {
		return x.DataPoint
	}
	return nil
}

// WriteResponse is the response contract for write
type WriteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WriteResponse) Reset() {
	*x = WriteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banyandb_measure_v1_write_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteResponse) ProtoMessage() {}

func (x *WriteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_banyandb_measure_v1_write_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteResponse.ProtoReflect.Descriptor instead.
func (*WriteResponse) Descriptor() ([]byte, []int) {
	return file_banyandb_measure_v1_write_proto_rawDescGZIP(), []int{2}
}

type InternalWriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShardId    uint32        `protobuf:"varint,1,opt,name=shard_id,json=shardId,proto3" json:"shard_id,omitempty"`
	SeriesHash []byte        `protobuf:"bytes,2,opt,name=series_hash,json=seriesHash,proto3" json:"series_hash,omitempty"`
	Request    *WriteRequest `protobuf:"bytes,3,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *InternalWriteRequest) Reset() {
	*x = InternalWriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banyandb_measure_v1_write_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InternalWriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InternalWriteRequest) ProtoMessage() {}

func (x *InternalWriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_banyandb_measure_v1_write_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InternalWriteRequest.ProtoReflect.Descriptor instead.
func (*InternalWriteRequest) Descriptor() ([]byte, []int) {
	return file_banyandb_measure_v1_write_proto_rawDescGZIP(), []int{3}
}

func (x *InternalWriteRequest) GetShardId() uint32 {
	if x != nil {
		return x.ShardId
	}
	return 0
}

func (x *InternalWriteRequest) GetSeriesHash() []byte {
	if x != nil {
		return x.SeriesHash
	}
	return nil
}

func (x *InternalWriteRequest) GetRequest() *WriteRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

var File_banyandb_measure_v1_write_proto protoreflect.FileDescriptor

var file_banyandb_measure_v1_write_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62, 0x2f, 0x6d, 0x65, 0x61, 0x73, 0x75,
	0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x13, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62, 0x2e, 0x6d, 0x65, 0x61, 0x73,
	0x75, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e, 0x64,
	0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e,
	0x64, 0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xca, 0x01, 0x0a, 0x0e, 0x44, 0x61, 0x74,
	0x61, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x47, 0x0a, 0x0c, 0x74, 0x61, 0x67, 0x5f, 0x66, 0x61, 0x6d,
	0x69, 0x6c, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x62, 0x61,
	0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x61, 0x67, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x46, 0x6f, 0x72, 0x57, 0x72, 0x69, 0x74,
	0x65, 0x52, 0x0b, 0x74, 0x61, 0x67, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x69, 0x65, 0x73, 0x12, 0x35,
	0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x22, 0x8c, 0x01, 0x0a, 0x0c, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x61, 0x6e, 0x79, 0x61,
	0x6e, 0x64, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x42, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62, 0x2e,
	0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x8f, 0x01, 0x0a, 0x14, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x73, 0x68, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x73, 0x68, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72,
	0x69, 0x65, 0x73, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a,
	0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x48, 0x61, 0x73, 0x68, 0x12, 0x3b, 0x0a, 0x07, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x62, 0x61,
	0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62, 0x2e, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x70, 0x0a, 0x29, 0x6f, 0x72, 0x67, 0x2e, 0x61,
	0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62, 0x2e, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2f, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69,
	0x6e, 0x67, 0x2d, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62, 0x2f, 0x6d,
	0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_banyandb_measure_v1_write_proto_rawDescOnce sync.Once
	file_banyandb_measure_v1_write_proto_rawDescData = file_banyandb_measure_v1_write_proto_rawDesc
)

func file_banyandb_measure_v1_write_proto_rawDescGZIP() []byte {
	file_banyandb_measure_v1_write_proto_rawDescOnce.Do(func() {
		file_banyandb_measure_v1_write_proto_rawDescData = protoimpl.X.CompressGZIP(file_banyandb_measure_v1_write_proto_rawDescData)
	})
	return file_banyandb_measure_v1_write_proto_rawDescData
}

var file_banyandb_measure_v1_write_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_banyandb_measure_v1_write_proto_goTypes = []interface{}{
	(*DataPointValue)(nil),        // 0: banyandb.measure.v1.DataPointValue
	(*WriteRequest)(nil),          // 1: banyandb.measure.v1.WriteRequest
	(*WriteResponse)(nil),         // 2: banyandb.measure.v1.WriteResponse
	(*InternalWriteRequest)(nil),  // 3: banyandb.measure.v1.InternalWriteRequest
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
	(*v1.TagFamilyForWrite)(nil),  // 5: banyandb.model.v1.TagFamilyForWrite
	(*v1.FieldValue)(nil),         // 6: banyandb.model.v1.FieldValue
	(*v11.Metadata)(nil),          // 7: banyandb.common.v1.Metadata
}
var file_banyandb_measure_v1_write_proto_depIdxs = []int32{
	4, // 0: banyandb.measure.v1.DataPointValue.timestamp:type_name -> google.protobuf.Timestamp
	5, // 1: banyandb.measure.v1.DataPointValue.tag_families:type_name -> banyandb.model.v1.TagFamilyForWrite
	6, // 2: banyandb.measure.v1.DataPointValue.fields:type_name -> banyandb.model.v1.FieldValue
	7, // 3: banyandb.measure.v1.WriteRequest.metadata:type_name -> banyandb.common.v1.Metadata
	0, // 4: banyandb.measure.v1.WriteRequest.data_point:type_name -> banyandb.measure.v1.DataPointValue
	1, // 5: banyandb.measure.v1.InternalWriteRequest.request:type_name -> banyandb.measure.v1.WriteRequest
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_banyandb_measure_v1_write_proto_init() }
func file_banyandb_measure_v1_write_proto_init() {
	if File_banyandb_measure_v1_write_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_banyandb_measure_v1_write_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataPointValue); i {
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
		file_banyandb_measure_v1_write_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteRequest); i {
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
		file_banyandb_measure_v1_write_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteResponse); i {
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
		file_banyandb_measure_v1_write_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InternalWriteRequest); i {
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
			RawDescriptor: file_banyandb_measure_v1_write_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_banyandb_measure_v1_write_proto_goTypes,
		DependencyIndexes: file_banyandb_measure_v1_write_proto_depIdxs,
		MessageInfos:      file_banyandb_measure_v1_write_proto_msgTypes,
	}.Build()
	File_banyandb_measure_v1_write_proto = out.File
	file_banyandb_measure_v1_write_proto_rawDesc = nil
	file_banyandb_measure_v1_write_proto_goTypes = nil
	file_banyandb_measure_v1_write_proto_depIdxs = nil
}
