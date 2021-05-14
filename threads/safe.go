package threads

import (
	"log"
)

var SafeMutex iSafeSync = new(mutex)
var SafeGroup iSafeSync = new(group)

type iSafeSync interface {
	Run(func())
	SafeRun(func())
	Wait()
}

func GoSafe(call func()) {
	go RunSafe(call)
}

func RunSafe(call func()) {
	defer Recover()
	call()
}

func Recover(calls ...func()) {
	for _, call := range calls {
		call()
	}
	if err := recover(); err != nil {
		log.Println(err)
	}
}
