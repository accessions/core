package buffer

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"testing"
)

type ConfigurableBenchMarker struct {
	Action string
	GetBench func(input []byte) func(b *testing.B)
}

func bench ()  {
	marker := ConfigurableBenchMarker{
		Action:   "",
		GetBench: func(input []byte) func(b *testing.B) {
			return func(b *testing.B) {
				for i:=0; i < b.N; i++ {
					c := NewBufferedScanner(nil)
					bytes.NewReader(input)
					_, _ = io.Copy(ioutil.Discard, c)
					_ = c.Close()
				}
			}
		},
	}
	fmt.Println(marker)
}