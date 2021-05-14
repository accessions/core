package times

import (
	"fmt"
	"time"
)

func test()  {
	start := time.Now()
	format := time.Now().Format(time.RFC3339)
	println(format)

	time.Sleep(1e9) // 1秒
	time.Sleep(3e9) // 3秒
	time.Sleep(1e6) // 1毫秒
	time.Sleep(3e6) // 3毫秒

	sub := time.Now().Sub(start)
	fmt.Println("use time : ", sub)
}
