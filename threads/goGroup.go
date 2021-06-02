package threads

import "sync"

type group struct {
	group sync.WaitGroup
}

func NewGroup() *group {
	return &group{}
}

func (g *group) Run(call func()) {
	g.group.Add(1)
	go func() {
		defer g.group.Done()
		call()
	}()
}

func (g *group) SafeRun(call func()) {
	g.group.Add(1)
	GoSafe(func() {
		defer g.group.Done()
		call()
	})
}

func (g *group) Wait() {
	g.group.Wait()
}
