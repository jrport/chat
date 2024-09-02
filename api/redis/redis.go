package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

func GetClient() (*redis.Client, error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
	defer cancel()

	rdb := redis.NewClient(&redis.Options{
		Addr: ":6379",
	})

	status := rdb.Ping(ctx)
	if err := status.Err(); err != nil {
		return nil, err
	}

	return rdb, nil
}
