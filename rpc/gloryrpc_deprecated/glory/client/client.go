package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	hessian "github.com/apache/dubbo-go-hessian2"

	"github.com/glory-go/glory/glory"
	"github.com/glory-go/glory/metrics"

	_ "github.com/glory-go/glory/config"
	_ "github.com/glory-go/glory/load_balance/round_robin"
	_ "github.com/glory-go/glory/metrics"
	_ "github.com/glory-go/glory/protocol/glory"
	_ "github.com/glory-go/glory/registry/nacos"
)

// 定义
type GloryClient struct {
	//SayHello func(ctx context.Context, s string) (string, error)
	SayHello func(ctx context.Context, req *ReqBody, str2 string) (*RspBody, error)
	SayHi    func(ctx context.Context, req string) (string, error)
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

var gloryClient GloryClient

func init() {
	hessian.RegisterPOJO(&ReqBody{})
	hessian.RegisterPOJO(&RspBody{})
}

func main() {
	metrics.CounterInsc("main_metrics")

	glory.NewClient(context.Background(), "gloryClient", &gloryClient)
	// test sayHi
	//rsp, err := gloryClient.SayHi(context.Background(), "say hi request")
	//
	//log.Info(rsp, err)

	// test say hello
	wg := sync.WaitGroup{}
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			timeStart := time.Now()
			rsp, err := gloryClient.SayHello(context.Background(), &ReqBody{
				SeqNum: 1000,
				Value:  "payload string",
				ID:     "24234",
			}, "req2")
			fmt.Println("----timeCost = ", time.Now().Sub(timeStart))
			fmt.Printf("get rsp = %+v, err = %+v\n", *rsp, err)
			fmt.Println(rsp.TimeStamp.String())
			wg.Done()
		}()
	}
	wg.Wait()
}
