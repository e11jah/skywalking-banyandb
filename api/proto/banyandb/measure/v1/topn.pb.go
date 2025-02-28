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
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: banyandb/measure/v1/topn.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	v1 "github.com/apache/skywalking-banyandb/api/proto/banyandb/common/v1"
	v11 "github.com/apache/skywalking-banyandb/api/proto/banyandb/model/v1"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//TopNList contains a series of topN items
type TopNList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// timestamp is in the timeunit of milliseconds.
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// items contains top-n items in a list
	Items []*TopNList_Item `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *TopNList) Reset() {
	*x = TopNList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banyandb_measure_v1_topn_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopNList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopNList) ProtoMessage() {}

func (x *TopNList) ProtoReflect() protoreflect.Message {
	mi := &file_banyandb_measure_v1_topn_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopNList.ProtoReflect.Descriptor instead.
func (*TopNList) Descriptor() ([]byte, []int) {
	return file_banyandb_measure_v1_topn_proto_rawDescGZIP(), []int{0}
}

func (x *TopNList) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *TopNList) GetItems() []*TopNList_Item {
	if x != nil {
		return x.Items
	}
	return nil
}

// TopNResponse is the response for a query to the Query module.
type TopNResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// lists contain a series topN lists ranked by timestamp
	// if agg_func in query request is specified, lists' size should be one.
	Lists []*TopNList `protobuf:"bytes,1,rep,name=lists,proto3" json:"lists,omitempty"`
}

func (x *TopNResponse) Reset() {
	*x = TopNResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banyandb_measure_v1_topn_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopNResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopNResponse) ProtoMessage() {}

func (x *TopNResponse) ProtoReflect() protoreflect.Message {
	mi := &file_banyandb_measure_v1_topn_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopNResponse.ProtoReflect.Descriptor instead.
func (*TopNResponse) Descriptor() ([]byte, []int) {
	return file_banyandb_measure_v1_topn_proto_rawDescGZIP(), []int{1}
}

func (x *TopNResponse) GetLists() []*TopNList {
	if x != nil {
		return x.Lists
	}
	return nil
}

// TopNRequest is the request contract for query.
type TopNRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// metadata is required
	Metadata *v1.Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// time_range is a range query with begin/end time of entities in the timeunit of milliseconds.
	TimeRange *v11.TimeRange `protobuf:"bytes,2,opt,name=time_range,json=timeRange,proto3" json:"time_range,omitempty"`
	// top_n set the how many items should be returned in each list.
	TopN int32 `protobuf:"varint,3,opt,name=top_n,json=topN,proto3" json:"top_n,omitempty"`
	// agg aggregates lists grouped by field names in the time_range
	Agg v11.AggregationFunction `protobuf:"varint,4,opt,name=agg,proto3,enum=banyandb.model.v1.AggregationFunction" json:"agg,omitempty"`
	// criteria select counters.
	Conditions []*v11.Condition `protobuf:"bytes,5,rep,name=conditions,proto3" json:"conditions,omitempty"`
	// field_value_sort indicates how to sort fields
	FieldValueSort v11.Sort `protobuf:"varint,6,opt,name=field_value_sort,json=fieldValueSort,proto3,enum=banyandb.model.v1.Sort" json:"field_value_sort,omitempty"`
}

func (x *TopNRequest) Reset() {
	*x = TopNRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banyandb_measure_v1_topn_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopNRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopNRequest) ProtoMessage() {}

func (x *TopNRequest) ProtoReflect() protoreflect.Message {
	mi := &file_banyandb_measure_v1_topn_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopNRequest.ProtoReflect.Descriptor instead.
func (*TopNRequest) Descriptor() ([]byte, []int) {
	return file_banyandb_measure_v1_topn_proto_rawDescGZIP(), []int{2}
}

func (x *TopNRequest) GetMetadata() *v1.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *TopNRequest) GetTimeRange() *v11.TimeRange {
	if x != nil {
		return x.TimeRange
	}
	return nil
}

func (x *TopNRequest) GetTopN() int32 {
	if x != nil {
		return x.TopN
	}
	return 0
}

func (x *TopNRequest) GetAgg() v11.AggregationFunction {
	if x != nil {
		return x.Agg
	}
	return v11.AggregationFunction(0)
}

func (x *TopNRequest) GetConditions() []*v11.Condition {
	if x != nil {
		return x.Conditions
	}
	return nil
}

func (x *TopNRequest) GetFieldValueSort() v11.Sort {
	if x != nil {
		return x.FieldValueSort
	}
	return v11.Sort(0)
}

type TopNList_Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value *v11.FieldValue `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *TopNList_Item) Reset() {
	*x = TopNList_Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banyandb_measure_v1_topn_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopNList_Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopNList_Item) ProtoMessage() {}

func (x *TopNList_Item) ProtoReflect() protoreflect.Message {
	mi := &file_banyandb_measure_v1_topn_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopNList_Item.ProtoReflect.Descriptor instead.
func (*TopNList_Item) Descriptor() ([]byte, []int) {
	return file_banyandb_measure_v1_topn_proto_rawDescGZIP(), []int{0, 0}
}

func (x *TopNList_Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TopNList_Item) GetValue() *v11.FieldValue {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_banyandb_measure_v1_topn_proto protoreflect.FileDescriptor

var file_banyandb_measure_v1_topn_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62, 0x2f, 0x6d, 0x65, 0x61, 0x73, 0x75,
	0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x70, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x13, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62, 0x2e, 0x6d, 0x65, 0x61, 0x73, 0x75,
	0x72, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e, 0x64,
	0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e, 0x64,
	0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcf, 0x01, 0x0a, 0x08, 0x54, 0x6f, 0x70, 0x4e, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x38, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x62,
	0x61, 0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62, 0x2e, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x6f, 0x70, 0x4e, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x4f, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x43, 0x0a, 0x0c, 0x54, 0x6f, 0x70, 0x4e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x05, 0x6c, 0x69, 0x73, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e,
	0x64, 0x62, 0x2e, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f,
	0x70, 0x4e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x05, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x22, 0xd4, 0x02,
	0x0a, 0x0b, 0x54, 0x6f, 0x70, 0x4e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3b, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x61,
	0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x52,
	0x61, 0x6e, 0x67, 0x65, 0x12, 0x13, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x5f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x6f, 0x70, 0x4e, 0x12, 0x38, 0x0a, 0x03, 0x61, 0x67, 0x67,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e, 0x64,
	0x62, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03,
	0x61, 0x67, 0x67, 0x12, 0x3c, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e,
	0x64, 0x62, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x41, 0x0a, 0x10, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x5f, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x62, 0x61,
	0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x6f, 0x72, 0x74, 0x52, 0x0e, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x53, 0x6f, 0x72, 0x74, 0x42, 0x70, 0x0a, 0x29, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63,
	0x68, 0x65, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x62, 0x61,
	0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62, 0x2e, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x2e, 0x76,
	0x31, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70,
	0x61, 0x63, 0x68, 0x65, 0x2f, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2d,
	0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62, 0x2f, 0x6d, 0x65, 0x61, 0x73,
	0x75, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_banyandb_measure_v1_topn_proto_rawDescOnce sync.Once
	file_banyandb_measure_v1_topn_proto_rawDescData = file_banyandb_measure_v1_topn_proto_rawDesc
)

func file_banyandb_measure_v1_topn_proto_rawDescGZIP() []byte {
	file_banyandb_measure_v1_topn_proto_rawDescOnce.Do(func() {
		file_banyandb_measure_v1_topn_proto_rawDescData = protoimpl.X.CompressGZIP(file_banyandb_measure_v1_topn_proto_rawDescData)
	})
	return file_banyandb_measure_v1_topn_proto_rawDescData
}

var file_banyandb_measure_v1_topn_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_banyandb_measure_v1_topn_proto_goTypes = []interface{}{
	(*TopNList)(nil),              // 0: banyandb.measure.v1.TopNList
	(*TopNResponse)(nil),          // 1: banyandb.measure.v1.TopNResponse
	(*TopNRequest)(nil),           // 2: banyandb.measure.v1.TopNRequest
	(*TopNList_Item)(nil),         // 3: banyandb.measure.v1.TopNList.Item
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
	(*v1.Metadata)(nil),           // 5: banyandb.common.v1.Metadata
	(*v11.TimeRange)(nil),         // 6: banyandb.model.v1.TimeRange
	(v11.AggregationFunction)(0),  // 7: banyandb.model.v1.AggregationFunction
	(*v11.Condition)(nil),         // 8: banyandb.model.v1.Condition
	(v11.Sort)(0),                 // 9: banyandb.model.v1.Sort
	(*v11.FieldValue)(nil),        // 10: banyandb.model.v1.FieldValue
}
var file_banyandb_measure_v1_topn_proto_depIdxs = []int32{
	4,  // 0: banyandb.measure.v1.TopNList.timestamp:type_name -> google.protobuf.Timestamp
	3,  // 1: banyandb.measure.v1.TopNList.items:type_name -> banyandb.measure.v1.TopNList.Item
	0,  // 2: banyandb.measure.v1.TopNResponse.lists:type_name -> banyandb.measure.v1.TopNList
	5,  // 3: banyandb.measure.v1.TopNRequest.metadata:type_name -> banyandb.common.v1.Metadata
	6,  // 4: banyandb.measure.v1.TopNRequest.time_range:type_name -> banyandb.model.v1.TimeRange
	7,  // 5: banyandb.measure.v1.TopNRequest.agg:type_name -> banyandb.model.v1.AggregationFunction
	8,  // 6: banyandb.measure.v1.TopNRequest.conditions:type_name -> banyandb.model.v1.Condition
	9,  // 7: banyandb.measure.v1.TopNRequest.field_value_sort:type_name -> banyandb.model.v1.Sort
	10, // 8: banyandb.measure.v1.TopNList.Item.value:type_name -> banyandb.model.v1.FieldValue
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_banyandb_measure_v1_topn_proto_init() }
func file_banyandb_measure_v1_topn_proto_init() {
	if File_banyandb_measure_v1_topn_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_banyandb_measure_v1_topn_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopNList); i {
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
		file_banyandb_measure_v1_topn_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopNResponse); i {
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
		file_banyandb_measure_v1_topn_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopNRequest); i {
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
		file_banyandb_measure_v1_topn_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopNList_Item); i {
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
			RawDescriptor: file_banyandb_measure_v1_topn_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_banyandb_measure_v1_topn_proto_goTypes,
		DependencyIndexes: file_banyandb_measure_v1_topn_proto_depIdxs,
		MessageInfos:      file_banyandb_measure_v1_topn_proto_msgTypes,
	}.Build()
	File_banyandb_measure_v1_topn_proto = out.File
	file_banyandb_measure_v1_topn_proto_rawDesc = nil
	file_banyandb_measure_v1_topn_proto_goTypes = nil
	file_banyandb_measure_v1_topn_proto_depIdxs = nil
}
