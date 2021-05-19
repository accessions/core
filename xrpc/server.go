package xrpc

import "github.com/accessions/core/internal"

type (
	RpcServer struct {
		server internal.IServer
		register internal.FRegister
	}

)

func NewServer()  {

}
