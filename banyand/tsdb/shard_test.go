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
package tsdb_test

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/apache/skywalking-banyandb/api/common"
	"github.com/apache/skywalking-banyandb/banyand/tsdb"
	"github.com/apache/skywalking-banyandb/pkg/test"
	"github.com/apache/skywalking-banyandb/pkg/timestamp"
)

var _ = Describe("Shard", func() {
	Describe("Generate segments and blocks", func() {
		var tmp string
		var deferFn func()
		var shard tsdb.Shard
		var clock timestamp.MockClock

		BeforeEach(func() {
			var err error
			tmp, deferFn, err = test.NewSpace()
			Expect(err).NotTo(HaveOccurred())
			clock = timestamp.NewMockClock()
			clock.Set(time.Date(1970, 01, 01, 0, 0, 0, 0, time.Local))
			Expect(err).NotTo(HaveOccurred())
		})
		AfterEach(func() {
			shard.Close()
			deferFn()
		})
		It("generates several segments and blocks", func() {
			By("open 4 blocks")
			var err error
			shard, err = tsdb.OpenShard(timestamp.SetClock(context.Background(), clock), common.ShardID(0), tmp,
				tsdb.IntervalRule{
					Unit: tsdb.DAY,
					Num:  1,
				},
				tsdb.IntervalRule{
					Unit: tsdb.HOUR,
					Num:  12,
				},
				2,
			)
			Expect(err).NotTo(HaveOccurred())
			By("01/01 00:00 1st block is opened")
			t1 := clock.Now()
			Eventually(func() []tsdb.BlockState {
				return shard.State().Blocks
			}, defaultEventallyTimeout).Should(Equal([]tsdb.BlockState{
				{
					ID: tsdb.BlockID{
						SegID:   tsdb.GenerateInternalID(tsdb.DAY, 19700101),
						BlockID: tsdb.GenerateInternalID(tsdb.HOUR, 00),
					},
					TimeRange: timestamp.NewTimeRangeDuration(t1, 12*time.Hour, true, false),
				},
			}))
			By("01/01 10:00 2nd block is opened")
			clock.Add(10 * time.Hour)
			t2 := clock.Now().Add(2 * time.Hour)
			Eventually(func() []tsdb.BlockState {
				return shard.State().Blocks
			}, defaultEventallyTimeout).Should(Equal([]tsdb.BlockState{
				{
					ID: tsdb.BlockID{
						SegID:   tsdb.GenerateInternalID(tsdb.DAY, 19700101),
						BlockID: tsdb.GenerateInternalID(tsdb.HOUR, 00),
					},
					TimeRange: timestamp.NewTimeRangeDuration(t1, 12*time.Hour, true, false),
				},
				{
					ID: tsdb.BlockID{
						SegID:   tsdb.GenerateInternalID(tsdb.DAY, 19700101),
						BlockID: tsdb.GenerateInternalID(tsdb.HOUR, 12),
					},
					TimeRange: timestamp.NewTimeRangeDuration(t2, 12*time.Hour, true, false),
				},
			}))
			Eventually(func() []tsdb.BlockID {
				return shard.State().OpenBlocks
			}, defaultEventallyTimeout).Should(Equal([]tsdb.BlockID{}))
			By("01/01 13:00 moves to the 2nd block")
			clock.Add(3 * time.Hour)
			Eventually(func() []tsdb.BlockID {
				return shard.State().OpenBlocks
			}, defaultEventallyTimeout).Should(Equal([]tsdb.BlockID{
				{
					SegID:   tsdb.GenerateInternalID(tsdb.DAY, 19700101),
					BlockID: tsdb.GenerateInternalID(tsdb.HOUR, 00),
				},
			}))
			By("01/01 22:00 3rd block is opened")
			clock.Add(9 * time.Hour)
			t3 := clock.Now().Add(2 * time.Hour)
			Eventually(func() []tsdb.BlockState {
				return shard.State().Blocks
			}, defaultEventallyTimeout).Should(Equal([]tsdb.BlockState{
				{
					ID: tsdb.BlockID{
						SegID:   tsdb.GenerateInternalID(tsdb.DAY, 19700101),
						BlockID: tsdb.GenerateInternalID(tsdb.HOUR, 00),
					},
					TimeRange: timestamp.NewTimeRangeDuration(t1, 12*time.Hour, true, false),
				},
				{
					ID: tsdb.BlockID{
						SegID:   tsdb.GenerateInternalID(tsdb.DAY, 19700101),
						BlockID: tsdb.GenerateInternalID(tsdb.HOUR, 12),
					},
					TimeRange: timestamp.NewTimeRangeDuration(t2, 12*time.Hour, true, false),
				},
				{
					ID: tsdb.BlockID{
						SegID:   tsdb.GenerateInternalID(tsdb.DAY, 19700102),
						BlockID: tsdb.GenerateInternalID(tsdb.HOUR, 00),
					},
					TimeRange: timestamp.NewTimeRangeDuration(t3, 12*time.Hour, true, false),
				},
			}))
			By("01/02 01:00 moves to 3rd block")
			clock.Add(3 * time.Hour)
			Eventually(func() []tsdb.BlockID {
				return shard.State().OpenBlocks
			}).Should(Equal([]tsdb.BlockID{
				{
					SegID:   tsdb.GenerateInternalID(tsdb.DAY, 19700101),
					BlockID: tsdb.GenerateInternalID(tsdb.HOUR, 00),
				},
				{
					SegID:   tsdb.GenerateInternalID(tsdb.DAY, 19700101),
					BlockID: tsdb.GenerateInternalID(tsdb.HOUR, 12),
				},
			}))
			By("01/02 10:00 4th block is opened")
			clock.Add(9 * time.Hour)
			t4 := clock.Now().Add(2 * time.Hour)
			Eventually(func() []tsdb.BlockState {
				return shard.State().Blocks
			}, defaultEventallyTimeout).Should(Equal([]tsdb.BlockState{
				{
					ID: tsdb.BlockID{
						SegID:   tsdb.GenerateInternalID(tsdb.DAY, 19700101),
						BlockID: tsdb.GenerateInternalID(tsdb.HOUR, 00),
					},
					TimeRange: timestamp.NewTimeRangeDuration(t1, 12*time.Hour, true, false),
				},
				{
					ID: tsdb.BlockID{
						SegID:   tsdb.GenerateInternalID(tsdb.DAY, 19700101),
						BlockID: tsdb.GenerateInternalID(tsdb.HOUR, 12),
					},
					TimeRange: timestamp.NewTimeRangeDuration(t2, 12*time.Hour, true, false),
				},
				{
					ID: tsdb.BlockID{
						SegID:   tsdb.GenerateInternalID(tsdb.DAY, 19700102),
						BlockID: tsdb.GenerateInternalID(tsdb.HOUR, 00),
					},
					TimeRange: timestamp.NewTimeRangeDuration(t3, 12*time.Hour, true, false),
				},
				{
					ID: tsdb.BlockID{
						SegID:   tsdb.GenerateInternalID(tsdb.DAY, 19700102),
						BlockID: tsdb.GenerateInternalID(tsdb.HOUR, 12),
					},
					TimeRange: timestamp.NewTimeRangeDuration(t4, 12*time.Hour, true, false),
				},
			}))
			By("01/02 13:00 moves to 4th block")
			clock.Add(3 * time.Hour)
			Eventually(func() []tsdb.BlockID {
				return shard.State().OpenBlocks
			}, defaultEventallyTimeout).Should(Equal([]tsdb.BlockID{
				{
					SegID:   tsdb.GenerateInternalID(tsdb.DAY, 19700101),
					BlockID: tsdb.GenerateInternalID(tsdb.HOUR, 12),
				},
				{
					SegID:   tsdb.GenerateInternalID(tsdb.DAY, 19700102),
					BlockID: tsdb.GenerateInternalID(tsdb.HOUR, 00),
				},
				{
					SegID:   tsdb.GenerateInternalID(tsdb.DAY, 19700101),
					BlockID: tsdb.GenerateInternalID(tsdb.HOUR, 00),
				},
			}))
			By("01/02 22:00 5th block is opened")
			clock.Add(9 * time.Hour)
			t5 := clock.Now().Add(2 * time.Hour)
			Eventually(func() []tsdb.BlockState {
				return shard.State().Blocks
			}, defaultEventallyTimeout).Should(Equal([]tsdb.BlockState{
				{
					ID: tsdb.BlockID{
						SegID:   tsdb.GenerateInternalID(tsdb.DAY, 19700101),
						BlockID: tsdb.GenerateInternalID(tsdb.HOUR, 00),
					},
					TimeRange: timestamp.NewTimeRangeDuration(t1, 12*time.Hour, true, false),
				},
				{
					ID: tsdb.BlockID{
						SegID:   tsdb.GenerateInternalID(tsdb.DAY, 19700101),
						BlockID: tsdb.GenerateInternalID(tsdb.HOUR, 12),
					},
					TimeRange: timestamp.NewTimeRangeDuration(t2, 12*time.Hour, true, false),
				},
				{
					ID: tsdb.BlockID{
						SegID:   tsdb.GenerateInternalID(tsdb.DAY, 19700102),
						BlockID: tsdb.GenerateInternalID(tsdb.HOUR, 00),
					},
					TimeRange: timestamp.NewTimeRangeDuration(t3, 12*time.Hour, true, false),
				},
				{
					ID: tsdb.BlockID{
						SegID:   tsdb.GenerateInternalID(tsdb.DAY, 19700102),
						BlockID: tsdb.GenerateInternalID(tsdb.HOUR, 12),
					},
					TimeRange: timestamp.NewTimeRangeDuration(t4, 12*time.Hour, true, false),
				},
				{
					ID: tsdb.BlockID{
						SegID:   tsdb.GenerateInternalID(tsdb.DAY, 19700103),
						BlockID: tsdb.GenerateInternalID(tsdb.HOUR, 00),
					},
					TimeRange: timestamp.NewTimeRangeDuration(t5, 12*time.Hour, true, false),
				},
			}))
			By("01/03 01:00 close 1st block by adding 5th block")
			clock.Add(3 * time.Hour)
			Eventually(func() []tsdb.BlockID {
				return shard.State().OpenBlocks
			}, defaultEventallyTimeout).Should(Equal([]tsdb.BlockID{
				{
					SegID:   tsdb.GenerateInternalID(tsdb.DAY, 19700102),
					BlockID: tsdb.GenerateInternalID(tsdb.HOUR, 00),
				},
				{
					SegID:   tsdb.GenerateInternalID(tsdb.DAY, 19700102),
					BlockID: tsdb.GenerateInternalID(tsdb.HOUR, 12),
				},
				{
					SegID:   tsdb.GenerateInternalID(tsdb.DAY, 19700101),
					BlockID: tsdb.GenerateInternalID(tsdb.HOUR, 12),
				},
			}))
			Eventually(func() []tsdb.BlockState {
				return shard.State().Blocks
			}, defaultEventallyTimeout).Should(Equal([]tsdb.BlockState{
				{
					ID: tsdb.BlockID{
						SegID:   tsdb.GenerateInternalID(tsdb.DAY, 19700101),
						BlockID: tsdb.GenerateInternalID(tsdb.HOUR, 00),
					},
					TimeRange: timestamp.NewTimeRangeDuration(t1, 12*time.Hour, true, false),
					Closed:    true,
				},
				{
					ID: tsdb.BlockID{
						SegID:   tsdb.GenerateInternalID(tsdb.DAY, 19700101),
						BlockID: tsdb.GenerateInternalID(tsdb.HOUR, 12),
					},
					TimeRange: timestamp.NewTimeRangeDuration(t2, 12*time.Hour, true, false),
				},
				{
					ID: tsdb.BlockID{
						SegID:   tsdb.GenerateInternalID(tsdb.DAY, 19700102),
						BlockID: tsdb.GenerateInternalID(tsdb.HOUR, 00),
					},
					TimeRange: timestamp.NewTimeRangeDuration(t3, 12*time.Hour, true, false),
				},
				{
					ID: tsdb.BlockID{
						SegID:   tsdb.GenerateInternalID(tsdb.DAY, 19700102),
						BlockID: tsdb.GenerateInternalID(tsdb.HOUR, 12),
					},
					TimeRange: timestamp.NewTimeRangeDuration(t4, 12*time.Hour, true, false),
				},
				{
					ID: tsdb.BlockID{
						SegID:   tsdb.GenerateInternalID(tsdb.DAY, 19700103),
						BlockID: tsdb.GenerateInternalID(tsdb.HOUR, 00),
					},
					TimeRange: timestamp.NewTimeRangeDuration(t5, 12*time.Hour, true, false),
				},
			}))
			By("reopen 1st block")
			series, err := shard.Series().GetByID(common.SeriesID(11))
			Expect(err).NotTo(HaveOccurred())
			t1Range := timestamp.NewInclusiveTimeRangeDuration(t1, 1*time.Hour)
			span, err := series.Span(t1Range)
			Expect(err).NotTo(HaveOccurred())
			defer span.Close()
			writer, err := span.WriterBuilder().Family([]byte("test"), []byte("test")).Time(t1Range.End).Build()
			Expect(err).NotTo(HaveOccurred())
			_, err = writer.Write()
			Expect(err).NotTo(HaveOccurred())
			Eventually(func() []tsdb.BlockID {
				return shard.State().OpenBlocks
			}, defaultEventallyTimeout).Should(Equal([]tsdb.BlockID{
				{
					SegID:   tsdb.GenerateInternalID(tsdb.DAY, 19700102),
					BlockID: tsdb.GenerateInternalID(tsdb.HOUR, 12),
				},
				{
					SegID:   tsdb.GenerateInternalID(tsdb.DAY, 19700101),
					BlockID: tsdb.GenerateInternalID(tsdb.HOUR, 00),
				},
				{
					SegID:   tsdb.GenerateInternalID(tsdb.DAY, 19700102),
					BlockID: tsdb.GenerateInternalID(tsdb.HOUR, 00),
				},
			}))
			Eventually(func() []tsdb.BlockState {
				return shard.State().Blocks
			}, defaultEventallyTimeout).Should(Equal([]tsdb.BlockState{
				{
					ID: tsdb.BlockID{
						SegID:   tsdb.GenerateInternalID(tsdb.DAY, 19700101),
						BlockID: tsdb.GenerateInternalID(tsdb.HOUR, 00),
					},
					TimeRange: timestamp.NewTimeRangeDuration(t1, 12*time.Hour, true, false),
				},
				{
					ID: tsdb.BlockID{
						SegID:   tsdb.GenerateInternalID(tsdb.DAY, 19700101),
						BlockID: tsdb.GenerateInternalID(tsdb.HOUR, 12),
					},
					TimeRange: timestamp.NewTimeRangeDuration(t2, 12*time.Hour, true, false),
					Closed:    true,
				},
				{
					ID: tsdb.BlockID{
						SegID:   tsdb.GenerateInternalID(tsdb.DAY, 19700102),
						BlockID: tsdb.GenerateInternalID(tsdb.HOUR, 00),
					},
					TimeRange: timestamp.NewTimeRangeDuration(t3, 12*time.Hour, true, false),
				},
				{
					ID: tsdb.BlockID{
						SegID:   tsdb.GenerateInternalID(tsdb.DAY, 19700102),
						BlockID: tsdb.GenerateInternalID(tsdb.HOUR, 12),
					},
					TimeRange: timestamp.NewTimeRangeDuration(t4, 12*time.Hour, true, false),
				},
				{
					ID: tsdb.BlockID{
						SegID:   tsdb.GenerateInternalID(tsdb.DAY, 19700103),
						BlockID: tsdb.GenerateInternalID(tsdb.HOUR, 00),
					},
					TimeRange: timestamp.NewTimeRangeDuration(t5, 12*time.Hour, true, false),
				},
			}))
		})
	})
})
