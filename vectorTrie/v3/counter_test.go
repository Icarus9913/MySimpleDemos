package v3

import (
	"fmt"
	"testing"
)

func TestCounter(t *testing.T) {
	for i := 1; i <= 10; i++ {
		if n := nextId(); n != uint64(i) {
			t.Errorf("Expect next id to be %d, but got %d", i, n)
		}else {
			fmt.Println(n)
		}
	}
}

