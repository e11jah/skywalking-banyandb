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
package logical

import (
	"container/heap"
	"fmt"
	"sort"

	measurev1 "github.com/apache/skywalking-banyandb/api/proto/banyandb/measure/v1"
)

type TopElement struct {
	dp    *measurev1.DataPoint
	value int64
}

func NewTopElement(dp *measurev1.DataPoint, value int64) TopElement {
	return TopElement{
		dp:    dp,
		value: value,
	}
}

func (e TopElement) Val() int64 {
	return e.value
}

type topSortedList struct {
	elements []TopElement
	reverted bool
}

func (h topSortedList) Len() int {
	return len(h.elements)
}

func (h topSortedList) Less(i, j int) bool {
	if h.reverted {
		return h.elements[i].value < h.elements[j].value
	}
	return h.elements[i].value > h.elements[j].value
}

func (h *topSortedList) Swap(i, j int) {
	h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
}

type topHeap struct {
	elements []TopElement
	reverted bool
}

func (h topHeap) Len() int {
	return len(h.elements)
}

func (h topHeap) Less(i, j int) bool {
	if h.reverted {
		return h.elements[i].value > h.elements[j].value
	}
	return h.elements[i].value < h.elements[j].value
}

func (h *topHeap) Swap(i, j int) {
	h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
}

func (h *topHeap) Push(x interface{}) {
	h.elements = append(h.elements, x.(TopElement))
}

func (h *topHeap) Pop() interface{} {
	var e TopElement
	e, h.elements = h.elements[len(h.elements)-1], h.elements[:len(h.elements)-1]
	return e
}

type TopQueue struct {
	n  int
	th topHeap
}

func NewTopQueue(n int, reverted bool) *TopQueue {
	return &TopQueue{
		n: n,
		th: topHeap{
			elements: make([]TopElement, 0, n),
			reverted: reverted,
		},
	}
}

func (s *TopQueue) Insert(element TopElement) bool {
	if len(s.th.elements) < s.n {
		heap.Push(&s.th, element)
		return true
	}
	minElement := heap.Pop(&s.th).(TopElement)
	if s.th.reverted {
		if minElement.value < element.value {
			heap.Push(&s.th, minElement)
			return false
		}
	} else {
		if minElement.value > element.value {
			heap.Push(&s.th, minElement)
			return false
		}
	}
	heap.Push(&s.th, element)
	return true
}

func (s *TopQueue) Purge() {
	s.th.elements = s.th.elements[:0]
}

func (s *TopQueue) Elements() []TopElement {
	l := &topSortedList{
		elements: append([]TopElement(nil), s.th.elements...),
		reverted: s.th.reverted,
	}
	sort.Sort(l)
	return l.elements
}

func (s TopQueue) String() string {
	if s.th.reverted {
		return fmt.Sprintf("bottom(%d)", s.n)
	}
	return fmt.Sprintf("top(%d)", s.n)
}

func (s *TopQueue) Equal(other *TopQueue) bool {
	return s.th.reverted == other.th.reverted && s.n == other.n
}
