// +build go1.16

//  Copyright hyperjumptech/grule-rule-engine Authors
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package pkg

import (
	"embed"
	"strings"
	"testing"
)

//go:embed test
var rules embed.FS

func TestEmbeddedResourceBundle_Load(t *testing.T) {
	erb := NewEmbeddedResourceBundle(rules, ".", "/**/*.grl")
	resources := erb.MustLoad()
	if len(resources) != 6 {
		t.Errorf("Expected 6 but get %d", len(resources))
		t.FailNow()
	}
	if !strings.HasSuffix(resources[0].String(), "GrlFile11.grl") {
		t.Errorf("Expect [0] to have suffix GrlFile11.grl. But %s", resources[0].String())
	}
	if !strings.HasSuffix(resources[1].String(), "GrlFile12.grl") {
		t.Errorf("Expect [1] to have suffix GrlFile12.grl. But %s", resources[0].String())
	}
	if !strings.HasSuffix(resources[2].String(), "GrlFile21.grl") {
		t.Errorf("Expect [2] to have suffix GrlFile11.grl. But %s", resources[0].String())
	}
	if !strings.HasSuffix(resources[3].String(), "GrlFile22.grl") {
		t.Errorf("Expect [3] to have suffix GrlFile11.grl. But %s", resources[0].String())
	}
	if !strings.HasSuffix(resources[4].String(), "GrlFile211.grl") {
		t.Errorf("Expect [4] to have suffix GrlFile11.grl. But %s", resources[0].String())
	}
	if !strings.HasSuffix(resources[5].String(), "GrlFile212.grl") {
		t.Errorf("Expect [5] to have suffix GrlFile11.grl. But %s", resources[0].String())
	}
}
