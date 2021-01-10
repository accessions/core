package gredis

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-redis/redis"
)

var Cluster *redis.ClusterClient
var Client *redis.Client

var ServiceRedis *Redis

type Redis struct {
	cluster *redis.ClusterClient
	client *redis.Client
	option *Options
}
type Options struct {
	Addrs string
	Password string
	ReadTimeout time.Duration
	WriteTimeout time.Duration
}

type newOption func(options *Options)

func init ()  {
	ServiceRedis = &Redis{}
}

func (rs *Redis) New (opts ...newOption) *Redis {
	opt := &Options{
		Addrs:        "",
		Password:     "",
		ReadTimeout:  50 * time.Millisecond,
		WriteTimeout: 50 * time.Millisecond,
	}
	for _, o := range opts {
		o(opt)
	}
	rs.option = opt
	return rs
}

func (rs *Redis) Go (i int)  {
	if i > 0 {
		clusters(rs.option)
	} else {
		client(rs.option)
	}
}

func clusters (option *Options) {
	Cluster = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: strings.Split(option.Addrs, "|"),
		Password: option.Password,
	})
	output("Cluster")
}

func client (option *Options)  {
	Client = redis.NewClient(&redis.Options{
		Addr:         option.Addrs,
		Password:     option.Password,
	})
	output("Client")
}

func output (name string)  {
	fmt.Println("****************[ Redis Start ]********************************")
	fmt.Println("****************[ "+name+" - "+time.Now().String()+"]*****************")
	fmt.Println(ServiceRedis.option)
}