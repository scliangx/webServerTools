package cache

import (
	"github.com/go-redis/redis"
	"github.com/scliang-strive/webServerTools/config"
	"time"
)

func NewRedisClient() *redis.Client {
	cfg := config.GetConfig().Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:         cfg.Address,
		Password:     cfg.Password,
		DB:           cfg.DB,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})

	if err := rdb.Ping(); err != nil {
		return nil
	}
	return rdb
}
