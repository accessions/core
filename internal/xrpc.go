package internal

import (
	"github.com/accessions/core/stat"
	"google.golang.org/grpc"
)

type (

	FRegister func(*grpc.Server)

	IServer interface {
		AddOptions(options ...grpc.ServerOption)
		AddStreamInterceptors(interceptors ...grpc.StreamServerInterceptor)
		AddUnarySvcInterceptors(interceptors ...grpc.UnaryServerInterceptor)
		AddUnaryCliInterceptors(interceptors ...grpc.UnaryClientInterceptor)
		SetName(string)
		Start(register FRegister) error
	}

	baseRpcServer struct {
		address string
		metrics *stat.Metrics
		options []grpc.ServerOption
		streamInterceptors []grpc.StreamServerInterceptor
		unarySvcInterceptors []grpc.UnaryServerInterceptor
		unaryCliInterceptors []grpc.UnaryClientInterceptor
	}
)


func newBaseRpcServer(address string, metrics *stat.Metrics) *baseRpcServer {
	return &baseRpcServer{
		address:              address,
		metrics:              metrics,
	}
}


func (s *baseRpcServer) AddOptions(options ...grpc.ServerOption) {
	s.options = append(s.options, options...)
}

func (s *baseRpcServer) AddStreamInterceptors(interceptors ...grpc.StreamServerInterceptor) {
	s.streamInterceptors = append(s.streamInterceptors, interceptors...)
}

func (s *baseRpcServer) AddUnarySvcInterceptors(interceptors ...grpc.UnaryServerInterceptor) {
	s.unarySvcInterceptors = append(s.unarySvcInterceptors, interceptors...)
}

func (s *baseRpcServer) AddUnaryCliInterceptors(interceptors ...grpc.UnaryClientInterceptor) {
	s.unaryCliInterceptors = append(s.unaryCliInterceptors, interceptors...)
}

func (s *baseRpcServer) SetName(name string) {
	s.metrics.Name = name
}

func (s *baseRpcServer) Start(register FRegister) error {
	return nil
}
