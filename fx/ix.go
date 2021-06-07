package fx

import "fmt"

type IIx interface {
	add() bool
	save() bool
}

type Ix struct {
}

func (i Ix) adds(ix IIx) {

	fmt.Println(ix.add())

}
