package example

import (
	"fmt"
	"time"
)

type ticket struct {}

type Mutex struct {
	ch chan ticket
}
// 创建一个缓冲区为1的通道
func newMutex () *Mutex  {
	return &Mutex{ch: make(chan ticket, 1)}
}
// 能在缓冲区放入数据，就获取了锁
func (m *Mutex) lock ()  {
	m.ch <- struct{}{}
}
// 解锁把数据取出
func (m *Mutex) unLock()  {
	select {
	case <-m.ch:
	default:
		panic("已经解锁了")
	}
}


// StartMutex .互斥锁
// 我们也可以通过 channel 实现一个小小的互斥锁。通过设置一个缓冲区为1的通道，如果成功地往通道发送数据，说明拿到锁，否则锁被别人拿了，等待他人解锁。
func StartMutex ()  {
	mutex := newMutex()

	for i:=0; i< 10;i++ {
		go func(task int) {
			mutex.lock()
			fmt.Printf("-任务 %d 拿到锁了\n", task)
			time.Sleep(1 * time.Second)
			mutex.unLock()
		}(i)
		go func(task int) {
			mutex.lock()
			fmt.Printf("=任务 %d 拿到锁了\n", task)
			mutex.unLock()
		}(i)
		time.Sleep(500 * time.Millisecond)
	}
	mutex.lock()
	mutex.unLock()
	close(mutex.ch)
}
