package shorten

import (
	"github.com/go-redis/redis"
	"strconv"
)

const (
	redisKey      = "_shortUrl:id"
	redisKeyDebug = "_shortUrl:id:debug"
)

// RedisSerialNoGenerator 通过Redis获取唯一ID
type RedisSerialNoGenerator struct {
	C *redis.Client
	K string
}

// NewRedisSerialNoGenerator Get instance
func NewRedisSerialNoGenerator(c *redis.Client, debug bool, idStart uint64) *RedisSerialNoGenerator {
	g := &RedisSerialNoGenerator{C: c, K: redisKey}
	if debug {
		g.K = redisKeyDebug
	}
	if idStart > 0 {
		v, err := c.Get(g.K).Int64()
		if (err != nil && err == redis.Nil) || (err == nil && uint64(v) < idStart) {
			c.Set(g.K, strconv.FormatUint(idStart, 10), 0)
		}
	}
	return g
}

// ID 通过Redis获取全局唯一自增ID
func (g *RedisSerialNoGenerator) ID() uint64 {
	i, err := g.C.Incr(g.K).Result()
	if err != nil {
		panic(err)
	}
	return uint64(i)
}
