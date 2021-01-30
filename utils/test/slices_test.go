package test

import (
	"fmt"
	"github.com/accessions/core/utils"
	"testing"
)

func TestShift (t *testing.T) {
	shift, i := utils.Shift([]interface{}{1, 2, 3, 4})
	fmt.Println(shift, i)
}

func TestPop (t *testing.T) {
	shift, i := utils.Pop([]interface{}{1, 2, 3, 4})
	fmt.Println(shift, i)
}


