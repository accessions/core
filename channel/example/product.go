package example

import (
	"fmt"
	"time"
)

// 生产消费模型: 生产者只需要关注生产，而不用去理会消费者的消费行为，更不用关心消费者是否执行完毕。而消费者只关心消费任务，而不需要关注如何生产。
func StartProduct ()  {
	ch := make(chan int, 10)
	producer(ch)
	consumer(ch)
	time.Sleep(3 * time.Second)
}
// 1个生产者
func producer(ch chan int)  {
	for i:=0; i< 10 ;i++ {
		ch <- i
		fmt.Println("消费金", i)
	}

}
// 消费者
func consumer(task <-chan int)  {
	// 5个消费者
	for i:=0; i < 5; i++ {
		go func(id int) {
			for {
				item, ok := <- task
				// 如果false, 说明通道关闭
				if !ok {
					//fmt.Printf("关闭：%d : %d", task, item)
					return
				}
				fmt.Printf("消费者:%d, 消费了:%d \n", id, item)
				// 给别人一点机会不吃亏
				time.Sleep(50 * time.Millisecond)
			}
		}(i)
	}
}
