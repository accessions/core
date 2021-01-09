package test

import (
	"fmt"
	"github.com/accessions/core/apollo"
	"testing"
)

func TestApollo_InitAgollo(t *testing.T)  {
	apollo.ApolloService.New(func(options *apollo.Options) {
		options.AppId = "payfun.client.mq"
		options.AppName = "payfun.client.mq"
		options.Path = "39.98.214.1:8080"
	}).Go()

}

func Test_GetConfig(t *testing.T)  {
	f := apollo.GetSetting("redisConn", "cp.common")
	fmt.Println(f)
}
