// +build !go1.11

package pkg

import (
	"fmt"
)

// Load will load the file from your git repository
func (bundle *GITResourceBundle) Load() ([]Resource, error) {
	return nil, fmt.Errorf("GIT resources are not supported with Go 1.10 or below")
}
