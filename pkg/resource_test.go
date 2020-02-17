package pkg

import (
	"reflect"
	"testing"
)

const (
	loremipsum = `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum`
)

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
