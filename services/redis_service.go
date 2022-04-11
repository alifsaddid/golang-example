package services

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisService interface {
	GetCache(key string) (string, bool)
	WriteCache(key string, value interface{}) bool
}

type redisService struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisService(client *redis.Client) RedisService {
	return &redisService{
		client: client,
		ctx:    context.Background(),
	}
}

func (r redisService) GetCache(key string) (string, bool) {
	data, err := r.client.Get(r.ctx, key).Result()
	if err != nil {
		return "", false
	}
	return data, true
}

func (r redisService) WriteCache(key string, value interface{}) bool {
	err := r.client.Set(r.ctx, key, value, time.Duration(10*time.Minute)).Err()
	return err == nil
}
