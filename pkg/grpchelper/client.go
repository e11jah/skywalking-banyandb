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
package grpchelper

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/apache/skywalking-banyandb/pkg/logger"
)

var l = logger.GetLogger()

func Conn(addr string, connTimeout time.Duration) (*grpc.ClientConn, error) {
	opts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	connStart := time.Now()
	dialCtx, dialCancel := context.WithTimeout(context.Background(), connTimeout)
	defer dialCancel()
	conn, err := grpc.DialContext(dialCtx, addr, opts...)
	if err != nil {
		if err == context.DeadlineExceeded {
			l.Warn().Str("addr", addr).Dur("timeout", connTimeout).Msg("timeout: failed to connect service")
		} else {
			l.Warn().Str("addr", addr).Err(err).Msg("error: failed to connect service")
		}
		return nil, err
	}
	connDuration := time.Since(connStart)
	l.Info().Dur("conn", connDuration).Msg("time elapsed")
	return conn, nil
}

func Request(ctx context.Context, rpcTimeout time.Duration, fn func(rpcCtx context.Context) error) error {
	rpcStart := time.Now()
	rpcCtx, rpcCancel := context.WithTimeout(ctx, rpcTimeout)
	defer rpcCancel()
	rpcCtx = metadata.NewOutgoingContext(rpcCtx, make(metadata.MD))

	err := fn(rpcCtx)
	if err != nil {
		if stat, ok := status.FromError(err); ok && stat.Code() == codes.Unimplemented {
			l.Warn().Str("stat", stat.Message()).Msg("error: this server does not implement the service")
		} else if stat, ok := status.FromError(err); ok && stat.Code() == codes.DeadlineExceeded {
			l.Warn().Dur("rpcTimeout", rpcTimeout).Msg("timeout: rpc did not complete within")
		} else {
			l.Error().Err(err).Msg("error: rpc failed:")
		}
		return err
	}
	rpcDuration := time.Since(rpcStart)
	l.Info().Dur("rpc", rpcDuration).Msg("time elapsed")
	return nil
}
