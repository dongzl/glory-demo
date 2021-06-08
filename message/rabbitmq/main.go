package main

import (
	"fmt"
	"time"

	"github.com/glory-go/glory/mq"
	_ "github.com/glory-go/glory/registry/redis"
)

func main() {

	if err := mq.StartOnMQMsgHandler("testmq", "testqueue", func(msg []byte) error {
		// handler async func, developers can define it as they like
		fmt.Println("handler get data = ", string(msg))
		// if return error, it only means that this receive action break up.
		// the handler is still work the next time it receive msg.
		return nil
	}); err != nil {
		panic(err)
	}

	time.Sleep(time.Second)
	// 因为启动handler是异步操作，不是立刻就handler的，需要等一秒钟handler开启了监听，防止运行demo时出现丢失mq数据丢失
	// 默认不是持久化的，持久化以及其他mq配置可以之后以option的形式写入，看开发者的需要吧。
	for i := 0; i < 10; i++ {
		if err := mq.SendMQMessage("testmq", "testqueue", []byte("hello world")); err != nil {
			panic(err)
		}
	}
	time.Sleep(time.Second * 10)
}
