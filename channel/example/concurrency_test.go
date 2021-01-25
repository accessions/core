package example

import "testing"

func TestStartConcurrency(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	StartConcurrency()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_job(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
