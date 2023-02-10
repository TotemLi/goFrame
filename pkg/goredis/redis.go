package goredis

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	"goFrame/pkg/stringx"
	"time"
)

var ErrNotFind = redis.Nil
var ErrMarshal = errors.New("marshal error")

type RedisClient struct {
	client *redis.Client
}

func NewClient(addr, password string, db int) RedisClient {
	var r = RedisClient{}
	r.client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return r
}

func (r *RedisClient) Set(ctx context.Context, key string, data any, tm ...int64) error {
	var expiration time.Duration = 0
	if tm != nil && tm[0] > 0 {
		expiration = time.Duration(tm[0]) * time.Second
	}
	value, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = r.client.Set(ctx, key, value, expiration).Err()
	return err
}

func (r *RedisClient) Get(ctx context.Context, key string, ptr any) error {
	valStr, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return err
	}
	return stringx.StrCopyToPtr(valStr, ptr)
}

func (r *RedisClient) Incr(ctx context.Context, key string, tm ...int64) (int64, error) {
	val, err := r.client.Incr(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	// 设置过期时间
	var expiration time.Duration = 0
	if tm != nil && tm[0] > 0 {
		expiration = time.Duration(tm[0]) * time.Second
		err = r.client.Expire(ctx, key, expiration).Err()
	}
	return val, err
}
