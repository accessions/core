package test

import (
	"fmt"
	"github.com/accessions/core/kong"
	"github.com/kevholditch/gokong"
	"strings"
	"testing"
)

func TestKong_init (t *testing.T)  {
	kong.ServiceKong.New("smoke", func(options *kong.Options) {
		options = &kong.Options{
			KongHost:        "http://kong-ingress-controller.kong:8001",
			UpStreamName:    "upStreamName",
			TargetPath:      "svc-targetPath-http.payfun",
			TargetPort:      "1405",
			TargetWeight:    100,
			ServiceName:     "serviceName",
			ServiceProtocol: "http",
			ServicePort:     80,
			RouteProtocol:   strings.Split("http,https", ","),
			RouteHost:       strings.Split("gateway.pay.fun,api.t.dev.pay.fun", ","),
			RoutePath: "routePath",
		}
	}).Go()
}

func TestName(t *testing.T) {
	c := gokong.Config{
		HostAddress:"http://kong-ingress-controller.kong:8001",
		Username:"wangcheng",
		Password:"W3411c0509",
		InsecureSkipVerify: false,
	}
	client := gokong.NewClient(&c)
	get,_:= client.Snis().List()
	fmt.Println(get)
}
