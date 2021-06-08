package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/glory-go/glory/service"

	hessian "github.com/apache/dubbo-go-hessian2"

	_ "github.com/glory-go/glory/config"
	"github.com/glory-go/glory/glory"
	"github.com/glory-go/glory/log"

	//_ "github.com/glory-go/glory/metrics"
	_ "github.com/glory-go/glory/protocol/glory"
	_ "github.com/glory-go/glory/registry/k8s"
)

type GloryProvider struct {
}

func (g GloryProvider) SayHello(ctx context.Context, req *ReqBody, str2 string) (*RspBody, error) {
	log.Info("req = ", *req, "+", str2)
	fmt.Println(time.Now().String())
	return &RspBody{
		Value:     req.Value,
		SeqNum:    req.SeqNum + 1,
		TimeStamp: time.Now(),
	}, errors.New("user defined error")
}

func (g GloryProvider) SayHi(ctx context.Context, req string) (string, error) {
	log.Info("SayHi get req = ", req)
	return "rsp", nil
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
}

func main() {
	gloryServer := glory.NewServer()
	gloryService := service.NewGloryService("gloryService", &GloryProvider{})
	gloryServer.RegisterService(gloryService)
	gloryServer.Run()
}
