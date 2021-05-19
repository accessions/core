package collection

import (
	"fmt"
	"testing"
)

func TestFifo(t *testing.T) {
	elements := [][]byte{
		[]byte("hello"),
		[]byte("world"),
		[]byte("again"),
	}
	queue := NewQueue(8)

	for i := range elements {
		queue.Put(elements[i])
	}
	for range elements {
		take, b := queue.Take()
		fmt.Println(string(take.([]uint8)), b)
	}

	/*for _, element := range elements {
		body, ok := queue.Take()
		fmt.Println(body, ok, string(element))
	}*/
}
