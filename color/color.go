package color

import (
	"fmt"
	"github.com/mazznoer/colorgrad"
)

func init()  {
	build, err := colorgrad.NewGradient().Build()
	if err != nil {
		panic(err)
	}
	fmt.Println(build)
}

func handler ()  {

}
