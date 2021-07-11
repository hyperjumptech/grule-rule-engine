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
	"path/filepath"
	"strings"

	"github.com/bmatcuk/doublestar"
	"github.com/hyperjumptech/grule-rule-engine/logger"
)

// EmbeddedResource is a struct that will load an embedded file from an embed.FS struct.
// Note: EmbeddedResource is only available when using Go 1.16 or above
type EmbeddedResource struct {
	Path   string
	Source embed.FS
	Bytes  []byte
}

// NewEmbeddedResource will create a new instance of EmbeddedResource.
// source is an embed.FS struct
// path is the path to the embedded resource.
func NewEmbeddedResource(source embed.FS, path string) Resource {
	return &EmbeddedResource{
		Path:   path,
		Source: source,
	}
}

// Load will load the resource into a byte array from the embedded source.
func (res *EmbeddedResource) Load() ([]byte, error) {
	if res.Bytes != nil {
		return res.Bytes, nil
	}

	var err error
	res.Bytes, err = res.Source.ReadFile(res.Path)
	return res.Bytes, err
}

func (res *EmbeddedResource) String() string {
	return "From embedded path: " + res.Path
}

// EmbeddedResourceBundle is a helper struct to load multiple embedded resources
// all at once by specifying the root location of the file and the file pattern
// to look for. It will look into sub-directories for the file with pattern matching.
type EmbeddedResourceBundle struct {
	// The base path for the embedded resources
	BasePath string
	// List Glob like file pattern.
	// *.grl           <- matches abc.grl but not /anyfolder/abc.grl
	// **/*.grl        <- matches abc/def.grl or abc/def/ghi.grl or abc/def/.grl
	// /abc/**/*.grl   <- matches /abc/def.grl or /abc/def/ghi.drl
	PathPattern []string
	Source      embed.FS
}

// NewEmbeddedResourceBundle creates new instance of EmbeddedResourceBundle struct
// source is the embed.FS from which to load files
// basePath denotes the directory location where the file is located.
// pathPattern are list of paths that filters the files. Its important that
// the pattern will include the base path as it filter.
// For example, if the base path is "/some/base/path"
// The pattern to accept all GRL file is "/some/base/path/**/*.grl".
// This will accept all *.grl files under /some/base/path and its directories.
func NewEmbeddedResourceBundle(source embed.FS, basePath string, pathPattern ...string) *EmbeddedResourceBundle {
	return &EmbeddedResourceBundle{
		Source:      source,
		BasePath:    strings.TrimLeft(basePath, "/"),
		PathPattern: pathPattern,
	}
}

// Load all embedded file resources that located under BasePath that conform to the PathPattern.
func (bundle *EmbeddedResourceBundle) Load() ([]Resource, error) {
	return bundle.loadPath(bundle.BasePath)
}

// MustLoad function is the same as Load with difference that it will panic if any error is raised
func (bundle *EmbeddedResourceBundle) MustLoad() []Resource {
	resources, err := bundle.Load()
	if err != nil {
		panic(err)
	}
	return resources
}

func (bundle *EmbeddedResourceBundle) loadPath(path string) ([]Resource, error) {
	logger.Log.Tracef("Enter embedded directory %s", path)

	finfos, err := bundle.Source.ReadDir(path)
	if err != nil {
		return nil, err
	}
	if path == "." || path == "./" {
		path = ""
	}
	ret := make([]Resource, 0)
	for _, finfo := range finfos {
		fulPath := filepath.Join(path, finfo.Name())
		if finfo.IsDir() {
			gres, err := bundle.loadPath(fulPath)
			if err != nil {
				return nil, err
			}
			ret = append(ret, gres...)
		} else {
			for _, pattern := range bundle.PathPattern {
				matched, err := doublestar.PathMatch(pattern, "/"+fulPath)
				if err != nil {
					return nil, err
				}
				if matched {
					logger.Log.Debugf("Loading embedded file %s", fulPath)
					gress := NewEmbeddedResource(bundle.Source, fulPath)
					_, err := gress.Load()
					if err != nil {
						return nil, err
					}
					ret = append(ret, gress)
					break
				}
			}
		}
	}
	return ret, nil
}
