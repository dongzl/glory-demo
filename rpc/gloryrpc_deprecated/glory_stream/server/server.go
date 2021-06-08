package main

import (
	"context"
	"fmt"
	"time"

	"github.com/glory-go/glory/service"

	hessian "github.com/apache/dubbo-go-hessian2"

	_ "github.com/glory-go/glory/config"
	"github.com/glory-go/glory/glory"
	"github.com/glory-go/glory/log"

	_ "github.com/glory-go/glory/protocol/glory"
	_ "github.com/glory-go/glory/registry/nacos"
)

type GloryProvider struct {
}

//func (g GloryProvider) SayHello(ctx context.Context, s string) (string, error) {
//	return fmt.Sprintf("hello %s", s), nil
//}

func (g GloryProvider) SayHello(ctx context.Context, req *ReqBody, str2 string) (*RspBody, error) {
	log.Info("req = ", *req, "+", str2)
	fmt.Println(time.Now().String())
	return &RspBody{
		Value:     req.Value,
		SeqNum:    req.SeqNum + 1,
		TimeStamp: time.Now(),
	}, nil
}

type ReqStruct struct {
	Input string
}

func (ReqStruct) JavaClassName() string {
	return "ReqStruct"
}

type RspStruct struct {
	Output string
}

func (RspStruct) JavaClassName() string {
	return "RspStruct"
}

// 最后一个参数一定为回包channel
func (g GloryProvider) SayHi(ctx context.Context, sClient chan ReqStruct, sClient2 chan ReqStruct, sRsp chan RspStruct) error {
	//for {
	req := <-sClient
	req2 := <-sClient2
	log.Info("SayHi get req = ", req)
	sRsp <- RspStruct{
		Output: req.Input + req2.Input + " pong",
	}
	//}
	return nil
}

type ReqBody struct {
	ID     string
	SeqNum int
	Value  string
}

func (ReqBody) JavaClassName() string {
	return "ReqBody"
}

type RspBody struct {
	SeqNum    int
	Value     string
	TimeStamp time.Time
}

func (RspBody) JavaClassName() string {
	return "RspBody"
}
func init() {
	hessian.RegisterPOJO(&ReqBody{})
	hessian.RegisterPOJO(&RspBody{})
	hessian.RegisterPOJO(&ReqStruct{})
	hessian.RegisterPOJO(&RspStruct{})
}

func main() {
	gloryServer := glory.NewServer()
	gloryService := service.NewGloryService("gloryService", &GloryProvider{})
	gloryServer.RegisterService(gloryService)
	gloryServer.Run()
}
