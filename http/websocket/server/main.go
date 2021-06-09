package main

import (
	"github.com/glory-go/glory/glory"
	ghttp "github.com/glory-go/glory/http"
	_ "github.com/glory-go/glory/registry/redis"
	"github.com/glory-go/glory/service"
)

type gloryWSReq struct {
	BodyStr  string `json:"body_str"`   // body 参数使用json标签
	BodyStr2 string `json:"body_str_2"` // body 参数使用json标签
}

type gloryWSRsp struct {
	BodyStr  string `json:"body_str"`   // body 参数使用json标签
	BodyStr2 string `json:"body_str_2"` // body 参数使用json标签
}

// 自定义业务逻辑处理 handler
func testWSHandler(controller *ghttp.GRegisterWSController) {
	req := &gloryWSReq{}
	rsp := &gloryWSRsp{}
	controller.WSConn.ReadJSON(req)
	/*
		your code
	*/
	controller.WSConn.WriteJSON(rsp)
}

func main() {
	gloryServer := glory.NewServer()
	httpService := service.NewHttpService("httpDemo")
	httpService.RegisterWSRouter("/test/{hello}/{hello2}", testWSHandler)
	gloryServer.RegisterService(httpService)
	gloryServer.Run()
}
