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
	ServiceAction = func(svc *services) IServiceOptions {
		return svc.opt.opts
	}
}

type IServiceOptions interface {
	config () error
}

type services struct {
	opt serviceOptions
}

type serviceOptions struct {
	opts IServiceOptions
}

func (*serviceOptions) config () error {
	return nil
}

func test()  {
	f := ServiceAction.(func(svc *services) IServiceOptions)
	_ = f(new(services)).config()

}
