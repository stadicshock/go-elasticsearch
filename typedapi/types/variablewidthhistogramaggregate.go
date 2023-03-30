// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

package types

import (
	"bytes"
	"errors"
	"io"

	"encoding/json"
)

// VariableWidthHistogramAggregate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/_types/aggregations/Aggregate.ts#L361-L363
type VariableWidthHistogramAggregate struct {
	Buckets BucketsVariableWidthHistogramBucket `json:"buckets"`
	Meta    map[string]json.RawMessage          `json:"meta,omitempty"`
}

func (s *VariableWidthHistogramAggregate) UnmarshalJSON(data []byte) error {
	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "buckets":

			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			source := bytes.NewReader(rawMsg)
			localDec := json.NewDecoder(source)
			switch rawMsg[0] {

			case '{':
				o := make(map[string]VariableWidthHistogramBucket, 0)
				localDec.Decode(&o)
				s.Buckets = o

			case '[':
				o := []VariableWidthHistogramBucket{}
				localDec.Decode(&o)
				s.Buckets = o
			}

		case "meta":
			if err := dec.Decode(&s.Meta); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewVariableWidthHistogramAggregate returns a VariableWidthHistogramAggregate.
func NewVariableWidthHistogramAggregate() *VariableWidthHistogramAggregate {
	r := &VariableWidthHistogramAggregate{}

	return r
}