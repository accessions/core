package test

import (
	"github.com/accessions/core/kong"
	"strings"
	"testing"
)

func TestKong_init (t *testing.T)  {
	kong.ServiceKong.New("smoke", func(options *kong.Options) {
		options = &kong.Options{
			KongHost:        "http://kong-proxy:8001",
			UpStreamName:    "UpStreamName",
			TargetPath:      "节点名",
			TargetPort:      "节点端口",
			TargetWeight:    100,
			ServiceName:     "ServiceName",
			ServiceProtocol: "http",
			ServicePort:     80,
			RouteProtocol:   strings.Split("http,https", ","),
			RouteHost:       strings.Split("域名1,域名2", ","),
		}
	}).Go()
}
