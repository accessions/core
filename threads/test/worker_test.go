package test

import (
	"fmt"
	"github.com/accessions/core/threads"
	"testing"
)

func TestWorker(t *testing.T) {
	threads.NewWorker(func() {
		fmt.Println("11111")
	}, 3).GRun()
}
