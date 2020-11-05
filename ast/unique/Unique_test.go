package unique

import "testing"

func TestNewId(t *testing.T) {
	checkMap := make(map[string]int)
	for i := 0; i < 100000; i++ {
		id := NewID()
		if _, ok := checkMap[id]; ok {
			t.FailNow()
		}
		checkMap[id] = 1
	}
}
