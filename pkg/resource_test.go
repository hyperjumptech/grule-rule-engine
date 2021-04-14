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
	"os"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

const (
	loremipsum = `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum`
)

func TestFileResourceBundle_Load(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	t.Logf("OS is : %s", runtime.GOOS)
	t.Logf("Path : %s", path)
	var frb *FileResourceBundle
	if runtime.GOOS == "windows" {
		frb = NewFileResourceBundle(path, "**\\*.grl")
	} else {
		frb = NewFileResourceBundle(path, "/**/*.grl")
	}
	resources := frb.MustLoad()
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

func TestNewBytesResource(t *testing.T) {
	resource := NewBytesResource([]byte(loremipsum))
	loaded, err := resource.Load()
	if err != nil {
		t.Error("Failed to load byte resource", err)
		t.Fail()
	}
	if !reflect.DeepEqual([]byte(loremipsum), loaded) {
		t.Error("Loaded array are not equal to origin array")
		t.Fail()
	}
}

func TestNewURLResource(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping URL resource test in short mode")
	}
	urlResource := NewURLResource("https://raw.githubusercontent.com/hyperjumptech/grule-rule-engine/master/LICENSE-2.0.txt")
	loadedURL, err := urlResource.Load()
	if err != nil {
		t.Error("Failed to load url resource", err)
		t.Fail()
	}

	fileResource := NewFileResource("../LICENSE-2.0.txt")
	loadedFile, err := fileResource.Load()
	if err != nil {
		t.Error("Failed to load file resource", err)
		t.Fail()
	}
	if !reflect.DeepEqual(loadedURL, loadedFile) {
		t.Error("Loaded array are not equal to origin array")
		t.Fail()
	}
}

func TestGitResource(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping GIT resource test in short mode")
	}
	gitRb := &GITResourceBundle{
		URL: "https://github.com/hyperjumptech/grule-rule-engine.git",
		PathPattern: []string{
			"/antlr/*.grl",
		},
	}
	resources := gitRb.MustLoad()
	if len(resources) != 2 {
		t.Logf("Expected 2 drl but %d", len(resources))
	}
	for _, r := range resources {
		bytes, _ := r.Load()
		t.Logf("Loaded %s . %d bytes", r.String(), len(bytes))
		if len(bytes) == 0 {
			t.FailNow()
		}
	}
}
