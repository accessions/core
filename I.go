package core

type I interface {
	Igob
}

type Igob interface {
	EnCoder(interface{}) ([]byte, error)
	DeCoder([]byte) (interface{}, error)
}

var (

)
