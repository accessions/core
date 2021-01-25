package example

import (
	"fmt"
	"unsafe"
)
// Start 经常会有这样的场景，当信息收集完成，通知下游开始计算数据。
func Start ()  {
	isOver := make(chan struct{})
	fmt.Println("isOver:", unsafe.Sizeof(isOver))
	go func() {
		collectMsg(isOver)
		fmt.Println(3)
		close(isOver)
	}()
	fmt.Println(1)
	<- isOver
	fmt.Println(2)

	calculateMsg()
}

func collectMsg (isOver chan struct{})  {
	//time.Sleep(2 * time.Second)
	fmt.Println("collect finish")
	isOver <- struct{}{}
}

func calculateMsg ()  {
	fmt.Println("start calculate.")
}
