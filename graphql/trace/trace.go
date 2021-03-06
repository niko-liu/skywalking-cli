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

package trace

import (
	"github.com/machinebox/graphql"
	"github.com/urfave/cli"

	"github.com/apache/skywalking-cli/assets"
	"github.com/apache/skywalking-cli/graphql/client"
	"github.com/apache/skywalking-cli/graphql/schema"
)

func Trace(ctx *cli.Context, traceID string) schema.Trace {
	var response map[string]schema.Trace

	request := graphql.NewRequest(assets.Read("graphqls/trace/Trace.graphql"))
	request.Var("traceId", traceID)

	client.ExecuteQueryOrFail(ctx, request, &response)

	return response["result"]
}

func Traces(ctx *cli.Context, condition *schema.TraceQueryCondition) schema.TraceBrief {
	var response map[string]schema.TraceBrief

	request := graphql.NewRequest(assets.Read("graphqls/trace/Traces.graphql"))
	request.Var("condition", condition)

	client.ExecuteQueryOrFail(ctx, request, &response)

	return response["result"]
}
