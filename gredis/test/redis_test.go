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
	i := gredis.Client.Publish("channel-data_sync", "test")
	fmt.Println(i)
}


func TestRedis_Cluster(t *testing.T)  {
	gredis.ServiceRedis.New(func(options *gredis.Options) {
		options.Addrs = "192.168.20.96:6379|192.168.20.46:6379|192.168.20.69:6379|192.168.20.96:6389|192.168.20.46:6389|192.168.20.69:6389"
	}).Go(1)
	test := RedisTestStruct{
		Name:  "eros",
		Value: "smoke",
	}
	o := gredis.Cluster.Set("key", &test, 10 * time.Second)
	s := gredis.Cluster.Get("key")
	fmt.Println(o,s)

}

type RedisTestStruct struct {
	Name string `json:"name"`
	Value string `json:"value"`
}
