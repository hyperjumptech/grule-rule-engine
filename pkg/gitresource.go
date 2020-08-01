// +build go1.11

package pkg

import (
	"fmt"

	"gopkg.in/src-d/go-billy.v4/memfs"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	http2 "gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

// Load will load the file from your git repository
func (bundle *GITResourceBundle) Load() ([]Resource, error) {
	fs := memfs.New()
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

	_, err := git.Clone(memory.NewStorage(), fs, CloneOpts)
	if err != nil {
		return nil, err
	}

	return bundle.loadPath(bundle.URL, "/", fs)
}
