package internal

import "math"

const (
	MaxBufSize = math.MaxInt32
)

type service interface {}
var (
	ServiceAction service
)
// grpc internal.go:67
func init()  {
	ServiceAction = func(svc *Service) IServiceOptions {
		return svc.opt.opts
	}
}

type IServiceOptions interface {
	config () error
}

type Service struct {
	opt serviceOptions
}

type serviceOptions struct {
	opts IServiceOptions
}

func (*serviceOptions) config () error {
	return nil
}

func test()  {
	f := ServiceAction.(func(svc *Service) IServiceOptions)
	_ = f(new(Service)).config()

}
