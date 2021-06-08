package main

import (
	"context"
	"log"

	"github.com/glory-go/glory/glory"
	_ "github.com/glory-go/glory/registry/nacos"
	"github.com/glory-go/glory/service"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	gloryServer := glory.NewServer()
	gloryService := service.NewGrpcService("gloryGrpcService")
	RegisterGreeterServer(gloryService.GetGrpcServer(), &server{})
	gloryServer.RegisterService(gloryService)
	gloryServer.Run()
}
