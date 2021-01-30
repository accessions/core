package aiot

import "sync"

type lightState int

const (
	on = lightState(iota)
	off
)

type light struct {
	state lightState
	mux sync.Mutex
}

func newLight() light  {
	return light{
		state: off,
	}
}

func (l *light) setState(s lightState)  {
	l.mux.Lock()
	defer l.mux.Unlock()
	l.state = s
}

func (l *light) getState () lightState {
	return l.state
}
