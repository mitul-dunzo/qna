package mocks

import (
	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis"
)

func NewMockRedis() (*redis.Client, *miniredis.Miniredis) {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})
	return client, mr
}
