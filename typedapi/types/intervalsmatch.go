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

// IntervalsMatch type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/_types/query_dsl/fulltext.ts#L99-L108
type IntervalsMatch struct {
	Analyzer *string          `json:"analyzer,omitempty"`
	Filter   *IntervalsFilter `json:"filter,omitempty"`
	MaxGaps  *int             `json:"max_gaps,omitempty"`
	Ordered  *bool            `json:"ordered,omitempty"`
	Query    string           `json:"query"`
	UseField *string          `json:"use_field,omitempty"`
}

// NewIntervalsMatch returns a IntervalsMatch.
func NewIntervalsMatch() *IntervalsMatch {
	r := &IntervalsMatch{}

	return r
}