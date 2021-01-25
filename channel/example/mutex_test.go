package example

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMutex_lock(t *testing.T) {
	type fields struct {
		ch chan ticket
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	StartMutex()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mutex{
				ch: tt.fields.ch,
			}
			fmt.Println("--", m)
		})
	}
}

func TestMutex_unLock(t *testing.T) {
	type fields struct {
		ch chan ticket
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mutex{
				ch: tt.fields.ch,
			}
			fmt.Println("====", m)
		})
	}
}

func TestStartMutex(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_newMutex(t *testing.T) {
	tests := []struct {
		name string
		want *Mutex
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newMutex(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newMutex() = %v, want %v", got, tt.want)
			}
		})
	}
}
