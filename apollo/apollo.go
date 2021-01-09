package apollo

import (
	"fmt"
	"github.com/shima-park/agollo"
	"time"
)

var (
	ApolloService *Apollo
	ApolloConfigs agollo.Agollo
	defaultOptionAppName = "Smoke"
	defaultOptionPath    = "192.168.20.12:8080"
	defaultOptionAppId   = "someApp"
)


type Apollo struct {
	config agollo.Agollo
	action byte
	option *Options
}
type Options struct {
	AppName string
	Path    string
	AppId   string
}
type newOption func(options *Options)

func init()  {
	ApolloService = &Apollo{}
}
func (apollo *Apollo) New (opts ...newOption) *Apollo {
	option := &Options{
		AppName: defaultOptionAppName,
		Path:    defaultOptionPath,
		AppId:   defaultOptionAppId,
	}
	for _, o := range opts {
		o(option)
	}
	apollo.option = option
	return apollo
}

func (apollo *Apollo) Go () {

	ApolloConfigs, err := agollo.New(apollo.option.Path, apollo.option.AppId, agollo.AutoFetchOnCacheMiss())
	if err != nil {
		fmt.Println(0,err.Error())
	} else {
		fmt.Println("****************[ Apollo Start ]********************************")
		fmt.Println("****************[ "+apollo.option.AppName+" - "+time.Now().String()+"]*****************")
		fmt.Println(ApolloConfigs)
	}

}

func (apollo *Apollo) Get (key, ns string) string {
	return apollo.config.Get(key, agollo.WithNamespace(ns))
}

func  GetSetting(key, ns string) string {

	return ApolloConfigs.Get(key, agollo.WithNamespace(ns))
}

