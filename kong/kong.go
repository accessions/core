package kong

import (
	"encoding/json"
	"fmt"
	"github.com/kevholditch/gokong"
	"strings"
	"time"
)

var ServiceKong *Kong

type Kong struct {
	action string
	option *Options
}
type Options struct {
	KongHost        string
	UpStreamName    string
	TargetPath      string
	TargetPort      string
	TargetWeight    int
	ServiceName     string
	ServiceProtocol string
	ServicePort     int
	RouteProtocol   []string
	RouteHost       []string
	RoutePath       string
}

type newOption func(options *Options)

func init()  {
	ServiceKong = &Kong{}
}

func (kong *Kong) New (action string, opts ...newOption) *Kong {
	option := &Options{
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
	for _, o := range opts {
		o(option)
	}
	kong.action = action
	kong.option = option
	return kong
}

func (kong *Kong) Go ()  {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error: ",err)
			panic(err)
		}
	}()
	config := gokong.Config{HostAddress: kong.option.KongHost}
	client := gokong.NewClient(&config)
	// Upstream
	updateUpstreamRequest := &gokong.UpstreamRequest{}
	_ = json.Unmarshal(upstreamJson, updateUpstreamRequest)
	updateUpstreamRequest.Name = kong.option.UpStreamName
	updateUpstreamRequest.HealthChecks.Active.HttpPath = fmt.Sprintf("/%s/kong/healthchecks", kong.option.RoutePath)
	var updatedUpstream *gokong.Upstream
	if upstream, _ := client.Upstreams().GetByName(kong.option.UpStreamName); upstream != nil {
		updatedUpstream, _ = client.Upstreams().UpdateByName(kong.option.UpStreamName, updateUpstreamRequest)
	} else {
		updatedUpstream, _ = client.Upstreams().Create(updateUpstreamRequest)
	}
	fmt.Printf("Upstream: %s，%+v", updatedUpstream.Id, updatedUpstream)
	fmt.Println("-----------------------------------------------------")

	// Target
	targetRequest := &gokong.TargetRequest{
		Target: kong.option.TargetPath + ":" + kong.option.TargetPort,
		Weight: kong.option.TargetWeight,
	}
	createdTarget, err := client.Targets().CreateFromUpstreamId(updatedUpstream.Id, targetRequest)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Target: %s，%+v", *createdTarget.Id, createdTarget)
	fmt.Println("-----------------------------------------------------")

	// Service
	serviceRequest := &gokong.ServiceRequest{
		Name:     &kong.option.ServiceName,
		Protocol: &kong.option.ServiceProtocol,
		Host:     &kong.option.UpStreamName,
		Port:     &kong.option.ServicePort,
	}

	var updatedService *gokong.Service
	if service, _ := client.Services().GetServiceByName(kong.option.ServiceName); service != nil {
		updatedService, _ = client.Services().UpdateServiceByName(kong.option.ServiceName, serviceRequest)
	} else {
		updatedService, _ = client.Services().Create(serviceRequest)
	}
	fmt.Printf("Service: %s，%+v", *updatedService.Id, updatedService)
	fmt.Println("-----------------------------------------------------")
	// Route
	routeRequest := &gokong.RouteRequest{
		Name:         gokong.String(*updatedService.Name + "-Route"),
		Protocols:    gokong.StringSlice(kong.option.RouteProtocol),
		Methods:      gokong.StringSlice([]string{"POST", "GET", "PUT", "DELETE", "OPTIONS", "HEAD", "TRACE", "CONNECT"}),
		Hosts:        gokong.StringSlice(kong.option.RouteHost),
		Paths:        gokong.StringSlice([]string{fmt.Sprintf("/%s/(?i)", kong.option.RoutePath)}),
		StripPath:    gokong.Bool(false),
		PreserveHost: gokong.Bool(false),
		Service:      gokong.ToId(*updatedService.Id),
	}
	var updatedRoute *gokong.Route
	if route, _ := client.Routes().GetByName(*routeRequest.Name); route != nil {
		updatedRoute, _ = client.Routes().UpdateByName(*routeRequest.Name, routeRequest)
	} else {
		updatedRoute, _ = client.Routes().Create(routeRequest)
	}

	fmt.Printf("Route: %s，%+v", *updatedRoute.Id, updatedRoute)
	fmt.Println("-----------------------------------------------------")
	fmt.Println("-----------------[ Kong Start ]-----------------------")
	fmt.Println("-----------------["+kong.action+":"+time.Now().String()+"]------------------------------------")
	fmt.Println(ServiceKong.option)
}

