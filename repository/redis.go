package repository

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type RedisRepositoryInterface interface {
	Keys(pattern string) []string
	Get(key string) string
	Set(key string, value string)
	Delete(key string)
}

type RedisRepository struct{}

func (repo *RedisRepository) client() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	return client
}

func (repo *RedisRepository) Keys(pattern string) []string {
	vals, _ := repo.client().Keys(context.Background(), pattern).Result()
	return vals
}

func (repo *RedisRepository) Get(key string) string {
	val, err := repo.client().Get(context.Background(), key).Result()
	if err != nil {
		val = ""
	}
	return val
}

func (repo *RedisRepository) Set(key string, value string) {
	err := repo.client().Set(context.Background(), key, value, 0).Err()
	if err != nil {
		fmt.Printf("%-v", err)
	}
}

func (repo *RedisRepository) Delete(key string) {
	err := repo.client().Del(context.Background(), key).Err()
	if err != nil {
		fmt.Printf("%-v", err)
	}
}
