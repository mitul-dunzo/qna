package config

import (
	r "github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"time"
)

type Redis struct {
	*r.Client
}

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

func Save(key string, value string, t time.Duration) error {
	err := redis.Set(key, value, t).Err()
	if err != nil {
		logrus.Error("Redis: Couldn't save ", value, " for ", key)
		return err
	}
	return nil
}
