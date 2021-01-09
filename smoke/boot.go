package smoke

import (
	"fmt"
	"github.com/accessions/core/apollo"
)

func init()  {

}

func Setup ()  {
	var err error
	defer func() {
		if err != nil {
			fmt.Println("boot, error: %v", err)
		}
	}()
	_apollo()

}

func _apollo ()  {

	apollo.ApolloService.New(func(options *apollo.Options) {
		options = Options.Apollo
	}).Go()
}