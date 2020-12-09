package unique

import (
	"fmt"
	"sync"
	"time"
)

var (
	offset int64 = 0
	lastMS int64 = 0
	mutex  sync.Mutex
)

// NewID will create a new unique ID string for this runtime.
// Uniqueness between system or apps is not necessary.
func NewID() string {
	mutex.Lock()
	defer mutex.Unlock()
	MS := time.Now().Unix()
	if lastMS == MS {
		offset++
		return fmt.Sprintf("%d-%d", lastMS, offset)
	}
	lastMS = MS
	offset = 0
	return fmt.Sprintf("%d-%d", lastMS, offset)
}
