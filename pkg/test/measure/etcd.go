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
	"context"
	"embed"
	"fmt"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/multierr"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	commonv1 "github.com/apache/skywalking-banyandb/api/proto/banyandb/common/v1"
	databasev1 "github.com/apache/skywalking-banyandb/api/proto/banyandb/database/v1"
	"github.com/apache/skywalking-banyandb/banyand/metadata/schema"
	"github.com/apache/skywalking-banyandb/pkg/grpchelper"
)

const (
	groupDir            = "testdata/groups"
	measureDir          = "testdata/measures"
	indexRuleDir        = "testdata/index_rules"
	indexRuleBindingDir = "testdata/index_rule_bindings"
)

var (
	//go:embed testdata/*
	store embed.FS
)

func PreloadSchema(e schema.Registry) error {
	if err := loadSchema(groupDir, &commonv1.Group{}, func(group proto.Message) error {
		return e.CreateGroup(context.TODO(), group.(*commonv1.Group))
	}); err != nil {
		return errors.WithStack(err)
	}
	if err := loadSchema(measureDir, &databasev1.Measure{}, func(measure proto.Message) error {
		return e.CreateMeasure(context.TODO(), measure.(*databasev1.Measure))
	}); err != nil {
		return errors.WithStack(err)
	}
	if err := loadSchema(indexRuleDir, &databasev1.IndexRule{}, func(indexRule proto.Message) error {
		return e.CreateIndexRule(context.TODO(), indexRule.(*databasev1.IndexRule))
	}); err != nil {
		return errors.WithStack(err)
	}
	if err := loadSchema(indexRuleBindingDir, &databasev1.IndexRuleBinding{}, func(indexRuleBinding proto.Message) error {
		return e.CreateIndexRuleBinding(context.TODO(), indexRuleBinding.(*databasev1.IndexRuleBinding))
	}); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func RandomTempDir() string {
	return path.Join(os.TempDir(), fmt.Sprintf("banyandb-embed-etcd-%s", uuid.New().String()))
}

func loadSchema(dir string, resource proto.Message, loadFn func(resource proto.Message) error) error {
	entries, err := store.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		data, err := store.ReadFile(path.Join(dir, entry.Name()))
		if err != nil {
			return err
		}
		resource.ProtoReflect().Descriptor().RequiredNumbers()
		if err := protojson.Unmarshal(data, resource); err != nil {
			return err
		}
		if err := loadFn(resource); err != nil {
			return err
		}
	}
	return nil
}

var rpcTimeout = 10 * time.Second

func RegisterForNew(addr string, metricNum int) error {
	conn, err := grpchelper.Conn(addr, 1*time.Second)
	if err != nil {
		return err
	}
	defer conn.Close()

	ctx := context.Background()

	if err := loadSchema(groupDir, &commonv1.Group{}, func(group proto.Message) error {
		return grpchelper.Request(ctx, rpcTimeout, func(rpcCtx context.Context) (err error) {
			_, err = databasev1.NewGroupRegistryServiceClient(conn).
				Create(rpcCtx, &databasev1.GroupRegistryServiceCreateRequest{
					Group: group.(*commonv1.Group),
				})
			return err
		})
	}); err != nil {
		return errors.WithStack(err)
	}
	if err := loadSchema(measureDir, &databasev1.Measure{}, func(measure proto.Message) error {
		var err error
		name := measure.(*databasev1.Measure).Metadata.Name
		num := metricNum
		if name != "service_cpm_minute" {
			num = 1
		}
		for i := 0; i < num; i++ {
			err = multierr.Append(err, grpchelper.Request(ctx, rpcTimeout, func(rpcCtx context.Context) (err error) {
				m := proto.Clone(measure).(*databasev1.Measure)
				if i > 0 {
					m.Metadata.Name = m.Metadata.Name + "_" + strconv.Itoa(i)
				}
				_, err = databasev1.NewMeasureRegistryServiceClient(conn).
					Create(rpcCtx, &databasev1.MeasureRegistryServiceCreateRequest{
						Measure: m,
					})
				return err
			}))
		}
		return err
	}); err != nil {
		return errors.WithStack(err)
	}

	if err := loadSchema(indexRuleDir, &databasev1.IndexRule{}, func(indexRule proto.Message) error {
		return grpchelper.Request(ctx, rpcTimeout, func(rpcCtx context.Context) (err error) {
			_, err = databasev1.NewIndexRuleRegistryServiceClient(conn).
				Create(rpcCtx, &databasev1.IndexRuleRegistryServiceCreateRequest{
					IndexRule: indexRule.(*databasev1.IndexRule),
				})
			return err
		})
	}); err != nil {
		return errors.WithStack(err)
	}
	if err := loadSchema(indexRuleBindingDir, &databasev1.IndexRuleBinding{}, func(indexRuleBinding proto.Message) error {
		return grpchelper.Request(ctx, rpcTimeout, func(rpcCtx context.Context) (err error) {
			_, err = databasev1.NewIndexRuleBindingRegistryServiceClient(conn).
				Create(rpcCtx, &databasev1.IndexRuleBindingRegistryServiceCreateRequest{
					IndexRuleBinding: indexRuleBinding.(*databasev1.IndexRuleBinding),
				})
			return err
		})
	}); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
