package rediskeeper

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"secretservice/storage/keeper"
	"time"
)

//var Keep = getRedisKeeper()

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

func (d RedisKeeper) Set(key, message string, ttl int) error {
	err := d.cn.Set(d.ctx, key, message, time.Duration(ttl)*time.Second).Err()
	return err
}

func (d RedisKeeper) Delete(key string) error {
	return d.cn.Del(d.ctx, key).Err()
}
