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
// https://github.com/elastic/elasticsearch-specification/tree/24afbdf78c21fde141eb2cad34491d952bd6daa8

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// Influence type.
//
// https://github.com/elastic/elasticsearch-specification/blob/24afbdf78c21fde141eb2cad34491d952bd6daa8/specification/ml/_types/Anomaly.ts#L140-L143
type Influence struct {
	InfluencerFieldName   string   `json:"influencer_field_name"`
	InfluencerFieldValues []string `json:"influencer_field_values"`
}

func (s *Influence) UnmarshalJSON(data []byte) error {

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

		case "influencer_field_name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.InfluencerFieldName = o

		case "influencer_field_values":
			if err := dec.Decode(&s.InfluencerFieldValues); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewInfluence returns a Influence.
func NewInfluence() *Influence {
	r := &Influence{}

	return r
}
