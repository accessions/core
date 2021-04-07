package main

import (
	"fmt"
	"time"
)

func main() {

	timeout := timer(time.Second * 4)
	for  {
		select {
		case <-timeout:
			fmt.Println("超时了")
			return //加return
		case <-time.After(1* time.Second):
			fmt.Println("超时")
			return
		}
	}
}

func timer(duration time.Duration) chan bool  {
	ch := make(chan bool)
	go func() {
		time.Sleep(duration)
		ch<-true
	}()
	return ch
}
