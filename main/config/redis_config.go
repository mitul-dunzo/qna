package config

import (
	r "github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"os"
)

var redis *r.Client

func SetupRedis() {
	address := os.Getenv("RedisAddress")

	client := r.NewClient(&r.Options{
		Addr: address,
	})
	pong, err := client.Ping().Result()
	if err != nil {
		logrus.Fatal("Failed to connect to redis: ", err.Error())
	}
	logrus.Info("Successfully connected to redis: ", pong)
	redis = client
}

func GetRedis() *r.Client {
	return redis
}
