package test

import (
	"github.com/accessions/core/rabbitmq"
	"testing"
)

func TestName(t *testing.T) {
	rabbitmq.ServiceMq.New("smoke", func(options *rabbitmq.Options) {
		options.Pwd = "pay.Media@2020"
	}).Go()
	//rabbitmq.SendMsg("", []byte("test"))
}
