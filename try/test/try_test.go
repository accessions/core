package test

import (
	"errors"
	"fmt"
	"github.com/accessions/core/try"
	"testing"
)

func TestTry(t *testing.T) {
	var err interface{}
	try.Try(func() {
		defer func() {
			if err = recover(); err != nil {
				try.Throw(1, err.(string))
			}
		}()
		fmt.Println("test try chatch")
	}).Catch(1, func(e try.Exception) {
		err = errors.New(e.Msg)
	}).Finally(func() {})
}

