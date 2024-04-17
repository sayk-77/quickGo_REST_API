package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
)

var cts = context.Background()

func ConnectRedis() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDRESS"),
		DB:   0,
	})
	_, err := rdb.Ping(cts).Result()
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к Redis: %v", err)
	}

	return rdb, nil
}
