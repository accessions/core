package trace

import "testing"

func TestTrace(t *testing.T) {

	Trace()
	testA()

}

func testA()  {
	testB()
}

func testB()  {
	Trace()
}
