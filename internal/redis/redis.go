package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/scliang-strive/webServerTools/config"
	"github.com/sirupsen/logrus"
)

var redisCli *redis.Client

func NewRedisClient()  {
	cfg := config.GetConfig().Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d",cfg.Address,cfg.Port),
		Password:     cfg.Password,
		DB:           cfg.IndexDb,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})

	if err := rdb.Ping().Err(); err != nil {
		logrus.Error("redis connection failed.",err)
		return
	}
	redisCli = rdb
	return
}

func GetClient() *redis.Client{
	return redisCli
}
