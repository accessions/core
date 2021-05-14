package threads

import "sync"

type mutex struct {
	m sync.Mutex
}

func (g *mutex) Run(call func()) {
	g.m.Lock()
	go func() {
		defer g.m.Unlock()
		call()
	}()
}

func (g *mutex) SafeRun(call func()) {
	g.m.Lock()
	GoSafe(func() {
		defer g.m.Unlock()
		call()
	})
}

func (g *mutex) Wait() {
	g.m.Lock()
}
