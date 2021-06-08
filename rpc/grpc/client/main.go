package main

import (
	"context"
	"github.com/glory-go/glory/grpc"
	"github.com/glory-go/glory/log"
)

func main() {
	client := grpc.NewGrpcClient("grpc-helloworld-demo")
	greeterClient := NewGreeterClient(client.GetConn())
	reply, err := greeterClient.SayHello(context.Background(), &HelloRequest{
		Name: "grpcDemo",
	})
	if err != nil {
		panic(err)
	}
	log.Infof("reply = %+v", reply)
}
