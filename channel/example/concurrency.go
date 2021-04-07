package example

import (
	"fmt"
	"sync"
)

// StartConcurrency 并发控制
// 在凌晨的时候对内或者对外拉取数据，但是如果不对并发请求加以控制，往往会导致 groutine 泛滥，进而打满 CPU 资源。往往不能控制的东西意味着不好的事情将要发生。对于我们来说，可以通过 channel 来控制并发数。
func StartConcurrency ()  {
	//limit := make(chan struct{}, 10)
	var wg sync.WaitGroup
	jobCount := 100
	for i:=0; i< jobCount; i++ {
		go func(index int) {
			defer wg.Done()
			wg.Add(1)
			//limit<- struct{}{}
			job(index)
			//<-limit
		}(i)
	}
	wg.Wait()
	//select {}
}

func job(index int)  {
	//耗时任务
	//time.Sleep(1 * time.Second)
	fmt.Printf("任务 %d 已完成. \n", index)
}
