package Redis

import (
	"github.com/go-redis/redis"
)

//定义redis连接池
//初始化
func NewRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		PoolSize: 10,
	})
}
