package smoke

import (
	"github.com/accessions/core/apollo"
	"github.com/go-ini/ini"
	"log"
	"os"
)

var (
	Env string
	cfg *ini.File
	Options struct{
		Apollo *apollo.Options
		Kong *interface{}
	}
)

func init()  {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("set, fail to parse 'app.ini': %v", err)
	}
	Env = os.Getenv("GOENV")
	if len(Env) > 0 {
		_ini(Env, Options.Apollo)
		_ini(Env, Options.Kong)
	}

}


func _ini (section string, v interface{})  {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("cfg.mapTo %s err: %v", section, err)
	}
}

