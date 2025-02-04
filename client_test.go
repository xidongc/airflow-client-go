// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package client_test

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xidongc/airflow-client-go/airflow"
)

func TestBasicAuth(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(
			t,
			fmt.Sprintf(
				"Basic %s",
				base64.StdEncoding.EncodeToString([]byte("username:password")),
			),
			req.Header.Get("Authorization"))
		assert.Equal(t, "/api/v1/variables/foo", req.URL.String())
		rw.Header().Set("Content-Type", "application/json")
		rw.Write([]byte(`{"key": "foo", "value": "bar"}`))
	}))
	// Close the server when test finishes
	defer server.Close()
	url, err := url.Parse(server.URL)
	assert.Equal(t, nil, err)

	conf := airflow.NewConfiguration()
	conf.Host = url.Host
	conf.Scheme = "http"
	cli := airflow.NewAPIClient(conf)

	cred := airflow.BasicAuth{
		UserName: "username",
		Password: "password",
	}
	ctx := context.WithValue(context.Background(), airflow.ContextBasicAuth, cred)

	variable, _, err := cli.VariableApi.GetVariable(ctx, "foo").Execute()
	assert.Equal(t, nil, err)
	assert.Equal(t, *variable.Key, "foo")
	assert.Equal(t, *variable.Value, "bar")
}
