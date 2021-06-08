package main

import (
	"github.com/glory-go/glory/log"
	"github.com/glory-go/glory/redis"
	_ "github.com/glory-go/glory/registry/nacos"
)

func main() {
	redisClient, err := redis.NewRedisClient("redisKey", 0)
	if err != nil {
		log.Errorf("GetIMRedisClientInstance err = %s", err)
	}
	redisClient.Set("name", "laurence", 0)
	log.Infof("set name to laurence")
	val, err := redisClient.Get("name").Result()
	if err != nil {
		panic(err)
	}
	log.Infof("get name %s", val)
}
