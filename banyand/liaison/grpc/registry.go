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

package grpc

import (
	"context"

	databasev1 "github.com/apache/skywalking-banyandb/api/proto/banyandb/database/v1"
	"github.com/apache/skywalking-banyandb/banyand/metadata"
	"github.com/apache/skywalking-banyandb/banyand/metadata/schema"
)

type streamRegistryServer struct {
	schemaRegistry metadata.Service
	databasev1.UnimplementedStreamRegistryServiceServer
}

func (rs *streamRegistryServer) Create(ctx context.Context,
	req *databasev1.StreamRegistryServiceCreateRequest) (*databasev1.StreamRegistryServiceCreateResponse, error) {
	if err := rs.schemaRegistry.StreamRegistry().CreateStream(ctx, req.GetStream()); err != nil {
		return nil, err
	}
	return &databasev1.StreamRegistryServiceCreateResponse{}, nil
}

func (rs *streamRegistryServer) Update(ctx context.Context,
	req *databasev1.StreamRegistryServiceUpdateRequest) (*databasev1.StreamRegistryServiceUpdateResponse, error) {
	if err := rs.schemaRegistry.StreamRegistry().UpdateStream(ctx, req.GetStream()); err != nil {
		return nil, err
	}
	return &databasev1.StreamRegistryServiceUpdateResponse{}, nil
}

func (rs *streamRegistryServer) Delete(ctx context.Context,
	req *databasev1.StreamRegistryServiceDeleteRequest) (*databasev1.StreamRegistryServiceDeleteResponse, error) {
	ok, err := rs.schemaRegistry.StreamRegistry().DeleteStream(ctx, req.GetMetadata())
	if err != nil {
		return nil, err
	}
	return &databasev1.StreamRegistryServiceDeleteResponse{
		Deleted: ok,
	}, nil
}

func (rs *streamRegistryServer) Get(ctx context.Context,
	req *databasev1.StreamRegistryServiceGetRequest) (*databasev1.StreamRegistryServiceGetResponse, error) {
	entity, err := rs.schemaRegistry.StreamRegistry().GetStream(ctx, req.GetMetadata())
	if err != nil {
		return nil, err
	}
	return &databasev1.StreamRegistryServiceGetResponse{
		Stream: entity,
	}, nil
}

func (rs *streamRegistryServer) List(ctx context.Context,
	req *databasev1.StreamRegistryServiceListRequest) (*databasev1.StreamRegistryServiceListResponse, error) {
	entities, err := rs.schemaRegistry.StreamRegistry().ListStream(ctx, schema.ListOpt{Group: req.GetGroup()})
	if err != nil {
		return nil, err
	}
	return &databasev1.StreamRegistryServiceListResponse{
		Stream: entities,
	}, nil
}

type indexRuleBindingRegistryServer struct {
	schemaRegistry metadata.Service
	databasev1.UnimplementedIndexRuleBindingRegistryServiceServer
}

func (rs *indexRuleBindingRegistryServer) Create(ctx context.Context,
	req *databasev1.IndexRuleBindingRegistryServiceCreateRequest) (
	*databasev1.IndexRuleBindingRegistryServiceCreateResponse, error) {
	if err := rs.schemaRegistry.IndexRuleBindingRegistry().CreateIndexRuleBinding(ctx, req.GetIndexRuleBinding()); err != nil {
		return nil, err
	}
	return &databasev1.IndexRuleBindingRegistryServiceCreateResponse{}, nil
}

func (rs *indexRuleBindingRegistryServer) Update(ctx context.Context,
	req *databasev1.IndexRuleBindingRegistryServiceUpdateRequest) (
	*databasev1.IndexRuleBindingRegistryServiceUpdateResponse, error) {
	if err := rs.schemaRegistry.IndexRuleBindingRegistry().UpdateIndexRuleBinding(ctx, req.GetIndexRuleBinding()); err != nil {
		return nil, err
	}
	return &databasev1.IndexRuleBindingRegistryServiceUpdateResponse{}, nil
}

func (rs *indexRuleBindingRegistryServer) Delete(ctx context.Context,
	req *databasev1.IndexRuleBindingRegistryServiceDeleteRequest) (
	*databasev1.IndexRuleBindingRegistryServiceDeleteResponse, error) {
	ok, err := rs.schemaRegistry.IndexRuleBindingRegistry().DeleteIndexRuleBinding(ctx, req.GetMetadata())
	if err != nil {
		return nil, err
	}
	return &databasev1.IndexRuleBindingRegistryServiceDeleteResponse{
		Deleted: ok,
	}, nil
}

func (rs *indexRuleBindingRegistryServer) Get(ctx context.Context,
	req *databasev1.IndexRuleBindingRegistryServiceGetRequest) (
	*databasev1.IndexRuleBindingRegistryServiceGetResponse, error) {
	entity, err := rs.schemaRegistry.IndexRuleBindingRegistry().GetIndexRuleBinding(ctx, req.GetMetadata())
	if err != nil {
		return nil, err
	}
	return &databasev1.IndexRuleBindingRegistryServiceGetResponse{
		IndexRuleBinding: entity,
	}, nil
}

func (rs *indexRuleBindingRegistryServer) List(ctx context.Context,
	req *databasev1.IndexRuleBindingRegistryServiceListRequest) (
	*databasev1.IndexRuleBindingRegistryServiceListResponse, error) {
	entities, err := rs.schemaRegistry.IndexRuleBindingRegistry().
		ListIndexRuleBinding(ctx, schema.ListOpt{Group: req.GetGroup()})
	if err != nil {
		return nil, err
	}
	return &databasev1.IndexRuleBindingRegistryServiceListResponse{
		IndexRuleBinding: entities,
	}, nil
}

type indexRuleRegistryServer struct {
	schemaRegistry metadata.Service
	databasev1.UnimplementedIndexRuleRegistryServiceServer
}

func (rs *indexRuleRegistryServer) Create(ctx context.Context, req *databasev1.IndexRuleRegistryServiceCreateRequest) (
	*databasev1.IndexRuleRegistryServiceCreateResponse, error) {
	if err := rs.schemaRegistry.IndexRuleRegistry().CreateIndexRule(ctx, req.GetIndexRule()); err != nil {
		return nil, err
	}
	return &databasev1.IndexRuleRegistryServiceCreateResponse{}, nil
}

func (rs *indexRuleRegistryServer) Update(ctx context.Context, req *databasev1.IndexRuleRegistryServiceUpdateRequest) (
	*databasev1.IndexRuleRegistryServiceUpdateResponse, error) {
	if err := rs.schemaRegistry.IndexRuleRegistry().UpdateIndexRule(ctx, req.GetIndexRule()); err != nil {
		return nil, err
	}
	return &databasev1.IndexRuleRegistryServiceUpdateResponse{}, nil
}

func (rs *indexRuleRegistryServer) Delete(ctx context.Context, req *databasev1.IndexRuleRegistryServiceDeleteRequest) (
	*databasev1.IndexRuleRegistryServiceDeleteResponse, error) {
	ok, err := rs.schemaRegistry.IndexRuleRegistry().DeleteIndexRule(ctx, req.GetMetadata())
	if err != nil {
		return nil, err
	}
	return &databasev1.IndexRuleRegistryServiceDeleteResponse{
		Deleted: ok,
	}, nil
}

func (rs *indexRuleRegistryServer) Get(ctx context.Context, req *databasev1.IndexRuleRegistryServiceGetRequest) (
	*databasev1.IndexRuleRegistryServiceGetResponse, error) {
	entity, err := rs.schemaRegistry.IndexRuleRegistry().GetIndexRule(ctx, req.GetMetadata())
	if err != nil {
		return nil, err
	}
	return &databasev1.IndexRuleRegistryServiceGetResponse{
		IndexRule: entity,
	}, nil
}

func (rs *indexRuleRegistryServer) List(ctx context.Context, req *databasev1.IndexRuleRegistryServiceListRequest) (
	*databasev1.IndexRuleRegistryServiceListResponse, error) {
	entities, err := rs.schemaRegistry.IndexRuleRegistry().ListIndexRule(ctx, schema.ListOpt{Group: req.GetGroup()})
	if err != nil {
		return nil, err
	}
	return &databasev1.IndexRuleRegistryServiceListResponse{
		IndexRule: entities,
	}, nil
}

type measureRegistryServer struct {
	schemaRegistry metadata.Service
	databasev1.UnimplementedMeasureRegistryServiceServer
}

func (rs *measureRegistryServer) Create(ctx context.Context, req *databasev1.MeasureRegistryServiceCreateRequest) (
	*databasev1.MeasureRegistryServiceCreateResponse, error) {
	if err := rs.schemaRegistry.MeasureRegistry().CreateMeasure(ctx, req.GetMeasure()); err != nil {
		return nil, err
	}
	return &databasev1.MeasureRegistryServiceCreateResponse{}, nil
}

func (rs *measureRegistryServer) Update(ctx context.Context, req *databasev1.MeasureRegistryServiceUpdateRequest) (
	*databasev1.MeasureRegistryServiceUpdateResponse, error) {
	if err := rs.schemaRegistry.MeasureRegistry().UpdateMeasure(ctx, req.GetMeasure()); err != nil {
		return nil, err
	}
	return &databasev1.MeasureRegistryServiceUpdateResponse{}, nil
}

func (rs *measureRegistryServer) Delete(ctx context.Context, req *databasev1.MeasureRegistryServiceDeleteRequest) (
	*databasev1.MeasureRegistryServiceDeleteResponse, error) {
	ok, err := rs.schemaRegistry.MeasureRegistry().DeleteMeasure(ctx, req.GetMetadata())
	if err != nil {
		return nil, err
	}
	return &databasev1.MeasureRegistryServiceDeleteResponse{
		Deleted: ok,
	}, nil
}

func (rs *measureRegistryServer) Get(ctx context.Context, req *databasev1.MeasureRegistryServiceGetRequest) (
	*databasev1.MeasureRegistryServiceGetResponse, error) {
	entity, err := rs.schemaRegistry.MeasureRegistry().GetMeasure(ctx, req.GetMetadata())
	if err != nil {
		return nil, err
	}
	return &databasev1.MeasureRegistryServiceGetResponse{
		Measure: entity,
	}, nil
}

func (rs *measureRegistryServer) List(ctx context.Context, req *databasev1.MeasureRegistryServiceListRequest) (
	*databasev1.MeasureRegistryServiceListResponse, error) {
	entities, err := rs.schemaRegistry.MeasureRegistry().ListMeasure(ctx, schema.ListOpt{Group: req.GetGroup()})
	if err != nil {
		return nil, err
	}
	return &databasev1.MeasureRegistryServiceListResponse{
		Measure: entities,
	}, nil
}

type groupRegistryServer struct {
	schemaRegistry metadata.Service
	databasev1.UnimplementedGroupRegistryServiceServer
}

func (rs *groupRegistryServer) Create(ctx context.Context, req *databasev1.GroupRegistryServiceCreateRequest) (
	*databasev1.GroupRegistryServiceCreateResponse, error) {
	if err := rs.schemaRegistry.GroupRegistry().CreateGroup(ctx, req.GetGroup()); err != nil {
		return nil, err
	}
	return &databasev1.GroupRegistryServiceCreateResponse{}, nil
}

func (rs *groupRegistryServer) Update(ctx context.Context, req *databasev1.GroupRegistryServiceUpdateRequest) (
	*databasev1.GroupRegistryServiceUpdateResponse, error) {
	if err := rs.schemaRegistry.GroupRegistry().UpdateGroup(ctx, req.GetGroup()); err != nil {
		return nil, err
	}
	return &databasev1.GroupRegistryServiceUpdateResponse{}, nil
}

func (rs *groupRegistryServer) Delete(ctx context.Context, req *databasev1.GroupRegistryServiceDeleteRequest) (
	*databasev1.GroupRegistryServiceDeleteResponse, error) {
	deleted, err := rs.schemaRegistry.GroupRegistry().DeleteGroup(ctx, req.GetGroup())
	if err != nil {
		return nil, err
	}
	return &databasev1.GroupRegistryServiceDeleteResponse{
		Deleted: deleted,
	}, nil
}

func (rs *groupRegistryServer) Get(ctx context.Context, req *databasev1.GroupRegistryServiceGetRequest) (
	*databasev1.GroupRegistryServiceGetResponse, error) {
	g, err := rs.schemaRegistry.GroupRegistry().GetGroup(ctx, req.GetGroup())
	if err != nil {
		return nil, err
	}
	return &databasev1.GroupRegistryServiceGetResponse{
		Group: g,
	}, nil
}

func (rs *groupRegistryServer) List(ctx context.Context, req *databasev1.GroupRegistryServiceListRequest) (
	*databasev1.GroupRegistryServiceListResponse, error) {
	groups, err := rs.schemaRegistry.GroupRegistry().ListGroup(ctx)
	if err != nil {
		return nil, err
	}
	return &databasev1.GroupRegistryServiceListResponse{
		Group: groups,
	}, nil
}

type topNAggregationRegistryServer struct {
	schemaRegistry metadata.Service
	databasev1.UnimplementedTopNAggregationRegistryServiceServer
}

func (ts *topNAggregationRegistryServer) Create(ctx context.Context,
	req *databasev1.TopNAggregationRegistryServiceCreateRequest) (*databasev1.TopNAggregationRegistryServiceCreateResponse, error) {
	if err := ts.schemaRegistry.TopNAggregationRegistry().CreateTopNAggregation(ctx, req.GetTopNAggregation()); err != nil {
		return nil, err
	}
	return &databasev1.TopNAggregationRegistryServiceCreateResponse{}, nil
}

func (ts *topNAggregationRegistryServer) Update(ctx context.Context,
	req *databasev1.TopNAggregationRegistryServiceUpdateRequest) (*databasev1.TopNAggregationRegistryServiceUpdateResponse, error) {
	if err := ts.schemaRegistry.TopNAggregationRegistry().UpdateTopNAggregation(ctx, req.GetTopNAggregation()); err != nil {
		return nil, err
	}
	return &databasev1.TopNAggregationRegistryServiceUpdateResponse{}, nil
}

func (ts *topNAggregationRegistryServer) Delete(ctx context.Context,
	req *databasev1.TopNAggregationRegistryServiceDeleteRequest) (*databasev1.TopNAggregationRegistryServiceDeleteResponse, error) {
	ok, err := ts.schemaRegistry.TopNAggregationRegistry().DeleteTopNAggregation(ctx, req.GetMetadata())
	if err != nil {
		return nil, err
	}
	return &databasev1.TopNAggregationRegistryServiceDeleteResponse{
		Deleted: ok,
	}, nil
}

func (ts *topNAggregationRegistryServer) Get(ctx context.Context,
	req *databasev1.TopNAggregationRegistryServiceGetRequest) (*databasev1.TopNAggregationRegistryServiceGetResponse, error) {
	entity, err := ts.schemaRegistry.TopNAggregationRegistry().GetTopNAggregation(ctx, req.GetMetadata())
	if err != nil {
		return nil, err
	}
	return &databasev1.TopNAggregationRegistryServiceGetResponse{
		TopNAggregation: entity,
	}, nil
}

func (ts *topNAggregationRegistryServer) List(ctx context.Context,
	req *databasev1.TopNAggregationRegistryServiceListRequest) (*databasev1.TopNAggregationRegistryServiceListResponse, error) {
	entities, err := ts.schemaRegistry.TopNAggregationRegistry().ListTopNAggregation(ctx, schema.ListOpt{Group: req.GetGroup()})
	if err != nil {
		return nil, err
	}
	return &databasev1.TopNAggregationRegistryServiceListResponse{
		TopNAggregation: entities,
	}, nil
}
