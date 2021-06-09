package main

import (
	"context"
	"github.com/glory-go/glory/grpc"
	"github.com/glory-go/glory/log"
	_ "github.com/glory-go/glory/registry/nacos"
	"time"
)

func main() {
	client := grpc.NewGrpcClient("grpc-helloworld-demo")
	greeterClient := NewGreeterClient(client.GetConn())
	for {
		reply, err := greeterClient.SayHello(context.Background(), &HelloRequest{
			Name: "grpcDemo",
		})
		if err != nil {
			panic(err)
		}
		log.Infof("reply = %+v", reply)
		time.Sleep(time.Second)
	}

}
