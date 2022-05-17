package rediskeeper

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"secretservice/storage/keeper"
	"time"
)

//var Keep = getRedisKeeper()

const TTL = 100 * time.Second

type RedisKeeper struct {
	cn  redis.Client
	ctx context.Context
}

func GetRedisKeeper() keeper.Keeper {
	return RedisKeeper{*redis.NewClient(
		&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		}), context.Background()}
}

func (d RedisKeeper) Get(key string) (string, error) {
	val, err := d.cn.GetDel(d.ctx, key).Result()
	if err == redis.Nil {
		return "", errors.New("message not found")
	}
	return val, err
}

func (d RedisKeeper) Set(key, message string) error {
	return d.cn.Set(d.ctx, key, message, TTL).Err()
}

func (d RedisKeeper) Delete(key string) error {
	return d.cn.Del(d.ctx, key).Err()
}
