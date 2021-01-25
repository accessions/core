package example

import (
	"reflect"
	"testing"
)

func TestStartOutTime(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"StartOutTime"},
	}
	StartOutTime()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_doWork(t *testing.T) {
	tests := []struct {
		name string
		want <-chan struct{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doWork(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("doWork() = %v, want %v", got, tt.want)
			}
		})
	}
}
