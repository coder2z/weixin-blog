package Redis

import (
	"github.com/go-redis/redis"
	"wx-blog/config"
)

//定义redis连接池
//初始化
func NewRedis(c *config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.Redis.Host,
		Password: c.Redis.Password,
		DB:       c.Redis.DB,
		PoolSize: c.Redis.PoolSize,
	})
}
