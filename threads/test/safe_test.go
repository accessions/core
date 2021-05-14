package test

import (
	"fmt"
	"github.com/accessions/core/threads"
	"testing"
	"time"
)

func TestSafe(t *testing.T) {
	threads.SafeMutex.SafeRun(func() {
		fmt.Println(11)
	})
	time.Sleep(1e9)
}
