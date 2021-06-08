package main

import (
	context "context"
	"log"

	"github.com/glory-go/glory/metrics"

	"github.com/glory-go/glory/glory"
	_ "github.com/glory-go/glory/registry/redis"
	"github.com/glory-go/glory/service"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	metrics.CounterInsc("in main")
	gloryServer := glory.NewServer()
	gloryService := service.NewGrpcService("gloryGrpcService")
	RegisterGreeterServer(gloryService.GetGrpcServer(), &server{})
	gloryServer.RegisterService(gloryService)
	gloryServer.Run()
}
