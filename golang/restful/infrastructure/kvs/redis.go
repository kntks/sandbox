package kvs

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"golang.org/x/xerrors"
)

type Redis struct {
	client *redis.Client
}

func NewRedis() (*Redis, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, xerrors.Errorf("fail to connect redis:%w", err)
	}

	return &Redis{client}, nil
}

var ctx = context.Background()

func (r *Redis) Set(key string, value interface{}, expire time.Duration) error {
	return r.client.Set(ctx, key, value, expire).Err()
}

func (r *Redis) Get(key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *Redis) Del(key string) error {
	return r.client.Del(ctx, key).Err()
}
