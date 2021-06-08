package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/glory-go/glory/glory"
	ghttp "github.com/glory-go/glory/http"
	"github.com/glory-go/glory/log"
	"github.com/glory-go/glory/metrics"

	//"github.com/glory-go/glory/mongodb"
	//"github.com/glory-go/glory/mysql"
	_ "github.com/glory-go/glory/registry/redis"
	"github.com/glory-go/glory/service"
)

// 定义 request 结构体，支持 validate 标签校验，具体语法参考：https://godoc.org/github.com/go-playground/validator
type gloryHttpReq struct {
	Input    []int  `schema:"input" validate:"required"` // query参数使用schema 标签
	BodyStr  string `json:"body_str"`                    // body 参数使用json标签
	BodyStr2 string `json:"body_str_2"`                  // body 参数使用json标签
}

// 定义 response 结构体
type gloryHttpRsp struct {
	Output int `schema:"output"`
}

// 自定义mysqlModel
type myMysqlModel struct {
	UserID     uint32     `gorm:"user_id;primary_key"`
	Username   string     `gorm:"username"`
	Password   string     `gorm:"password"`
	createTime *time.Time `gorm:"create_time"`
}

func (m *myMysqlModel) TableName() string {
	return "user"
}

// 自定义业务逻辑处理 handler
func testHandler(controller *ghttp.GRegisterController) error {
	metrics.CounterInsc("test_handler_metrics")
	req := controller.Req.(*gloryHttpReq)
	rsp := controller.Rsp.(*gloryHttpRsp)
	log.Info("req = ", fmt.Sprintf("%+v", req))
	rsp.Output = req.Input[0] + 1

	// url注册的/req/{var} 变量都会解析到controller.VarsMap[]里面
	log.Info(controller.VarsMap["hello"])
	controller.RspCode = 408
	for _, v := range req.Input {
		fmt.Println(v)
	}
	fmt.Printf("reseive request : %+v", req)
	return nil
}

// 测试用filter 如果input字段为-1则报错
func myFilter1(controller *ghttp.GRegisterController, f ghttp.HandleFunc) error {
	req, ok := controller.Req.(*gloryHttpReq)
	if !ok {
		log.Error("req type err")
		return errors.New("req type err")
	}

	if req.Input[0] == -1 {
		log.Error("filting because input == -1")
		return errors.New("filting because input == -1")
	}
	err := f(controller)
	return err
}

// 测试用filter 如果input字段为-2则报错
func myFilter2(controller *ghttp.GRegisterController, f ghttp.HandleFunc) error {
	req, ok := controller.Req.(*gloryHttpReq)
	if !ok {
		log.Error("req type err")
		return errors.New("req type err")
	}

	if req.Input[0] == -2 {
		log.Error("filting because input == -2")
		return errors.New("filting because input == -2")
	}
	err := f(controller)
	return err
}

func main() {
	gloryServer := glory.NewServer()
	//// 这些都可以写在配置里面的，但目前不希望配置过重。
	// create http service example
	httpService := service.NewHttpService("httpDemo")
	httpService.RegisterRouter("/testwithfilter/{hello}/{hello2}", testHandler, &gloryHttpReq{}, &gloryHttpRsp{}, "POST", myFilter1, myFilter2)
	gloryServer.RegisterService(httpService)
	gloryServer.Run()
}
