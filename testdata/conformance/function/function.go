// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package function contains test functions to validate the framework.
package function

import (
	"net/http"
	"time"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/GoogleCloudPlatform/functions-framework-go/testdata/conformance/nondeclarative"
)

const (
	outputFile = "function_output.json"
)

// Register declarative functions
func init() {
	functions.HTTP("declarativeHTTP", nondeclarative.HTTP)
	functions.HTTP("concurrentHTTP", concurrentHTTP)
	functions.CloudEvent("declarativeCloudEvent", nondeclarative.CloudEvent)
}

func concurrentHTTP(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
}
