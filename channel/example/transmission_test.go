package example

import "testing"

func TestStartTransmission(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	StartTransmission()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_worker(t *testing.T) {
	type args struct {
		id   int
		ch   chan token
		next chan token
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
