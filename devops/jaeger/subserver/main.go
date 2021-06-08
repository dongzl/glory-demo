package main

import (
	"context"
	"time"

	"github.com/glory-go/glory/grpc"

	// 开启框架服务必须引入
	"github.com/glory-go/glory/glory"
	// 使用日志组件必须引入
	"github.com/glory-go/glory/log"
	// 注册service必须引入
	"github.com/glory-go/glory/service"
)

type server struct {
}

var greeter GreeterClient

func (s *server) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	log.Info("Received: %v", in.GetName())
	time.Sleep(5 * time.Millisecond)
	// 发起rpc调用，传递参数。本质上是确保ctx被正确传递下去
	reply, err := greeter.SayHello(ctx, &HelloRequest{
		Name: "server",
	})
	if err != nil {
		panic(err)
	}
	// 打印subserver结果
	log.Infof("reply = %+v", reply)
	return &HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	client := grpc.NewGrpcClient("jaeger-demo-server2")
	greeter = NewGreeterClient(client.GetConn())

	gloryServer := glory.NewServer()
	gloryService := service.NewGrpcService("GrpcSubService")
	RegisterGreeterServer(gloryService.GetGrpcServer(), &server{})
	gloryServer.RegisterService(gloryService)
	gloryServer.Run()
}
