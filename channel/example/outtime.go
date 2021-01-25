package example

import (
	"fmt"
	"time"
)

// Start 执行任务超时: 我们在做任务处理的时候，并不能保证任务的处理时间，通常会加上一些超时控制做异常的处理。
func StartOutTime ()  {
	select {
	case <-doWork():
		fmt.Println("finish")
	case <-time.After(3 * time.Second):
		fmt.Println("任务处理超时")
	}
}

func doWork () <-chan struct{}  {
	ch := make(chan struct{})
	go func() {
		fmt.Println("do work")
		//time.Sleep(1 * time.Second)
	}()
	return ch
}
