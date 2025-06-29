//go:build go1.11
// +build go1.11

package pkg

import (
	"fmt"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	http2 "github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"
)

// Load will load the file from your git repository
func (bundle *GITResourceBundle) Load() ([]Resource, error) {
	fileSystem := memfs.New()
	CloneOpts := &git.CloneOptions{}
	if len(bundle.URL) == 0 {

		return nil, fmt.Errorf("GIT URL is not specified")
	}
	CloneOpts.URL = bundle.URL

	if len(bundle.RefName) == 0 {
		CloneOpts.ReferenceName = plumbing.ReferenceName("refs/heads/master")
	} else {
		CloneOpts.ReferenceName = plumbing.ReferenceName(bundle.RefName)
	}

	if len(bundle.Remote) == 0 {
		CloneOpts.RemoteName = "origin"
	} else {
		CloneOpts.RemoteName = bundle.Remote
	}

	if len(bundle.PathPattern) == 0 {
		return nil, fmt.Errorf("no path pattern specified")
	}

	if len(bundle.User) != 0 {
		CloneOpts.Auth = &http2.BasicAuth{
			Username: bundle.User,
			Password: bundle.Password,
		}
	}

	_, err := git.Clone(memory.NewStorage(), fileSystem, CloneOpts)
	if err != nil {

		return nil, err
	}

	return bundle.loadPath(bundle.URL, "/", fileSystem)
}
