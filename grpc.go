package grpc

import (
	"time"
	"fmt"

	"github.com/micro/go-grpc/broker"
	"github.com/micro/go-grpc/client"
	"github.com/micro/go-grpc/server"
	"github.com/micro/go-micro"
)

// NewService returns a grpc service compatible with go-micro.Service
func NewService(opts ...micro.Option) micro.Service {
	// our grpc client
	c := client.NewClient()
	// our grpc server
	s := server.NewServer()
	// our grpc broker
	b := broker.NewBroker()

	// create options with priority for our opts
	options := []micro.Option{
		micro.Client(c),
		micro.Server(s),
		micro.Broker(b),
	}

	// append passed in opts
	options = append(options, opts...)

	go func() {
		for {
			fmt.Println("THIS REPO IS DEPRECATED. Please migrate to github.com/micro/go-micro/service/grpc.")
			time.Sleep(time.Minute)
		}
	}()

	// generate and return a service
	return micro.NewService(options...)
}

// NewFunction returns a grpc service compatible with go-micro.Function
func NewFunction(opts ...micro.Option) micro.Function {
	// our grpc client
	c := client.NewClient()
	// our grpc server
	s := server.NewServer()
	// our grpc broker
	b := broker.NewBroker()

	// create options with priority for our opts
	options := []micro.Option{
		micro.Client(c),
		micro.Server(s),
		micro.Broker(b),
		micro.RegisterTTL(time.Minute),
		micro.RegisterInterval(time.Second * 30),
	}

	// append passed in opts
	options = append(options, opts...)

	// generate and return a function
	return micro.NewFunction(options...)
}
