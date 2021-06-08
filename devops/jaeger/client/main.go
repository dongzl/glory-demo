package main

import (
	"context"
	"time"

	// grpc客户端需引入
	"github.com/glory-go/glory/grpc"
	// 框架日志组件引入
	"github.com/glory-go/glory/log"
)

func main() {
	for {
		ctx := context.Background()

		// 调用server
		// 从配置生成grpc客户端，与配置中serviceName对应
		client := grpc.NewGrpcClient("client_server")

		// 与协议文件结合，拿到greeterClient
		greeterClient := NewGreeterClient(client.GetConn())

		// 发起rpc调用，传递参数
		reply, err := greeterClient.SayHello(ctx, &HelloRequest{
			Name: "client",
		})
		if err != nil {
			panic(err)
		}
		// 打印server结果
		log.Infof("reply = %+v", reply)

		//// 调用subserver
		//client = grpc.NewGrpcClient("client_subserver")
		//greeterClient = NewGreeterClient(client.GetConn())
		//
		//// 发起rpc调用，传递参数
		//reply, err = greeterClient.SayHello(ctx, &HelloRequest{
		//	Name: "client",
		//})
		//if err != nil {
		//	panic(err)
		//}
		//// 打印server结果
		//log.Infof("reply = %+v", reply)

		time.Sleep(1 * time.Second)
	}

}
