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
	"io"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/apache/skywalking-banyandb/api/data"
	measurev1 "github.com/apache/skywalking-banyandb/api/proto/banyandb/measure/v1"
	"github.com/apache/skywalking-banyandb/banyand/tsdb"
	"github.com/apache/skywalking-banyandb/pkg/bus"
	"github.com/apache/skywalking-banyandb/pkg/timestamp"
)

type measureService struct {
	*discoveryService
	measurev1.UnimplementedMeasureServiceServer
}

func (ms *measureService) Write(measure measurev1.MeasureService_WriteServer) error {
	reply := func() error {
		if err := measure.Send(&measurev1.WriteResponse{}); err != nil {
			return err
		}
		return nil
	}
	for {
		writeRequest, err := measure.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if errTime := timestamp.CheckPb(writeRequest.DataPoint.Timestamp); errTime != nil {
			ms.log.Error().Err(errTime).Msg("the data point time is invalid")
			if errResp := reply(); errResp != nil {
				return errResp
			}
			continue
		}
		entity, shardID, err := ms.navigate(writeRequest.GetMetadata(), writeRequest.GetDataPoint().GetTagFamilies())
		if err != nil {
			ms.log.Error().Err(err).Msg("failed to navigate to the write target")
			if errResp := reply(); errResp != nil {
				return errResp
			}
			continue
		}
		message := bus.NewMessage(bus.MessageID(time.Now().UnixNano()), &measurev1.InternalWriteRequest{
			Request:    writeRequest,
			ShardId:    uint32(shardID),
			SeriesHash: tsdb.HashEntity(entity),
		})
		_, errWritePub := ms.pipeline.Publish(data.TopicMeasureWrite, message)
		if errWritePub != nil {
			ms.log.Error().Err(errWritePub).Msg("failed to send a message")
			if errResp := reply(); errResp != nil {
				return errResp
			}
			continue
		}
		if errSend := reply(); errSend != nil {
			return errSend
		}
	}
}

func (ms *measureService) Query(_ context.Context, entityCriteria *measurev1.QueryRequest) (*measurev1.QueryResponse, error) {
	if err := timestamp.CheckTimeRange(entityCriteria.GetTimeRange()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v is invalid :%s", entityCriteria.GetTimeRange(), err)
	}
	message := bus.NewMessage(bus.MessageID(time.Now().UnixNano()), entityCriteria)
	feat, errQuery := ms.pipeline.Publish(data.TopicMeasureQuery, message)
	if errQuery != nil {
		return nil, errQuery
	}
	msg, errFeat := feat.Get()
	if errFeat != nil {
		return nil, errFeat
	}
	queryMsg, ok := msg.Data().([]*measurev1.DataPoint)
	if !ok {
		return nil, ErrQueryMsg
	}
	return &measurev1.QueryResponse{DataPoints: queryMsg}, nil
}

//TODO: implement topN
