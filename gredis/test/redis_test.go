package test

import (
	"fmt"
	"github.com/accessions/core/gredis"
	"testing"
	"time"
)

func TestRedis_Client(t *testing.T)  {
	gredis.ServiceRedis.New(func(options *gredis.Options) {
		options.Addrs = "127.0.0.1:6379"
	}).Go(0)
	o := gredis.Client.Set("name", "smoke", 10 * time.Second)
	s := gredis.Client.Get("name")
	fmt.Println(o,s)
}


func TestRedis_Cluster(t *testing.T)  {
	gredis.ServiceRedis.New(func(options *gredis.Options) {
		options.Addrs = "127.0.0.1:6379,127.0.0.1:6379"
	}).Go(1)
	/*o := gredis.Client.Set("name", "eros", 10 * time.Second)
	s := gredis.Client.Get("name")
	fmt.Println(o,s)*/
}
