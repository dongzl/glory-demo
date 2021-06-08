package main

import (
	"time"

	"github.com/glory-go/glory/log"

	_ "github.com/glory-go/glory/registry/redis"
)

func main() {
	for {
		time.Sleep(time.Second)
		log.Debug("debug level log")
		log.Info("info level log")
		log.Error("hello glory go framework!")
	}
}
