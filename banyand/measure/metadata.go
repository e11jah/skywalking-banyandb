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
//
package measure

import (
	"context"
	"path"
	"time"

	"github.com/apache/skywalking-banyandb/api/common"
	"github.com/apache/skywalking-banyandb/api/event"
	commonv1 "github.com/apache/skywalking-banyandb/api/proto/banyandb/common/v1"
	databasev1 "github.com/apache/skywalking-banyandb/api/proto/banyandb/database/v1"
	"github.com/apache/skywalking-banyandb/banyand/discovery"
	"github.com/apache/skywalking-banyandb/banyand/metadata"
	"github.com/apache/skywalking-banyandb/banyand/metadata/schema"
	"github.com/apache/skywalking-banyandb/banyand/tsdb"
	"github.com/apache/skywalking-banyandb/pkg/logger"
	resourceSchema "github.com/apache/skywalking-banyandb/pkg/schema"
)

type schemaRepo struct {
	resourceSchema.Repository
	l        *logger.Logger
	metadata metadata.Repo
}

func newSchemaRepo(path string, metadata metadata.Repo, repo discovery.ServiceRepo, l *logger.Logger) schemaRepo {
	return schemaRepo{
		l:        l,
		metadata: metadata,
		Repository: resourceSchema.NewRepository(
			metadata,
			repo,
			l,
			newSupplier(path, metadata, l),
			event.MeasureTopicShardEvent,
			event.MeasureTopicEntityEvent,
		),
	}
}

func (sr *schemaRepo) OnAddOrUpdate(m schema.Metadata) {
	switch m.Kind {
	case schema.KindGroup:
		g := m.Spec.(*commonv1.Group)
		if g.Catalog != commonv1.Catalog_CATALOG_MEASURE {
			return
		}
		sr.SendMetadataEvent(resourceSchema.MetadataEvent{
			Typ:      resourceSchema.EventAddOrUpdate,
			Kind:     resourceSchema.EventKindGroup,
			Metadata: g.GetMetadata(),
		})
	case schema.KindMeasure:
		sr.SendMetadataEvent(resourceSchema.MetadataEvent{
			Typ:      resourceSchema.EventAddOrUpdate,
			Kind:     resourceSchema.EventKindResource,
			Metadata: m.Spec.(*databasev1.Measure).GetMetadata(),
		})
	case schema.KindIndexRuleBinding:
		irb, ok := m.Spec.(*databasev1.IndexRuleBinding)
		if !ok {
			sr.l.Warn().Msg("fail to convert message to IndexRuleBinding")
			return
		}
		if irb.GetSubject().Catalog == commonv1.Catalog_CATALOG_MEASURE {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			stm, err := sr.metadata.MeasureRegistry().GetMeasure(ctx, &commonv1.Metadata{
				Name:  irb.GetSubject().GetName(),
				Group: m.Group,
			})
			cancel()
			if err != nil {
				sr.l.Error().Err(err).Msg("fail to get subject")
				return
			}
			sr.SendMetadataEvent(resourceSchema.MetadataEvent{
				Typ:      resourceSchema.EventAddOrUpdate,
				Kind:     resourceSchema.EventKindResource,
				Metadata: stm.GetMetadata(),
			})
		}
	case schema.KindIndexRule:
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		subjects, err := sr.metadata.Subjects(ctx, m.Spec.(*databasev1.IndexRule), commonv1.Catalog_CATALOG_MEASURE)
		cancel()
		if err != nil {
			sr.l.Error().Err(err).Msg("fail to get subjects(measure)")
			return
		}
		for _, sub := range subjects {
			sr.SendMetadataEvent(resourceSchema.MetadataEvent{
				Typ:      resourceSchema.EventAddOrUpdate,
				Kind:     resourceSchema.EventKindResource,
				Metadata: sub.(*databasev1.Measure).GetMetadata(),
			})
		}
	default:
	}
}

func (sr *schemaRepo) OnDelete(m schema.Metadata) {
	switch m.Kind {
	case schema.KindGroup:
		g := m.Spec.(*commonv1.Group)
		if g.Catalog != commonv1.Catalog_CATALOG_MEASURE {
			return
		}
		sr.SendMetadataEvent(resourceSchema.MetadataEvent{
			Typ:      resourceSchema.EventDelete,
			Kind:     resourceSchema.EventKindGroup,
			Metadata: g.GetMetadata(),
		})
	case schema.KindMeasure:
		sr.SendMetadataEvent(resourceSchema.MetadataEvent{
			Typ:      resourceSchema.EventDelete,
			Kind:     resourceSchema.EventKindResource,
			Metadata: m.Spec.(*databasev1.Measure).GetMetadata(),
		})
	case schema.KindIndexRuleBinding:
		if m.Spec.(*databasev1.IndexRuleBinding).GetSubject().Catalog == commonv1.Catalog_CATALOG_MEASURE {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			stm, err := sr.metadata.MeasureRegistry().GetMeasure(ctx, &commonv1.Metadata{
				Name:  m.Name,
				Group: m.Group,
			})
			cancel()
			if err != nil {
				sr.l.Error().Err(err).Msg("fail to get subject")
				return
			}
			sr.SendMetadataEvent(resourceSchema.MetadataEvent{
				Typ:      resourceSchema.EventDelete,
				Kind:     resourceSchema.EventKindResource,
				Metadata: stm.GetMetadata(),
			})
		}
	case schema.KindIndexRule:
	default:
	}
}

func (sr *schemaRepo) loadMeasure(metadata *commonv1.Metadata) (*measure, bool) {
	r, ok := sr.LoadResource(metadata)
	if !ok {
		return nil, false
	}
	s, ok := r.(*measure)
	return s, ok
}

var _ resourceSchema.ResourceSupplier = (*supplier)(nil)

type supplier struct {
	path     string
	metadata metadata.Repo
	l        *logger.Logger
}

func newSupplier(path string, metadata metadata.Repo, l *logger.Logger) *supplier {
	return &supplier{
		path:     path,
		metadata: metadata,
		l:        l,
	}
}

func (s *supplier) OpenResource(shardNum uint32, db tsdb.Supplier, spec resourceSchema.ResourceSpec) (resourceSchema.Resource, error) {
	measureSchema := spec.Schema.(*databasev1.Measure)
	return openMeasure(shardNum, db, measureSpec{
		schema:     measureSchema,
		indexRules: spec.IndexRules,
	}, s.l)
}
func (s *supplier) ResourceSchema(repo metadata.Repo, md *commonv1.Metadata) (resourceSchema.ResourceSchema, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.metadata.MeasureRegistry().GetMeasure(ctx, md)
}

func (s *supplier) OpenDB(groupSchema *commonv1.Group) (tsdb.Database, error) {
	return tsdb.OpenDatabase(
		context.WithValue(context.Background(), common.PositionKey, common.Position{
			Module:   "measure",
			Database: groupSchema.Metadata.Name,
		}),
		tsdb.DatabaseOpts{
			Location: path.Join(s.path, groupSchema.Metadata.Name),
			ShardNum: groupSchema.ResourceOpts.ShardNum,
			EncodingMethod: tsdb.EncodingMethod{
				EncoderPool: newEncoderPool(plainChunkSize, intChunkSize, s.l),
				DecoderPool: newDecoderPool(plainChunkSize, intChunkSize, s.l),
			},
		})
}
