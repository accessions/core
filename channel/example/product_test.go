package example

import "testing"

func TestStartProduct(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	StartProduct()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_consumer(t *testing.T) {
	type args struct {
		task <-chan int
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
