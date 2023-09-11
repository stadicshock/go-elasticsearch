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
//
// Code generated from specification version 8.10.0: DO NOT EDIT

package esapi

import (
	"context"
	"io"
	"net/http"
	"strings"
)

func newFleetPostSecretFunc(t Transport) FleetPostSecret {
	return func(body io.Reader, o ...func(*FleetPostSecretRequest)) (*Response, error) {
		var r = FleetPostSecretRequest{Body: body}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// FleetPostSecret creates a secret stored by Fleet.
//
// This API is experimental.
type FleetPostSecret func(body io.Reader, o ...func(*FleetPostSecretRequest)) (*Response, error)

// FleetPostSecretRequest configures the Fleet Post Secret API request.
type FleetPostSecretRequest struct {
	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do executes the request and returns response or error.
func (r FleetPostSecretRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "POST"

	path.Grow(7 + len("/_fleet/secret"))
	path.WriteString("http://")
	path.WriteString("/_fleet/secret")

	params = make(map[string]string)

	if r.Pretty {
		params["pretty"] = "true"
	}

	if r.Human {
		params["human"] = "true"
	}

	if r.ErrorTrace {
		params["error_trace"] = "true"
	}

	if len(r.FilterPath) > 0 {
		params["filter_path"] = strings.Join(r.FilterPath, ",")
	}

	req, err := newRequest(method, path.String(), r.Body)
	if err != nil {
		return nil, err
	}

	if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if len(r.Header) > 0 {
		if len(req.Header) == 0 {
			req.Header = r.Header
		} else {
			for k, vv := range r.Header {
				for _, v := range vv {
					req.Header.Add(k, v)
				}
			}
		}
	}

	if r.Body != nil && req.Header.Get(headerContentType) == "" {
		req.Header[headerContentType] = headerContentTypeJSON
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := transport.Perform(req)
	if err != nil {
		return nil, err
	}

	response := Response{
		StatusCode: res.StatusCode,
		Body:       res.Body,
		Header:     res.Header,
	}

	return &response, nil
}

// WithContext sets the request context.
func (f FleetPostSecret) WithContext(v context.Context) func(*FleetPostSecretRequest) {
	return func(r *FleetPostSecretRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f FleetPostSecret) WithPretty() func(*FleetPostSecretRequest) {
	return func(r *FleetPostSecretRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f FleetPostSecret) WithHuman() func(*FleetPostSecretRequest) {
	return func(r *FleetPostSecretRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f FleetPostSecret) WithErrorTrace() func(*FleetPostSecretRequest) {
	return func(r *FleetPostSecretRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f FleetPostSecret) WithFilterPath(v ...string) func(*FleetPostSecretRequest) {
	return func(r *FleetPostSecretRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f FleetPostSecret) WithHeader(h map[string]string) func(*FleetPostSecretRequest) {
	return func(r *FleetPostSecretRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f FleetPostSecret) WithOpaqueID(s string) func(*FleetPostSecretRequest) {
	return func(r *FleetPostSecretRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}