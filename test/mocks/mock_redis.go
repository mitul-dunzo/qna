package mocks

import (
	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis"
)

var mr *miniredis.Miniredis

func NewMockRedis() *redis.Client {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	client := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})
	mr = s
	return client
}

func StopMockRedis() {
	mr.Close()
}
