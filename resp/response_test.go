package resp

import (
	"github.com/accessions/core/resp/err"
	"reflect"
	"testing"
)

func TestFailure(t *testing.T) {
	type args struct {
		errCode int
	}
	tests := []struct {
		name string
		args args
		want *ResponseFailureBean
	}{
		{"eros", args{errCode: err.USER_NOT_FOUND}, &ResponseFailureBean{
			Code: 400,
			Msg:  "",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Failure(tt.args.errCode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Failure() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSuccess(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name string
		args args
		want *ResponseSuccessBean
	}{
		// TODO: Add test cases.
		{"eros", args{data: "hello"}, &ResponseSuccessBean{
			Code: 200,
			Msg:  "",
			Data: "hello",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Success(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Success() = %v, want %v", got, tt.want)
			}
		})
	}
}
