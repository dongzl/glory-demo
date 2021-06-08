package main

import (
	"context"
	"time"

	"github.com/glory-go/glory/log"

	"github.com/glory-go/glory/grpc"
	_ "github.com/glory-go/glory/registry/k8s"
)

func main() {
	client := grpc.NewGrpcClient("grpc-helloworld-demo")
	greeterClient := NewGreeterClient(client.GetConn())
	for {
		time.Sleep(time.Second * 5)
		reply, err := greeterClient.SayHello(context.Background(), &HelloRequest{
			Name: "grpcDemo",
		})
		if err != nil {
			log.Error(err)
			continue
		}
		log.Infof("reply = %+v", reply)

	}

}
