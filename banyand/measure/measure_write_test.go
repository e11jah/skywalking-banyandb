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

package measure_test

import (
	"embed"
	"encoding/json"
	"time"

	"github.com/golang/protobuf/jsonpb"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"google.golang.org/protobuf/types/known/timestamppb"

	commonv1 "github.com/apache/skywalking-banyandb/api/proto/banyandb/common/v1"
	measurev1 "github.com/apache/skywalking-banyandb/api/proto/banyandb/measure/v1"
	modelv1 "github.com/apache/skywalking-banyandb/api/proto/banyandb/model/v1"
	"github.com/apache/skywalking-banyandb/banyand/measure"
	"github.com/apache/skywalking-banyandb/banyand/tsdb"
	"github.com/apache/skywalking-banyandb/pkg/timestamp"
)

var _ = Describe("Write and Update service_cpm_minute", func() {
	var svcs *services
	var deferFn func()
	var measure measure.Measure
	var baseTime time.Time

	var count = func() (num int) {
		//Retrieve all shards
		shards, err := measure.Shards(nil)
		Expect(err).ShouldNot(HaveOccurred())
		for _, shard := range shards {
			sl, err := shard.Series().List(tsdb.NewPath([]tsdb.Entry{tsdb.AnyEntry}))
			Expect(err).ShouldNot(HaveOccurred())
			for _, series := range sl {
				seriesSpan, err := series.Span(timestamp.NewInclusiveTimeRangeDuration(baseTime, 1*time.Hour))
				defer func(seriesSpan tsdb.SeriesSpan) {
					_ = seriesSpan.Close()
				}(seriesSpan)
				Expect(err).ShouldNot(HaveOccurred())
				seeker, err := seriesSpan.SeekerBuilder().OrderByTime(modelv1.Sort_SORT_ASC).Build()
				Expect(err).ShouldNot(HaveOccurred())
				iters, err := seeker.Seek()
				Expect(err).ShouldNot(HaveOccurred())
				for _, iter := range iters {
					defer func(iterator tsdb.Iterator) {
						Expect(iterator.Close()).ShouldNot(HaveOccurred())
					}(iter)
					for iter.Next() {
						num++
					}
				}
			}
		}
		return num
	}

	BeforeEach(func() {
		svcs, deferFn = setUp()
		var err error
		measure, err = svcs.measure.Measure(&commonv1.Metadata{
			Name:  "service_cpm_minute",
			Group: "sw_metric",
		})
		Expect(err).ShouldNot(HaveOccurred())
	})
	AfterEach(func() {
		deferFn()
	})
	DescribeTable("writes", func(metadata *commonv1.Metadata, dataFile string, expectDatapointsNum int) {
		var err error
		measure, err = svcs.measure.Measure(metadata)
		Expect(err).ShouldNot(HaveOccurred())
		baseTime = writeData(dataFile, measure)
		Expect(count()).To(Equal(expectDatapointsNum))

	},
		Entry("service_cpm_minute", &commonv1.Metadata{
			Name:  "service_cpm_minute",
			Group: "sw_metric",
		}, "service_cpm_minute_data.json", 3),
		Entry("service_traffic", &commonv1.Metadata{
			Name:  "service_traffic",
			Group: "sw_metric",
		}, "service_traffic_data.json", 3),
	)
	It("updates", func() {
		metadata := &commonv1.Metadata{
			Name:  "service_cpm_minute",
			Group: "sw_metric",
		}
		var err error
		measure, err = svcs.measure.Measure(metadata)
		Expect(err).ShouldNot(HaveOccurred())
		baseTime = writeData("service_cpm_minute_data.json", measure)
		Expect(count()).To(Equal(3))
		writeDataWithBaseTime(baseTime, "service_cpm_minute_data1.json", measure)
		Expect(count()).To(Equal(3))
	})

})

//go:embed testdata/*.json
var dataFS embed.FS

func writeData(dataFile string, measure measure.Measure) (baseTime time.Time) {
	baseTime = timestamp.NowMilli()
	writeDataWithBaseTime(baseTime, dataFile, measure)
	return baseTime
}

func writeDataWithBaseTime(baseTime time.Time, dataFile string, measure measure.Measure) time.Time {
	var templates []interface{}
	content, err := dataFS.ReadFile("testdata/" + dataFile)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(json.Unmarshal(content, &templates)).ShouldNot(HaveOccurred())
	for i, template := range templates {
		rawDataPointValue, errMarshal := json.Marshal(template)
		Expect(errMarshal).ShouldNot(HaveOccurred())
		dataPointValue := &measurev1.DataPointValue{}
		dataPointValue.Timestamp = timestamppb.New(baseTime.Add(time.Duration(i) * time.Minute))
		Expect(jsonpb.UnmarshalString(string(rawDataPointValue), dataPointValue)).ShouldNot(HaveOccurred())
		errInner := measure.Write(dataPointValue)
		Expect(errInner).ShouldNot(HaveOccurred())
	}
	return baseTime
}
