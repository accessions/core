package example

import (
	"fmt"
	"time"
)

// StartTransmission 数据传递
// 极客上一道有意思的题，假设有4个 goroutine，编号为1，2，3，4。每秒钟会有一个 goroutine 打印出它自己的编号。现在让你写一个程序，要求输出的编号总是按照1，2，3，4这样的顺序打印。类似下图，
type token struct {}
func StartTransmission()  {
	num := 4
	var chs []chan token
	// 4个work
	for i:=0; i<num; i++ {
		chs = append(chs, make(chan token))
	}
	fmt.Println("4个work", chs)
	for j:=0; j< num; j++ {
		go worker(j, chs[j], chs[(j+1)%num])
	}
	// 令牌交给第一个
	chs[0] <- struct{}{}
	select {}
}

func worker(id int, ch chan token, next chan token)  {
	for  {
		//对应work,取得令牌
		token := <-ch
		fmt.Println(id+1)
		time.Sleep(1 * time.Second)
		// 传递给下一个
		next <- token
	}
}
