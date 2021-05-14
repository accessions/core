package trace

import (
	"log"
	"runtime"
)

func Trace() {
	defer func() {
		caller, file, line, _ := runtime.Caller(1)
		log.Printf("%v - %s:%d", caller, file, line)
	}()
}

