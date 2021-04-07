package main

import (
	"fmt"
	"time"
)

func do_stuff(i  int) int {
	time.Sleep(1 * time.Second)
	return i
}

func branch(x int) chan int{
	ch := make(chan int)
	go func() {
		ch <- do_stuff(x)
	}()
	return ch
}

func fanIn(branch ...chan int) chan int {
	ch := make(chan int)
	/*for _, c := range branch {
		//开多个goroutine 明确传值
		go func(c chan int) {
			ch <- <-c
		}(c)
	}*/

	// 开一个goroutine

	go func() {
		for i := 0; i < len(branch); i++ {
			select {
			case v := <-branch[i]:
				ch <- v
			}
		}
	}()
	return ch
}

func ff()  {
	ins := fanIn(branch(1), branch(2), branch(3))

	for i := 0; i < 3; i++ {
		fmt.Println(<-ins)
	}
}
