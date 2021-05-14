package enum

import (
	"fmt"
	"sync/atomic"
)

type Enum  string

const testFun Enum = "string"

type State int

// const step State = -iota
const (
	_ State = (1 << iota) / 4
	stepA
	stepB
	stepC
	stepD
	stepE
)

func f()  {
	id := int32(10)
	addInt32 := atomic.AddInt32(&id, 1)
	fmt.Println(addInt32)
}
