package redis

import (
	"github.com/Mobo140/auth/internal/client/cache"
	"github.com/Mobo140/auth/internal/config"
	"github.com/gomodule/redigo/redis"
)

var _ cache.Client = (*redisClient)(nil)

type redisClient struct {
	masterRedis cache.Cache
}

func NewClient(pool *redis.Pool, config config.RedisConfig) *redisClient { //nolint:revive // it's ok
	return &redisClient{
		masterRedis: NewRD(pool, config),
	}
}

func (c *redisClient) Cache() cache.Cache {
	return c.masterRedis
}

func (c *redisClient) Close() error {
	if c.masterRedis != nil {
		c.masterRedis.Close()
	}

	return nil
}
