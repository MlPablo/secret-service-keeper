package rediskeeper

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"secretservice/keeper"
)

var Keep = getRedisKeeper()

const TTL = 0

type RedisKeeper struct {
	cn  redis.Client
	ctx context.Context
}

func getRedisKeeper() keeper.Keeper {
	return RedisKeeper{*redis.NewClient(
		&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		}), context.Background()}
}

func (d RedisKeeper) Get(key string) (string, error) {
	val, err := d.cn.Get(d.ctx, key).Result()
	if err == redis.Nil {
		return "", errors.New("Not found")
	}
	return val, err
}

func (d RedisKeeper) Set(key, message string) error {
	return d.cn.Set(d.ctx, key, message, TTL).Err()
}

func (d RedisKeeper) Delete(key string) error {
	return d.cn.Del(d.ctx, key).Err()
}
