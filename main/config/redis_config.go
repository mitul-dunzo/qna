package config

import (
	r "github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

var redis *r.Client

func SetupRedis() {
	address := "127.0.0.1:6379"

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
