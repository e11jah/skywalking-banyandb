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

package measure

import (
	"io"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"

	"github.com/apache/skywalking-banyandb/api/common"
	commonv1 "github.com/apache/skywalking-banyandb/api/proto/banyandb/common/v1"
	databasev1 "github.com/apache/skywalking-banyandb/api/proto/banyandb/database/v1"
	measurev1 "github.com/apache/skywalking-banyandb/api/proto/banyandb/measure/v1"
	modelv1 "github.com/apache/skywalking-banyandb/api/proto/banyandb/model/v1"
	"github.com/apache/skywalking-banyandb/banyand/tsdb"
	"github.com/apache/skywalking-banyandb/pkg/partition"
	resourceSchema "github.com/apache/skywalking-banyandb/pkg/schema"
)

var (
	ErrTagFamilyNotExist = errors.New("tag family doesn't exist")
)

type Query interface {
	LoadGroup(name string) (resourceSchema.Group, bool)
	Measure(measure *commonv1.Metadata) (Measure, error)
}

type Measure interface {
	io.Closer
	Write(value *measurev1.DataPointValue) error
	Shards(entity tsdb.Entity) ([]tsdb.Shard, error)
	Shard(id common.ShardID) (tsdb.Shard, error)
	ParseTagFamily(family string, item tsdb.Item) (*modelv1.TagFamily, error)
	ParseField(name string, item tsdb.Item) (*measurev1.DataPoint_Field, error)
	GetSchema() *databasev1.Measure
	GetIndexRules() []*databasev1.IndexRule
}

var _ Measure = (*measure)(nil)

func (s *measure) Shards(entity tsdb.Entity) ([]tsdb.Shard, error) {
	wrap := func(shards []tsdb.Shard) []tsdb.Shard {
		result := make([]tsdb.Shard, len(shards))
		for i := 0; i < len(shards); i++ {
			result[i] = tsdb.NewScopedShard(tsdb.Entry(s.name), shards[i])
		}
		return result
	}
	db := s.db.SupplyTSDB()
	if len(entity) < 1 {
		return wrap(db.Shards()), nil
	}
	for _, e := range entity {
		if e == nil {
			return wrap(db.Shards()), nil
		}
	}
	shardID, err := partition.ShardID(entity.Prepend(tsdb.Entry(s.name)).Marshal(), s.shardNum)
	if err != nil {
		return nil, err
	}
	shard, err := db.Shard(common.ShardID(shardID))
	if err != nil {
		return nil, err
	}
	return []tsdb.Shard{tsdb.NewScopedShard(tsdb.Entry(s.name), shard)}, nil
}

func (s *measure) Shard(id common.ShardID) (tsdb.Shard, error) {
	shard, err := s.db.SupplyTSDB().Shard(id)
	if err != nil {
		return nil, err
	}
	return tsdb.NewScopedShard(tsdb.Entry(s.name), shard), nil
}

func (s *measure) ParseTagFamily(family string, item tsdb.Item) (*modelv1.TagFamily, error) {
	familyRawBytes, err := item.Family(familyIdentity(family, TagFlag))
	if err != nil {
		return nil, err
	}
	tagFamily := &modelv1.TagFamilyForWrite{}
	err = proto.Unmarshal(familyRawBytes, tagFamily)
	if err != nil {
		return nil, err
	}
	tags := make([]*modelv1.Tag, len(tagFamily.GetTags()))
	var tagSpec []*databasev1.TagSpec
	for _, tf := range s.schema.GetTagFamilies() {
		if tf.GetName() == family {
			tagSpec = tf.GetTags()
		}
	}
	if tagSpec == nil {
		return nil, ErrTagFamilyNotExist
	}
	for i, tag := range tagFamily.GetTags() {
		tags[i] = &modelv1.Tag{
			Key: tagSpec[i].GetName(),
			Value: &modelv1.TagValue{
				Value: tag.GetValue(),
			},
		}
	}
	return &modelv1.TagFamily{
		Name: family,
		Tags: tags,
	}, err
}

func (s *measure) ParseField(name string, item tsdb.Item) (*measurev1.DataPoint_Field, error) {
	var fieldSpec *databasev1.FieldSpec
	for _, spec := range s.schema.GetFields() {
		if spec.GetName() == name {
			fieldSpec = spec
			break
		}
	}
	bytes, err := item.Family(familyIdentity(name, encoderFieldFlag(fieldSpec, s.interval)))
	if err != nil {
		return nil, err
	}
	fieldValue := decodeFieldValue(bytes, fieldSpec)
	return &measurev1.DataPoint_Field{
		Name:  name,
		Value: fieldValue,
	}, err
}
