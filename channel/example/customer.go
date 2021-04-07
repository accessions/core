package example

import (
	"fmt"
	"sync"
)

func test()  {
	// 生产
	deal := make(chan string)
	go func() {
		defer func() {
			close(deal)
		}()
		for {
			deal<-string("123")
		}
	}()

	//消费
	var wg sync.WaitGroup
	count := 10
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer func() {wg.Done()}()
			for s := range deal {
				jobs(s)
			}
		}()
	}
	wg.Wait()
}

func jobs(s string)  {
	fmt.Println(s)
}
