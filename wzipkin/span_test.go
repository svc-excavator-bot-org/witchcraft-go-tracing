// Copyright (c) 2018 Palantir Technologies. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package wzipkin_test

import (
	"testing"

	"github.com/palantir/witchcraft-go-tracing/wtracing"
	"github.com/palantir/witchcraft-go-tracing/wtracing/wtracingtests"
	"github.com/palantir/witchcraft-go-tracing/wzipkin"
)

var implProvider = wtracingtests.ImplProvider{
	Name: "wzipkin",
	TracerCreator: func(reporter wtracing.Reporter, opts ...wtracing.TracerOption) (wtracing.Tracer, error) {
		return wzipkin.NewTracer(reporter, opts...)
	},
}

func TestWZipkinImpl(t *testing.T) {
	wtracingtests.RunTests(t, implProvider)
}
