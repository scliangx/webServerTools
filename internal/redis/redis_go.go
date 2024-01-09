package redis

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/scliangx/webServerTools/config"
)

var redisPool *redis.Pool

// InitRedisClientPool 处于程序底层的包，init 初始化的代码段的执行会优先于上层代码，因此这里读取配置项不能使用全局配置项变量
func InitRedisClientPool(cfg config.RedisConfig) *redis.Pool {
	redisPool = &redis.Pool{
		MaxIdle:     cfg.MaxIdleConns,                             //最大空闲数
		MaxActive:   cfg.MaxActive,                                //最大活跃数
		IdleTimeout: time.Duration(cfg.ConnTimeout) * time.Second, //最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭
		Dial: func() (redis.Conn, error) {
			//此处对应redis ip及端口号
			conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", cfg.Address, cfg.Port))
			if err != nil {
				return nil, err
			}
			auth := cfg.Password //通过配置项设置redis密码
			if len(auth) >= 1 {
				if _, err := conn.Do("AUTH", auth); err != nil {
					_ = conn.Close()
				}
			}
			_, _ = conn.Do("select", cfg.IndexDb)
			return conn, err
		},
	}

	return redisPool
}

// GetOneRedisClient 从连接池获取一个redis连接
func GetOneRedisClient() *RedisClient {
	maxRetryTimes := config.GetConfig().Redis.MaxRetryTimes
	var oneConn redis.Conn
	for i := 1; i <= maxRetryTimes; i++ {
		oneConn = redisPool.Get()
		if oneConn.Err() != nil {
			//variable.ZapLog.Error("Redis：网络中断,开始重连进行中..." , zap.Error(oneConn.Err()))
			if i == maxRetryTimes {
				return nil
			}
			//如果出现网络短暂的抖动，短暂休眠后，支持自动重连
			time.Sleep(time.Second * time.Duration(config.GetConfig().Redis.ReConnectInterval))
		} else {
			break
		}
	}
	return &RedisClient{oneConn}
}

// RedisClient 定义一个redis客户端结构体
type RedisClient struct {
	client redis.Conn
}

// Execute 为redis-go 客户端封装统一操作函数入口
func (r *RedisClient) Execute(cmd string, args ...interface{}) (interface{}, error) {
	return r.client.Do(cmd, args...)
}

// ReleaseOneRedisClient 释放连接到连接池
func (r *RedisClient) ReleaseOneRedisClient() {
	_ = r.client.Close()
}

// Bool bool 类型转换
func (r *RedisClient) Bool(reply interface{}, err error) (bool, error) {
	return redis.Bool(reply, err)
}

// string 类型转换
func (r *RedisClient) String(reply interface{}, err error) (string, error) {
	return redis.String(reply, err)
}

// StringMap string map 类型转换
func (r *RedisClient) StringMap(reply interface{}, err error) (map[string]string, error) {
	return redis.StringMap(reply, err)
}

// Strings strings 类型转换
func (r *RedisClient) Strings(reply interface{}, err error) ([]string, error) {
	return redis.Strings(reply, err)
}

// Float64 Float64 类型转换
func (r *RedisClient) Float64(reply interface{}, err error) (float64, error) {
	return redis.Float64(reply, err)
}

// Int int 类型转换
func (r *RedisClient) Int(reply interface{}, err error) (int, error) {
	return redis.Int(reply, err)
}

// Int64 int64 类型转换
func (r *RedisClient) Int64(reply interface{}, err error) (int64, error) {
	return redis.Int64(reply, err)
}

// IntMap int map 类型转换
func (r *RedisClient) IntMap(reply interface{}, err error) (map[string]int, error) {
	return redis.IntMap(reply, err)
}

// Int64Map int64map类型转换
func (r *RedisClient) Int64Map(reply interface{}, err error) (map[string]int64, error) {
	return redis.Int64Map(reply, err)
}

// Int64s int64s 类型转换
func (r *RedisClient) Int64s(reply interface{}, err error) ([]int64, error) {
	return redis.Int64s(reply, err)
}

// Uint64 uint64 类型转换
func (r *RedisClient) Uint64(reply interface{}, err error) (uint64, error) {
	return redis.Uint64(reply, err)
}

//Bytes bytes类型转换
func (r *RedisClient) Bytes(reply interface{}, err error) ([]byte, error) {
	return redis.Bytes(reply, err)
}

func (r *RedisClient) Get(key string) interface{} {
	val, err := r.client.Do("GET", key)
	if err != nil {
		return nil
	}
	return val
}

func (r *RedisClient) Set(key, val string) bool {
	_, err := r.client.Do("SET", key, val)
	if err != nil {
		return false
	}
	return true
}

// 以上封装了很多最常见类型转换函数，其他您可以参考以上格式自行封装
