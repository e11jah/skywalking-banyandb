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
// source: banyandb/database/v1/event.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	v1 "github.com/apache/skywalking-banyandb/api/proto/banyandb/common/v1"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Action int32

const (
	Action_ACTION_UNSPECIFIED Action = 0
	Action_ACTION_PUT         Action = 1
	Action_ACTION_DELETE      Action = 2
)

// Enum value maps for Action.
var (
	Action_name = map[int32]string{
		0: "ACTION_UNSPECIFIED",
		1: "ACTION_PUT",
		2: "ACTION_DELETE",
	}
	Action_value = map[string]int32{
		"ACTION_UNSPECIFIED": 0,
		"ACTION_PUT":         1,
		"ACTION_DELETE":      2,
	}
)

func (x Action) Enum() *Action {
	p := new(Action)
	*p = x
	return p
}

func (x Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Action) Descriptor() protoreflect.EnumDescriptor {
	return file_banyandb_database_v1_event_proto_enumTypes[0].Descriptor()
}

func (Action) Type() protoreflect.EnumType {
	return &file_banyandb_database_v1_event_proto_enumTypes[0]
}

func (x Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Action.Descriptor instead.
func (Action) EnumDescriptor() ([]byte, []int) {
	return file_banyandb_database_v1_event_proto_rawDescGZIP(), []int{0}
}

type ShardEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shard  *Shard                 `protobuf:"bytes,1,opt,name=shard,proto3" json:"shard,omitempty"`
	Action Action                 `protobuf:"varint,2,opt,name=action,proto3,enum=banyandb.database.v1.Action" json:"action,omitempty"`
	Time   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *ShardEvent) Reset() {
	*x = ShardEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banyandb_database_v1_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShardEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShardEvent) ProtoMessage() {}

func (x *ShardEvent) ProtoReflect() protoreflect.Message {
	mi := &file_banyandb_database_v1_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShardEvent.ProtoReflect.Descriptor instead.
func (*ShardEvent) Descriptor() ([]byte, []int) {
	return file_banyandb_database_v1_event_proto_rawDescGZIP(), []int{0}
}

func (x *ShardEvent) GetShard() *Shard {
	if x != nil {
		return x.Shard
	}
	return nil
}

func (x *ShardEvent) GetAction() Action {
	if x != nil {
		return x.Action
	}
	return Action_ACTION_UNSPECIFIED
}

func (x *ShardEvent) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

type EntityEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subject       *v1.Metadata              `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	EntityLocator []*EntityEvent_TagLocator `protobuf:"bytes,2,rep,name=entity_locator,json=entityLocator,proto3" json:"entity_locator,omitempty"`
	Action        Action                    `protobuf:"varint,3,opt,name=action,proto3,enum=banyandb.database.v1.Action" json:"action,omitempty"`
	Time          *timestamppb.Timestamp    `protobuf:"bytes,4,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *EntityEvent) Reset() {
	*x = EntityEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banyandb_database_v1_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityEvent) ProtoMessage() {}

func (x *EntityEvent) ProtoReflect() protoreflect.Message {
	mi := &file_banyandb_database_v1_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityEvent.ProtoReflect.Descriptor instead.
func (*EntityEvent) Descriptor() ([]byte, []int) {
	return file_banyandb_database_v1_event_proto_rawDescGZIP(), []int{1}
}

func (x *EntityEvent) GetSubject() *v1.Metadata {
	if x != nil {
		return x.Subject
	}
	return nil
}

func (x *EntityEvent) GetEntityLocator() []*EntityEvent_TagLocator {
	if x != nil {
		return x.EntityLocator
	}
	return nil
}

func (x *EntityEvent) GetAction() Action {
	if x != nil {
		return x.Action
	}
	return Action_ACTION_UNSPECIFIED
}

func (x *EntityEvent) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

type EntityEvent_TagLocator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FamilyOffset uint32 `protobuf:"varint,1,opt,name=family_offset,json=familyOffset,proto3" json:"family_offset,omitempty"`
	TagOffset    uint32 `protobuf:"varint,2,opt,name=tag_offset,json=tagOffset,proto3" json:"tag_offset,omitempty"`
}

func (x *EntityEvent_TagLocator) Reset() {
	*x = EntityEvent_TagLocator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banyandb_database_v1_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityEvent_TagLocator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityEvent_TagLocator) ProtoMessage() {}

func (x *EntityEvent_TagLocator) ProtoReflect() protoreflect.Message {
	mi := &file_banyandb_database_v1_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityEvent_TagLocator.ProtoReflect.Descriptor instead.
func (*EntityEvent_TagLocator) Descriptor() ([]byte, []int) {
	return file_banyandb_database_v1_event_proto_rawDescGZIP(), []int{1, 0}
}

func (x *EntityEvent_TagLocator) GetFamilyOffset() uint32 {
	if x != nil {
		return x.FamilyOffset
	}
	return 0
}

func (x *EntityEvent_TagLocator) GetTagOffset() uint32 {
	if x != nil {
		return x.TagOffset
	}
	return 0
}

var File_banyandb_database_v1_event_proto protoreflect.FileDescriptor

var file_banyandb_database_v1_event_proto_rawDesc = []byte{
	0x0a, 0x20, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x14, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x62, 0x61, 0x6e, 0x79, 0x61,
	0x6e, 0x64, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x62, 0x61, 0x6e, 0x79,
	0x61, 0x6e, 0x64, 0x62, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xa5, 0x01, 0x0a, 0x0a, 0x53, 0x68, 0x61, 0x72, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x31,
	0x0a, 0x05, 0x73, 0x68, 0x61, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x64, 0x52, 0x05, 0x73, 0x68, 0x61, 0x72,
	0x64, 0x12, 0x34, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1c, 0x2e, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0xd2, 0x02, 0x0a, 0x0b, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x36, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x61, 0x6e, 0x79, 0x61,
	0x6e, 0x64, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x53, 0x0a, 0x0e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e,
	0x64, 0x62, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x61, 0x67, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x0d, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x6f, 0x72, 0x12, 0x34, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x1a, 0x50, 0x0a, 0x0a, 0x54, 0x61,
	0x67, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x61, 0x6d, 0x69,
	0x6c, 0x79, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0c, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x74, 0x61, 0x67, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x74, 0x61, 0x67, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x2a, 0x43, 0x0a, 0x06,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x12, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e,
	0x0a, 0x0a, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x55, 0x54, 0x10, 0x01, 0x12, 0x11,
	0x0a, 0x0d, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10,
	0x02, 0x42, 0x72, 0x0a, 0x2a, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e,
	0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x62, 0x61, 0x6e, 0x79, 0x61,
	0x6e, 0x64, 0x62, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x5a,
	0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x61, 0x63,
	0x68, 0x65, 0x2f, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2d, 0x62, 0x61,
	0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_banyandb_database_v1_event_proto_rawDescOnce sync.Once
	file_banyandb_database_v1_event_proto_rawDescData = file_banyandb_database_v1_event_proto_rawDesc
)

func file_banyandb_database_v1_event_proto_rawDescGZIP() []byte {
	file_banyandb_database_v1_event_proto_rawDescOnce.Do(func() {
		file_banyandb_database_v1_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_banyandb_database_v1_event_proto_rawDescData)
	})
	return file_banyandb_database_v1_event_proto_rawDescData
}

var file_banyandb_database_v1_event_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_banyandb_database_v1_event_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_banyandb_database_v1_event_proto_goTypes = []interface{}{
	(Action)(0),                    // 0: banyandb.database.v1.Action
	(*ShardEvent)(nil),             // 1: banyandb.database.v1.ShardEvent
	(*EntityEvent)(nil),            // 2: banyandb.database.v1.EntityEvent
	(*EntityEvent_TagLocator)(nil), // 3: banyandb.database.v1.EntityEvent.TagLocator
	(*Shard)(nil),                  // 4: banyandb.database.v1.Shard
	(*timestamppb.Timestamp)(nil),  // 5: google.protobuf.Timestamp
	(*v1.Metadata)(nil),            // 6: banyandb.common.v1.Metadata
}
var file_banyandb_database_v1_event_proto_depIdxs = []int32{
	4, // 0: banyandb.database.v1.ShardEvent.shard:type_name -> banyandb.database.v1.Shard
	0, // 1: banyandb.database.v1.ShardEvent.action:type_name -> banyandb.database.v1.Action
	5, // 2: banyandb.database.v1.ShardEvent.time:type_name -> google.protobuf.Timestamp
	6, // 3: banyandb.database.v1.EntityEvent.subject:type_name -> banyandb.common.v1.Metadata
	3, // 4: banyandb.database.v1.EntityEvent.entity_locator:type_name -> banyandb.database.v1.EntityEvent.TagLocator
	0, // 5: banyandb.database.v1.EntityEvent.action:type_name -> banyandb.database.v1.Action
	5, // 6: banyandb.database.v1.EntityEvent.time:type_name -> google.protobuf.Timestamp
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_banyandb_database_v1_event_proto_init() }
func file_banyandb_database_v1_event_proto_init() {
	if File_banyandb_database_v1_event_proto != nil {
		return
	}
	file_banyandb_database_v1_database_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_banyandb_database_v1_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShardEvent); i {
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
		file_banyandb_database_v1_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntityEvent); i {
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
		file_banyandb_database_v1_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntityEvent_TagLocator); i {
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
			RawDescriptor: file_banyandb_database_v1_event_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_banyandb_database_v1_event_proto_goTypes,
		DependencyIndexes: file_banyandb_database_v1_event_proto_depIdxs,
		EnumInfos:         file_banyandb_database_v1_event_proto_enumTypes,
		MessageInfos:      file_banyandb_database_v1_event_proto_msgTypes,
	}.Build()
	File_banyandb_database_v1_event_proto = out.File
	file_banyandb_database_v1_event_proto_rawDesc = nil
	file_banyandb_database_v1_event_proto_goTypes = nil
	file_banyandb_database_v1_event_proto_depIdxs = nil
}
