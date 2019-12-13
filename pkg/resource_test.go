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
	urlResource := NewURLResource("https://www.apache.org/licenses/LICENSE-2.0.txt")
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
