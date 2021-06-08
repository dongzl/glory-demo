package main

import (
	"context"
	"time"

	// 开启框架服务必须引入
	"github.com/glory-go/glory/glory"
	// 使用日志组件必须引入
	"github.com/glory-go/glory/log"
	// 注册service必须引入
	"github.com/glory-go/glory/service"
)

type server struct {
}

func (s *server) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	log.Info("Received: %v", in.GetName())
	time.Sleep(5 * time.Millisecond)
	return &HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	gloryServer := glory.NewServer()
	gloryService := service.NewGrpcService("GrpcSubService")
	RegisterGreeterServer(gloryService.GetGrpcServer(), &server{})
	gloryServer.RegisterService(gloryService)
	gloryServer.Run()
}
