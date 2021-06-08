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

// server is used to implement helloworld.GreeterServer.
type server struct {
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	log.CtxInfof(ctx, "Received: %v", in.GetName())
	//time.Sleep(5 * time.Millisecond)
	time.Sleep(time.Second)
	// 发起rpc调用，传递参数。本质上是确保ctx被正确传递下去
	reply, err := greeter.SayHello(ctx, &HelloRequest{
		Name: "server",
	})
	if err != nil {
		panic(err)
	}
	// 打印subserver结果
	log.Infof("reply = %+v", reply)

	//// 再次发起rpc调用，传递参数
	//reply, err = greeter.SayHello(ctx, &HelloRequest{
	//	Name: "server",
	//})
	//if err != nil {
	//	panic(err)
	//}
	//// 打印subserver结果
	//log.Infof("reply = %+v", reply)

	return &HelloReply{Message: "Hello " + in.GetName()}, nil
}

var greeter GreeterClient

func main() {
	// client远程调用server
	client := grpc.NewGrpcClient("jaeger-demo-server1")
	greeter = NewGreeterClient(client.GetConn())

	// server远程调用subserver
	gloryServer := glory.NewServer()
	gloryService := service.NewGrpcService("GrpcService")
	RegisterGreeterServer(gloryService.GetGrpcServer(), &server{})
	gloryServer.RegisterService(gloryService)
	gloryServer.Run()
}
