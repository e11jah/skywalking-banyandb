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

package lsm

import (
	"go.uber.org/multierr"

	"github.com/apache/skywalking-banyandb/api/common"
	"github.com/apache/skywalking-banyandb/banyand/kv"
	"github.com/apache/skywalking-banyandb/banyand/observability"
	"github.com/apache/skywalking-banyandb/pkg/convert"
	"github.com/apache/skywalking-banyandb/pkg/index"
	"github.com/apache/skywalking-banyandb/pkg/index/metadata"
	"github.com/apache/skywalking-banyandb/pkg/logger"
)

var _ index.Store = (*store)(nil)

type store struct {
	lsm          kv.Store
	termMetadata metadata.Term
	l            *logger.Logger
}

func (*store) Flush() error {
	panic("do not call flush here. LSM index is using its own controller to flush memory data")
}

func (s *store) Stats() observability.Statistics {
	return s.lsm.Stats()
}

func (s *store) Close() error {
	return multierr.Combine(s.lsm.Close(), s.termMetadata.Close())
}

func (s *store) Write(field index.Field, itemID common.ItemID) error {
	f, err := field.Marshal(s.termMetadata)
	if err != nil {
		return err
	}
	itemIDInt := uint64(itemID)
	return s.lsm.PutWithVersion(f, convert.Uint64ToBytes(itemIDInt), itemIDInt)
}

type StoreOpts struct {
	Path   string
	Logger *logger.Logger
}

func NewStore(opts StoreOpts) (index.Store, error) {
	var err error
	var lsm kv.Store
	if lsm, err = kv.OpenStore(0, opts.Path+"/lsm", kv.StoreWithLogger(opts.Logger)); err != nil {
		return nil, err
	}
	var md metadata.Term
	if md, err = metadata.NewTerm(metadata.TermOpts{
		Path:   opts.Path + "/tmd",
		Logger: opts.Logger,
	}); err != nil {
		return nil, err
	}
	return &store{
		lsm:          lsm,
		termMetadata: md,
		l:            opts.Logger,
	}, nil
}
