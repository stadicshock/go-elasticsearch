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

package clearcachedservicetokens

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package clearcachedservicetokens
//
// https://github.com/elastic/elasticsearch-specification/blob/24afbdf78c21fde141eb2cad34491d952bd6daa8/specification/security/clear_cached_service_tokens/ClearCachedServiceTokensResponse.ts#L25-L32

type Response struct {
	ClusterName string                       `json:"cluster_name"`
	NodeStats   types.NodeStatistics         `json:"_nodes"`
	Nodes       map[string]types.ClusterNode `json:"nodes"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{
		Nodes: make(map[string]types.ClusterNode, 0),
	}
	return r
}
