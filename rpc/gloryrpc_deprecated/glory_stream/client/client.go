package main

import (
	"context"
	"sync"
	"time"

	"github.com/glory-go/glory/log"

	"github.com/glory-go/glory/glory"
	"github.com/glory-go/glory/metrics"

	hessian "github.com/apache/dubbo-go-hessian2"

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
	SayHi    func(ctx context.Context) (chan ReqStruct, chan ReqStruct, chan RspStruct, error)
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
	hessian.RegisterPOJO(&ReqStruct{})
	hessian.RegisterPOJO(&RspStruct{})
}

func main() {
	metrics.CounterInsc("main_metrics")

	glory.NewClient(context.Background(), "gloryClient", &gloryClient)

	// test sayHi
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			sreq, sreq2, srsp, _ := gloryClient.SayHi(context.Background())
			//if err != nil {
			//	fmt.Println(err)
			//	wg.Done()
			//	return
			//}
			if sreq == nil || srsp == nil {
				log.Errorf("sreq == nil or srsp == nil")
				return
			}
			sreq2 <- ReqStruct{
				Input: "ping2",
			}
			sreq <- ReqStruct{
				Input: "ping",
			}
			rsp := <-srsp
			log.Info(rsp)
			wg.Done()
		}()

	}
	wg.Wait()

	// todo 高并发场景下存在安全问题 test say hello
	//for i := 0; i < 5; i++ {
	//	go func() {
	//		timeStart := time.Now()
	//		rsp, _ := gloryClient.SayHello(context.Background(), &ReqBody{
	//			SeqNum: 1000,
	//			Value:  "payload string",
	//			ID:     "24234",
	//		}, "req2")
	//		fmt.Println("----timeCost = ", time.Now().Sub(timeStart))
	//		fmt.Printf("get rsp = %+v", *rsp)
	//		fmt.Println(rsp.TimeStamp.String())
	//	}()
	//}
	//time.Sleep(time.Second * 10)
}
