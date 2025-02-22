//  Copyright (c) 2020 Cisco and/or its affiliates.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at:
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

// +build integration

package vppapi_test

import (
	"encoding/json"
	"testing"

	"go.fd.io/govpp/binapigen/vppapi"
)

func TestParse(t *testing.T) {
	files, err := vppapi.Parse()
	if err != nil {
		t.Fatal(err)
	}

	for _, file := range files {
		b, err := json.MarshalIndent(file, "\t", "  ")
		if err != nil {
			t.Fatal(err)
		}
		t.Logf(" - %s:\n%s", file.Name, b)
	}

	t.Logf("parsed %d files", len(files))
}
